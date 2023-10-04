package service

import (
	"errors"
	"fmt"

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
	fmt.Println(c)
	if err != nil {
		return nil, err
	}
	if c != nil {
		return nil, errors.New("同じ名前のカテゴリがすでに登録されています")
	}
	c, err = model.NewCategory(name, displayOrder)
	return c, err
}

