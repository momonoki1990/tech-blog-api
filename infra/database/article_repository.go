package database

import (
	"context"
	"errors"
	"fmt"

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
	if err != nil {
		return nil, err
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
	fmt.Println("😈Insert called")
	dbArticle, err := toDbArticle(c)
	fmt.Println("😈After toDbArticle")
	if err != nil {
		return err
	}
	err = dbArticle.Insert(r.ctx, r.exec, boil.Infer())
	if err != nil {
		return err
	}

	dbTags1, err := dbModel.Tags().All(r.ctx, r.exec)
	for _, v := range dbTags1 {
		fmt.Println("😈dbTags1[i]", v)
		fmt.Println("😈dbTags1[i].ID", v.ID)
		fmt.Println("😈dbTags1[i].Name", v.Name)
	}
	dbTags:= toDbTags(c)
	for _, v := range dbTags {
		fmt.Println("😈dbTags[i]", v)
		err = v.Upsert(r.ctx, r.exec, boil.Infer(), boil.Infer())
		if err != nil {
			fmt.Println("👹エラー起きてる", err)
			return err
		}
	}
	dbTags2, err := dbModel.Tags().All(r.ctx, r.exec)
	for _, v := range dbTags2 {
		fmt.Println("😈dbTags2[i]", v)
		fmt.Println("😈dbTags2[i].ID", v.ID)
		fmt.Println("😈dbTags2[i].Name", v.Name)
	}
	
	dbTaggings := toDbTaggings(c)
	for _, v := range dbTaggings {
		fmt.Println("😈dbTaggings[i]", v)
	}
	for _, v := range dbTaggings {
		fmt.Println(v.TagID)
		err = v.Insert(r.ctx, r.exec, boil.Infer())
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (r *ArticleRepository) Update(a *model.Article) (error) {
	dbArticle, err := dbModel.FindArticle(r.ctx, r.exec, a.Id.String())
	if err != nil {
		return err
	}
	if dbArticle == nil {
		return errors.New("変更対象の記事が見つかりませんでした")
	}
	dbArticle.Title = a.Title
	dbArticle.Content = a.Content
	dbArticle.CategoryID = a.CategoryId.String()
	dbArticle.Status = a.Status.String()
	dbArticle.PublishedAt = null.TimeFrom(*a.PublishedAt)
	dbArticle.CreatedAt = a.CreatedAt
	dbArticle.UpdatedAt = a.UpdatedAt
	dbArticle.Update(r.ctx, r.exec, boil.Infer())

	// タグの処理
	foundDbTaggings, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(a.Id.String())).All(r.ctx, r.exec)
	if err != nil {
		return err
	}

	// 元々のタグ付けと比較して、追加・削除されたタグごとに処理(タグマスタも必要に応じて追加・削除)
	var foundTaggingTagIds []string
	for _, v := range foundDbTaggings {
		foundTaggingTagIds = append(foundTaggingTagIds, v.TagID)
	}

	var tagIds []string
	for _, v := range a.Tags {
		tagIds = append(tagIds, v.Id.String())
	}

	var addedTagIds []string
	var removedTagIds []string

	for _, v := range tagIds {
		included := false
		for _, v2 := range foundTaggingTagIds {
			if v == v2 {
				included = true
			}
		}
		if !included {
			addedTagIds = append(addedTagIds, v)
		}
	}

	for _, v := range foundTaggingTagIds {
		included := false
		for _, v2 := range tagIds {
			if v == v2 {
				included = true
			}
		}
		if !included {
			removedTagIds = append(removedTagIds, v)
		}
	}

	for _, v := range addedTagIds {
		foundTag, err := dbModel.Tags(dbModel.TagWhere.ID.EQ(v)).One(r.ctx, r.exec)
		if err != nil {
			return err
		}
		if foundTag == nil {
			var found model.Tag
			for _, v2 := range a.Tags {
				if v == v2.Id.String() {
					found = v2
				}
			}
			tag := &dbModel.Tag{
				ID: v,
				Name: found.Name,
			}
			tag.Insert(r.ctx, r.exec, boil.Infer())
		}
	}

	for _, v := range removedTagIds {
		foundTagging, err := dbModel.Taggings(dbModel.TaggingWhere.TagID.EQ(v), dbModel.TaggingWhere.ArticleID.NEQ(a.Id.String())).One(r.ctx, r.exec)
		if err != nil {
			return err
		}
		if foundTagging == nil {
			foundTag, err := dbModel.Tags(dbModel.TagWhere.ID.EQ(v)).One(r.ctx, r.exec)
			if err != nil {
				return err
			}
			foundTag.Delete(r.ctx, r.exec)
		}
	}

	// タグ付けは洗い替え
	for _, v := range foundDbTaggings {
		v.Delete(r.ctx, r.exec)
	}

	for _, v := range a.Tags {
		dbTagging := &dbModel.Tagging{
			ArticleID: a.Id.String(),
			TagID: v.Id.String(),
		}
		err := dbTagging.Insert(r.ctx, r.exec, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ArticleRepository) Delete(id uuid.UUID) (error) {
	dbArticle, err := dbModel.FindArticle(r.ctx, r.exec, id.String())
	if err != nil {
		return err
	}
	if dbArticle == nil {
		return errors.New("対象の記事が見つかりません")
	}
	dbArticle.Delete(r.ctx, r.exec)
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
		qm.InnerJoin("taggings on taggings.tag_id = tags.id"),
		qm.Where("taggings.article_id = ?", articleId),
	).All(r.ctx, r.exec)
	if err != nil {
		return nil, err
	}
	var tags []model.Tag
	for _, v := range dbTags {
		tagId, err := uuid.Parse(v.ID)
		if err != nil {
			return nil, err
		}
		tag := &model.Tag{
			Id: tagId,
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
	
	article := &model.Article{
		Id: id,
		Title: d.Title,
		Content: d.Content,
		CategoryId: categoryId,
		Tags: tags,
		PublishedAt: &d.PublishedAt.Time,
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
	fmt.Println("😈After toDbStatus called")
	if err != nil {
		return nil, err
	}

	dbArticle := &dbModel.Article{
		ID: e.Id.String(),
		Title: e.Title,
		Content: e.Content,
		CategoryID: e.CategoryId.String(),
		PublishedAt: null.TimeFromPtr(e.PublishedAt),
		Status: status,
	}
	fmt.Println("😈After toDbStatus called2", dbArticle)
	return dbArticle, nil
}

func toDbTaggings(a *model.Article) ([]*dbModel.Tagging) {
	var dbTaggings []*dbModel.Tagging
	for _, v := range a.Tags {
		t := &dbModel.Tagging{
			ArticleID: a.Id.String(),
			TagID: v.Id.String(),
		}
		dbTaggings = append(dbTaggings, t)
	}
	return dbTaggings
}


func toDbTags(a *model.Article) ([]*dbModel.Tag) {
	var dbTags []*dbModel.Tag
	for _, v := range a.Tags {
		t := &dbModel.Tag{
			ID:  v.Id.String(),
			Name: v.Name,
		}
		dbTags = append(dbTags, t)
	}
	return dbTags
}