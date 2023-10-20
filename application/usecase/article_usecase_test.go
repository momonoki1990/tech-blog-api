package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	mock_repo "github.com/momonoki1990/tech-blog-api/infra/mock"
	"go.uber.org/mock/gomock"
)

func TestGetArticle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Prepare1
	mockArticleRepository := mock_repo.NewMockArticleRepository(mockCtrl)
	
	// Expected & Mock
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	article, err := model.NewArticle("Title1", "Content1", categoryId1, []string{"Tag1", "Tag2"}, false)
	if err != nil {
		panic(err)
	}
	mockArticleRepository.EXPECT().FindOneById(article.Id).Return(article, nil)
	
	// Execute
	u := NewArticleUseCase(mockArticleRepository)
	actual, err := u.GetArticle(article.Id)
	if err != nil {
		panic(err)
	}
	
	// Check
	if actual.Id != article.Id {
		t.Errorf("actual.Id: Expected %s, but got %s", article.Id, actual.Id)
	}
	if actual.Title != article.Title {
		t.Errorf("actual.Title: Expected %s, but got %s", article.Title, actual.Title)
	}
	if actual.Content != article.Content {
		t.Errorf("actual.Content: Expected %s, but got %s", article.Content, actual.Content)
	}
	if actual.CategoryId != article.CategoryId {
		t.Errorf("actual.CategoryId: Expected %s, but got %s", article.CategoryId, actual.CategoryId)
	}
	if actual.Tags[0].Name != article.Tags[0].Name {
		t.Errorf("actual.Tags[0].Name: Expected %s, but got %s", article.Tags[0].Name, actual.Tags[0].Name)
	}
	if actual.Tags[1].Name != article.Tags[1].Name {
		t.Errorf("actual.Tags[1].Name: Expected %s, but got %s", article.Tags[1].Name, actual.Tags[1].Name)
	}
	if actual.PublishedAt != article.PublishedAt {
		t.Errorf("actual.PublishedAt: Expected %v, but got %v", article.PublishedAt, actual.PublishedAt)
	}
	if actual.CreatedAt != article.CreatedAt {
		t.Errorf("actual.CreatedAt: Expected %v, but got %v", article.CreatedAt, actual.CreatedAt)
	}
	if actual.UpdatedAt != article.UpdatedAt {
		t.Errorf("actual.UpdatedAt: Expected %v, but got %v", article.UpdatedAt, actual.UpdatedAt)
	}
}

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