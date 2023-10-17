package database

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	"github.com/momonoki1990/tech-blog-api/domain/service"
	dbModel "github.com/momonoki1990/tech-blog-api/infra/database/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestCategoryFindOneById(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	dbCategory1 := &dbModel.Category{
		ID: "11111111-1111-1111-1111-111111111111",
		Name: "Category1",
		DisplayOrder: null.IntFrom(1),
	}
	err := dbCategory1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbCategory2 := &dbModel.Category{
		ID: "11111111-1111-1111-1111-111111111112",
		Name: "Category2",
		DisplayOrder: null.IntFrom(2),
	}
	err = dbCategory2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId2, err := uuid.Parse("11111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}

	// Execute
	r := NewCategoryRepository(ctx, tx)
	category1, err := r.FindOneById(categoryId1)
	if err != nil {
		panic(err)
	}
	category2, err := r.FindOneById(categoryId2)
	if err != nil {
		panic(err)
	}
	
	// Check
	if category1.Id != categoryId1 {
		t.Errorf("category1.Id: Expected %v, but got %v", categoryId1, category1.Id)
	}
	if category1.Name != "Category1" {
		t.Errorf("category1.Name: Expected %v, but got %v", "Category1", category1.Name)
	}
	if category1.DisplayOrder != 1 {
		t.Errorf("category1.DisplayOrder: Expected %v, but got %v", 1, category1.DisplayOrder)
	}

	if category2.Id != categoryId2 {
		t.Errorf("category2.Id: Expected %v, but got %v", categoryId2, category2.Id)
	}
	if category2.Name != "Category2" {
		t.Errorf("category2.Name: Expected %v, but got %v", "Category2", category2.Name)
	}
	if category2.DisplayOrder != 2 {
		t.Errorf("category2.DisplayOrder: Expected %v, but got %v", 2, category2.DisplayOrder)
	}
}

func TestCategoryFindOneByIdNotExisting(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}

	// Execute
	r := NewCategoryRepository(ctx, tx)
	category, err := r.FindOneById(categoryId1)
	
	// Check
	if err != nil {
		t.Errorf("err of r.FindOneById(categoryId): Expected %v, but got %v", nil, err)
	}
	if category != nil {
		t.Errorf("category: Expected %v, but got %v", nil, category)
	}
}

func TestCategoryFind(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	dbCategory1 := &dbModel.Category{
		ID: "11111111-1111-1111-1111-111111111111",
		Name: "Category1",
		DisplayOrder: null.IntFrom(1),
	}
	err := dbCategory1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbCategory2 := &dbModel.Category{
		ID: "11111111-1111-1111-1111-111111111112",
		Name: "Category2",
		DisplayOrder: null.IntFrom(2),
	}
	err = dbCategory2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId2, err := uuid.Parse("11111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}

	// Execute
	r := NewCategoryRepository(ctx, tx)
	categories, err := r.Find()
	if err != nil {
		panic(err)
	}
	
	// Check
	var category1 *model.Category
	var category2 *model.Category
	for _, v := range categories {
		if v.Id == categoryId1 {
			category1 = v
		}
		if v.Id == categoryId2 {
			category2 = v
		}
	}
	if category1.Id != categoryId1 {
		t.Errorf("category1.Id: Expected %v, but got %v", categoryId1, category1.Id)
	}
	if category1.Name != "Category1" {
		t.Errorf("category1.Name: Expected %v, but got %v", "Category1", category1.Name)
	}
	if category1.DisplayOrder != 1 {
		t.Errorf("category1.DisplayOrder: Expected %v, but got %v", 1, category1.DisplayOrder)
	}

	if category2.Id != categoryId2 {
		t.Errorf("category2.Id: Expected %v, but got %v", categoryId2, category2.Id)
	}
	if category2.Name != "Category2" {
		t.Errorf("category2.Name: Expected %v, but got %v", "Category2", category2.Name)
	}
	if category2.DisplayOrder != 2 {
		t.Errorf("category2.DisplayOrder: Expected %v, but got %v", 2, category2.DisplayOrder)
	}
}

func TestCategoryFindWithNoData(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	// NOTE: Can not delete only categories becase of foreign key constraint(ex. taggings references articles, articles references categories)
	_, err := dbModel.Taggings().DeleteAll(ctx, tx)
	if err != nil {
		panic(err)
	}
	_, err = dbModel.Articles().DeleteAll(ctx, tx)
	if err != nil {
		panic(err)
	}
	_, err = dbModel.Categories().DeleteAll(ctx, tx)
	if err != nil {
		panic(err)
	}

	// Execute
	r := NewCategoryRepository(ctx, tx)
	actuals, err := r.Find()
	if err != nil {
		panic(err)
	}

	// Check
	if actuals == nil {
		t.Errorf("actuals: Expected %s, but got %v", "not nil", actuals)
	}
	if len(actuals) != 0 {
		t.Errorf("len(actuals): Expected %d, but got %d", 0, len(actuals))
	}
}

func TestCategoryInsert(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	r := NewCategoryRepository(ctx, tx)
	creator := service.NewCategoryCreator(r)
	category, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}

	// Execute
	err = r.Insert(category)
	if err != nil {
		panic(err)
	}

	// Check
	dbCategory1, err := dbModel.FindCategory(ctx, tx, category.Id.String())
	if err != nil {
		panic(err)
	}
	if dbCategory1.Name != "Name1" {
		t.Errorf("dbCategory1.Name: Expected %s, but got %s", "Name1", dbCategory1.Name)
	}
	if dbCategory1.DisplayOrder.Int != 1 {
		t.Errorf("dbCategory1.DisplayOrder.Int: Expected %d, but got %d", 1, dbCategory1.DisplayOrder.Int)
	}
}

func TestCategoryInsertAlreadyExisting(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	r := NewCategoryRepository(ctx, tx)
	creator := service.NewCategoryCreator(r)
	category, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}
	err = r.Insert(category)
	if err != nil {
		panic(err)
	}

	// Execute
	err = r.Insert(category)

	// Check
	if err == nil {
		t.Errorf("err of r.Insert(category): Expected %v, but got %v", "not nil", err)
	}
}

func TestCategoryUpdate(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	r := NewCategoryRepository(ctx, tx)
	creator := service.NewCategoryCreator(r)
	category, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}
	err = r.Insert(category)
	if err != nil {
		panic(err)
	}

	// Execute
	category.Name = "Name1Changed"
	category.DisplayOrder = 11
	err = r.Update(category)
	if err != nil {
		panic(err)
	}

	// Check
	categoryCheck, err := r.FindOneById(category.Id)
	if err != nil {
		panic(err)
	}
	if categoryCheck.Id != category.Id {
		t.Errorf("categoryCheck.Id: Expected %v, but got %v", category.Id, categoryCheck.Id)
	}
	if categoryCheck.Name != "Name1Changed" {
		t.Errorf("categoryCheck.Name: Expected %v, but got %v", "Name1Changed", categoryCheck.Name)
	} 
	if categoryCheck.DisplayOrder != 11 {
		t.Errorf("categoryCheck.DisplayOrder: Expected %v, but got %v", 11, categoryCheck.DisplayOrder)
	}
}

func TestCategoryUpdateNotExisting(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	r := NewCategoryRepository(ctx, tx)
	creator := service.NewCategoryCreator(r)
	category, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}

	// Execute
	category.Name = "Name1Changed"
	category.DisplayOrder = 11
	err = r.Update(category)

	// Check
	if err == nil {
		t.Errorf("err of r.Update(category): Expected %v, but got %v", "not nil", err)
	}
	if err.Error() != "Category to update was not found" {
		t.Errorf("err of r.Update(article1): Expected %v, but got %v", "Category to update was not found", err)
	}
}

func TestCategoryDelete(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	r := NewCategoryRepository(ctx, tx)
	creator := service.NewCategoryCreator(r)
	category, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}
	err = r.Insert(category)
	if err != nil {
		panic(err)
	}

	// Execute
	err = r.Delete(category.Id)
	if err != nil {
		panic(err)
	}

	// Check
	categoryCheck, err := r.FindOneById(category.Id)
	if err != nil {
		panic(err)
	}
	if categoryCheck != nil {
		t.Errorf("categoryCheck: Expected %v, but got %v", nil, categoryCheck)
	}
}

func TestCategoryDeleteNotExisting(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	r := NewCategoryRepository(ctx, tx)
	creator := service.NewCategoryCreator(r)
	category, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}

	// Execute
	err = r.Delete(category.Id)

	// Check
	if err == nil {
		t.Errorf("err of r.Delete(category): Expected %v, but got %v", "not nil", err)
	}
	if err.Error() != "Category to delete was not found" {
		t.Errorf("err of r.Update(article1): Expected %v, but got %v", "Category to delete was not found", err)
	}
}