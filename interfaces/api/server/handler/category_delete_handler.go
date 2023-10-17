package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type CategoryDeleteHandler interface {
    DeleteCategory(c echo.Context) error
}

type categoryDeleteHandler struct {
    u usecase.CategoryUseCase
}

func NewCategoryDeleteHandler(u usecase.CategoryUseCase) CategoryDeleteHandler {
    return &categoryDeleteHandler{u}
}

func (h *categoryDeleteHandler) DeleteCategory(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
    if err := h.u.DeleteCategory(id); err != nil {
		fmt.Println("👹ここ")
        return err
    }
    return c.String(http.StatusOK, "Delete category ok")
}