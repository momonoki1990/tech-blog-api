package model

import "testing"

func TestNewCategory(t *testing.T) {
	// Execute
	category, err := NewCategory("Name1", 1)
	if err != nil {
		panic(err)
	}

	// Check
	if category == nil {
		t.Errorf("category: Expected %v, but got %v", "not nil", category)
	}
	if category.Name != "Name1" {
		t.Errorf("category.Name: Expected %v, but got %v", "Name1", category.Name)
	}
	if category.DisplayOrder != 1 {
		t.Errorf("category.DisplayOrder: Expected %d, but got %d", 1, category.DisplayOrder)
	}
}