package repository

import (
	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
)

type ArticleRepository interface {
	FindOneById(id uuid.UUID) (*model.Article, error)
	Find() ([]*model.Article, error)
	Insert(*model.Article) (error)
	Update(*model.Article) (error)
	Delete(id uuid.UUID) (error)
}