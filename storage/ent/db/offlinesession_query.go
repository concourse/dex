// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/offlinesession"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// OfflineSessionQuery is the builder for querying OfflineSession entities.
type OfflineSessionQuery struct {
	config
	ctx        *QueryContext
	order      []offlinesession.OrderOption
	inters     []Interceptor
	predicates []predicate.OfflineSession
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OfflineSessionQuery builder.
func (osq *OfflineSessionQuery) Where(ps ...predicate.OfflineSession) *OfflineSessionQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit the number of records to be returned by this query.
func (osq *OfflineSessionQuery) Limit(limit int) *OfflineSessionQuery {
	osq.ctx.Limit = &limit
	return osq
}

// Offset to start from.
func (osq *OfflineSessionQuery) Offset(offset int) *OfflineSessionQuery {
	osq.ctx.Offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OfflineSessionQuery) Unique(unique bool) *OfflineSessionQuery {
	osq.ctx.Unique = &unique
	return osq
}

// Order specifies how the records should be ordered.
func (osq *OfflineSessionQuery) Order(o ...offlinesession.OrderOption) *OfflineSessionQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// First returns the first OfflineSession entity from the query.
// Returns a *NotFoundError when no OfflineSession was found.
func (osq *OfflineSessionQuery) First(ctx context.Context) (*OfflineSession, error) {
	nodes, err := osq.Limit(1).All(setContextOp(ctx, osq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{offlinesession.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OfflineSessionQuery) FirstX(ctx context.Context) *OfflineSession {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OfflineSession ID from the query.
// Returns a *NotFoundError when no OfflineSession ID was found.
func (osq *OfflineSessionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = osq.Limit(1).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{offlinesession.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OfflineSessionQuery) FirstIDX(ctx context.Context) string {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OfflineSession entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OfflineSession entity is found.
// Returns a *NotFoundError when no OfflineSession entities are found.
func (osq *OfflineSessionQuery) Only(ctx context.Context) (*OfflineSession, error) {
	nodes, err := osq.Limit(2).All(setContextOp(ctx, osq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{offlinesession.Label}
	default:
		return nil, &NotSingularError{offlinesession.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OfflineSessionQuery) OnlyX(ctx context.Context) *OfflineSession {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OfflineSession ID in the query.
// Returns a *NotSingularError when more than one OfflineSession ID is found.
// Returns a *NotFoundError when no entities are found.
func (osq *OfflineSessionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = osq.Limit(2).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{offlinesession.Label}
	default:
		err = &NotSingularError{offlinesession.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OfflineSessionQuery) OnlyIDX(ctx context.Context) string {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OfflineSessions.
func (osq *OfflineSessionQuery) All(ctx context.Context) ([]*OfflineSession, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryAll)
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OfflineSession, *OfflineSessionQuery]()
	return withInterceptors[[]*OfflineSession](ctx, osq, qr, osq.inters)
}

// AllX is like All, but panics if an error occurs.
func (osq *OfflineSessionQuery) AllX(ctx context.Context) []*OfflineSession {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OfflineSession IDs.
func (osq *OfflineSessionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if osq.ctx.Unique == nil && osq.path != nil {
		osq.Unique(true)
	}
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryIDs)
	if err = osq.Select(offlinesession.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OfflineSessionQuery) IDsX(ctx context.Context) []string {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OfflineSessionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryCount)
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, osq, querierCount[*OfflineSessionQuery](), osq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OfflineSessionQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OfflineSessionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryExist)
	switch _, err := osq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OfflineSessionQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OfflineSessionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OfflineSessionQuery) Clone() *OfflineSessionQuery {
	if osq == nil {
		return nil
	}
	return &OfflineSessionQuery{
		config:     osq.config,
		ctx:        osq.ctx.Clone(),
		order:      append([]offlinesession.OrderOption{}, osq.order...),
		inters:     append([]Interceptor{}, osq.inters...),
		predicates: append([]predicate.OfflineSession{}, osq.predicates...),
		// clone intermediate query.
		sql:  osq.sql.Clone(),
		path: osq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OfflineSession.Query().
//		GroupBy(offlinesession.FieldUserID).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (osq *OfflineSessionQuery) GroupBy(field string, fields ...string) *OfflineSessionGroupBy {
	osq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OfflineSessionGroupBy{build: osq}
	grbuild.flds = &osq.ctx.Fields
	grbuild.label = offlinesession.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.OfflineSession.Query().
//		Select(offlinesession.FieldUserID).
//		Scan(ctx, &v)
func (osq *OfflineSessionQuery) Select(fields ...string) *OfflineSessionSelect {
	osq.ctx.Fields = append(osq.ctx.Fields, fields...)
	sbuild := &OfflineSessionSelect{OfflineSessionQuery: osq}
	sbuild.label = offlinesession.Label
	sbuild.flds, sbuild.scan = &osq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OfflineSessionSelect configured with the given aggregations.
func (osq *OfflineSessionQuery) Aggregate(fns ...AggregateFunc) *OfflineSessionSelect {
	return osq.Select().Aggregate(fns...)
}

func (osq *OfflineSessionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range osq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, osq); err != nil {
				return err
			}
		}
	}
	for _, f := range osq.ctx.Fields {
		if !offlinesession.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	return nil
}

func (osq *OfflineSessionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OfflineSession, error) {
	var (
		nodes = []*OfflineSession{}
		_spec = osq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OfflineSession).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OfflineSession{config: osq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (osq *OfflineSessionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	_spec.Node.Columns = osq.ctx.Fields
	if len(osq.ctx.Fields) > 0 {
		_spec.Unique = osq.ctx.Unique != nil && *osq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OfflineSessionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(offlinesession.Table, offlinesession.Columns, sqlgraph.NewFieldSpec(offlinesession.FieldID, field.TypeString))
	_spec.From = osq.sql
	if unique := osq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if osq.path != nil {
		_spec.Unique = true
	}
	if fields := osq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, offlinesession.FieldID)
		for i := range fields {
			if fields[i] != offlinesession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OfflineSessionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(offlinesession.Table)
	columns := osq.ctx.Fields
	if len(columns) == 0 {
		columns = offlinesession.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osq.ctx.Unique != nil && *osq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OfflineSessionGroupBy is the group-by builder for OfflineSession entities.
type OfflineSessionGroupBy struct {
	selector
	build *OfflineSessionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OfflineSessionGroupBy) Aggregate(fns ...AggregateFunc) *OfflineSessionGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the selector query and scans the result into the given value.
func (osgb *OfflineSessionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, osgb.build.ctx, ent.OpQueryGroupBy)
	if err := osgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OfflineSessionQuery, *OfflineSessionGroupBy](ctx, osgb.build, osgb, osgb.build.inters, v)
}

func (osgb *OfflineSessionGroupBy) sqlScan(ctx context.Context, root *OfflineSessionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*osgb.flds)+len(osgb.fns))
		for _, f := range *osgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*osgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OfflineSessionSelect is the builder for selecting fields of OfflineSession entities.
type OfflineSessionSelect struct {
	*OfflineSessionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oss *OfflineSessionSelect) Aggregate(fns ...AggregateFunc) *OfflineSessionSelect {
	oss.fns = append(oss.fns, fns...)
	return oss
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OfflineSessionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oss.ctx, ent.OpQuerySelect)
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OfflineSessionQuery, *OfflineSessionSelect](ctx, oss.OfflineSessionQuery, oss, oss.inters, v)
}

func (oss *OfflineSessionSelect) sqlScan(ctx context.Context, root *OfflineSessionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oss.fns))
	for _, fn := range oss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
