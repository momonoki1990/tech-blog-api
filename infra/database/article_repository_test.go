package database

import (
	"context"
	"database/sql"
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

	dbCategory1 := &dbModel.Category{
		ID: "21111111-1111-1111-1111-111111111111",
		Name: "Category1",
		DisplayOrder: null.IntFrom(99),
	}
	err := dbCategory1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	
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
	err = dbArticle1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbTag1 := &dbModel.Tag{
		Name: "Tag1",
	}
	err = dbTag1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbTagging1 := &dbModel.Tagging{
		ArticleID: "11111111-1111-1111-1111-111111111111",
		TagName: "Tag1",
	}
	err = dbTagging1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbArticle2 := &dbModel.Article{
		ID: "11111111-1111-1111-1111-111111111112",
		Title: "Title2",
		Content: "Content2",
		CategoryID: "21111111-1111-1111-1111-111111111111",
		PublishedAt: null.TimeFromPtr(nil),
		Status: "Draft",
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = dbArticle2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbTag2 := &dbModel.Tag{
		Name: "Tag2",
	}
	err = dbTag2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbTagging2 := &dbModel.Tagging{
		ArticleID: "11111111-1111-1111-1111-111111111112",
		TagName: "Tag2",
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
	categoryId, err := uuid.Parse("21111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	tags1 := []model.Tag{
		{Name: "Tag1"},
	}
	
	expected1 := &model.Article{
		Id: articleId1,
		Title: "Title1",
		Content: "Content1",
		CategoryId: categoryId,
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
	tags2 := []model.Tag{
		{Name: "Tag2"},
	}
	
	expected2 := &model.Article{
		Id: articleId2,
		Title: "Title2",
		Content: "Content2",
		CategoryId: categoryId,
		Tags: tags2,
		PublishedAt: nil,
		Status: model.Draft,
		CreatedAt: now,
		UpdatedAt: now,
	}
	

	// Execute
	r := NewArticleRepository(ctx, tx)
	actual1, err := r.FindOneById(articleId1)
	if err != nil {
		panic(err)
	}
	actual2, err := r.FindOneById(articleId2)
	if err != nil {
		panic(err)
	}

	// Check
	if actual1 == nil {
		t.Errorf("actual1: Expected not nil, but got nil")
	}
	if actual1.Id != expected1.Id {
		t.Errorf("actual1.Id: Expected %s, but got %s", expected1.Id, actual1.Id)
	}
	if actual1.Title != expected1.Title {
		t.Errorf("actual1.Title: Expected %s, but got %s", expected1.Title, actual1.Title)
	}
	if actual1.Content != expected1.Content {
		t.Errorf("actual1.Content: Expected %s, but got %s", expected1.Content, actual1.Content)
	}
	if actual1.Tags[0].Name != expected1.Tags[0].Name {
		t.Errorf("actual1.Tags[0].Name: Expected %s, but got %s", expected1.Tags[0].Name, actual1.Tags[0].Name)
	}
	if !actual1.PublishedAt.Equal(*expected1.PublishedAt) {
		t.Errorf("actual1.PublishedAt: Expected %s, but got %s", expected1.PublishedAt, actual1.PublishedAt)
	}
	if actual1.Status != expected1.Status {
		t.Errorf("actual1.Status: Expected %s, but got %s", expected1.Status, actual1.Status)
	}
	if !actual1.CreatedAt.Equal(expected1.CreatedAt) {
		t.Errorf("actual1.CreatedAt: Expected %s, but got %s", actual1.CreatedAt, expected1.CreatedAt)
	}
	if !actual1.UpdatedAt.Equal(expected1.UpdatedAt) {
		t.Errorf("actual1.UpdatedAt: Expected %s, but got %s", actual1.UpdatedAt, expected1.UpdatedAt)
	}

	if actual2 == nil {
		t.Errorf("actual2: Expected not nil, but got nil")
	}
	if actual2.Id != expected2.Id {
		t.Errorf("actual2.Id: Expected %s, but got %s", expected2.Id, actual2.Id)
	}
	if actual2.Title != expected2.Title {
		t.Errorf("actual2.Title: Expected %s, but got %s", expected2.Title, actual2.Title)
	}
	if actual2.Content != expected2.Content {
		t.Errorf("actual2.Content: Expected %s, but got %s", expected2.Content, actual2.Content)
	}
	if actual2.Tags[0].Name != expected2.Tags[0].Name {
		t.Errorf("actual2.Tags[0].Name: Expected %s, but got %s", expected2.Tags[0].Name, actual2.Tags[0].Name)
	}
	if actual2.PublishedAt != nil {
		t.Errorf("actual2.PublishedAt: Expected %v, but got %s", nil, actual2.PublishedAt)
	}
	if actual2.Status != expected2.Status {
		t.Errorf("actual2.Status: Expected %s, but got %s", expected2.Status, actual2.Status)
	}
	if !actual2.CreatedAt.Equal(expected2.CreatedAt) {
		t.Errorf("actual2.CreatedAt: Expected %s, but got %s", actual2.CreatedAt, expected2.CreatedAt)
	}
	if !actual2.UpdatedAt.Equal(expected2.UpdatedAt) {
		t.Errorf("actual2.UpdatedAt: Expected %s, but got %s", actual2.UpdatedAt, expected2.UpdatedAt)
	}
} 

func TestFind(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	n := time.Now()
	now := n.Truncate(time.Second)

	dbCategory1 := &dbModel.Category{
		ID: "21111111-1111-1111-1111-111111111111",
		Name: "Category1",
		DisplayOrder: null.IntFrom(99),
	}
	err := dbCategory1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

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
	err = dbArticle1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTag1 := &dbModel.Tag{
		Name: "Tag1",
	}
	err = dbTag1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbTagging1 := &dbModel.Tagging{
		ArticleID: "11111111-1111-1111-1111-111111111111",
		TagName: "Tag1",
	}
	err = dbTagging1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbArticle2 := &dbModel.Article{
		ID: "11111111-1111-1111-1111-111111111112",
		Title: "Title2",
		Content: "Content2",
		CategoryID: "21111111-1111-1111-1111-111111111111",
		PublishedAt: null.TimeFromPtr(nil),
		Status: "Draft",
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = dbArticle2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbTag2 := &dbModel.Tag{
		Name: "Tag2",
	}
	err = dbTag2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}
	dbTagging2 := &dbModel.Tagging{
		ArticleID: "11111111-1111-1111-1111-111111111112",
		TagName: "Tag2",
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
	tags1 := []model.Tag{
		{Name: "Tag1"},
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
	tags2 := []model.Tag{
		{Name: "Tag2"},
	}
	expected2 := &model.Article{
		Id: articleId2,
		Title: "Title2",
		Content: "Content2",
		CategoryId: categoryId2,
		Tags: tags2,
		PublishedAt: nil,
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
		t.Errorf("actual1: Expected not nil, but got nil")
	}
	if actual1.Id != expected1.Id {
		t.Errorf("actual1.Id: Expected %s, but got %s", expected1.Id, actual1.Id)
	}
	if actual1.Title != expected1.Title {
		t.Errorf("actual1.Title: Expected %s, but got %s", expected1.Title, actual1.Title)
	}
	if actual1.Content != expected1.Content {
		t.Errorf("actual1.Content: Expected %s, but got %s", expected1.Content, actual1.Content)
	}
	if actual1.Tags[0].Name != expected1.Tags[0].Name {
		t.Errorf("actual1.Tags[0].Name: Expected %s, but got %s", expected1.Tags[0].Name, actual1.Tags[0].Name)
	}
	if !actual1.PublishedAt.Equal(*expected1.PublishedAt) {
		t.Errorf("actual1.PublishedAt: Expected %s, but got %s", expected1.PublishedAt, actual1.PublishedAt)
	}
	if actual1.Status != expected1.Status {
		t.Errorf("actual1.Status: Expected %s, but got %s", expected1.Status, actual1.Status)
	}
	if !actual1.CreatedAt.Equal(expected1.CreatedAt) {
		t.Errorf("actual1.CreatedAt: Expected %s, but got %s", actual1.CreatedAt, expected1.CreatedAt)
	}
	if !actual1.UpdatedAt.Equal(expected1.UpdatedAt) {
		t.Errorf("actual1.UpdatedAt: Expected %s, but got %s", actual1.UpdatedAt, expected1.UpdatedAt)
	}

	if actual2 == nil {
		t.Errorf("actual2: Expected not nil, but got nil")
	}
	if actual2.Id != expected2.Id {
		t.Errorf("actual2.Id: Expected %s, but got %s", expected2.Id, actual2.Id)
	}
	if actual2.Title != expected2.Title {
		t.Errorf("actual1.Title: Expected %s, but got %s", expected2.Title, actual1.Title)
	}
	if actual2.Content != expected2.Content {
		t.Errorf("actual1.Content: Expected %s, but got %s", expected2.Content, actual2.Content)
	}
	if actual2.Tags[0].Name != expected2.Tags[0].Name {
		t.Errorf("actual1.Tags[0].Name: Expected %s, but got %s", expected2.Tags[0].Name, actual2.Tags[0].Name)
	}
	if actual2.PublishedAt != expected2.PublishedAt {
		t.Errorf("actual1.PublishedAt: Expected %s, but got %s", expected2.PublishedAt, actual2.PublishedAt)
	}
	if actual2.Status != expected2.Status {
		t.Errorf("actual1.Id: Expected %s, but got %s", expected2.Status, actual2.Status)
	}
	if !actual2.CreatedAt.Equal(expected2.CreatedAt) {
		t.Errorf("actual1.CreatedAt: Expected %s, but got %s", actual2.CreatedAt, expected2.CreatedAt)
	}
	if !actual2.UpdatedAt.Equal(expected2.UpdatedAt) {
		t.Errorf("actual1.UpdatedAt: Expected %s, but got %s", actual2.UpdatedAt, expected2.UpdatedAt)
	}
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
	dbCategory1 := &dbModel.Category{
		ID: "11111111-1111-1111-1111-111111111111",
		Name: "Category1",
		DisplayOrder: null.IntFrom(99),
	}
	err = dbCategory1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	article1, err := model.NewArticle("Title1", "Content1", categoryId, []string{"Tag1", "Tag2"}, true)
	if err != nil {
		panic(err)
	}
	article2, err := model.NewArticle("Title2", "Content2", categoryId, []string{"Tag1", "Tag3"}, false)
	if err != nil {
		panic(err)
	}
	
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
	dbArticle1, err := dbModel.FindArticle(ctx, tx, article1.Id.String())
	if err != nil {
		panic(err)
	}
	if dbArticle1.Title != "Title1" {
		t.Errorf("dbArticle1.Title: Expected %s, but got %s", "Title1", dbArticle1.Title)
	}
	if dbArticle1.Content != "Content1" {
		t.Errorf("dbArticle1.Content: Expected %s, but got %s", "Content1", dbArticle1.Content)
	}
	if dbArticle1.CategoryID != "11111111-1111-1111-1111-111111111111" {
		t.Errorf("dbArticle1.CategoryId: Expected %s, but got %s", "11111111-1111-1111-1111-111111111111", dbArticle1.CategoryID)
	}
	if dbArticle1.Status != "Published" {
		t.Errorf("dbArticle1.Status: Expected %s, but got %s", "Published", dbArticle1.Status)
	}
	if !dbArticle1.PublishedAt.Time.Equal(article1.PublishedAt.Round(time.Second)) {
		t.Errorf("dbArticle1.PublishedAt: Expected %v, but got %v", *article1.PublishedAt, dbArticle1.PublishedAt)
	}

	dbArticle2, err := dbModel.FindArticle(ctx, tx, article2.Id.String())
	if err != nil {
		panic(err)
	}
	if dbArticle2.Title != "Title2" {
		t.Errorf("dbArticle2.Title: Expected %s, but got %s", "Title2", dbArticle2.Title)
	}
	if dbArticle2.Content != "Content2" {
		t.Errorf("dbArticle2.Content: Expected %s, but got %s", "Content2", dbArticle2.Content)
	}
	if dbArticle2.CategoryID != "11111111-1111-1111-1111-111111111111" {
		t.Errorf("dbArticle2.CategoryId: Expected %s, but got %s", "11111111-1111-1111-1111-111111111111", dbArticle2.CategoryID)
	}
	if dbArticle2.Status != "Draft" {
		t.Errorf("dbArticle2.Status: Expected %s, but got %s", "Draft", dbArticle2.Status)
	}
	if dbArticle2.PublishedAt.Valid {
		t.Errorf("dbArticle2.PublishedAt: Expected nil, but got %v", dbArticle2.PublishedAt)
	}

	dbTag1, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag1")).One(ctx, tx)
	if err != nil {
		panic(err)
	}
	dbTag2, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag2")).One(ctx, tx)
	if err != nil {
		panic(err)
	}
	dbTag3, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag3")).One(ctx, tx)
	if err != nil {
		panic(err)
	}
	if &dbTag1 == nil {
		t.Errorf("Tag1: Expected not nil, but got %v", &dbTag1)
	}
	if &dbTag2 == nil {
		t.Errorf("Tag2: Expected not nil, but got %v", &dbTag2)
	}
	if &dbTag3 == nil {
		t.Errorf("Tag3: Expected not nil, but got %v", &dbTag2)
	}
	
	dbTaggings1, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(article1.Id.String())).All(ctx, tx)
	if err != nil {
		panic(err)
	}
	if len(dbTaggings1) != 2 {
		t.Errorf("len(dbTaggings2): Expected count %d, but got %d", 2, len(dbTaggings1))
	}
	if dbTaggings1[0].TagName != "Tag1" {
		t.Errorf("dbTaggings1[0].TagName: Expected %s, but got %s", "Tag1", dbTaggings1[0].TagName)
	}
	if dbTaggings1[1].TagName != "Tag2" {
		t.Errorf("dbTaggings1[1].TagName: Expected %s, but got %s", "Tag2", dbTaggings1[1].TagName)
	}
	dbTaggings2, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(article2.Id.String())).All(ctx, tx)
	if err != nil {
		panic(err)
	}
	if len(dbTaggings2) != 2 {
		t.Errorf("len(dbTaggings2): Expected count %d, but got %d", 2, len(dbTaggings2))
	}
	if dbTaggings2[0].TagName != "Tag1" {
		t.Errorf("dbTaggings2[0].TagName: Expected %s, but got %s", "Tag1", dbTaggings2[0].TagName)
	}
	if dbTaggings2[1].TagName != "Tag3" {
		t.Errorf("dbTaggings2[1].TagName: Expected %s, but got %s", "Tag3", dbTaggings2[1].TagName)
	}
}

func TestUpdate(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	dbCategory1 := &dbModel.Category{
		ID: "21111111-1111-1111-1111-111111111111",
		Name: "Category1",
		DisplayOrder: null.IntFrom(99),
	}
	err := dbCategory1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbCategory2 := &dbModel.Category{
		ID: "21111111-1111-1111-1111-111111111112",
		Name: "Category2",
		DisplayOrder: null.IntFrom(99),
	}
	err = dbCategory2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	categoryId1, err := uuid.Parse("21111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId2, err := uuid.Parse("21111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}

	article1, err := model.NewArticle("Title1", "Content1", categoryId1, []string{"Tag1"}, false)
	if err != nil {
		panic(err)
	}
	article2, err := model.NewArticle("Title2", "Content2", categoryId2, []string{"Tag2"}, false)
	if err != nil {
		panic(err)
	}

	r := NewArticleRepository(ctx, tx)
	err = r.Insert(article1)
	if err != nil {
		panic(err)
	}
	err = r.Insert(article2)
	if err != nil {
		panic(err)
	}

	// Execute
	article1.Title = "Title1Changed"
	article1.Content = "Content1Changed"
	article1.CategoryId = categoryId2
	article1.ChangeTags([]string{"Tag2", "Tag3"})
	err = article1.ChangeStatus(model.Published)
	if err != nil {
		panic(err)
	}
	err = r.Update(article1)
	if err != nil {
		panic(err)
	}

	// Check
	dbArticle1Check, err := dbModel.FindArticle(ctx, tx, article1.Id.String())
	if err != nil {
		panic(err)
	}
	if dbArticle1Check.Title != "Title1Changed" {
		t.Errorf("dbArticle1Check.Title: Expected %s, but got %s", "Title1Changed", dbArticle1Check.Title)
	}
	if dbArticle1Check.Content != "Content1Changed" {
		t.Errorf("dbArticle1Check.Content: Expected %s, but got %s", "Content1Changed", dbArticle1Check.Content)
	}
	if dbArticle1Check.CategoryID != "21111111-1111-1111-1111-111111111112" {
		t.Errorf("dbArticle1Check.CategoryId: Expected %s, but got %s", "11111111-1111-1111-1111-111111111112", dbArticle1Check.CategoryID)
	}
	if dbArticle1Check.Status != "Published" {
		t.Errorf("dbArticle1Check.Status: Expected %s, but got %s", "Published", dbArticle1Check.Status)
	}
	if !dbArticle1Check.PublishedAt.Valid {
		t.Errorf("dbArticle1Check.PublishedAt: Expected %v, but got %v", "Valid is true", dbArticle1Check.PublishedAt)
	}

	dbTag1Check, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag1")).One(ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	dbTag2Check, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag2")).One(ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	dbTag3Check, err := dbModel.Tags(dbModel.TagWhere.Name.EQ("Tag3")).One(ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if dbTag1Check != nil {
		t.Errorf("Tag1: Expected %v, but got %v", nil, &dbTag1Check)
	}
	if dbTag2Check == nil {
		t.Errorf("Tag2: Expected not nil, but got %v", &dbTag2Check)
	}
	if dbTag3Check == nil {
		t.Errorf("Tag3: Expected not nil, but got %v", &dbTag3Check)
	}
	
	dbTaggings1, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(article1.Id.String())).All(ctx, tx)
	if err != nil {
		panic(err)
	}
	if len(dbTaggings1) != 2 {
		t.Errorf("len(dbTaggings2): Expected count %d, but got %d", 2, len(dbTaggings1))
	}
	if dbTaggings1[0].TagName != "Tag2" {
		t.Errorf("dbTaggings1[0].TagName: Expected %s, but got %s", "Tag2", dbTaggings1[0].TagName)
	}
	if dbTaggings1[1].TagName != "Tag3" {
		t.Errorf("dbTaggings1[1].TagName: Expected %s, but got %s", "Tag3", dbTaggings1[1].TagName)
	}

	// Execute2
	err = article1.ChangeStatus(model.Draft)
	if err != nil {
		panic(err)
	}
	err = r.Update(article1)
	if err != nil {
		panic(err)
	}
	
	// Check2
	dbArticle1Check2, err := dbModel.FindArticle(ctx, tx, article1.Id.String())
	if dbArticle1Check2.Status != "Draft" {
		t.Errorf("dbArticle1Check2.Status: Expected %s, but got %s", "Draft", dbArticle1Check2.Status)
	}
	if dbArticle1Check2.PublishedAt.Valid {
		t.Errorf("dbArticle1Check2.PublishedAt: Expected %v, but got %v", "Valid is false", dbArticle1Check2.PublishedAt)
	}
}

func TestDelete(t *testing.T) {
	db := GetTestConnection()
	ctx := context.TODO()
	tx := GetTestTransaction(db, ctx)
	defer tx.Rollback()

	// Prepare data
	dbCategory1 := &dbModel.Category{
		ID: "21111111-1111-1111-1111-111111111111",
		Name: "Category1",
		DisplayOrder: null.IntFrom(99),
	}
	err := dbCategory1.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	dbCategory2 := &dbModel.Category{
		ID: "21111111-1111-1111-1111-111111111112",
		Name: "Category2",
		DisplayOrder: null.IntFrom(99),
	}
	err = dbCategory2.Insert(ctx, tx, boil.Infer())
	if err != nil {
		panic(err)
	}

	categoryId1, err := uuid.Parse("21111111-1111-1111-1111-111111111111")
	if err != nil {
		panic(err)
	}
	categoryId2, err := uuid.Parse("21111111-1111-1111-1111-111111111112")
	if err != nil {
		panic(err)
	}

	article1, err := model.NewArticle("Title1", "Content1", categoryId1, []string{"Tag1", "Tag2"}, false)
	if err != nil {
		panic(err)
	}
	article2, err := model.NewArticle("Title2", "Content2", categoryId2, []string{"Tag2"}, false)
	if err != nil {
		panic(err)
	}

	r := NewArticleRepository(ctx, tx)
	err = r.Insert(article1)
	if err != nil {
		panic(err)
	}
	err = r.Insert(article2)
	if err != nil {
		panic(err)
	}

	// Execute
	err = r.Delete(article1.Id)
	if err != nil {
		panic(err)
	}

	// Check
	dbArticle1, err := dbModel.FindArticle(ctx, tx, article1.Id.String())
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if dbArticle1 != nil {
		t.Errorf("dbArticle1: Expected %v, but got %v", nil, dbArticle1)
	}

	dbTag1, err := dbModel.FindTag(ctx, tx, "Tag1")
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if dbTag1 != nil {
		t.Errorf("dbTag1: Expected %v, but got %v", nil, dbTag1)
	}
	dbTag2, err := dbModel.FindTag(ctx, tx, "Tag2")
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if dbTag2 == nil {
		t.Errorf("dbTag2: Expected %v, but got %v", "not nil", dbTag2)
	}

	dbTaggings, err := dbModel.Taggings(dbModel.TaggingWhere.ArticleID.EQ(article1.Id.String())).All(ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if len(dbTaggings) != 0 {
		t.Errorf("len(dbTaggings): Expected %v, but got %v", 0, len(dbTaggings))
	}
}