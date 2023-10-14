package model

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	Name string
}

type Status int

const (
	Draft Status = iota
    Published 
)

func (s Status) String() string {
    switch s {
    case Draft:
        return "Draft"
    case Published :
        return "Published"
    default:
        return "Unknown"
    }
}

type Article struct {
	Id uuid.UUID
	Title string
	Content string
	CategoryId uuid.UUID
	Tags []Tag
	PublishedAt *time.Time
	Status Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle (title string, content string, categoryId uuid.UUID, tagNames []string, shouldPublish bool) (*Article, error) {
	var tags []Tag
	for i := 0; i < len(tagNames); i++ {
		tag := Tag{
			Name: tagNames[i],
		}
		tags = append(tags, tag)
	}
	var publishedAt *time.Time
	status := Draft
	if (shouldPublish == true) {
		now := time.Now()
		publishedAt = &now
		status = Published
	}

	article := &Article{
		Id: uuid.New(),
		Title: title,
		Content: content,
		CategoryId: categoryId,
		Tags: tags,
		Status: status,
		PublishedAt: publishedAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return article, nil
}
