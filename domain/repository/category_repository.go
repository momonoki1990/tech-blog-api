package repository

import (
	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
)

type CategoryRepository interface {
	FindOneByName(name string) (*model.Category, error)
	FindOneById(id uuid.UUID) (*model.Category, error)
	Find() ([]*model.Category, error)
	Insert(*model.Category) (error)
	Update(*model.Category) (error)
	Delete(id uuid.UUID) (error)
}