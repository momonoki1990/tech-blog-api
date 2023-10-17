package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewArticle(t *testing.T) {
	// Prepare data
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId2, err := uuid.Parse("11111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}
	
	// Execute
	article1, err := NewArticle("Title1", "Content1", categoryId1, []string{"Tag1", "Tag2"}, false)
	if err != nil {
		panic(err)
	}
	article2, err := NewArticle("Title2", "Content2", categoryId2, []string{}, true)
	if err != nil {
		panic(err)
	}

	// Check
	if article1 == nil {
		t.Errorf("article1: Expected %v, but got %v", "not nil", article1)
	}
	if article1.Title != "Title1" {
		t.Errorf("article1.Title: Expected %v, but got %v", "Title1", article1.Title)
	}
	if article1.Content != "Content1" {
		t.Errorf("article1.Content: Expected %v, but got %v", "Content1", article1.Content)
	}
	if article1.CategoryId != categoryId1 {
		t.Errorf("article1.CategoryId: Expected %v, but got %v", categoryId1, article1.CategoryId)
	}
	if len(article1.Tags) != 2 {
		t.Errorf("len(article1.Tags): Expected %d, but got %d", 2, len(article1.Tags))
	}
	if article1.Tags[0].Name != "Tag1" {
		t.Errorf("article1.Tags[0].Name: Expected %v, but got %v", "Tag1", article1.Tags[0].Name)
	}
	if article1.Tags[1].Name != "Tag2" {
		t.Errorf("article1.Tags[1].Name: Expected %v, but got %v", "Tag2", article1.Tags[1].Name)
	}
	if article1.PublishedAt != nil {
		t.Errorf("article1.PublishedAt: Expected %v, but got %v", nil, article1.PublishedAt)
	}
	if article1.Status != Draft {
		t.Errorf("article1.Status: Expected %s, but got %s", Draft, article1.Status)
	}
	if &article1.CreatedAt == nil {
		t.Errorf("&article1.CreatedAt: Expected %s, but got %v", "not nil", &article1.CreatedAt)
	}
	if &article1.UpdatedAt == nil {
		t.Errorf("&article1.UpdatedAt: Expected %s, but got %v", "not nil", &article1.UpdatedAt)
	}

	if article2 == nil {
		t.Errorf("article2: Expected %v, but got %v", "not nil", article2)
	}
	if article2.Title != "Title2" {
		t.Errorf("article2.Title: Expected %v, but got %v", "Title2", article2.Title)
	}
	if article2.Content != "Content2" {
		t.Errorf("article2.Content: Expected %v, but got %v", "Content2", article2.Content)
	}
	if article2.CategoryId != categoryId2 {
		t.Errorf("article2.CategoryId: Expected %v, but got %v", categoryId2, article2.CategoryId)
	}
	if len(article2.Tags) != 0 {
		t.Errorf("len(article2.Tags): Expected %d, but got %d", 0, len(article2.Tags))
	}
	if article2.PublishedAt == nil {
		t.Errorf("article2.PublishedAt: Expected %v, but got %v", "not nil", article2.PublishedAt)
	}
	if article2.Status != Published {
		t.Errorf("article2.Status: Expected %s, but got %s", Published, article2.Status)
	}
	if &article2.CreatedAt == nil {
		t.Errorf("&article2.CreatedAt: Expected %v, but got %v", "not nil", &article2.CreatedAt)
	}
	if &article2.UpdatedAt == nil {
		t.Errorf("&article2.UpdatedAt: Expected %v, but got %v", "not nil", &article2.UpdatedAt)
	}
}

func TestSetTags(t *testing.T) {
	// Prepare
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	article1, err := NewArticle("Title1", "Content1", categoryId1, []string{"Tag1", "Tag2"}, false)
	if err != nil {
		panic(err)
	}

	// Execute
	article1.SetTags([]string{"Tag3", "Tag4"})

	// Check
	if len(article1.Tags) != 2 {
		t.Errorf("len(article1.Tags): Expected %d, but got %d", 2, len(article1.Tags))
	}
	if article1.Tags[0].Name != "Tag3" {
		t.Errorf("article1.Tags[0].Name: Expected %v, but got %v", "Tag3", article1.Tags[0].Name)
	}
	if article1.Tags[1].Name != "Tag4" {
		t.Errorf("article1.Tags[1].Name: Expected %v, but got %v", "Tag4", article1.Tags[1].Name)
	}
}

func TestSetStatus(t *testing.T) {
	// Prepare
	categoryId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	article1, err := NewArticle("Title1", "Content1", categoryId1, []string{"Tag1", "Tag2"}, false)
	if err != nil {
		panic(err)
	}

	// Execute1
	article1.SetStatus(Published)

	// Check1
	if article1.Status != Published {
		t.Errorf("article1.Status: Expected %d, but got %d", Published, article1.Status)
	}
	if article1.PublishedAt == nil {
		t.Errorf("article1.PublishedAt: Expected %v, but got %v", "not nil", article1.PublishedAt)
	}
	firstPublishedAt := *article1.PublishedAt

	// Execute2
	article1.SetStatus(Draft)

	// Check2
	if article1.Status != Draft {
		t.Errorf("article1.Status: Expected %d, but got %d", Draft, article1.Status)
	}
	if article1.PublishedAt == nil {
		t.Errorf("article1.PublishedAt: Expected %s, but got %v", "not nil", article1.PublishedAt)
	}

	// Execute3
	article1.SetStatus(Published)

	// Check3
	if !article1.PublishedAt.Equal(firstPublishedAt) {
		t.Errorf("article1.PublishedAt: Expected %s, but got %v", firstPublishedAt, *article1.PublishedAt)
	}
}