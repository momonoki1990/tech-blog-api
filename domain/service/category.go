package service

import (
	"errors"

	"github.com/momonoki1990/tech-blog-api/domain/model"
	"github.com/momonoki1990/tech-blog-api/domain/repository"
)

type CategoryCreator interface {
	Create(name string, displayOrder int) (*model.Category, error)
}

type categoryCreator struct {
	repository.CategoryRepository
}

func NewCategoryCreator(r repository.CategoryRepository) CategoryCreator {
	return &categoryCreator{r}
}

func (s *categoryCreator) Create(name string, displayOrder int) (*model.Category, error) {
	c, err := s.CategoryRepository.FindOneByName(name)
	if err != nil {
		return nil, err
	}
	if c != nil {
		return nil, errors.New("Category name is already registered")
	}
	c, err = model.NewCategory(name, displayOrder)
	return c, err
}

