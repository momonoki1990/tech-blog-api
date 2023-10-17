package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type ArticleListHandler interface {
    ArticleList(c echo.Context) error
}

type articleListHandler struct {
    u usecase.ArticleUseCase
}

func NewArticleListHandler(u usecase.ArticleUseCase) ArticleListHandler {
    return &articleListHandler{u}
}

func (h *articleListHandler) ArticleList(c echo.Context) error {
    articles, err := h.u.GetArticleList()
	if err != nil {
		return err
	}
    return c.JSON(http.StatusOK, articles)
}