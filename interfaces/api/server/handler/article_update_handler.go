package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type UpdateArticleBody struct {
    Title string `json:"title"`
    Content string `json:"content"`
    CategoryId string `json:"categoryId"`
	TagNames []string `json:"tagNames"`
	ShouldPublish bool `json:"shouldPublish"`
}

type ArticleUpdateHandler interface {
    UpdateArticle(c echo.Context) error
}

type articleUpdateHandler struct {
    u usecase.ArticleUseCase
}

func NewArticleUpdateHandler(u usecase.ArticleUseCase) ArticleUpdateHandler {
    return &articleUpdateHandler{u}
}

func (h *articleUpdateHandler) UpdateArticle(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
    body := new(UpdateArticleBody)
    if err := c.Bind(body); err != nil {
        fmt.Print(err)
        return c.String(http.StatusBadRequest, "Bad request")
    }
	categoryId, err := uuid.Parse(body.CategoryId)
	if err != nil {
		fmt.Print(err)
        return c.String(http.StatusBadRequest, "Bad request")
	}
    if err := h.u.UpdateArticle(id, body.Title, body.Content, categoryId, body.TagNames, body.ShouldPublish); err != nil {
        return err
    }
    return c.String(http.StatusOK, "Update article ok")
}