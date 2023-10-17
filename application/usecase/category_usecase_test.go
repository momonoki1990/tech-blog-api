package usecase

import (
	"testing"

	"github.com/momonoki1990/tech-blog-api/domain/model"
	mock_service "github.com/momonoki1990/tech-blog-api/domain/service/mock"
	mock_repo "github.com/momonoki1990/tech-blog-api/infra/mock"
	"go.uber.org/mock/gomock"
)

func TestGetCategoryList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockCategoryRepository := mock_repo.NewMockCategoryRepository(mockCtrl)
	mockCategoryCreator := mock_service.NewMockCategoryCreator(mockCtrl)
	
	// Expected & Mock
	categories := []*model.Category{}
	mockCategoryRepository.EXPECT().Find().Return(categories, nil)
	
	// Execute
	u := NewCategoryUseCase(mockCategoryRepository, mockCategoryCreator)
	actual, err := u.GetCategoryList()
	if err != nil {
		panic(err)
	}
	
	// Check
	if len(actual) != len(categories) {
		t.Errorf("len(actual): Expected %d, but got %d", len(actual), len(categories))
	}
}

func TestRegisterCategory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockCategoryRepository := mock_repo.NewMockCategoryRepository(mockCtrl)
	mockCategoryCreator := mock_service.NewMockCategoryCreator(mockCtrl)
	category, err := model.NewCategory("Name1", 1)
	if err != nil {
		panic(err)
	}

	// Expected & Mock
	mockCategoryCreator.EXPECT().Create("Name1", 1).Return(category, nil)
	mockCategoryRepository.EXPECT().Insert(category).Return(nil)

	// Execute
	u := NewCategoryUseCase(mockCategoryRepository, mockCategoryCreator)
	err = u.RegisterCategory("Name1", 1)

	// Check
	if err != nil {
		t.Errorf("err of u.RegisterCategory('Name1', 1): Expected %v, but got %v", nil, err)
	}
}

func TestUpdateCategory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockCategoryRepository := mock_repo.NewMockCategoryRepository(mockCtrl)
	mockCategoryCreator := mock_service.NewMockCategoryCreator(mockCtrl)
	category, err := model.NewCategory("Name1", 1)
	if err != nil {
		panic(err)
	}
	categoryId := category.Id

	// Expected & Mock
	mockCategoryRepository.EXPECT().FindOneById(categoryId).Return(category, nil)
	mockCategoryRepository.EXPECT().Update(category).Return(nil)

	// Execute
	u := NewCategoryUseCase(mockCategoryRepository, mockCategoryCreator)
	err = u.UpdateCategory(categoryId, "Name1Changed", 101)

	// Check
	if err != nil {
		t.Errorf("err of u.UpdateCategory(categoryId, 'Name1Changed', 101): Expected %v, but got %v", nil, err)
	}
}

func TestUpdateCategoryNotFoundError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockCategoryRepository := mock_repo.NewMockCategoryRepository(mockCtrl)
	mockCategoryCreator := mock_service.NewMockCategoryCreator(mockCtrl)
	category, err := model.NewCategory("Name1", 1)
	if err != nil {
		panic(err)
	}
	categoryId := category.Id

	// Expected & Mock
	mockCategoryRepository.EXPECT().FindOneById(categoryId).Return(nil, nil)

	// Execute
	u := NewCategoryUseCase(mockCategoryRepository, mockCategoryCreator)
	err = u.UpdateCategory(categoryId, "Name1Changed", 101)

	// Check
	if err.Error() != "Category to update was not found" {
		t.Errorf("err.Error() of u.UpdateCategory(categoryId, 'Name1Changed', 101): Expected %s, but got %v", "Category to update was not found", err)
	}
}

func TestDeleteCategory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockCategoryRepository := mock_repo.NewMockCategoryRepository(mockCtrl)
	mockCategoryCreator := mock_service.NewMockCategoryCreator(mockCtrl)
	category, err := model.NewCategory("Name1", 1)
	if err != nil {
		panic(err)
	}
	categoryId := category.Id

	// Expected & Mock
	mockCategoryRepository.EXPECT().Delete(categoryId).Return(nil)

	// Execute
	u := NewCategoryUseCase(mockCategoryRepository, mockCategoryCreator)
	err = u.DeleteCategory(categoryId)

	// Check
	if err != nil {
		t.Errorf("err of u.DeleteCategory(categoryId): Expected %v, but got %v", nil, err)
	}
}