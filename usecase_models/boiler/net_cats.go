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

// NetCat is an object representing the database table.
type NetCat struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Data      null.String `boil:"data" json:"data,omitempty" toml:"data" yaml:"data,omitempty"`
	ProjectID int         `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *netCatR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L netCatL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NetCatColumns = struct {
	ID        string
	Data      string
	ProjectID string
	UpdatedAt string
	CreatedAt string
	DeletedAt string
}{
	ID:        "id",
	Data:      "data",
	ProjectID: "project_id",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

var NetCatTableColumns = struct {
	ID        string
	Data      string
	ProjectID string
	UpdatedAt string
	CreatedAt string
	DeletedAt string
}{
	ID:        "net_cats.id",
	Data:      "net_cats.data",
	ProjectID: "net_cats.project_id",
	UpdatedAt: "net_cats.updated_at",
	CreatedAt: "net_cats.created_at",
	DeletedAt: "net_cats.deleted_at",
}

// Generated where

var NetCatWhere = struct {
	ID        whereHelperint
	Data      whereHelpernull_String
	ProjectID whereHelperint
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"net_cats\".\"id\""},
	Data:      whereHelpernull_String{field: "\"net_cats\".\"data\""},
	ProjectID: whereHelperint{field: "\"net_cats\".\"project_id\""},
	UpdatedAt: whereHelpertime_Time{field: "\"net_cats\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"net_cats\".\"created_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"net_cats\".\"deleted_at\""},
}

// NetCatRels is where relationship names are stored.
var NetCatRels = struct {
	Project string
}{
	Project: "Project",
}

// netCatR is where relationships are stored.
type netCatR struct {
	Project *Project `boil:"Project" json:"Project" toml:"Project" yaml:"Project"`
}

// NewStruct creates a new relationship struct
func (*netCatR) NewStruct() *netCatR {
	return &netCatR{}
}

// netCatL is where Load methods for each relationship are stored.
type netCatL struct{}

var (
	netCatAllColumns            = []string{"id", "data", "project_id", "updated_at", "created_at", "deleted_at"}
	netCatColumnsWithoutDefault = []string{"project_id", "created_at"}
	netCatColumnsWithDefault    = []string{"id", "data", "updated_at", "deleted_at"}
	netCatPrimaryKeyColumns     = []string{"id"}
	netCatGeneratedColumns      = []string{}
)

type (
	// NetCatSlice is an alias for a slice of pointers to NetCat.
	// This should almost always be used instead of []NetCat.
	NetCatSlice []*NetCat
	// NetCatHook is the signature for custom NetCat hook methods
	NetCatHook func(context.Context, boil.ContextExecutor, *NetCat) error

	netCatQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	netCatType                 = reflect.TypeOf(&NetCat{})
	netCatMapping              = queries.MakeStructMapping(netCatType)
	netCatPrimaryKeyMapping, _ = queries.BindMapping(netCatType, netCatMapping, netCatPrimaryKeyColumns)
	netCatInsertCacheMut       sync.RWMutex
	netCatInsertCache          = make(map[string]insertCache)
	netCatUpdateCacheMut       sync.RWMutex
	netCatUpdateCache          = make(map[string]updateCache)
	netCatUpsertCacheMut       sync.RWMutex
	netCatUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var netCatAfterSelectHooks []NetCatHook

var netCatBeforeInsertHooks []NetCatHook
var netCatAfterInsertHooks []NetCatHook

var netCatBeforeUpdateHooks []NetCatHook
var netCatAfterUpdateHooks []NetCatHook

var netCatBeforeDeleteHooks []NetCatHook
var netCatAfterDeleteHooks []NetCatHook

var netCatBeforeUpsertHooks []NetCatHook
var netCatAfterUpsertHooks []NetCatHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NetCat) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NetCat) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NetCat) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NetCat) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NetCat) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NetCat) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NetCat) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NetCat) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NetCat) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range netCatAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNetCatHook registers your hook function for all future operations.
func AddNetCatHook(hookPoint boil.HookPoint, netCatHook NetCatHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		netCatAfterSelectHooks = append(netCatAfterSelectHooks, netCatHook)
	case boil.BeforeInsertHook:
		netCatBeforeInsertHooks = append(netCatBeforeInsertHooks, netCatHook)
	case boil.AfterInsertHook:
		netCatAfterInsertHooks = append(netCatAfterInsertHooks, netCatHook)
	case boil.BeforeUpdateHook:
		netCatBeforeUpdateHooks = append(netCatBeforeUpdateHooks, netCatHook)
	case boil.AfterUpdateHook:
		netCatAfterUpdateHooks = append(netCatAfterUpdateHooks, netCatHook)
	case boil.BeforeDeleteHook:
		netCatBeforeDeleteHooks = append(netCatBeforeDeleteHooks, netCatHook)
	case boil.AfterDeleteHook:
		netCatAfterDeleteHooks = append(netCatAfterDeleteHooks, netCatHook)
	case boil.BeforeUpsertHook:
		netCatBeforeUpsertHooks = append(netCatBeforeUpsertHooks, netCatHook)
	case boil.AfterUpsertHook:
		netCatAfterUpsertHooks = append(netCatAfterUpsertHooks, netCatHook)
	}
}

// One returns a single netCat record from the query.
func (q netCatQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NetCat, error) {
	o := &NetCat{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for net_cats")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NetCat records from the query.
func (q netCatQuery) All(ctx context.Context, exec boil.ContextExecutor) (NetCatSlice, error) {
	var o []*NetCat

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NetCat slice")
	}

	if len(netCatAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NetCat records in the query.
func (q netCatQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count net_cats rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q netCatQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if net_cats exists")
	}

	return count > 0, nil
}

// Project pointed to by the foreign key.
func (o *NetCat) Project(mods ...qm.QueryMod) projectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	return Projects(queryMods...)
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (netCatL) LoadProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNetCat interface{}, mods queries.Applicator) error {
	var slice []*NetCat
	var object *NetCat

	if singular {
		object = maybeNetCat.(*NetCat)
	} else {
		slice = *maybeNetCat.(*[]*NetCat)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &netCatR{}
		}
		args = append(args, object.ProjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &netCatR{}
			}

			for _, a := range args {
				if a == obj.ProjectID {
					continue Outer
				}
			}

			args = append(args, obj.ProjectID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`projects`),
		qm.WhereIn(`projects.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Project")
	}

	var resultSlice []*Project
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Project")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projects")
	}

	if len(netCatAfterSelectHooks) != 0 {
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
		object.R.Project = foreign
		if foreign.R == nil {
			foreign.R = &projectR{}
		}
		foreign.R.NetCats = append(foreign.R.NetCats, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProjectID == foreign.ID {
				local.R.Project = foreign
				if foreign.R == nil {
					foreign.R = &projectR{}
				}
				foreign.R.NetCats = append(foreign.R.NetCats, local)
				break
			}
		}
	}

	return nil
}

// SetProject of the netCat to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.NetCats.
func (o *NetCat) SetProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Project) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"net_cats\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
		strmangle.WhereClause("\"", "\"", 2, netCatPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProjectID = related.ID
	if o.R == nil {
		o.R = &netCatR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &projectR{
			NetCats: NetCatSlice{o},
		}
	} else {
		related.R.NetCats = append(related.R.NetCats, o)
	}

	return nil
}

// NetCats retrieves all the records using an executor.
func NetCats(mods ...qm.QueryMod) netCatQuery {
	mods = append(mods, qm.From("\"net_cats\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"net_cats\".*"})
	}

	return netCatQuery{NewQuery(mods...)}
}

// FindNetCat retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNetCat(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*NetCat, error) {
	netCatObj := &NetCat{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"net_cats\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, netCatObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from net_cats")
	}

	if err = netCatObj.doAfterSelectHooks(ctx, exec); err != nil {
		return netCatObj, err
	}

	return netCatObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NetCat) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no net_cats provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(netCatColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	netCatInsertCacheMut.RLock()
	cache, cached := netCatInsertCache[key]
	netCatInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			netCatAllColumns,
			netCatColumnsWithDefault,
			netCatColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(netCatType, netCatMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(netCatType, netCatMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"net_cats\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"net_cats\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into net_cats")
	}

	if !cached {
		netCatInsertCacheMut.Lock()
		netCatInsertCache[key] = cache
		netCatInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NetCat.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NetCat) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	netCatUpdateCacheMut.RLock()
	cache, cached := netCatUpdateCache[key]
	netCatUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			netCatAllColumns,
			netCatPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update net_cats, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"net_cats\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, netCatPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(netCatType, netCatMapping, append(wl, netCatPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update net_cats row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for net_cats")
	}

	if !cached {
		netCatUpdateCacheMut.Lock()
		netCatUpdateCache[key] = cache
		netCatUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q netCatQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for net_cats")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for net_cats")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NetCatSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), netCatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"net_cats\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, netCatPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in netCat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all netCat")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NetCat) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no net_cats provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(netCatColumnsWithDefault, o)

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

	netCatUpsertCacheMut.RLock()
	cache, cached := netCatUpsertCache[key]
	netCatUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			netCatAllColumns,
			netCatColumnsWithDefault,
			netCatColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			netCatAllColumns,
			netCatPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert net_cats, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(netCatPrimaryKeyColumns))
			copy(conflict, netCatPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"net_cats\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(netCatType, netCatMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(netCatType, netCatMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert net_cats")
	}

	if !cached {
		netCatUpsertCacheMut.Lock()
		netCatUpsertCache[key] = cache
		netCatUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single NetCat record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NetCat) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NetCat provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), netCatPrimaryKeyMapping)
	sql := "DELETE FROM \"net_cats\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from net_cats")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for net_cats")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q netCatQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no netCatQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from net_cats")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for net_cats")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NetCatSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(netCatBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), netCatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"net_cats\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, netCatPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from netCat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for net_cats")
	}

	if len(netCatAfterDeleteHooks) != 0 {
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
func (o *NetCat) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNetCat(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NetCatSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NetCatSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), netCatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"net_cats\".* FROM \"net_cats\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, netCatPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NetCatSlice")
	}

	*o = slice

	return nil
}

// NetCatExists checks if the NetCat row exists.
func NetCatExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"net_cats\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if net_cats exists")
	}

	return exists, nil
}
