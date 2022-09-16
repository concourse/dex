// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/devicetoken"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// DeviceTokenQuery is the builder for querying DeviceToken entities.
type DeviceTokenQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeviceToken
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceTokenQuery builder.
func (dtq *DeviceTokenQuery) Where(ps ...predicate.DeviceToken) *DeviceTokenQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit adds a limit step to the query.
func (dtq *DeviceTokenQuery) Limit(limit int) *DeviceTokenQuery {
	dtq.limit = &limit
	return dtq
}

// Offset adds an offset step to the query.
func (dtq *DeviceTokenQuery) Offset(offset int) *DeviceTokenQuery {
	dtq.offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DeviceTokenQuery) Unique(unique bool) *DeviceTokenQuery {
	dtq.unique = &unique
	return dtq
}

// Order adds an order step to the query.
func (dtq *DeviceTokenQuery) Order(o ...OrderFunc) *DeviceTokenQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// First returns the first DeviceToken entity from the query.
// Returns a *NotFoundError when no DeviceToken was found.
func (dtq *DeviceTokenQuery) First(ctx context.Context) (*DeviceToken, error) {
	nodes, err := dtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{devicetoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DeviceTokenQuery) FirstX(ctx context.Context) *DeviceToken {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceToken ID from the query.
// Returns a *NotFoundError when no DeviceToken ID was found.
func (dtq *DeviceTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{devicetoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DeviceTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeviceToken entity is found.
// Returns a *NotFoundError when no DeviceToken entities are found.
func (dtq *DeviceTokenQuery) Only(ctx context.Context) (*DeviceToken, error) {
	nodes, err := dtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{devicetoken.Label}
	default:
		return nil, &NotSingularError{devicetoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DeviceTokenQuery) OnlyX(ctx context.Context) *DeviceToken {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceToken ID in the query.
// Returns a *NotSingularError when more than one DeviceToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DeviceTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{devicetoken.Label}
	default:
		err = &NotSingularError{devicetoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DeviceTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceTokens.
func (dtq *DeviceTokenQuery) All(ctx context.Context) ([]*DeviceToken, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DeviceTokenQuery) AllX(ctx context.Context) []*DeviceToken {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceToken IDs.
func (dtq *DeviceTokenQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dtq.Select(devicetoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DeviceTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DeviceTokenQuery) Count(ctx context.Context) (int, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DeviceTokenQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DeviceTokenQuery) Exist(ctx context.Context) (bool, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DeviceTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DeviceTokenQuery) Clone() *DeviceTokenQuery {
	if dtq == nil {
		return nil
	}
	return &DeviceTokenQuery{
		config:     dtq.config,
		limit:      dtq.limit,
		offset:     dtq.offset,
		order:      append([]OrderFunc{}, dtq.order...),
		predicates: append([]predicate.DeviceToken{}, dtq.predicates...),
		// clone intermediate query.
		sql:    dtq.sql.Clone(),
		path:   dtq.path,
		unique: dtq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeviceCode string `json:"device_code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeviceToken.Query().
//		GroupBy(devicetoken.FieldDeviceCode).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (dtq *DeviceTokenQuery) GroupBy(field string, fields ...string) *DeviceTokenGroupBy {
	grbuild := &DeviceTokenGroupBy{config: dtq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dtq.sqlQuery(ctx), nil
	}
	grbuild.label = devicetoken.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DeviceCode string `json:"device_code,omitempty"`
//	}
//
//	client.DeviceToken.Query().
//		Select(devicetoken.FieldDeviceCode).
//		Scan(ctx, &v)
//
func (dtq *DeviceTokenQuery) Select(fields ...string) *DeviceTokenSelect {
	dtq.fields = append(dtq.fields, fields...)
	selbuild := &DeviceTokenSelect{DeviceTokenQuery: dtq}
	selbuild.label = devicetoken.Label
	selbuild.flds, selbuild.scan = &dtq.fields, selbuild.Scan
	return selbuild
}

func (dtq *DeviceTokenQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dtq.fields {
		if !devicetoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DeviceTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DeviceToken, error) {
	var (
		nodes = []*DeviceToken{}
		_spec = dtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DeviceToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DeviceToken{config: dtq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dtq *DeviceTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	_spec.Node.Columns = dtq.fields
	if len(dtq.fields) > 0 {
		_spec.Unique = dtq.unique != nil && *dtq.unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DeviceTokenQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %w", err)
	}
	return n > 0, nil
}

func (dtq *DeviceTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicetoken.Table,
			Columns: devicetoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicetoken.FieldID,
			},
		},
		From:   dtq.sql,
		Unique: true,
	}
	if unique := dtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicetoken.FieldID)
		for i := range fields {
			if fields[i] != devicetoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DeviceTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(devicetoken.Table)
	columns := dtq.fields
	if len(columns) == 0 {
		columns = devicetoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.unique != nil && *dtq.unique {
		selector.Distinct()
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceTokenGroupBy is the group-by builder for DeviceToken entities.
type DeviceTokenGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DeviceTokenGroupBy) Aggregate(fns ...AggregateFunc) *DeviceTokenGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dtgb *DeviceTokenGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dtgb.path(ctx)
	if err != nil {
		return err
	}
	dtgb.sql = query
	return dtgb.sqlScan(ctx, v)
}

func (dtgb *DeviceTokenGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dtgb.fields {
		if !devicetoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dtgb *DeviceTokenGroupBy) sqlQuery() *sql.Selector {
	selector := dtgb.sql.Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dtgb.fields)+len(dtgb.fns))
		for _, f := range dtgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dtgb.fields...)...)
}

// DeviceTokenSelect is the builder for selecting fields of DeviceToken entities.
type DeviceTokenSelect struct {
	*DeviceTokenQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DeviceTokenSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	dts.sql = dts.DeviceTokenQuery.sqlQuery(ctx)
	return dts.sqlScan(ctx, v)
}

func (dts *DeviceTokenSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dts.sql.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
