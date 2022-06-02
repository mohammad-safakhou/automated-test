// Code generated by SQLBoiler 4.9.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Datacenter is an object representing the database table.
type Datacenter struct {
	ID             int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Baseurl        string    `boil:"baseurl" json:"baseurl" toml:"baseurl" yaml:"baseurl"`
	Title          string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	ConnectionRate null.Int  `boil:"connection_rate" json:"connection_rate,omitempty" toml:"connection_rate" yaml:"connection_rate,omitempty"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt      null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *datacenterR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L datacenterL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DatacenterColumns = struct {
	ID             string
	Baseurl        string
	Title          string
	ConnectionRate string
	UpdatedAt      string
	CreatedAt      string
	DeletedAt      string
}{
	ID:             "id",
	Baseurl:        "baseurl",
	Title:          "title",
	ConnectionRate: "connection_rate",
	UpdatedAt:      "updated_at",
	CreatedAt:      "created_at",
	DeletedAt:      "deleted_at",
}

var DatacenterTableColumns = struct {
	ID             string
	Baseurl        string
	Title          string
	ConnectionRate string
	UpdatedAt      string
	CreatedAt      string
	DeletedAt      string
}{
	ID:             "datacenters.id",
	Baseurl:        "datacenters.baseurl",
	Title:          "datacenters.title",
	ConnectionRate: "datacenters.connection_rate",
	UpdatedAt:      "datacenters.updated_at",
	CreatedAt:      "datacenters.created_at",
	DeletedAt:      "datacenters.deleted_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var DatacenterWhere = struct {
	ID             whereHelperint
	Baseurl        whereHelperstring
	Title          whereHelperstring
	ConnectionRate whereHelpernull_Int
	UpdatedAt      whereHelpertime_Time
	CreatedAt      whereHelpertime_Time
	DeletedAt      whereHelpernull_Time
}{
	ID:             whereHelperint{field: "\"datacenters\".\"id\""},
	Baseurl:        whereHelperstring{field: "\"datacenters\".\"baseurl\""},
	Title:          whereHelperstring{field: "\"datacenters\".\"title\""},
	ConnectionRate: whereHelpernull_Int{field: "\"datacenters\".\"connection_rate\""},
	UpdatedAt:      whereHelpertime_Time{field: "\"datacenters\".\"updated_at\""},
	CreatedAt:      whereHelpertime_Time{field: "\"datacenters\".\"created_at\""},
	DeletedAt:      whereHelpernull_Time{field: "\"datacenters\".\"deleted_at\""},
}

// DatacenterRels is where relationship names are stored.
var DatacenterRels = struct {
	RelationDatacenters string
}{
	RelationDatacenters: "RelationDatacenters",
}

// datacenterR is where relationships are stored.
type datacenterR struct {
	RelationDatacenters RelationDatacenterSlice `boil:"RelationDatacenters" json:"RelationDatacenters" toml:"RelationDatacenters" yaml:"RelationDatacenters"`
}

// NewStruct creates a new relationship struct
func (*datacenterR) NewStruct() *datacenterR {
	return &datacenterR{}
}

// datacenterL is where Load methods for each relationship are stored.
type datacenterL struct{}

var (
	datacenterAllColumns            = []string{"id", "baseurl", "title", "connection_rate", "updated_at", "created_at", "deleted_at"}
	datacenterColumnsWithoutDefault = []string{"baseurl", "title", "created_at"}
	datacenterColumnsWithDefault    = []string{"id", "connection_rate", "updated_at", "deleted_at"}
	datacenterPrimaryKeyColumns     = []string{"id"}
	datacenterGeneratedColumns      = []string{}
)

type (
	// DatacenterSlice is an alias for a slice of pointers to Datacenter.
	// This should almost always be used instead of []Datacenter.
	DatacenterSlice []*Datacenter
	// DatacenterHook is the signature for custom Datacenter hook methods
	DatacenterHook func(context.Context, boil.ContextExecutor, *Datacenter) error

	datacenterQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	datacenterType                 = reflect.TypeOf(&Datacenter{})
	datacenterMapping              = queries.MakeStructMapping(datacenterType)
	datacenterPrimaryKeyMapping, _ = queries.BindMapping(datacenterType, datacenterMapping, datacenterPrimaryKeyColumns)
	datacenterInsertCacheMut       sync.RWMutex
	datacenterInsertCache          = make(map[string]insertCache)
	datacenterUpdateCacheMut       sync.RWMutex
	datacenterUpdateCache          = make(map[string]updateCache)
	datacenterUpsertCacheMut       sync.RWMutex
	datacenterUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var datacenterAfterSelectHooks []DatacenterHook

var datacenterBeforeInsertHooks []DatacenterHook
var datacenterAfterInsertHooks []DatacenterHook

var datacenterBeforeUpdateHooks []DatacenterHook
var datacenterAfterUpdateHooks []DatacenterHook

var datacenterBeforeDeleteHooks []DatacenterHook
var datacenterAfterDeleteHooks []DatacenterHook

var datacenterBeforeUpsertHooks []DatacenterHook
var datacenterAfterUpsertHooks []DatacenterHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Datacenter) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Datacenter) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Datacenter) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Datacenter) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Datacenter) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Datacenter) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Datacenter) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Datacenter) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Datacenter) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datacenterAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDatacenterHook registers your hook function for all future operations.
func AddDatacenterHook(hookPoint boil.HookPoint, datacenterHook DatacenterHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		datacenterAfterSelectHooks = append(datacenterAfterSelectHooks, datacenterHook)
	case boil.BeforeInsertHook:
		datacenterBeforeInsertHooks = append(datacenterBeforeInsertHooks, datacenterHook)
	case boil.AfterInsertHook:
		datacenterAfterInsertHooks = append(datacenterAfterInsertHooks, datacenterHook)
	case boil.BeforeUpdateHook:
		datacenterBeforeUpdateHooks = append(datacenterBeforeUpdateHooks, datacenterHook)
	case boil.AfterUpdateHook:
		datacenterAfterUpdateHooks = append(datacenterAfterUpdateHooks, datacenterHook)
	case boil.BeforeDeleteHook:
		datacenterBeforeDeleteHooks = append(datacenterBeforeDeleteHooks, datacenterHook)
	case boil.AfterDeleteHook:
		datacenterAfterDeleteHooks = append(datacenterAfterDeleteHooks, datacenterHook)
	case boil.BeforeUpsertHook:
		datacenterBeforeUpsertHooks = append(datacenterBeforeUpsertHooks, datacenterHook)
	case boil.AfterUpsertHook:
		datacenterAfterUpsertHooks = append(datacenterAfterUpsertHooks, datacenterHook)
	}
}

// One returns a single datacenter record from the query.
func (q datacenterQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Datacenter, error) {
	o := &Datacenter{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for datacenters")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Datacenter records from the query.
func (q datacenterQuery) All(ctx context.Context, exec boil.ContextExecutor) (DatacenterSlice, error) {
	var o []*Datacenter

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Datacenter slice")
	}

	if len(datacenterAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Datacenter records in the query.
func (q datacenterQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count datacenters rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q datacenterQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if datacenters exists")
	}

	return count > 0, nil
}

// RelationDatacenters retrieves all the relation_datacenter's RelationDatacenters with an executor.
func (o *Datacenter) RelationDatacenters(mods ...qm.QueryMod) relationDatacenterQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"relation_datacenters\".\"datacenter_id\"=?", o.ID),
	)

	return RelationDatacenters(queryMods...)
}

// LoadRelationDatacenters allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (datacenterL) LoadRelationDatacenters(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDatacenter interface{}, mods queries.Applicator) error {
	var slice []*Datacenter
	var object *Datacenter

	if singular {
		object = maybeDatacenter.(*Datacenter)
	} else {
		slice = *maybeDatacenter.(*[]*Datacenter)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &datacenterR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &datacenterR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`relation_datacenters`),
		qm.WhereIn(`relation_datacenters.datacenter_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load relation_datacenters")
	}

	var resultSlice []*RelationDatacenter
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice relation_datacenters")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on relation_datacenters")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for relation_datacenters")
	}

	if len(relationDatacenterAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RelationDatacenters = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &relationDatacenterR{}
			}
			foreign.R.Datacenter = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.DatacenterID) {
				local.R.RelationDatacenters = append(local.R.RelationDatacenters, foreign)
				if foreign.R == nil {
					foreign.R = &relationDatacenterR{}
				}
				foreign.R.Datacenter = local
				break
			}
		}
	}

	return nil
}

// AddRelationDatacenters adds the given related objects to the existing relationships
// of the datacenter, optionally inserting them as new records.
// Appends related to o.R.RelationDatacenters.
// Sets related.R.Datacenter appropriately.
func (o *Datacenter) AddRelationDatacenters(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RelationDatacenter) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.DatacenterID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"relation_datacenters\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"datacenter_id"}),
				strmangle.WhereClause("\"", "\"", 2, relationDatacenterPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.DatacenterID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &datacenterR{
			RelationDatacenters: related,
		}
	} else {
		o.R.RelationDatacenters = append(o.R.RelationDatacenters, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &relationDatacenterR{
				Datacenter: o,
			}
		} else {
			rel.R.Datacenter = o
		}
	}
	return nil
}

// SetRelationDatacenters removes all previously related items of the
// datacenter replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Datacenter's RelationDatacenters accordingly.
// Replaces o.R.RelationDatacenters with related.
// Sets related.R.Datacenter's RelationDatacenters accordingly.
func (o *Datacenter) SetRelationDatacenters(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RelationDatacenter) error {
	query := "update \"relation_datacenters\" set \"datacenter_id\" = null where \"datacenter_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.RelationDatacenters {
			queries.SetScanner(&rel.DatacenterID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Datacenter = nil
		}
		o.R.RelationDatacenters = nil
	}

	return o.AddRelationDatacenters(ctx, exec, insert, related...)
}

// RemoveRelationDatacenters relationships from objects passed in.
// Removes related items from R.RelationDatacenters (uses pointer comparison, removal does not keep order)
// Sets related.R.Datacenter.
func (o *Datacenter) RemoveRelationDatacenters(ctx context.Context, exec boil.ContextExecutor, related ...*RelationDatacenter) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.DatacenterID, nil)
		if rel.R != nil {
			rel.R.Datacenter = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("datacenter_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.RelationDatacenters {
			if rel != ri {
				continue
			}

			ln := len(o.R.RelationDatacenters)
			if ln > 1 && i < ln-1 {
				o.R.RelationDatacenters[i] = o.R.RelationDatacenters[ln-1]
			}
			o.R.RelationDatacenters = o.R.RelationDatacenters[:ln-1]
			break
		}
	}

	return nil
}

// Datacenters retrieves all the records using an executor.
func Datacenters(mods ...qm.QueryMod) datacenterQuery {
	mods = append(mods, qm.From("\"datacenters\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"datacenters\".*"})
	}

	return datacenterQuery{NewQuery(mods...)}
}

// FindDatacenter retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDatacenter(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Datacenter, error) {
	datacenterObj := &Datacenter{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"datacenters\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, datacenterObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from datacenters")
	}

	if err = datacenterObj.doAfterSelectHooks(ctx, exec); err != nil {
		return datacenterObj, err
	}

	return datacenterObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Datacenter) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no datacenters provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(datacenterColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	datacenterInsertCacheMut.RLock()
	cache, cached := datacenterInsertCache[key]
	datacenterInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			datacenterAllColumns,
			datacenterColumnsWithDefault,
			datacenterColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(datacenterType, datacenterMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(datacenterType, datacenterMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"datacenters\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"datacenters\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into datacenters")
	}

	if !cached {
		datacenterInsertCacheMut.Lock()
		datacenterInsertCache[key] = cache
		datacenterInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Datacenter.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Datacenter) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	datacenterUpdateCacheMut.RLock()
	cache, cached := datacenterUpdateCache[key]
	datacenterUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			datacenterAllColumns,
			datacenterPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update datacenters, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"datacenters\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, datacenterPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(datacenterType, datacenterMapping, append(wl, datacenterPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update datacenters row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for datacenters")
	}

	if !cached {
		datacenterUpdateCacheMut.Lock()
		datacenterUpdateCache[key] = cache
		datacenterUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q datacenterQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for datacenters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for datacenters")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DatacenterSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), datacenterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"datacenters\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, datacenterPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in datacenter slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all datacenter")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Datacenter) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no datacenters provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(datacenterColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	datacenterUpsertCacheMut.RLock()
	cache, cached := datacenterUpsertCache[key]
	datacenterUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			datacenterAllColumns,
			datacenterColumnsWithDefault,
			datacenterColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			datacenterAllColumns,
			datacenterPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert datacenters, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(datacenterPrimaryKeyColumns))
			copy(conflict, datacenterPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"datacenters\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(datacenterType, datacenterMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(datacenterType, datacenterMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert datacenters")
	}

	if !cached {
		datacenterUpsertCacheMut.Lock()
		datacenterUpsertCache[key] = cache
		datacenterUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Datacenter record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Datacenter) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Datacenter provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), datacenterPrimaryKeyMapping)
	sql := "DELETE FROM \"datacenters\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from datacenters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for datacenters")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q datacenterQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no datacenterQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from datacenters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for datacenters")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DatacenterSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(datacenterBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), datacenterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"datacenters\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, datacenterPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from datacenter slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for datacenters")
	}

	if len(datacenterAfterDeleteHooks) != 0 {
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
func (o *Datacenter) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDatacenter(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DatacenterSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DatacenterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), datacenterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"datacenters\".* FROM \"datacenters\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, datacenterPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DatacenterSlice")
	}

	*o = slice

	return nil
}

// DatacenterExists checks if the Datacenter row exists.
func DatacenterExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"datacenters\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if datacenters exists")
	}

	return exists, nil
}
