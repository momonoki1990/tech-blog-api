package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	"github.com/momonoki1990/tech-blog-api/domain/repository"
)

type ArticleUseCase interface {
    GetArticle(id uuid.UUID) (*model.Article, error)
    GetArticleList() ([]*model.Article, error)
    RegisterArticle(title string, content string, categoryId uuid.UUID, tagNames []string, shouldPublish bool) (string, error)
	UpdateArticle(id uuid.UUID, title string, content string, categoryId uuid.UUID, tagNames []string, shouldPublish bool) (error)
	DeleteArticle(id uuid.UUID) (error)
}

type articleUseCase struct {
    repository.ArticleRepository
}

func NewArticleUseCase(r repository.ArticleRepository) ArticleUseCase {
    return &articleUseCase{r}
}

func (u *articleUseCase) GetArticle(id uuid.UUID) (*model.Article, error) {
    article, err := u.ArticleRepository.FindOneById(id)
	return article, err
}

func (u *articleUseCase) GetArticleList() ([]*model.Article, error) {
    articles, err := u.ArticleRepository.Find()
	return articles, err
}

func (u *articleUseCase) RegisterArticle(title string, content string, categoryId uuid.UUID, tagNames []string, shouldPublish bool) (string, error) {
	article, err := model.NewArticle(title, content, categoryId, tagNames, shouldPublish)
	if err != nil {
		return "", err
	}
	err = u.ArticleRepository.Insert(article)
	if err != nil {
		return "", err
	}
	articleId := article.Id.String()
	return articleId, nil
}

func (u *articleUseCase) UpdateArticle(id uuid.UUID, title string, content string, categoryId uuid.UUID, tagNames []string, shouldPublish bool) (error) {
	article, err := u.ArticleRepository.FindOneById(id)
	if err != nil {
		return err
	}
	if article == nil {
		return errors.New("Article to update was not found")
	}

	article.Title = title
	article.Content = content
	article.CategoryId = categoryId
	article.SetTags(tagNames)
	if shouldPublish {
		article.SetStatus(model.Published)
	} else {
		article.SetStatus(model.Draft)
	}
	err = u.ArticleRepository.Update(article)
	return err
}

func (u *articleUseCase) DeleteArticle(id uuid.UUID) (error) {
	err := u.ArticleRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}