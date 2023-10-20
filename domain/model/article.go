package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	Name string `json:"name"`
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
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CategoryId uuid.UUID `json:"categoryId"`
	Tags []Tag `json:"tags"`
	PublishedAt *time.Time `json:"publishedAt"`
	Status Status `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

func (a *Article) Equals(compared *Article) bool {
	return a.Id == compared.Id
}

func (a *Article) SetTags (tagNames []string) {
	tags := generateTags(tagNames)
	a.Tags = tags
}

func generateTags(tagNames []string) []Tag {
	var tags []Tag
	tagMap := make(map[string]bool)
	for _, v := range tagNames {
		if !tagMap[v] {
			tagMap[v] = true
			tag := Tag{
				Name: v,
			}
			tags = append(tags, tag)
		}
	}
	return tags
}

func (a *Article) SetStatus (s Status) error {
	switch s {
		case Draft:
			a.Status = Draft
		case Published:
			a.Status = Published
			if a.PublishedAt == nil {
				now := time.Now()
				a.PublishedAt = &now
			}
		default:
			return errors.New("Invalid status")
	}
	return nil
}