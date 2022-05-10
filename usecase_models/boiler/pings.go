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

// Ping is an object representing the database table.
type Ping struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Data      null.String `boil:"data" json:"data,omitempty" toml:"data" yaml:"data,omitempty"`
	ProjectID int         `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *pingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PingColumns = struct {
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

var PingTableColumns = struct {
	ID        string
	Data      string
	ProjectID string
	UpdatedAt string
	CreatedAt string
	DeletedAt string
}{
	ID:        "pings.id",
	Data:      "pings.data",
	ProjectID: "pings.project_id",
	UpdatedAt: "pings.updated_at",
	CreatedAt: "pings.created_at",
	DeletedAt: "pings.deleted_at",
}

// Generated where

var PingWhere = struct {
	ID        whereHelperint
	Data      whereHelpernull_String
	ProjectID whereHelperint
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"pings\".\"id\""},
	Data:      whereHelpernull_String{field: "\"pings\".\"data\""},
	ProjectID: whereHelperint{field: "\"pings\".\"project_id\""},
	UpdatedAt: whereHelpertime_Time{field: "\"pings\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"pings\".\"created_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"pings\".\"deleted_at\""},
}

// PingRels is where relationship names are stored.
var PingRels = struct {
	Project string
}{
	Project: "Project",
}

// pingR is where relationships are stored.
type pingR struct {
	Project *Project `boil:"Project" json:"Project" toml:"Project" yaml:"Project"`
}

// NewStruct creates a new relationship struct
func (*pingR) NewStruct() *pingR {
	return &pingR{}
}

// pingL is where Load methods for each relationship are stored.
type pingL struct{}

var (
	pingAllColumns            = []string{"id", "data", "project_id", "updated_at", "created_at", "deleted_at"}
	pingColumnsWithoutDefault = []string{"project_id", "created_at"}
	pingColumnsWithDefault    = []string{"id", "data", "updated_at", "deleted_at"}
	pingPrimaryKeyColumns     = []string{"id"}
	pingGeneratedColumns      = []string{}
)

type (
	// PingSlice is an alias for a slice of pointers to Ping.
	// This should almost always be used instead of []Ping.
	PingSlice []*Ping
	// PingHook is the signature for custom Ping hook methods
	PingHook func(context.Context, boil.ContextExecutor, *Ping) error

	pingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pingType                 = reflect.TypeOf(&Ping{})
	pingMapping              = queries.MakeStructMapping(pingType)
	pingPrimaryKeyMapping, _ = queries.BindMapping(pingType, pingMapping, pingPrimaryKeyColumns)
	pingInsertCacheMut       sync.RWMutex
	pingInsertCache          = make(map[string]insertCache)
	pingUpdateCacheMut       sync.RWMutex
	pingUpdateCache          = make(map[string]updateCache)
	pingUpsertCacheMut       sync.RWMutex
	pingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var pingAfterSelectHooks []PingHook

var pingBeforeInsertHooks []PingHook
var pingAfterInsertHooks []PingHook

var pingBeforeUpdateHooks []PingHook
var pingAfterUpdateHooks []PingHook

var pingBeforeDeleteHooks []PingHook
var pingAfterDeleteHooks []PingHook

var pingBeforeUpsertHooks []PingHook
var pingAfterUpsertHooks []PingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Ping) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Ping) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Ping) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Ping) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Ping) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Ping) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Ping) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Ping) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Ping) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPingHook registers your hook function for all future operations.
func AddPingHook(hookPoint boil.HookPoint, pingHook PingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		pingAfterSelectHooks = append(pingAfterSelectHooks, pingHook)
	case boil.BeforeInsertHook:
		pingBeforeInsertHooks = append(pingBeforeInsertHooks, pingHook)
	case boil.AfterInsertHook:
		pingAfterInsertHooks = append(pingAfterInsertHooks, pingHook)
	case boil.BeforeUpdateHook:
		pingBeforeUpdateHooks = append(pingBeforeUpdateHooks, pingHook)
	case boil.AfterUpdateHook:
		pingAfterUpdateHooks = append(pingAfterUpdateHooks, pingHook)
	case boil.BeforeDeleteHook:
		pingBeforeDeleteHooks = append(pingBeforeDeleteHooks, pingHook)
	case boil.AfterDeleteHook:
		pingAfterDeleteHooks = append(pingAfterDeleteHooks, pingHook)
	case boil.BeforeUpsertHook:
		pingBeforeUpsertHooks = append(pingBeforeUpsertHooks, pingHook)
	case boil.AfterUpsertHook:
		pingAfterUpsertHooks = append(pingAfterUpsertHooks, pingHook)
	}
}

// One returns a single ping record from the query.
func (q pingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Ping, error) {
	o := &Ping{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Ping records from the query.
func (q pingQuery) All(ctx context.Context, exec boil.ContextExecutor) (PingSlice, error) {
	var o []*Ping

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Ping slice")
	}

	if len(pingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Ping records in the query.
func (q pingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pings exists")
	}

	return count > 0, nil
}

// Project pointed to by the foreign key.
func (o *Ping) Project(mods ...qm.QueryMod) projectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	return Projects(queryMods...)
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pingL) LoadProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybePing interface{}, mods queries.Applicator) error {
	var slice []*Ping
	var object *Ping

	if singular {
		object = maybePing.(*Ping)
	} else {
		slice = *maybePing.(*[]*Ping)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pingR{}
		}
		args = append(args, object.ProjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pingR{}
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

	if len(pingAfterSelectHooks) != 0 {
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
		foreign.R.Pings = append(foreign.R.Pings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProjectID == foreign.ID {
				local.R.Project = foreign
				if foreign.R == nil {
					foreign.R = &projectR{}
				}
				foreign.R.Pings = append(foreign.R.Pings, local)
				break
			}
		}
	}

	return nil
}

// SetProject of the ping to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.Pings.
func (o *Ping) SetProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Project) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"pings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
		strmangle.WhereClause("\"", "\"", 2, pingPrimaryKeyColumns),
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
		o.R = &pingR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &projectR{
			Pings: PingSlice{o},
		}
	} else {
		related.R.Pings = append(related.R.Pings, o)
	}

	return nil
}

// Pings retrieves all the records using an executor.
func Pings(mods ...qm.QueryMod) pingQuery {
	mods = append(mods, qm.From("\"pings\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"pings\".*"})
	}

	return pingQuery{NewQuery(mods...)}
}

// FindPing retrieves a single record by SequenceId with an executor.
// If selectCols is empty Find will return all columns.
func FindPing(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Ping, error) {
	pingObj := &Ping{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pings\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pings")
	}

	if err = pingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return pingObj, err
	}

	return pingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Ping) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pings provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(pingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pingInsertCacheMut.RLock()
	cache, cached := pingInsertCache[key]
	pingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pingAllColumns,
			pingColumnsWithDefault,
			pingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pingType, pingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pingType, pingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pings\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pings\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into pings")
	}

	if !cached {
		pingInsertCacheMut.Lock()
		pingInsertCache[key] = cache
		pingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Ping.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Ping) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pingUpdateCacheMut.RLock()
	cache, cached := pingUpdateCache[key]
	pingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pingAllColumns,
			pingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pings\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pingType, pingMapping, append(wl, pingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update pings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pings")
	}

	if !cached {
		pingUpdateCacheMut.Lock()
		pingUpdateCache[key] = cache
		pingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ping slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ping")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Ping) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pings provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(pingColumnsWithDefault, o)

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

	pingUpsertCacheMut.RLock()
	cache, cached := pingUpsertCache[key]
	pingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pingAllColumns,
			pingColumnsWithDefault,
			pingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			pingAllColumns,
			pingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pings, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pingPrimaryKeyColumns))
			copy(conflict, pingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pings\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pingType, pingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pingType, pingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert pings")
	}

	if !cached {
		pingUpsertCacheMut.Lock()
		pingUpsertCache[key] = cache
		pingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Ping record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Ping) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Ping provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pingPrimaryKeyMapping)
	sql := "DELETE FROM \"pings\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(pingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ping slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pings")
	}

	if len(pingAfterDeleteHooks) != 0 {
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
func (o *Ping) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPing(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pings\".* FROM \"pings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PingSlice")
	}

	*o = slice

	return nil
}

// PingExists checks if the Ping row exists.
func PingExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pings\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pings exists")
	}

	return exists, nil
}
