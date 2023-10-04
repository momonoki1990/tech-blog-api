package model

import "time"

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
	Title string
	Content string
	Category Category
	Tags []Tag
	PublishedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Status Status
}

func NewArticle (title string, content string, category Category, tagNames []string, shouldPublish bool) (*Article, error) {
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

	a := &Article{
		Title: title,
		Content: content,
		Category: category,
		Tags: tags,
		PublishedAt: publishedAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: status,
	}
	return a, nil
}
