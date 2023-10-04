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
	db *sql.DB
}

func NewCategoryRepository(ctx context.Context,  db *sql.DB) repository.CategoryRepository {
    return &CategoryRepository{ctx, db}
}

func (r *CategoryRepository)FindOneByName(name string) (*model.Category, error) {
	dbCategories, err := dbModel.Categories(dbModel.CategoryWhere.Name.EQ(name)).One(r.ctx, r.db)
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
	dbCategory, err := dbModel.Categories(dbModel.CategoryWhere.ID.EQ(id.String())).One(r.ctx, r.db)
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
	dbCategories, err := dbModel.Categories().All(r.ctx, r.db)
	if err != nil {
		return nil, err
	}
	entities, err := toCategories(dbCategories)
	return entities, nil
}

func (r *CategoryRepository) Insert(c *model.Category) (error) {
	dbCategory := toDbCategory(c)
	err := dbCategory.Insert(r.ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) Update(c *model.Category) (error) {
	dbCategory, err := dbModel.FindCategory(r.ctx, r.db, c.Id.String())
	if err != nil {
		return err
	}
	if dbCategory == nil {
		return errors.New("変更対象のカテゴリが見つかりませんでした")
	}
	dbCategory.Name = c.Name
	dbCategory.DisplayOrder = null.IntFrom(c.DisplayOrder)
	dbCategory.Update(r.ctx, r.db, boil.Infer())
	return nil
}

func (r *CategoryRepository) Delete(id uuid.UUID) (error) {
	dbCategory, err := dbModel.FindCategory(r.ctx, r.db, id.String())
	if err != nil {
		return err
	}
	dbCategory.Delete(r.ctx, r.db)
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
	var entities []*model.Category
	for _, v := range dbCategories {
		category, err := toCategory(v)
		if err != nil {
			return nil, err
		}
		entities = append(entities, category)
	}
	return entities, nil
}

func toDbCategory(e *model.Category) (*dbModel.Category) {
	dbCategory := &dbModel.Category{
		ID: e.Id.String(),
		Name: e.Name,
		DisplayOrder: null.IntFrom(e.DisplayOrder),
	}
	return dbCategory
}