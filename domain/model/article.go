package model

import (
	"errors"
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
	tags := generateTags(tagNames)
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

func (a *Article) ChangeTags (tagNames []string) {
	tags := generateTags(tagNames)
	a.Tags = tags
}

func generateTags(tagNames []string) []Tag {
	var tags []Tag
	for i := 0; i < len(tagNames); i++ {
		tag := Tag{
			Name: tagNames[i],
		}
		tags = append(tags, tag)
	}
	return tags
}

func (a *Article) ChangeStatus (s Status) error {
	if s == Draft {
		a.Status = Draft
		a.PublishedAt = nil
	} else if s == Published {
		a.Status = Published
		now := time.Now()
		a.PublishedAt = &now
	} else {
		return errors.New("Invalid status")
	}
	return nil
}