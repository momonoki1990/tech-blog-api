package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	"github.com/momonoki1990/tech-blog-api/domain/repository"
	"github.com/momonoki1990/tech-blog-api/domain/service"
)

type CategoryUseCase interface {
    GetCategoryList() ([]*model.Category, error)
    RegisterCategory(name string, displayOrder int) (error)
	UpdateCategory(id uuid.UUID, name string, displayOrder int) (error)
	DeleteCategory(id uuid.UUID) (error)
}

type categoryUseCase struct {
    repository.CategoryRepository
	service.CategoryCreator
}

func NewCategoryUseCase(r repository.CategoryRepository, s service.CategoryCreator) CategoryUseCase {
    return &categoryUseCase{r, s}
}

func (u *categoryUseCase) GetCategoryList() ([]*model.Category, error) {
    categories, err := u.CategoryRepository.Find()
	return categories, err
}

func (u *categoryUseCase) RegisterCategory(name string, displayOrder int) (error) {
	c, err := u.CategoryCreator.Create(name, displayOrder)
	if err != nil {
		return err
	}
	err = u.CategoryRepository.Insert(c)
	return err
}

func (u *categoryUseCase) UpdateCategory(id uuid.UUID, name string, displayOrder int) (error) {
	c, err := u.CategoryRepository.FindOneById(id)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("Category to update was not found")
	}
	c.Name = name
	c.DisplayOrder = displayOrder
	err = u.CategoryRepository.Update(c)
	return err
}

func (u *categoryUseCase) DeleteCategory(id uuid.UUID) (error) {
	err := u.CategoryRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}