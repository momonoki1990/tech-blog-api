package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/momonoki1990/tech-blog-api/domain/model"
	dbModel "github.com/momonoki1990/tech-blog-api/infra/database/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetTestConnection() (*sql.DB) {
    dataSource := os.ExpandEnv("${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}?parseTime=true")
    db, err := sql.Open("mysql", dataSource)
    if err!= nil {
        panic(err.Error())
    }
    if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("db connected!!")
    return db

}

func GetTestTransaction(db *sql.DB, ctx context.Context) *sql.Tx {
    tx, _ := db.BeginTx(ctx, nil)
    return tx
}

func TestFindOneById(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()


	// Prepare data
	n := time.Now()
	now := n.Truncate(time.Second)

	dbArticle1 := &dbModel.Article{
		ID: "11111111-1111-1111-1111-111111111111",
		Title: "Title1",
		Content: "Content1",
		CategoryID: "21111111-1111-1111-1111-111111111111",
		PublishedAt: null.TimeFrom(now),
		Status: "Published",
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := dbArticle1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTag1 := &dbModel.Tag{
		ID: "31111111-1111-1111-1111-111111111111",
		Name: "Tag1",
	}
	err = dbTag1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTagging1 := &dbModel.Tagging{
		ArticleID: "11111111-1111-1111-1111-111111111111",
		TagID: "31111111-1111-1111-1111-111111111111",
	}
	err = dbTagging1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}


	// Expected
	articleId, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId, err := uuid.Parse("21111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	tagId, err := uuid.Parse("31111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	tags := []model.Tag{
		{Id: tagId, Name: "Tag1"},
	}
	
	expected := &model.Article{
		Id: articleId,
		Title: "Title1",
		Content: "Content1",
		CategoryId: categoryId,
		Tags: tags,
		PublishedAt: &now,
		Status: model.Published,
		CreatedAt: now,
		UpdatedAt: now,
	}
	

	// Execute
	r := NewArticleRepository(ctx, tx)
	actual, err := r.FindOneById(articleId)
	if err != nil {
		panic(err)
	}

	// Check
	if actual == nil {
		t.Errorf("article1: Expected not nil, but got nil")
	}
	if actual.Id != expected.Id {
		t.Errorf("Id: Expected %s, but got %s", expected.Id, actual.Id)
	}
	if actual.Title != expected.Title {
		t.Errorf("Title: Expected %s, but got %s", expected.Title, actual.Title)
	}
	if actual.Content != expected.Content {
		t.Errorf("Content: Expected %s, but got %s", expected.Content, actual.Content)
	}
	if actual.Tags[0].Id != expected.Tags[0].Id {
		t.Errorf("Tags[0].Id: Expected %s, but got %s", expected.Tags[0].Id, actual.Tags[0].Id)
	}
	if !actual.PublishedAt.Equal(*expected.PublishedAt) {
		t.Errorf("PublishedAt: Expected %s, but got %s", expected.PublishedAt, actual.PublishedAt)
	}
	if actual.Status != expected.Status {
		t.Errorf("Id: Expected %s, but got %s", expected.Status, actual.Status)
	}
	if !actual.CreatedAt.Equal(expected.CreatedAt) {
		t.Errorf("CreatedAt: Expected %s, but got %s", actual.CreatedAt, expected.CreatedAt)
	}
	if !actual.UpdatedAt.Equal(expected.UpdatedAt) {
		t.Errorf("UpdatedAt: Expected %s, but got %s", actual.UpdatedAt, expected.UpdatedAt)
	}
	tx.Rollback()
} 

func TestFind(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	n := time.Now()
	now := n.Truncate(time.Second)

	dbArticle1 := &dbModel.Article{
		ID: "11111111-1111-1111-1111-111111111111",
		Title: "Title1",
		Content: "Content1",
		CategoryID: "21111111-1111-1111-1111-111111111111",
		PublishedAt: null.TimeFrom(now),
		Status: "Published",
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := dbArticle1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTag1 := &dbModel.Tag{
		ID: "31111111-1111-1111-1111-111111111111",
		Name: "Tag1",
	}
	err = dbTag1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTagging1 := &dbModel.Tagging{
		ArticleID: "11111111-1111-1111-1111-111111111111",
		TagID: "31111111-1111-1111-1111-111111111111",
	}
	err = dbTagging1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbArticle2 := &dbModel.Article{
		ID: "11111111-1111-1111-1111-111111111112",
		Title: "Title2",
		Content: "Content2",
		CategoryID: "21111111-1111-1111-1111-111111111112",
		PublishedAt: null.TimeFrom(now),
		Status: "Draft",
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = dbArticle2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTag2 := &dbModel.Tag{
		ID: "31111111-1111-1111-1111-111111111112",
		Name: "Tag2",
	}
	err = dbTag2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTagging2 := &dbModel.Tagging{
		ArticleID: "11111111-1111-1111-1111-111111111112",
		TagID: "31111111-1111-1111-1111-111111111112",
	}
	err = dbTagging2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}


	// Expected
	articleId1, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId1, err := uuid.Parse("21111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	tagId1, err := uuid.Parse("31111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	tags1 := []model.Tag{
		{Id: tagId1, Name: "Tag1"},
	}
	
	expected1 := &model.Article{
		Id: articleId1,
		Title: "Title1",
		Content: "Content1",
		CategoryId: categoryId1,
		Tags: tags1,
		PublishedAt: &now,
		Status: model.Published,
		CreatedAt: now,
		UpdatedAt: now,
	}

	articleId2, err := uuid.Parse("11111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}
	categoryId2, err := uuid.Parse("21111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}
	tagId2, err := uuid.Parse("31111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}
	tags2 := []model.Tag{
		{Id: tagId2, Name: "Tag2"},
	}
	
	expected2 := &model.Article{
		Id: articleId2,
		Title: "Title2",
		Content: "Content2",
		CategoryId: categoryId2,
		Tags: tags2,
		PublishedAt: &now,
		Status: model.Draft,
		CreatedAt: now,
		UpdatedAt: now,
	}
	

	// Execute
	r := NewArticleRepository(ctx, tx)
	actuals, err := r.Find()
	if err != nil {
		panic(err)
	}

	actual1 := actuals[0]
	actual2 := actuals[1]

	// Check
	if actual1 == nil {
		t.Errorf("article1: Expected not nil, but got nil")
	}
	if actual1.Id != expected1.Id {
		t.Errorf("article1.Id: Expected %s, but got %s", expected1.Id, actual1.Id)
	}
	if actual1.Title != expected1.Title {
		t.Errorf("article1.Title: Expected %s, but got %s", expected1.Title, actual1.Title)
	}
	if actual1.Content != expected1.Content {
		t.Errorf("article1.Content: Expected %s, but got %s", expected1.Content, actual1.Content)
	}
	if actual1.Tags[0].Id != expected1.Tags[0].Id {
		t.Errorf("article1.Tags[0].Id: Expected %s, but got %s", expected1.Tags[0].Id, actual1.Tags[0].Id)
	}
	if !actual1.PublishedAt.Equal(*expected1.PublishedAt) {
		t.Errorf("article1.PublishedAt: Expected %s, but got %s", expected1.PublishedAt, actual1.PublishedAt)
	}
	if actual1.Status != expected1.Status {
		t.Errorf("article1.Id: Expected %s, but got %s", expected1.Status, actual1.Status)
	}
	if !actual1.CreatedAt.Equal(expected1.CreatedAt) {
		t.Errorf("article1.CreatedAt: Expected %s, but got %s", actual1.CreatedAt, expected1.CreatedAt)
	}
	if !actual1.UpdatedAt.Equal(expected1.UpdatedAt) {
		t.Errorf("article1.UpdatedAt: Expected %s, but got %s", actual1.UpdatedAt, expected1.UpdatedAt)
	}

	if actual2 == nil {
		t.Errorf("article2: Expected not nil, but got nil")
	}
	if actual2.Id != expected2.Id {
		t.Errorf("article2.Id: Expected %s, but got %s", expected2.Id, actual2.Id)
	}
	if actual2.Title != expected2.Title {
		t.Errorf("article1.Title: Expected %s, but got %s", expected2.Title, actual1.Title)
	}
	if actual2.Content != expected2.Content {
		t.Errorf("article1.Content: Expected %s, but got %s", expected2.Content, actual2.Content)
	}
	if actual2.Tags[0].Id != expected2.Tags[0].Id {
		t.Errorf("article1.Tags[0].Id: Expected %s, but got %s", expected2.Tags[0].Id, actual2.Tags[0].Id)
	}
	if !actual2.PublishedAt.Equal(*expected2.PublishedAt) {
		t.Errorf("article1.PublishedAt: Expected %s, but got %s", expected2.PublishedAt, actual2.PublishedAt)
	}
	if actual2.Status != expected2.Status {
		t.Errorf("article1.Id: Expected %s, but got %s", expected2.Status, actual2.Status)
	}
	if !actual2.CreatedAt.Equal(expected2.CreatedAt) {
		t.Errorf("article1.CreatedAt: Expected %s, but got %s", actual2.CreatedAt, expected2.CreatedAt)
	}
	if !actual2.UpdatedAt.Equal(expected2.UpdatedAt) {
		t.Errorf("article1.UpdatedAt: Expected %s, but got %s", actual2.UpdatedAt, expected2.UpdatedAt)
	}
	tx.Rollback()
}

func TestInsert(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	categoryId, err := uuid.Parse("11111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	// TODO: publishedがfalseの場合も用意する
	article1, err := model.NewArticle("Title1", "Content1", categoryId, []string{"Tag1", "Tag2"}, true)
	if err != nil {
		panic(err)
	}
	article2, err := model.NewArticle("Title2", "Content2", categoryId, []string{"Tag1", "Tag3"}, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("👹article2", article2)
	
	// Execute
	r := NewArticleRepository(ctx, tx)
	err = r.Insert(article1)
	if err != nil {
		panic(err)
	}
	err = r.Insert(article2)
	if err != nil {
		panic(err)
	}

	// Check
	dbArticle, err := dbModel.FindArticle(ctx, tx, article1.Id.String())
	if err != nil {
		panic(err)
	}
	if dbArticle.Title != "Title1" {
		t.Errorf("Title: Expected %s, but got %s", "Title1", dbArticle.Title)
	}
	if dbArticle.Content != "Content1" {
		t.Errorf("Content: Expected %s, but got %s", "Content1", dbArticle.Content)
	}
	if dbArticle.CategoryID != "11111111-1111-1111-1111-111111111111" {
		t.Errorf("CategoryId: Expected %s, but got %s", "11111111-1111-1111-1111-111111111111", dbArticle.CategoryID)
	}
	if dbArticle.Status != "Published" {
		t.Errorf("Status: Expected %s, but got %s", "Published", dbArticle.Status)
	}
	if &dbArticle.PublishedAt == nil {
		t.Errorf("PublishedAt: Expected not nil, but got %v", &dbArticle.PublishedAt)
	}

	dbTag1, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag1")).One(ctx, tx)
	if err != nil {
		panic(err)
	}
	dbTag2, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag2")).One(ctx, tx)
	if err != nil {
		panic(err)
	}
	if &dbTag1 == nil {
		t.Errorf("Tag1: Expected not nil, but got %v", &dbTag1)
	}
	if &dbTag2 == nil {
		t.Errorf("Tag2: Expected not nil, but got %v", &dbTag2)
	}
	
	dbTaggings, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(article1.Id.String()), dbModel.TaggingWhere.TagID.IN([]string{dbTag1.ID, dbTag2.ID})).All(ctx, tx)
	if err != nil {
		panic(err)
	}
	if len(dbTaggings) != 2 {
		t.Errorf("Tagging: Expected count %d, but got %d", 2, len(dbTaggings))
	}

}