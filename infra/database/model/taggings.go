// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Tagging is an object representing the database table.
type Tagging struct {
	ArticleID string    `boil:"article_id" json:"article_id" toml:"article_id" yaml:"article_id"`
	TagID     string    `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *taggingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taggingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaggingColumns = struct {
	ArticleID string
	TagID     string
	CreatedAt string
}{
	ArticleID: "article_id",
	TagID:     "tag_id",
	CreatedAt: "created_at",
}

var TaggingTableColumns = struct {
	ArticleID string
	TagID     string
	CreatedAt string
}{
	ArticleID: "taggings.article_id",
	TagID:     "taggings.tag_id",
	CreatedAt: "taggings.created_at",
}

// Generated where

var TaggingWhere = struct {
	ArticleID whereHelperstring
	TagID     whereHelperstring
	CreatedAt whereHelpertime_Time
}{
	ArticleID: whereHelperstring{field: "`taggings`.`article_id`"},
	TagID:     whereHelperstring{field: "`taggings`.`tag_id`"},
	CreatedAt: whereHelpertime_Time{field: "`taggings`.`created_at`"},
}

// TaggingRels is where relationship names are stored.
var TaggingRels = struct {
	Article string
	Tag     string
}{
	Article: "Article",
	Tag:     "Tag",
}

// taggingR is where relationships are stored.
type taggingR struct {
	Article *Article `boil:"Article" json:"Article" toml:"Article" yaml:"Article"`
	Tag     *Tag     `boil:"Tag" json:"Tag" toml:"Tag" yaml:"Tag"`
}

// NewStruct creates a new relationship struct
func (*taggingR) NewStruct() *taggingR {
	return &taggingR{}
}

func (r *taggingR) GetArticle() *Article {
	if r == nil {
		return nil
	}
	return r.Article
}

func (r *taggingR) GetTag() *Tag {
	if r == nil {
		return nil
	}
	return r.Tag
}

// taggingL is where Load methods for each relationship are stored.
type taggingL struct{}

var (
	taggingAllColumns            = []string{"article_id", "tag_id", "created_at"}
	taggingColumnsWithoutDefault = []string{"article_id", "tag_id"}
	taggingColumnsWithDefault    = []string{"created_at"}
	taggingPrimaryKeyColumns     = []string{"article_id", "tag_id"}
	taggingGeneratedColumns      = []string{}
)

type (
	// TaggingSlice is an alias for a slice of pointers to Tagging.
	// This should almost always be used instead of []Tagging.
	TaggingSlice []*Tagging
	// TaggingHook is the signature for custom Tagging hook methods
	TaggingHook func(context.Context, boil.ContextExecutor, *Tagging) error

	taggingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taggingType                 = reflect.TypeOf(&Tagging{})
	taggingMapping              = queries.MakeStructMapping(taggingType)
	taggingPrimaryKeyMapping, _ = queries.BindMapping(taggingType, taggingMapping, taggingPrimaryKeyColumns)
	taggingInsertCacheMut       sync.RWMutex
	taggingInsertCache          = make(map[string]insertCache)
	taggingUpdateCacheMut       sync.RWMutex
	taggingUpdateCache          = make(map[string]updateCache)
	taggingUpsertCacheMut       sync.RWMutex
	taggingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var taggingAfterSelectHooks []TaggingHook

var taggingBeforeInsertHooks []TaggingHook
var taggingAfterInsertHooks []TaggingHook

var taggingBeforeUpdateHooks []TaggingHook
var taggingAfterUpdateHooks []TaggingHook

var taggingBeforeDeleteHooks []TaggingHook
var taggingAfterDeleteHooks []TaggingHook

var taggingBeforeUpsertHooks []TaggingHook
var taggingAfterUpsertHooks []TaggingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Tagging) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Tagging) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Tagging) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Tagging) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Tagging) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Tagging) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Tagging) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Tagging) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Tagging) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taggingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTaggingHook registers your hook function for all future operations.
func AddTaggingHook(hookPoint boil.HookPoint, taggingHook TaggingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		taggingAfterSelectHooks = append(taggingAfterSelectHooks, taggingHook)
	case boil.BeforeInsertHook:
		taggingBeforeInsertHooks = append(taggingBeforeInsertHooks, taggingHook)
	case boil.AfterInsertHook:
		taggingAfterInsertHooks = append(taggingAfterInsertHooks, taggingHook)
	case boil.BeforeUpdateHook:
		taggingBeforeUpdateHooks = append(taggingBeforeUpdateHooks, taggingHook)
	case boil.AfterUpdateHook:
		taggingAfterUpdateHooks = append(taggingAfterUpdateHooks, taggingHook)
	case boil.BeforeDeleteHook:
		taggingBeforeDeleteHooks = append(taggingBeforeDeleteHooks, taggingHook)
	case boil.AfterDeleteHook:
		taggingAfterDeleteHooks = append(taggingAfterDeleteHooks, taggingHook)
	case boil.BeforeUpsertHook:
		taggingBeforeUpsertHooks = append(taggingBeforeUpsertHooks, taggingHook)
	case boil.AfterUpsertHook:
		taggingAfterUpsertHooks = append(taggingAfterUpsertHooks, taggingHook)
	}
}

// One returns a single tagging record from the query.
func (q taggingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Tagging, error) {
	o := &Tagging{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for taggings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Tagging records from the query.
func (q taggingQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaggingSlice, error) {
	var o []*Tagging

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Tagging slice")
	}

	if len(taggingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Tagging records in the query.
func (q taggingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count taggings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q taggingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if taggings exists")
	}

	return count > 0, nil
}

// Article pointed to by the foreign key.
func (o *Tagging) Article(mods ...qm.QueryMod) articleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ArticleID),
	}

	queryMods = append(queryMods, mods...)

	return Articles(queryMods...)
}

// Tag pointed to by the foreign key.
func (o *Tagging) Tag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.TagID),
	}

	queryMods = append(queryMods, mods...)

	return Tags(queryMods...)
}

// LoadArticle allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (taggingL) LoadArticle(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTagging interface{}, mods queries.Applicator) error {
	var slice []*Tagging
	var object *Tagging

	if singular {
		var ok bool
		object, ok = maybeTagging.(*Tagging)
		if !ok {
			object = new(Tagging)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTagging)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTagging))
			}
		}
	} else {
		s, ok := maybeTagging.(*[]*Tagging)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTagging)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTagging))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taggingR{}
		}
		args = append(args, object.ArticleID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taggingR{}
			}

			for _, a := range args {
				if a == obj.ArticleID {
					continue Outer
				}
			}

			args = append(args, obj.ArticleID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`articles`),
		qm.WhereIn(`articles.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Article")
	}

	var resultSlice []*Article
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Article")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for articles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for articles")
	}

	if len(articleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Article = foreign
		if foreign.R == nil {
			foreign.R = &articleR{}
		}
		foreign.R.Taggings = append(foreign.R.Taggings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ArticleID == foreign.ID {
				local.R.Article = foreign
				if foreign.R == nil {
					foreign.R = &articleR{}
				}
				foreign.R.Taggings = append(foreign.R.Taggings, local)
				break
			}
		}
	}

	return nil
}

// LoadTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (taggingL) LoadTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTagging interface{}, mods queries.Applicator) error {
	var slice []*Tagging
	var object *Tagging

	if singular {
		var ok bool
		object, ok = maybeTagging.(*Tagging)
		if !ok {
			object = new(Tagging)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTagging)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTagging))
			}
		}
	} else {
		s, ok := maybeTagging.(*[]*Tagging)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTagging)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTagging))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taggingR{}
		}
		args = append(args, object.TagID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taggingR{}
			}

			for _, a := range args {
				if a == obj.TagID {
					continue Outer
				}
			}

			args = append(args, obj.TagID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tags`),
		qm.WhereIn(`tags.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tags")
	}

	if len(tagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Tag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.Taggings = append(foreign.R.Taggings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TagID == foreign.ID {
				local.R.Tag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.Taggings = append(foreign.R.Taggings, local)
				break
			}
		}
	}

	return nil
}

// SetArticle of the tagging to the related item.
// Sets o.R.Article to related.
// Adds o to related.R.Taggings.
func (o *Tagging) SetArticle(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Article) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `taggings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"article_id"}),
		strmangle.WhereClause("`", "`", 0, taggingPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ArticleID, o.TagID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ArticleID = related.ID
	if o.R == nil {
		o.R = &taggingR{
			Article: related,
		}
	} else {
		o.R.Article = related
	}

	if related.R == nil {
		related.R = &articleR{
			Taggings: TaggingSlice{o},
		}
	} else {
		related.R.Taggings = append(related.R.Taggings, o)
	}

	return nil
}

// SetTag of the tagging to the related item.
// Sets o.R.Tag to related.
// Adds o to related.R.Taggings.
func (o *Tagging) SetTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `taggings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"tag_id"}),
		strmangle.WhereClause("`", "`", 0, taggingPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ArticleID, o.TagID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TagID = related.ID
	if o.R == nil {
		o.R = &taggingR{
			Tag: related,
		}
	} else {
		o.R.Tag = related
	}

	if related.R == nil {
		related.R = &tagR{
			Taggings: TaggingSlice{o},
		}
	} else {
		related.R.Taggings = append(related.R.Taggings, o)
	}

	return nil
}

// Taggings retrieves all the records using an executor.
func Taggings(mods ...qm.QueryMod) taggingQuery {
	mods = append(mods, qm.From("`taggings`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`taggings`.*"})
	}

	return taggingQuery{q}
}

// FindTagging retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTagging(ctx context.Context, exec boil.ContextExecutor, articleID string, tagID string, selectCols ...string) (*Tagging, error) {
	taggingObj := &Tagging{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `taggings` where `article_id`=? AND `tag_id`=?", sel,
	)

	q := queries.Raw(query, articleID, tagID)

	err := q.Bind(ctx, exec, taggingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from taggings")
	}

	if err = taggingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return taggingObj, err
	}

	return taggingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tagging) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no taggings provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taggingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taggingInsertCacheMut.RLock()
	cache, cached := taggingInsertCache[key]
	taggingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taggingAllColumns,
			taggingColumnsWithDefault,
			taggingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(taggingType, taggingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taggingType, taggingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `taggings` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `taggings` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `taggings` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, taggingPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into taggings")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ArticleID,
		o.TagID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for taggings")
	}

CacheNoHooks:
	if !cached {
		taggingInsertCacheMut.Lock()
		taggingInsertCache[key] = cache
		taggingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Tagging.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tagging) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	taggingUpdateCacheMut.RLock()
	cache, cached := taggingUpdateCache[key]
	taggingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taggingAllColumns,
			taggingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update taggings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `taggings` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, taggingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taggingType, taggingMapping, append(wl, taggingPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update taggings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for taggings")
	}

	if !cached {
		taggingUpdateCacheMut.Lock()
		taggingUpdateCache[key] = cache
		taggingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q taggingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for taggings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for taggings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaggingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taggingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `taggings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taggingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in tagging slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all tagging")
	}
	return rowsAff, nil
}

var mySQLTaggingUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tagging) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no taggings provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taggingColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTaggingUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	taggingUpsertCacheMut.RLock()
	cache, cached := taggingUpsertCache[key]
	taggingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taggingAllColumns,
			taggingColumnsWithDefault,
			taggingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			taggingAllColumns,
			taggingPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model: unable to upsert taggings, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`taggings`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `taggings` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(taggingType, taggingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taggingType, taggingMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for taggings")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(taggingType, taggingMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for taggings")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for taggings")
	}

CacheNoHooks:
	if !cached {
		taggingUpsertCacheMut.Lock()
		taggingUpsertCache[key] = cache
		taggingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Tagging record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tagging) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Tagging provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taggingPrimaryKeyMapping)
	sql := "DELETE FROM `taggings` WHERE `article_id`=? AND `tag_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from taggings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for taggings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q taggingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no taggingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from taggings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for taggings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaggingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(taggingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taggingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `taggings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taggingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from tagging slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for taggings")
	}

	if len(taggingAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tagging) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTagging(ctx, exec, o.ArticleID, o.TagID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaggingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaggingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taggingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `taggings`.* FROM `taggings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taggingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in TaggingSlice")
	}

	*o = slice

	return nil
}

// TaggingExists checks if the Tagging row exists.
func TaggingExists(ctx context.Context, exec boil.ContextExecutor, articleID string, tagID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `taggings` where `article_id`=? AND `tag_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, articleID, tagID)
	}
	row := exec.QueryRowContext(ctx, sql, articleID, tagID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if taggings exists")
	}

	return exists, nil
}

// Exists checks if the Tagging row exists.
func (o *Tagging) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TaggingExists(ctx, exec, o.ArticleID, o.TagID)
}