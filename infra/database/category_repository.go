package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	"github.com/momonoki1990/tech-blog-api/domain/repository"
	dbModel "github.com/momonoki1990/tech-blog-api/infra/database/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CategoryRepository struct {
	ctx context.Context
	exec boil.ContextExecutor
}

func NewCategoryRepository(ctx context.Context,  exec boil.ContextExecutor) repository.CategoryRepository {
    return &CategoryRepository{ctx, exec}
}

func (r *CategoryRepository)FindOneByName(name string) (*model.Category, error) {
	dbCategories, err := dbModel.Categories(dbModel.CategoryWhere.Name.EQ(name)).One(r.ctx, r.exec)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	category, err := toCategory(dbCategories)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository)FindOneById(id uuid.UUID) (*model.Category, error) {
	dbCategory, err := dbModel.Categories(dbModel.CategoryWhere.ID.EQ(id.String())).One(r.ctx, r.exec)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	category, err := toCategory(dbCategory)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) Find() ([]*model.Category, error) {
	dbCategories, err := dbModel.Categories().All(r.ctx, r.exec)
	if err == sql.ErrNoRows {
		return []*model.Category{}, nil
	}
	if err != nil {
		return nil, err
	}
	categories, err := toCategories(dbCategories)
	return categories, nil
}

func (r *CategoryRepository) Insert(c *model.Category) (error) {
	dbCategory := toDbCategory(c)
	err := dbCategory.Insert(r.ctx, r.exec, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) Update(c *model.Category) (error) {
	dbCategory, err := dbModel.FindCategory(r.ctx, r.exec, c.Id.String())
	if err == sql.ErrNoRows {
		return errors.New("Category to update was not found")
	}
	if err != nil {
		return err
	}
	dbCategory.Name = c.Name
	dbCategory.DisplayOrder = null.IntFrom(c.DisplayOrder)
	dbCategory.Update(r.ctx, r.exec, boil.Infer())
	return nil
}

func (r *CategoryRepository) Delete(id uuid.UUID) (error) {
	dbCategory, err := dbModel.FindCategory(r.ctx, r.exec, id.String())
	if err == sql.ErrNoRows {
		return errors.New("Category to delete was not found")
	}
	if err != nil {
		return err
	}
	dbCategory.Delete(r.ctx, r.exec)
	return nil
}

func toCategory(d *dbModel.Category) (*model.Category, error) {
	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}
	category := &model.Category{
		Id: id,
		Name: d.Name,
		DisplayOrder: d.DisplayOrder.Int,
	}
	return category, nil
}

func toCategories(dbCategories []*dbModel.Category) ([]*model.Category, error) {
	var categories []*model.Category
	for _, v := range dbCategories {
		category, err := toCategory(v)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func toDbCategory(e *model.Category) (*dbModel.Category) {
	dbCategory := &dbModel.Category{
		ID: e.Id.String(),
		Name: e.Name,
		DisplayOrder: null.IntFrom(e.DisplayOrder),
	}
	return dbCategory
}