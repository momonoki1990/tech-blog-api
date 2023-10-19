package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	mock_repo "github.com/momonoki1990/tech-blog-api/infra/mock"
	"go.uber.org/mock/gomock"
)

func TestGetArticleList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockArticleRepository := mock_repo.NewMockArticleRepository(mockCtrl)
	
	// Expected & Mock
	articles := []*model.Article{}
	mockArticleRepository.EXPECT().Find().Return(articles, nil)
	
	// Execute
	u := NewArticleUseCase(mockArticleRepository)
	actual, err := u.GetArticleList()
	if err != nil {
		panic(err)
	}
	
	// Check
	if len(actual) != len(articles) {
		t.Errorf("len(actual): Expected %d, but got %d", len(actual), len(articles))
	}
}

func TestRegisterArticle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockArticleRepository := mock_repo.NewMockArticleRepository(mockCtrl)
	categoryId, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}

	// Expected & Mock
	mockArticleRepository.EXPECT().Insert(gomock.Any()).Return(nil)

	// Execute
	u := NewArticleUseCase(mockArticleRepository)
	id, err := u.RegisterArticle("Title1", "Content1", categoryId, []string{"Tag1", "Tag2"}, false)

	// Check
	if err != nil {
		t.Errorf("err of u.RegisterArticle('Title1', 'Content1', categoryId, []string{'Tag1', 'Tag2'}, false): Expected %v, but got %v", nil, err)
	}
	if id == "" {
		t.Errorf("id of u.RegisterArticle('Title1', 'Content1', categoryId, []string{'Tag1', 'Tag2'}, false): Expected %s, but got %v", "not empty string", id)
	}
}

func TestUpdateArticle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockArticleRepository := mock_repo.NewMockArticleRepository(mockCtrl)
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId2, err := uuid.Parse("11111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}
	article, err := model.NewArticle("Title1", "Content1", categoryId1, []string{"Tag1", "Tag2"}, false)
	if err != nil {
		panic(err)
	}
	articleId := article.Id

	// Expected & Mock
	mockArticleRepository.EXPECT().FindOneById(articleId).Return(article, nil)
	mockArticleRepository.EXPECT().Update(article).Return(nil)

	// Execute
	u := NewArticleUseCase(mockArticleRepository)
	err = u.UpdateArticle(articleId, "Title1Changed", "Content1Changed", categoryId2, []string{"Tag3", "Tag4"}, true)

	// Check
	if err != nil {
		t.Errorf("err of u.UpdateArticle(articleId, 'Title1Changed', 'Content1Changed', categoryId2, []string{'Tag3', 'Tag4'}, true): Expected %v, but got %v", nil, err)
	}
}

func TestUpdateArticleNotFoundError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockArticleRepository := mock_repo.NewMockArticleRepository(mockCtrl)
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	articleId, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}

	// Expected & Mock
	mockArticleRepository.EXPECT().FindOneById(articleId).Return(nil, nil)

	// Execute
	u := NewArticleUseCase(mockArticleRepository)
	err = u.UpdateArticle(articleId, "Title1Changed", "Content1Changed", categoryId1, []string{"Tag3", "Tag4"}, true)

	// Check
	if err.Error() != "Article to update was not found" {
		t.Errorf("err.Error() of u.UpdateArticle(articleId, 'Name1Changed', 101): Expected %s, but got %v", "Article to update was not found", err)
	}
}

func TestDeleteArticle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockArticleRepository := mock_repo.NewMockArticleRepository(mockCtrl)
	articleId, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}

	// Expected & Mock
	mockArticleRepository.EXPECT().Delete(articleId).Return(nil)

	// Execute
	u := NewArticleUseCase(mockArticleRepository)
	err = u.DeleteArticle(articleId)

	// Check
	if err != nil {
		t.Errorf("err of u.DeleteArticle(articleId): Expected %v, but got %v", nil, err)
	}
}