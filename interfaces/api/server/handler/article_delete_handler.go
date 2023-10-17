package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type ArticleDeleteHandler interface {
    DeleteArticle(c echo.Context) error
}

type articleDeleteHandler struct {
    u usecase.ArticleUseCase
}

func NewArticleDeleteHandler(u usecase.ArticleUseCase) ArticleDeleteHandler {
    return &articleDeleteHandler{u}
}

func (h *articleDeleteHandler) DeleteArticle(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
    if err := h.u.DeleteArticle(id); err != nil {
        return err
    }
    return c.String(http.StatusOK, "Delete article ok")
}