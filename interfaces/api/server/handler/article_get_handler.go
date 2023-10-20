package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type ArticleGetHandler interface {
    ArticleGet(c echo.Context) error
}

type articleGetHandler struct {
    u usecase.ArticleUseCase
}

func NewArticleGetHandler(u usecase.ArticleUseCase) ArticleGetHandler {
    return &articleGetHandler{u}
}

func (h *articleGetHandler) ArticleGet(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
    articles, err := h.u.GetArticle(id)
	if err != nil {
		return err
	}
    return c.JSON(http.StatusOK, articles)
}