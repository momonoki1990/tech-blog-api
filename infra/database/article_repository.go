package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	"github.com/momonoki1990/tech-blog-api/domain/repository"
	dbModel "github.com/momonoki1990/tech-blog-api/infra/database/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ArticleRepository struct {
	ctx context.Context
	exec boil.ContextExecutor
}

func NewArticleRepository(ctx context.Context,  exec boil.ContextExecutor) repository.ArticleRepository {
    return &ArticleRepository{ctx, exec}
}

func (r *ArticleRepository)FindOneById(id uuid.UUID) (*model.Article, error) {
	dbArticle, err := dbModel.Articles(dbModel.ArticleWhere.ID.EQ(id.String())).One(r.ctx, r.exec)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if dbArticle == nil {
		return nil, nil
	}
	
	article, err := toArticle(dbArticle, r)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (r *ArticleRepository) Find() ([]*model.Article, error) {
	dbArticles, err := dbModel.Articles().All(r.ctx, r.exec)
	if err != nil {
		return nil, err
	}
	articles, err := toArticles(dbArticles, r)
	return articles, nil
}

func (r *ArticleRepository) Insert(c *model.Article) (error) {
	dbArticle, err := toDbArticle(c)
	if err != nil {
		return err
	}
	err = dbArticle.Insert(r.ctx, r.exec, boil.Infer())
	if err != nil {
		return err
	}

	dbTags:= toDbTags(c)
	for _, v := range dbTags {
		err = v.Upsert(r.ctx, r.exec, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
	}
	
	dbTaggings := toDbTaggings(c)
	for _, v := range dbTaggings {
		err = v.Insert(r.ctx, r.exec, boil.Infer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ArticleRepository) Update(a *model.Article) (error) {
	dbArticle, err := dbModel.FindArticle(r.ctx, r.exec, a.Id.String())
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if dbArticle == nil {
		return errors.New("Article to update was not found")
	}

	var publishedAt null.Time
	if a.PublishedAt == nil {
		publishedAt = null.TimeFromPtr(nil)
	} else {
		publishedAt = null.TimeFromPtr(a.PublishedAt)
	}
	dbArticle.Title = a.Title
	dbArticle.Content = a.Content
	dbArticle.CategoryID = a.CategoryId.String()
	dbArticle.Status = a.Status.String()
	dbArticle.PublishedAt = publishedAt
	dbArticle.CreatedAt = a.CreatedAt
	dbArticle.UpdatedAt = a.UpdatedAt

	rowsAff, err := dbArticle.Update(r.ctx, r.exec, boil.Infer())
	if err != nil {
		return err
	}
	// NOTE: タグだけ変更があった場合はdbArticle.UpdateのrowsAffは0になる
	if rowsAff != 0 && rowsAff != 1  {
		return errors.New(fmt.Sprintf("Number of rows affected by update is invalid %v", rowsAff))
	}

	// タグの処理
	foundDbTaggings, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(a.Id.String())).All(r.ctx, r.exec)
	if err != nil {
		return err
	}

	// 元々のタグ付けと比較して、追加・削除されたタグごとに処理(タグマスタも必要に応じて追加・削除)
	var foundTaggingTagNames []string
	for _, v := range foundDbTaggings {
		foundTaggingTagNames = append(foundTaggingTagNames, v.TagName)
	}

	var tagNames []string
	for _, v := range a.Tags {
		tagNames = append(tagNames, v.Name)
	}

	var addedtagNames []string
	var removedtagNames []string

	for _, v := range tagNames {
		included := false
		for _, v2 := range foundTaggingTagNames {
			if v == v2 {
				included = true
			}
		}
		if !included {
			addedtagNames = append(addedtagNames, v)
		}
	}

	for _, v := range foundTaggingTagNames {
		included := false
		for _, v2 := range tagNames {
			if v == v2 {
				included = true
			}
		}
		if !included {
			removedtagNames = append(removedtagNames, v)
		}
	}

	for _, v := range addedtagNames {
		foundTag, err := dbModel.Tags(dbModel.TagWhere.Name.EQ(v)).One(r.ctx, r.exec)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if foundTag == nil {
			tag := &dbModel.Tag{
				Name: v,
			}
			err = tag.Insert(r.ctx, r.exec, boil.Infer())
			if err != nil {
				return err
			}
		}
	}

	// タグ付けは洗い替え
	for _, v := range foundDbTaggings {
		v.Delete(r.ctx, r.exec)
	}

	for _, v := range a.Tags {
		dbTagging := &dbModel.Tagging{
			ArticleID: a.Id.String(),
			TagName: v.Name,
		}
		err := dbTagging.Insert(r.ctx, r.exec, boil.Infer())
		if err != nil {
			return err
		}
	}

	// tagsの削除はtaggingsの処理の後で（外部キー制約に引っかかるので）
	for _, v := range removedtagNames {
		foundTagging, err := dbModel.Taggings(dbModel.TaggingWhere.TagName.EQ(v), dbModel.TaggingWhere.ArticleID.NEQ(a.Id.String())).One(r.ctx, r.exec)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if foundTagging == nil {
			foundTag, err := dbModel.Tags(dbModel.TagWhere.Name.EQ(v)).One(r.ctx, r.exec)
			if err != nil {
				return err
			}
			rowsAff, err := foundTag.Delete(r.ctx, r.exec)
			if err != nil {
				return err
			}
			if rowsAff != 1 {
				return errors.New(fmt.Sprintf("Number of rows affected by delete is invalid %v", rowsAff))
			}
		}
	}

	return nil
}

// taggingも削除、tagもチェック
func (r *ArticleRepository) Delete(id uuid.UUID) (error) {
	dbArticle, err := dbModel.FindArticle(r.ctx, r.exec, id.String())
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if dbArticle == nil {
		return errors.New("Article to delete was not found")
	}

	// タグの処理
	foundDbTaggings, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(dbArticle.ID)).All(r.ctx, r.exec)
	// TODO: タグが元々ない場合の処理も確認
	if err != nil {
		return err
	}
	for _, v := range foundDbTaggings {
		shouldDeleteTag := false
		foundDbTagging, err := dbModel.Taggings(dbModel.TaggingWhere.TagName.EQ(v.TagName), dbModel.TaggingWhere.ArticleID.NEQ(dbArticle.ID)).One(r.ctx, r.exec)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if foundDbTagging == nil {
			shouldDeleteTag = true
		}
		rowsAff, err := v.Delete(r.ctx, r.exec)
		if rowsAff != 1 {
			return errors.New(fmt.Sprintf("Number of rows affected by tagging delete is invalid %d", rowsAff))
		}
		if shouldDeleteTag {
			foundDbTag, err := dbModel.FindTag(r.ctx, r.exec, v.TagName)
			if err != nil {
				return err
			}
			rowsAff, err = foundDbTag.Delete(r.ctx, r.exec)
			if err != nil {
				return err
			}
			if rowsAff != 1 {
				return errors.New(fmt.Sprintf("Number of rows affected by tag delete is invalid %d", rowsAff))
			}
		}
	}

	rowsAff, err := dbArticle.Delete(r.ctx, r.exec)
	if err != nil {
		return err
	}
	if rowsAff != 1 {
		return errors.New(fmt.Sprintf("Number of rows affected by delete is invalid %v", rowsAff))
	}
	return nil
}

func toStatus(s string) (*model.Status, error) {
	var status model.Status
	switch s {
	case "Draft":
		status = model.Draft
		return &status, nil
	case "Published":
		status = model.Published
		return &status, nil
	default:
		return nil, errors.New("記事のステータスの値が不正です")
	}
}

func toDbStatus(s model.Status) (string, error) {
	switch s {
	case model.Draft:
		return "Draft", nil
	case model.Published:
		return "Published", nil
	default:
		return "", errors.New("記事のステータスの値が不正です")
	}
}

func findTags(articleId string, r *ArticleRepository) ([]model.Tag, error) {
	dbTags, err := dbModel.Tags(
		qm.InnerJoin("taggings on taggings.tag_name = tags.name"),
		qm.Where("taggings.article_id = ?", articleId),
	).All(r.ctx, r.exec)
	if err != nil {
		return nil, err
	}
	var tags []model.Tag
	for _, v := range dbTags {
		tag := &model.Tag{
			Name: v.Name,
		}
		tags = append(tags, *tag)
	}
	return tags, nil
}

func toArticle(d *dbModel.Article, r *ArticleRepository) (*model.Article, error) {
	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}
	categoryId, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}
	status, err := toStatus(d.Status)
	if err != nil {
		return nil, err
	}
	tags, err := findTags(d.ID, r)
	
	var publishedAt *time.Time
	if d.PublishedAt.Valid {
		publishedAt = &d.PublishedAt.Time
	} else {
		publishedAt = nil
	}
	article := &model.Article{
		Id: id,
		Title: d.Title,
		Content: d.Content,
		CategoryId: categoryId,
		Tags: tags,
		PublishedAt: publishedAt,
		Status: *status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
	
	
	return article, nil
}

func toArticles(dbArticles []*dbModel.Article, r *ArticleRepository) ([]*model.Article, error) {
	var articles []*model.Article
	for _, v := range dbArticles {
		article, err := toArticle(v, r)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func toDbArticle(e *model.Article) (*dbModel.Article, error) {
	status, err := toDbStatus(e.Status)
	if err != nil {
		return nil, err
	}

	var publishedAt null.Time
	if e.PublishedAt == nil {
		publishedAt = null.TimeFromPtr(nil)
	} else {
		publishedAt = null.TimeFromPtr(e.PublishedAt)
	}
	dbArticle := &dbModel.Article{
		ID: e.Id.String(),
		Title: e.Title,
		Content: e.Content,
		CategoryID: e.CategoryId.String(),
		PublishedAt: publishedAt,
		Status: status,
	}
	return dbArticle, nil
}

func toDbTaggings(a *model.Article) ([]*dbModel.Tagging) {
	var dbTaggings []*dbModel.Tagging
	for _, v := range a.Tags {
		t := &dbModel.Tagging{
			ArticleID: a.Id.String(),
			TagName: v.Name,
		}
		dbTaggings = append(dbTaggings, t)
	}
	return dbTaggings
}


func toDbTags(a *model.Article) ([]*dbModel.Tag) {
	var dbTags []*dbModel.Tag
	for _, v := range a.Tags {
		t := &dbModel.Tag{
			Name: v.Name,
		}
		dbTags = append(dbTags, t)
	}
	return dbTags
}