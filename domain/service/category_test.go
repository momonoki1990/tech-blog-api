package service

import (
	"testing"

	mock_repo "github.com/momonoki1990/tech-blog-api/infra/mock"
	"go.uber.org/mock/gomock"
)

func TestCategoryCreatorCreate (t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockCategoryRepository := mock_repo.NewMockCategoryRepository(mockCtrl)
	creator := NewCategoryCreator(mockCategoryRepository)

	// Expected
	mockCategoryRepository.EXPECT().FindOneByName("Name1").Return(nil, nil)

	// Execute1
	category1, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}
	
	// Check1
	if category1 == nil {
		t.Errorf("category1: Expected %s, but got %v", "not nil", category1)
	}
}

func TestCategoryCreatorCreateDuplicationError (t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockCategoryRepository := mock_repo.NewMockCategoryRepository(mockCtrl)
	creator := NewCategoryCreator(mockCategoryRepository)
	mockCategoryRepository.EXPECT().FindOneByName("Name1").Return(nil, nil)
	category1, err := creator.Create("Name1", 1)
	if err != nil {
		panic(err)
	}

	// Mock
	mockCategoryRepository.EXPECT().FindOneByName("Name1").Return(category1, nil)

	// Execute1
	category2, err := creator.Create("Name1", 2)
	
	// Check1
	if err.Error() != "Category name is already registered" {
		t.Errorf("err of creator.Create('Name1', 2): Expected %s, but got %s", "Category name is already registered", err.Error())
	}
	if category2 != nil {
		t.Errorf("category2: Expected %v, but got %v", nil, category2)
	}
}