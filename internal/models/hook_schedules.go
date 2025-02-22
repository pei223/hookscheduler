// Code generated by SQLBoiler 4.17.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// HookSchedule is an object representing the database table.
type HookSchedule struct {
	HookScheduleID       uuid.UUID `boil:"hook_schedule_id" json:"hook_schedule_id" toml:"hook_schedule_id" yaml:"hook_schedule_id"`
	HookID               uuid.UUID `boil:"hook_id" json:"hook_id" toml:"hook_id" yaml:"hook_id"`
	DisplayName          string    `boil:"display_name" json:"display_name" toml:"display_name" yaml:"display_name"`
	Description          string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	ScheduleIntervalUnit int16     `boil:"schedule_interval_unit" json:"schedule_interval_unit" toml:"schedule_interval_unit" yaml:"schedule_interval_unit"`
	ScheduleTimeMonth    int16     `boil:"schedule_time_month" json:"schedule_time_month" toml:"schedule_time_month" yaml:"schedule_time_month"`
	ScheduleTimeDay      int16     `boil:"schedule_time_day" json:"schedule_time_day" toml:"schedule_time_day" yaml:"schedule_time_day"`
	ScheduleTimeHour     int16     `boil:"schedule_time_hour" json:"schedule_time_hour" toml:"schedule_time_hour" yaml:"schedule_time_hour"`
	ScheduleTimeMinute   int16     `boil:"schedule_time_minute" json:"schedule_time_minute" toml:"schedule_time_minute" yaml:"schedule_time_minute"`
	ScheduleTimeSecond   int16     `boil:"schedule_time_second" json:"schedule_time_second" toml:"schedule_time_second" yaml:"schedule_time_second"`

	R *hookScheduleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L hookScheduleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HookScheduleColumns = struct {
	HookScheduleID       string
	HookID               string
	DisplayName          string
	Description          string
	ScheduleIntervalUnit string
	ScheduleTimeMonth    string
	ScheduleTimeDay      string
	ScheduleTimeHour     string
	ScheduleTimeMinute   string
	ScheduleTimeSecond   string
}{
	HookScheduleID:       "hook_schedule_id",
	HookID:               "hook_id",
	DisplayName:          "display_name",
	Description:          "description",
	ScheduleIntervalUnit: "schedule_interval_unit",
	ScheduleTimeMonth:    "schedule_time_month",
	ScheduleTimeDay:      "schedule_time_day",
	ScheduleTimeHour:     "schedule_time_hour",
	ScheduleTimeMinute:   "schedule_time_minute",
	ScheduleTimeSecond:   "schedule_time_second",
}

var HookScheduleTableColumns = struct {
	HookScheduleID       string
	HookID               string
	DisplayName          string
	Description          string
	ScheduleIntervalUnit string
	ScheduleTimeMonth    string
	ScheduleTimeDay      string
	ScheduleTimeHour     string
	ScheduleTimeMinute   string
	ScheduleTimeSecond   string
}{
	HookScheduleID:       "hook_schedules.hook_schedule_id",
	HookID:               "hook_schedules.hook_id",
	DisplayName:          "hook_schedules.display_name",
	Description:          "hook_schedules.description",
	ScheduleIntervalUnit: "hook_schedules.schedule_interval_unit",
	ScheduleTimeMonth:    "hook_schedules.schedule_time_month",
	ScheduleTimeDay:      "hook_schedules.schedule_time_day",
	ScheduleTimeHour:     "hook_schedules.schedule_time_hour",
	ScheduleTimeMinute:   "hook_schedules.schedule_time_minute",
	ScheduleTimeSecond:   "hook_schedules.schedule_time_second",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod    { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod   { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod   { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) SIMILAR(x string) qm.QueryMod { return qm.Where(w.field+" SIMILAR TO ?", x) }
func (w whereHelperstring) NSIMILAR(x string) qm.QueryMod {
	return qm.Where(w.field+" NOT SIMILAR TO ?", x)
}
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

var HookScheduleWhere = struct {
	HookScheduleID       whereHelperuuid_UUID
	HookID               whereHelperuuid_UUID
	DisplayName          whereHelperstring
	Description          whereHelperstring
	ScheduleIntervalUnit whereHelperint16
	ScheduleTimeMonth    whereHelperint16
	ScheduleTimeDay      whereHelperint16
	ScheduleTimeHour     whereHelperint16
	ScheduleTimeMinute   whereHelperint16
	ScheduleTimeSecond   whereHelperint16
}{
	HookScheduleID:       whereHelperuuid_UUID{field: "\"hook_schedules\".\"hook_schedule_id\""},
	HookID:               whereHelperuuid_UUID{field: "\"hook_schedules\".\"hook_id\""},
	DisplayName:          whereHelperstring{field: "\"hook_schedules\".\"display_name\""},
	Description:          whereHelperstring{field: "\"hook_schedules\".\"description\""},
	ScheduleIntervalUnit: whereHelperint16{field: "\"hook_schedules\".\"schedule_interval_unit\""},
	ScheduleTimeMonth:    whereHelperint16{field: "\"hook_schedules\".\"schedule_time_month\""},
	ScheduleTimeDay:      whereHelperint16{field: "\"hook_schedules\".\"schedule_time_day\""},
	ScheduleTimeHour:     whereHelperint16{field: "\"hook_schedules\".\"schedule_time_hour\""},
	ScheduleTimeMinute:   whereHelperint16{field: "\"hook_schedules\".\"schedule_time_minute\""},
	ScheduleTimeSecond:   whereHelperint16{field: "\"hook_schedules\".\"schedule_time_second\""},
}

// HookScheduleRels is where relationship names are stored.
var HookScheduleRels = struct {
	Hook string
}{
	Hook: "Hook",
}

// hookScheduleR is where relationships are stored.
type hookScheduleR struct {
	Hook *Hook `boil:"Hook" json:"Hook" toml:"Hook" yaml:"Hook"`
}

// NewStruct creates a new relationship struct
func (*hookScheduleR) NewStruct() *hookScheduleR {
	return &hookScheduleR{}
}

func (r *hookScheduleR) GetHook() *Hook {
	if r == nil {
		return nil
	}
	return r.Hook
}

// hookScheduleL is where Load methods for each relationship are stored.
type hookScheduleL struct{}

var (
	hookScheduleAllColumns            = []string{"hook_schedule_id", "hook_id", "display_name", "description", "schedule_interval_unit", "schedule_time_month", "schedule_time_day", "schedule_time_hour", "schedule_time_minute", "schedule_time_second"}
	hookScheduleColumnsWithoutDefault = []string{"hook_schedule_id", "hook_id", "display_name", "description", "schedule_interval_unit", "schedule_time_month", "schedule_time_day", "schedule_time_hour", "schedule_time_minute", "schedule_time_second"}
	hookScheduleColumnsWithDefault    = []string{}
	hookSchedulePrimaryKeyColumns     = []string{"hook_schedule_id"}
	hookScheduleGeneratedColumns      = []string{}
)

type (
	// HookScheduleSlice is an alias for a slice of pointers to HookSchedule.
	// This should almost always be used instead of []HookSchedule.
	HookScheduleSlice []*HookSchedule
	// HookScheduleHook is the signature for custom HookSchedule hook methods
	HookScheduleHook func(context.Context, boil.ContextExecutor, *HookSchedule) error

	hookScheduleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	hookScheduleType                 = reflect.TypeOf(&HookSchedule{})
	hookScheduleMapping              = queries.MakeStructMapping(hookScheduleType)
	hookSchedulePrimaryKeyMapping, _ = queries.BindMapping(hookScheduleType, hookScheduleMapping, hookSchedulePrimaryKeyColumns)
	hookScheduleInsertCacheMut       sync.RWMutex
	hookScheduleInsertCache          = make(map[string]insertCache)
	hookScheduleUpdateCacheMut       sync.RWMutex
	hookScheduleUpdateCache          = make(map[string]updateCache)
	hookScheduleUpsertCacheMut       sync.RWMutex
	hookScheduleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var hookScheduleAfterSelectMu sync.Mutex
var hookScheduleAfterSelectHooks []HookScheduleHook

var hookScheduleBeforeInsertMu sync.Mutex
var hookScheduleBeforeInsertHooks []HookScheduleHook
var hookScheduleAfterInsertMu sync.Mutex
var hookScheduleAfterInsertHooks []HookScheduleHook

var hookScheduleBeforeUpdateMu sync.Mutex
var hookScheduleBeforeUpdateHooks []HookScheduleHook
var hookScheduleAfterUpdateMu sync.Mutex
var hookScheduleAfterUpdateHooks []HookScheduleHook

var hookScheduleBeforeDeleteMu sync.Mutex
var hookScheduleBeforeDeleteHooks []HookScheduleHook
var hookScheduleAfterDeleteMu sync.Mutex
var hookScheduleAfterDeleteHooks []HookScheduleHook

var hookScheduleBeforeUpsertMu sync.Mutex
var hookScheduleBeforeUpsertHooks []HookScheduleHook
var hookScheduleAfterUpsertMu sync.Mutex
var hookScheduleAfterUpsertHooks []HookScheduleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HookSchedule) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HookSchedule) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HookSchedule) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HookSchedule) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HookSchedule) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HookSchedule) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HookSchedule) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HookSchedule) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HookSchedule) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookScheduleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHookScheduleHook registers your hook function for all future operations.
func AddHookScheduleHook(hookPoint boil.HookPoint, hookScheduleHook HookScheduleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		hookScheduleAfterSelectMu.Lock()
		hookScheduleAfterSelectHooks = append(hookScheduleAfterSelectHooks, hookScheduleHook)
		hookScheduleAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		hookScheduleBeforeInsertMu.Lock()
		hookScheduleBeforeInsertHooks = append(hookScheduleBeforeInsertHooks, hookScheduleHook)
		hookScheduleBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		hookScheduleAfterInsertMu.Lock()
		hookScheduleAfterInsertHooks = append(hookScheduleAfterInsertHooks, hookScheduleHook)
		hookScheduleAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		hookScheduleBeforeUpdateMu.Lock()
		hookScheduleBeforeUpdateHooks = append(hookScheduleBeforeUpdateHooks, hookScheduleHook)
		hookScheduleBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		hookScheduleAfterUpdateMu.Lock()
		hookScheduleAfterUpdateHooks = append(hookScheduleAfterUpdateHooks, hookScheduleHook)
		hookScheduleAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		hookScheduleBeforeDeleteMu.Lock()
		hookScheduleBeforeDeleteHooks = append(hookScheduleBeforeDeleteHooks, hookScheduleHook)
		hookScheduleBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		hookScheduleAfterDeleteMu.Lock()
		hookScheduleAfterDeleteHooks = append(hookScheduleAfterDeleteHooks, hookScheduleHook)
		hookScheduleAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		hookScheduleBeforeUpsertMu.Lock()
		hookScheduleBeforeUpsertHooks = append(hookScheduleBeforeUpsertHooks, hookScheduleHook)
		hookScheduleBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		hookScheduleAfterUpsertMu.Lock()
		hookScheduleAfterUpsertHooks = append(hookScheduleAfterUpsertHooks, hookScheduleHook)
		hookScheduleAfterUpsertMu.Unlock()
	}
}

// One returns a single hookSchedule record from the query.
func (q hookScheduleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HookSchedule, error) {
	o := &HookSchedule{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for hook_schedules")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HookSchedule records from the query.
func (q hookScheduleQuery) All(ctx context.Context, exec boil.ContextExecutor) (HookScheduleSlice, error) {
	var o []*HookSchedule

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to HookSchedule slice")
	}

	if len(hookScheduleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HookSchedule records in the query.
func (q hookScheduleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count hook_schedules rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q hookScheduleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if hook_schedules exists")
	}

	return count > 0, nil
}

// Hook pointed to by the foreign key.
func (o *HookSchedule) Hook(mods ...qm.QueryMod) hookQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"hook_id\" = ?", o.HookID),
	}

	queryMods = append(queryMods, mods...)

	return Hooks(queryMods...)
}

// LoadHook allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (hookScheduleL) LoadHook(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHookSchedule interface{}, mods queries.Applicator) error {
	var slice []*HookSchedule
	var object *HookSchedule

	if singular {
		var ok bool
		object, ok = maybeHookSchedule.(*HookSchedule)
		if !ok {
			object = new(HookSchedule)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeHookSchedule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeHookSchedule))
			}
		}
	} else {
		s, ok := maybeHookSchedule.(*[]*HookSchedule)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeHookSchedule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeHookSchedule))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &hookScheduleR{}
		}
		if !queries.IsNil(object.HookID) {
			args[object.HookID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &hookScheduleR{}
			}

			if !queries.IsNil(obj.HookID) {
				args[obj.HookID] = struct{}{}
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`hooks`),
		qm.WhereIn(`hooks.hook_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Hook")
	}

	var resultSlice []*Hook
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Hook")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for hooks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for hooks")
	}

	if len(hookAfterSelectHooks) != 0 {
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
		object.R.Hook = foreign
		if foreign.R == nil {
			foreign.R = &hookR{}
		}
		foreign.R.HookSchedules = append(foreign.R.HookSchedules, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.HookID, foreign.HookID) {
				local.R.Hook = foreign
				if foreign.R == nil {
					foreign.R = &hookR{}
				}
				foreign.R.HookSchedules = append(foreign.R.HookSchedules, local)
				break
			}
		}
	}

	return nil
}

// SetHook of the hookSchedule to the related item.
// Sets o.R.Hook to related.
// Adds o to related.R.HookSchedules.
func (o *HookSchedule) SetHook(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Hook) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"hook_schedules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"hook_id"}),
		strmangle.WhereClause("\"", "\"", 2, hookSchedulePrimaryKeyColumns),
	)
	values := []interface{}{related.HookID, o.HookScheduleID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.HookID, related.HookID)
	if o.R == nil {
		o.R = &hookScheduleR{
			Hook: related,
		}
	} else {
		o.R.Hook = related
	}

	if related.R == nil {
		related.R = &hookR{
			HookSchedules: HookScheduleSlice{o},
		}
	} else {
		related.R.HookSchedules = append(related.R.HookSchedules, o)
	}

	return nil
}

// HookSchedules retrieves all the records using an executor.
func HookSchedules(mods ...qm.QueryMod) hookScheduleQuery {
	mods = append(mods, qm.From("\"hook_schedules\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"hook_schedules\".*"})
	}

	return hookScheduleQuery{q}
}

// FindHookSchedule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHookSchedule(ctx context.Context, exec boil.ContextExecutor, hookScheduleID uuid.UUID, selectCols ...string) (*HookSchedule, error) {
	hookScheduleObj := &HookSchedule{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"hook_schedules\" where \"hook_schedule_id\"=$1", sel,
	)

	q := queries.Raw(query, hookScheduleID)

	err := q.Bind(ctx, exec, hookScheduleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from hook_schedules")
	}

	if err = hookScheduleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return hookScheduleObj, err
	}

	return hookScheduleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HookSchedule) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no hook_schedules provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hookScheduleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	hookScheduleInsertCacheMut.RLock()
	cache, cached := hookScheduleInsertCache[key]
	hookScheduleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			hookScheduleAllColumns,
			hookScheduleColumnsWithDefault,
			hookScheduleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(hookScheduleType, hookScheduleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(hookScheduleType, hookScheduleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"hook_schedules\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"hook_schedules\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into hook_schedules")
	}

	if !cached {
		hookScheduleInsertCacheMut.Lock()
		hookScheduleInsertCache[key] = cache
		hookScheduleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HookSchedule.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HookSchedule) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	hookScheduleUpdateCacheMut.RLock()
	cache, cached := hookScheduleUpdateCache[key]
	hookScheduleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			hookScheduleAllColumns,
			hookSchedulePrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update hook_schedules, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"hook_schedules\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, hookSchedulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(hookScheduleType, hookScheduleMapping, append(wl, hookSchedulePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update hook_schedules row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for hook_schedules")
	}

	if !cached {
		hookScheduleUpdateCacheMut.Lock()
		hookScheduleUpdateCache[key] = cache
		hookScheduleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q hookScheduleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for hook_schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for hook_schedules")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HookScheduleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hookSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"hook_schedules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, hookSchedulePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in hookSchedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all hookSchedule")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HookSchedule) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no hook_schedules provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hookScheduleColumnsWithDefault, o)

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

	hookScheduleUpsertCacheMut.RLock()
	cache, cached := hookScheduleUpsertCache[key]
	hookScheduleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			hookScheduleAllColumns,
			hookScheduleColumnsWithDefault,
			hookScheduleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			hookScheduleAllColumns,
			hookSchedulePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert hook_schedules, could not build update column list")
		}

		ret := strmangle.SetComplement(hookScheduleAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(hookSchedulePrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert hook_schedules, could not build conflict column list")
			}

			conflict = make([]string, len(hookSchedulePrimaryKeyColumns))
			copy(conflict, hookSchedulePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"hook_schedules\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(hookScheduleType, hookScheduleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(hookScheduleType, hookScheduleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert hook_schedules")
	}

	if !cached {
		hookScheduleUpsertCacheMut.Lock()
		hookScheduleUpsertCache[key] = cache
		hookScheduleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HookSchedule record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HookSchedule) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no HookSchedule provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), hookSchedulePrimaryKeyMapping)
	sql := "DELETE FROM \"hook_schedules\" WHERE \"hook_schedule_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from hook_schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for hook_schedules")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q hookScheduleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no hookScheduleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hook_schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hook_schedules")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HookScheduleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(hookScheduleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hookSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"hook_schedules\" WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 1, hookSchedulePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hookSchedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hook_schedules")
	}

	if len(hookScheduleAfterDeleteHooks) != 0 {
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
func (o *HookSchedule) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHookSchedule(ctx, exec, o.HookScheduleID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HookScheduleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HookScheduleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hookSchedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"hook_schedules\".* FROM \"hook_schedules\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hookSchedulePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HookScheduleSlice")
	}

	*o = slice

	return nil
}

// HookScheduleExists checks if the HookSchedule row exists.
func HookScheduleExists(ctx context.Context, exec boil.ContextExecutor, hookScheduleID uuid.UUID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"hook_schedules\" where \"hook_schedule_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, hookScheduleID)
	}
	row := exec.QueryRowContext(ctx, sql, hookScheduleID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if hook_schedules exists")
	}

	return exists, nil
}

// Exists checks if the HookSchedule row exists.
func (o *HookSchedule) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return HookScheduleExists(ctx, exec, o.HookScheduleID)
}
