package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/momonoki1990/tech-blog-api/application/usecase"
	"github.com/momonoki1990/tech-blog-api/infra/database"
	"github.com/momonoki1990/tech-blog-api/interfaces/api/server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func connectToDb() (*sql.DB) {
    dataSource := os.ExpandEnv("${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}?parseTime=true")
    db, err := sql.Open("mysql", dataSource)
    if err!= nil {
        log.Fatal(err)
    }
    if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected!!")
    return db

}

func main() {
    db := connectToDb()
    ctx := context.Background()
    e := echo.New()
    stage := flag.String("stage", "prd", "Stage in which the application runs")
    flag.Parse()
    if *stage == "local" {
        e.Debug = true
    }
    e.Use(middleware.Logger())
    e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    cr := database.NewCategoryRepository(ctx, db)
    cu := usecase.NewCategoryUseCase(cr)
    e.GET("/categories", handler.NewCategoryListHandler(cu).CategoryList)
    e.POST("/category", handler.NewCategoryCreateHandler(cu).CreateCategory)
    e.PUT("/category/:id", handler.NewCategoryUpdateHandler(cu).UpdateCategory)
    e.DELETE("/category/:id", handler.NewCategoryDeleteHandler(cu).DeleteCategory)

    e.Logger.Fatal(e.Start(":1323"))
}