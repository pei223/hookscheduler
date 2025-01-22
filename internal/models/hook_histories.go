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
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// HookHistory is an object representing the database table.
type HookHistory struct {
	HookHistoryID    uuid.UUID   `boil:"hook_history_id" json:"hook_history_id" toml:"hook_history_id" yaml:"hook_history_id"`
	HookID           uuid.UUID   `boil:"hook_id" json:"hook_id" toml:"hook_id" yaml:"hook_id"`
	HookScheduleID   uuid.UUID   `boil:"hook_schedule_id" json:"hook_schedule_id" toml:"hook_schedule_id" yaml:"hook_schedule_id"`
	Status           int16       `boil:"status" json:"status" toml:"status" yaml:"status"`
	StartedAt        time.Time   `boil:"started_at" json:"started_at" toml:"started_at" yaml:"started_at"`
	UpdatedAt        time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	EndedAt          null.Time   `boil:"ended_at" json:"ended_at,omitempty" toml:"ended_at" yaml:"ended_at,omitempty"`
	HookSnapshot     types.JSONB `boil:"hook_snapshot" json:"hook_snapshot,omitempty" toml:"hook_snapshot" yaml:"hook_snapshot,omitempty"`
	ScheduleSnapshot types.JSONB `boil:"schedule_snapshot" json:"schedule_snapshot,omitempty" toml:"schedule_snapshot" yaml:"schedule_snapshot,omitempty"`

	R *hookHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L hookHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HookHistoryColumns = struct {
	HookHistoryID    string
	HookID           string
	HookScheduleID   string
	Status           string
	StartedAt        string
	UpdatedAt        string
	EndedAt          string
	HookSnapshot     string
	ScheduleSnapshot string
}{
	HookHistoryID:    "hook_history_id",
	HookID:           "hook_id",
	HookScheduleID:   "hook_schedule_id",
	Status:           "status",
	StartedAt:        "started_at",
	UpdatedAt:        "updated_at",
	EndedAt:          "ended_at",
	HookSnapshot:     "hook_snapshot",
	ScheduleSnapshot: "schedule_snapshot",
}

var HookHistoryTableColumns = struct {
	HookHistoryID    string
	HookID           string
	HookScheduleID   string
	Status           string
	StartedAt        string
	UpdatedAt        string
	EndedAt          string
	HookSnapshot     string
	ScheduleSnapshot string
}{
	HookHistoryID:    "hook_histories.hook_history_id",
	HookID:           "hook_histories.hook_id",
	HookScheduleID:   "hook_histories.hook_schedule_id",
	Status:           "hook_histories.status",
	StartedAt:        "hook_histories.started_at",
	UpdatedAt:        "hook_histories.updated_at",
	EndedAt:          "hook_histories.ended_at",
	HookSnapshot:     "hook_histories.hook_snapshot",
	ScheduleSnapshot: "hook_histories.schedule_snapshot",
}

// Generated where

type whereHelperuuid_UUID struct{ field string }

func (w whereHelperuuid_UUID) EQ(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperuuid_UUID) NEQ(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperuuid_UUID) LT(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperuuid_UUID) LTE(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperuuid_UUID) GT(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperuuid_UUID) GTE(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint16) IN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint16) NIN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertypes_JSONB struct{ field string }

func (w whereHelpertypes_JSONB) EQ(x types.JSONB) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_JSONB) NEQ(x types.JSONB) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_JSONB) LT(x types.JSONB) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSONB) LTE(x types.JSONB) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSONB) GT(x types.JSONB) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSONB) GTE(x types.JSONB) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpertypes_JSONB) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_JSONB) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var HookHistoryWhere = struct {
	HookHistoryID    whereHelperuuid_UUID
	HookID           whereHelperuuid_UUID
	HookScheduleID   whereHelperuuid_UUID
	Status           whereHelperint16
	StartedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
	EndedAt          whereHelpernull_Time
	HookSnapshot     whereHelpertypes_JSONB
	ScheduleSnapshot whereHelpertypes_JSONB
}{
	HookHistoryID:    whereHelperuuid_UUID{field: "\"hook_histories\".\"hook_history_id\""},
	HookID:           whereHelperuuid_UUID{field: "\"hook_histories\".\"hook_id\""},
	HookScheduleID:   whereHelperuuid_UUID{field: "\"hook_histories\".\"hook_schedule_id\""},
	Status:           whereHelperint16{field: "\"hook_histories\".\"status\""},
	StartedAt:        whereHelpertime_Time{field: "\"hook_histories\".\"started_at\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"hook_histories\".\"updated_at\""},
	EndedAt:          whereHelpernull_Time{field: "\"hook_histories\".\"ended_at\""},
	HookSnapshot:     whereHelpertypes_JSONB{field: "\"hook_histories\".\"hook_snapshot\""},
	ScheduleSnapshot: whereHelpertypes_JSONB{field: "\"hook_histories\".\"schedule_snapshot\""},
}

// HookHistoryRels is where relationship names are stored.
var HookHistoryRels = struct {
}{}

// hookHistoryR is where relationships are stored.
type hookHistoryR struct {
}

// NewStruct creates a new relationship struct
func (*hookHistoryR) NewStruct() *hookHistoryR {
	return &hookHistoryR{}
}

// hookHistoryL is where Load methods for each relationship are stored.
type hookHistoryL struct{}

var (
	hookHistoryAllColumns            = []string{"hook_history_id", "hook_id", "hook_schedule_id", "status", "started_at", "updated_at", "ended_at", "hook_snapshot", "schedule_snapshot"}
	hookHistoryColumnsWithoutDefault = []string{"hook_history_id", "hook_id", "hook_schedule_id", "status", "started_at", "updated_at"}
	hookHistoryColumnsWithDefault    = []string{"ended_at", "hook_snapshot", "schedule_snapshot"}
	hookHistoryPrimaryKeyColumns     = []string{"hook_history_id"}
	hookHistoryGeneratedColumns      = []string{}
)

type (
	// HookHistorySlice is an alias for a slice of pointers to HookHistory.
	// This should almost always be used instead of []HookHistory.
	HookHistorySlice []*HookHistory
	// HookHistoryHook is the signature for custom HookHistory hook methods
	HookHistoryHook func(context.Context, boil.ContextExecutor, *HookHistory) error

	hookHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	hookHistoryType                 = reflect.TypeOf(&HookHistory{})
	hookHistoryMapping              = queries.MakeStructMapping(hookHistoryType)
	hookHistoryPrimaryKeyMapping, _ = queries.BindMapping(hookHistoryType, hookHistoryMapping, hookHistoryPrimaryKeyColumns)
	hookHistoryInsertCacheMut       sync.RWMutex
	hookHistoryInsertCache          = make(map[string]insertCache)
	hookHistoryUpdateCacheMut       sync.RWMutex
	hookHistoryUpdateCache          = make(map[string]updateCache)
	hookHistoryUpsertCacheMut       sync.RWMutex
	hookHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var hookHistoryAfterSelectMu sync.Mutex
var hookHistoryAfterSelectHooks []HookHistoryHook

var hookHistoryBeforeInsertMu sync.Mutex
var hookHistoryBeforeInsertHooks []HookHistoryHook
var hookHistoryAfterInsertMu sync.Mutex
var hookHistoryAfterInsertHooks []HookHistoryHook

var hookHistoryBeforeUpdateMu sync.Mutex
var hookHistoryBeforeUpdateHooks []HookHistoryHook
var hookHistoryAfterUpdateMu sync.Mutex
var hookHistoryAfterUpdateHooks []HookHistoryHook

var hookHistoryBeforeDeleteMu sync.Mutex
var hookHistoryBeforeDeleteHooks []HookHistoryHook
var hookHistoryAfterDeleteMu sync.Mutex
var hookHistoryAfterDeleteHooks []HookHistoryHook

var hookHistoryBeforeUpsertMu sync.Mutex
var hookHistoryBeforeUpsertHooks []HookHistoryHook
var hookHistoryAfterUpsertMu sync.Mutex
var hookHistoryAfterUpsertHooks []HookHistoryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HookHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HookHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HookHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HookHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HookHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HookHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HookHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HookHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HookHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hookHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHookHistoryHook registers your hook function for all future operations.
func AddHookHistoryHook(hookPoint boil.HookPoint, hookHistoryHook HookHistoryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		hookHistoryAfterSelectMu.Lock()
		hookHistoryAfterSelectHooks = append(hookHistoryAfterSelectHooks, hookHistoryHook)
		hookHistoryAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		hookHistoryBeforeInsertMu.Lock()
		hookHistoryBeforeInsertHooks = append(hookHistoryBeforeInsertHooks, hookHistoryHook)
		hookHistoryBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		hookHistoryAfterInsertMu.Lock()
		hookHistoryAfterInsertHooks = append(hookHistoryAfterInsertHooks, hookHistoryHook)
		hookHistoryAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		hookHistoryBeforeUpdateMu.Lock()
		hookHistoryBeforeUpdateHooks = append(hookHistoryBeforeUpdateHooks, hookHistoryHook)
		hookHistoryBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		hookHistoryAfterUpdateMu.Lock()
		hookHistoryAfterUpdateHooks = append(hookHistoryAfterUpdateHooks, hookHistoryHook)
		hookHistoryAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		hookHistoryBeforeDeleteMu.Lock()
		hookHistoryBeforeDeleteHooks = append(hookHistoryBeforeDeleteHooks, hookHistoryHook)
		hookHistoryBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		hookHistoryAfterDeleteMu.Lock()
		hookHistoryAfterDeleteHooks = append(hookHistoryAfterDeleteHooks, hookHistoryHook)
		hookHistoryAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		hookHistoryBeforeUpsertMu.Lock()
		hookHistoryBeforeUpsertHooks = append(hookHistoryBeforeUpsertHooks, hookHistoryHook)
		hookHistoryBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		hookHistoryAfterUpsertMu.Lock()
		hookHistoryAfterUpsertHooks = append(hookHistoryAfterUpsertHooks, hookHistoryHook)
		hookHistoryAfterUpsertMu.Unlock()
	}
}

// One returns a single hookHistory record from the query.
func (q hookHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HookHistory, error) {
	o := &HookHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for hook_histories")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HookHistory records from the query.
func (q hookHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (HookHistorySlice, error) {
	var o []*HookHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to HookHistory slice")
	}

	if len(hookHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HookHistory records in the query.
func (q hookHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count hook_histories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q hookHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if hook_histories exists")
	}

	return count > 0, nil
}

// HookHistories retrieves all the records using an executor.
func HookHistories(mods ...qm.QueryMod) hookHistoryQuery {
	mods = append(mods, qm.From("\"hook_histories\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"hook_histories\".*"})
	}

	return hookHistoryQuery{q}
}

// FindHookHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHookHistory(ctx context.Context, exec boil.ContextExecutor, hookHistoryID uuid.UUID, selectCols ...string) (*HookHistory, error) {
	hookHistoryObj := &HookHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"hook_histories\" where \"hook_history_id\"=$1", sel,
	)

	q := queries.Raw(query, hookHistoryID)

	err := q.Bind(ctx, exec, hookHistoryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from hook_histories")
	}

	if err = hookHistoryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return hookHistoryObj, err
	}

	return hookHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HookHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no hook_histories provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hookHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	hookHistoryInsertCacheMut.RLock()
	cache, cached := hookHistoryInsertCache[key]
	hookHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			hookHistoryAllColumns,
			hookHistoryColumnsWithDefault,
			hookHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(hookHistoryType, hookHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(hookHistoryType, hookHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"hook_histories\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"hook_histories\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into hook_histories")
	}

	if !cached {
		hookHistoryInsertCacheMut.Lock()
		hookHistoryInsertCache[key] = cache
		hookHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HookHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HookHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	hookHistoryUpdateCacheMut.RLock()
	cache, cached := hookHistoryUpdateCache[key]
	hookHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			hookHistoryAllColumns,
			hookHistoryPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update hook_histories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"hook_histories\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, hookHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(hookHistoryType, hookHistoryMapping, append(wl, hookHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update hook_histories row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for hook_histories")
	}

	if !cached {
		hookHistoryUpdateCacheMut.Lock()
		hookHistoryUpdateCache[key] = cache
		hookHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q hookHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for hook_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for hook_histories")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HookHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hookHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"hook_histories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, hookHistoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in hookHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all hookHistory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HookHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no hook_histories provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hookHistoryColumnsWithDefault, o)

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

	hookHistoryUpsertCacheMut.RLock()
	cache, cached := hookHistoryUpsertCache[key]
	hookHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			hookHistoryAllColumns,
			hookHistoryColumnsWithDefault,
			hookHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			hookHistoryAllColumns,
			hookHistoryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert hook_histories, could not build update column list")
		}

		ret := strmangle.SetComplement(hookHistoryAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(hookHistoryPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert hook_histories, could not build conflict column list")
			}

			conflict = make([]string, len(hookHistoryPrimaryKeyColumns))
			copy(conflict, hookHistoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"hook_histories\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(hookHistoryType, hookHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(hookHistoryType, hookHistoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert hook_histories")
	}

	if !cached {
		hookHistoryUpsertCacheMut.Lock()
		hookHistoryUpsertCache[key] = cache
		hookHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HookHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HookHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no HookHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), hookHistoryPrimaryKeyMapping)
	sql := "DELETE FROM \"hook_histories\" WHERE \"hook_history_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from hook_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for hook_histories")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q hookHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no hookHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hook_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hook_histories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HookHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(hookHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hookHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"hook_histories\" WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 1, hookHistoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hookHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hook_histories")
	}

	if len(hookHistoryAfterDeleteHooks) != 0 {
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
func (o *HookHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHookHistory(ctx, exec, o.HookHistoryID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HookHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HookHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hookHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"hook_histories\".* FROM \"hook_histories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hookHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HookHistorySlice")
	}

	*o = slice

	return nil
}

// HookHistoryExists checks if the HookHistory row exists.
func HookHistoryExists(ctx context.Context, exec boil.ContextExecutor, hookHistoryID uuid.UUID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"hook_histories\" where \"hook_history_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, hookHistoryID)
	}
	row := exec.QueryRowContext(ctx, sql, hookHistoryID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if hook_histories exists")
	}

	return exists, nil
}

// Exists checks if the HookHistory row exists.
func (o *HookHistory) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return HookHistoryExists(ctx, exec, o.HookHistoryID)
}
