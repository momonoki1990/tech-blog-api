package model

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Category struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	DisplayOrder int `json:"displayOrder"`
}

func NewCategory(name string, displayOrder int) (*Category, error) {
	const (
		displayOrderMin = 1
		displayOrderMax = 999
	)

	if (displayOrder < displayOrderMin || displayOrder > displayOrderMax) {
		return nil, errors.New(fmt.Sprintf("displayOrder should be from %d to %d", displayOrderMin, displayOrderMax))
	}

	c := &Category{
		Id: uuid.New(),
		Name: name,
		DisplayOrder: displayOrder,
	}

	return c, nil
}

func (c *Category) Equals(compared *Category) bool {
	return c.Id == compared.Id
}