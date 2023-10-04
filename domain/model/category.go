package model

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Category struct {
	Id uuid.UUID
	Name string
	DisplayOrder int
}

func NewCategory(name string, displayOrder int) (*Category, error) {
	const (
		displayOrderMin = 0
		displayOrderMax = 999
	)

	if (displayOrder < displayOrderMin || displayOrder > displayOrderMax) {
		return nil, errors.New(fmt.Sprintf("表示順序は%dから%dの間で設定してください", displayOrderMin, displayOrderMax))
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