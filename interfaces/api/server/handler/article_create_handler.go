package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/momonoki1990/tech-blog-api/application/usecase"
)

type CreateArticleBody struct {
    Title string `json:"title"`
    Content string `json:"content"`
    CategoryId string `json:"categoryId"`
	TagNames []string `json:"tagNames"`
	ShouldPublish bool `json:"shouldPublish"`
}

type CreateArticleResponseBody struct {
    ArticleId string `json:"articleId"`
}

type ArticleCreateHandler interface {
    CreateArticle(c echo.Context) error
}

type articleCreateHandler struct {
    u usecase.ArticleUseCase
}

func NewArticleCreateHandler(u usecase.ArticleUseCase) ArticleCreateHandler {
    return &articleCreateHandler{u}
}

func (h *articleCreateHandler) CreateArticle(c echo.Context) error {
    body := new(CreateArticleBody)
    if err := c.Bind(body); err != nil {
        fmt.Print(err)
        return c.String(http.StatusBadRequest, "Bad request")
    }
	categoryId, err := uuid.Parse(body.CategoryId)
	if err != nil {
		fmt.Print(err)
        return c.String(http.StatusBadRequest, "Bad request")
	}
    articleId, err := h.u.RegisterArticle(body.Title, body.Content, categoryId, body.TagNames, body.ShouldPublish)
    if err != nil {
        return err
    }
    responseBody := &CreateArticleResponseBody{ArticleId: articleId}
    return c.JSON(http.StatusCreated, responseBody)
}