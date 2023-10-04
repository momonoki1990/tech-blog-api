package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type UpdateCategoryBody struct {
    Name string `json:"name"`
    DisplayOrder int `json:"displayOrder"`
}

type CategoryUpdateHandler interface {
    UpdateCategory(c echo.Context) error
}

type categoryUpdateHandler struct {
    u usecase.CategoryUseCase
}

func NewCategoryUpdateHandler(u usecase.CategoryUseCase) CategoryUpdateHandler {
    return &categoryUpdateHandler{u}
}

func (h *categoryUpdateHandler) UpdateCategory(c echo.Context) error {
    fmt.Println("👹categoryUpdateHandler.UpdateCategory called")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
    body := new(UpdateCategoryBody)
    if err := c.Bind(body); err != nil {
        fmt.Print(err)
        return c.String(http.StatusBadRequest, "Bad request")
    }
    if err := h.u.UpdateCategory(id, body.Name, body.DisplayOrder); err != nil {
        return err
    }
    return c.String(http.StatusOK, "Update category ok")
}