package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type CategoryListHandler interface {
    CategoryList(c echo.Context) error
}

type categoryListHandler struct {
    u usecase.CategoryUseCase
}

func NewCategoryListHandler(u usecase.CategoryUseCase) CategoryListHandler {
    return &categoryListHandler{u}
}

func (h *categoryListHandler) CategoryList(c echo.Context) error {
    categories, err := h.u.GetCategoryList()
	if err != nil {
		return err
	}
    return c.JSON(http.StatusOK, categories)
}