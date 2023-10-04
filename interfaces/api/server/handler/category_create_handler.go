package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type CreateCategoryBody struct {
    Name string `json:"name"`
    DisplayOrder int `json:"displayOrder"`
}

type CategoryCreateHandler interface {
    CreateCategory(c echo.Context) error
}

type categoryCreateHandler struct {
    u usecase.CategoryUseCase
}

func NewCategoryCreateHandler(u usecase.CategoryUseCase) CategoryCreateHandler {
    return &categoryCreateHandler{u}
}

func (h *categoryCreateHandler) CreateCategory(c echo.Context) error {
    body := new(CreateCategoryBody)
    if err := c.Bind(body); err != nil {
        fmt.Print(err)
        return c.String(http.StatusBadRequest, "Bad request")
    }
    if err := h.u.RegisterCategory(body.Name, body.DisplayOrder); err != nil {
        return err
    }
    return c.String(http.StatusOK, "Create category ok")
}