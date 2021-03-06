// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/washanhanzi/ent-same-package-name/ent/predicate"
	"github.com/washanhanzi/ent-same-package-name/ent/tableb"
)

// TableBQuery is the builder for querying TableB entities.
type TableBQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TableB
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TableBQuery builder.
func (tb *TableBQuery) Where(ps ...predicate.TableB) *TableBQuery {
	tb.predicates = append(tb.predicates, ps...)
	return tb
}

// Limit adds a limit step to the query.
func (tb *TableBQuery) Limit(limit int) *TableBQuery {
	tb.limit = &limit
	return tb
}

// Offset adds an offset step to the query.
func (tb *TableBQuery) Offset(offset int) *TableBQuery {
	tb.offset = &offset
	return tb
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tb *TableBQuery) Unique(unique bool) *TableBQuery {
	tb.unique = &unique
	return tb
}

// Order adds an order step to the query.
func (tb *TableBQuery) Order(o ...OrderFunc) *TableBQuery {
	tb.order = append(tb.order, o...)
	return tb
}

// First returns the first TableB entity from the query.
// Returns a *NotFoundError when no TableB was found.
func (tb *TableBQuery) First(ctx context.Context) (*TableB, error) {
	nodes, err := tb.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tableb.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tb *TableBQuery) FirstX(ctx context.Context) *TableB {
	node, err := tb.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TableB ID from the query.
// Returns a *NotFoundError when no TableB ID was found.
func (tb *TableBQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tb.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tableb.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tb *TableBQuery) FirstIDX(ctx context.Context) int {
	id, err := tb.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TableB entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TableB entity is not found.
// Returns a *NotFoundError when no TableB entities are found.
func (tb *TableBQuery) Only(ctx context.Context) (*TableB, error) {
	nodes, err := tb.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tableb.Label}
	default:
		return nil, &NotSingularError{tableb.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tb *TableBQuery) OnlyX(ctx context.Context) *TableB {
	node, err := tb.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TableB ID in the query.
// Returns a *NotSingularError when exactly one TableB ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tb *TableBQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tb.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = &NotSingularError{tableb.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tb *TableBQuery) OnlyIDX(ctx context.Context) int {
	id, err := tb.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TableBs.
func (tb *TableBQuery) All(ctx context.Context) ([]*TableB, error) {
	if err := tb.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tb.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tb *TableBQuery) AllX(ctx context.Context) []*TableB {
	nodes, err := tb.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TableB IDs.
func (tb *TableBQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tb.Select(tableb.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tb *TableBQuery) IDsX(ctx context.Context) []int {
	ids, err := tb.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tb *TableBQuery) Count(ctx context.Context) (int, error) {
	if err := tb.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tb.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tb *TableBQuery) CountX(ctx context.Context) int {
	count, err := tb.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tb *TableBQuery) Exist(ctx context.Context) (bool, error) {
	if err := tb.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tb.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tb *TableBQuery) ExistX(ctx context.Context) bool {
	exist, err := tb.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TableBQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tb *TableBQuery) Clone() *TableBQuery {
	if tb == nil {
		return nil
	}
	return &TableBQuery{
		config:     tb.config,
		limit:      tb.limit,
		offset:     tb.offset,
		order:      append([]OrderFunc{}, tb.order...),
		predicates: append([]predicate.TableB{}, tb.predicates...),
		// clone intermediate query.
		sql:  tb.sql.Clone(),
		path: tb.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		B entity.B `json:"b,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TableB.Query().
//		GroupBy(tableb.FieldB).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tb *TableBQuery) GroupBy(field string, fields ...string) *TableBGroupBy {
	group := &TableBGroupBy{config: tb.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tb.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tb.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		B entity.B `json:"b,omitempty"`
//	}
//
//	client.TableB.Query().
//		Select(tableb.FieldB).
//		Scan(ctx, &v)
//
func (tb *TableBQuery) Select(field string, fields ...string) *TableBSelect {
	tb.fields = append([]string{field}, fields...)
	return &TableBSelect{TableBQuery: tb}
}

func (tb *TableBQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tb.fields {
		if !tableb.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tb.path != nil {
		prev, err := tb.path(ctx)
		if err != nil {
			return err
		}
		tb.sql = prev
	}
	return nil
}

func (tb *TableBQuery) sqlAll(ctx context.Context) ([]*TableB, error) {
	var (
		nodes = []*TableB{}
		_spec = tb.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TableB{config: tb.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tb.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tb *TableBQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tb.querySpec()
	return sqlgraph.CountNodes(ctx, tb.driver, _spec)
}

func (tb *TableBQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tb.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tb *TableBQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tableb.Table,
			Columns: tableb.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tableb.FieldID,
			},
		},
		From:   tb.sql,
		Unique: true,
	}
	if unique := tb.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tb.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tableb.FieldID)
		for i := range fields {
			if fields[i] != tableb.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tb.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tb.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tb.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tb.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tb *TableBQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tb.driver.Dialect())
	t1 := builder.Table(tableb.Table)
	selector := builder.Select(t1.Columns(tableb.Columns...)...).From(t1)
	if tb.sql != nil {
		selector = tb.sql
		selector.Select(selector.Columns(tableb.Columns...)...)
	}
	for _, p := range tb.predicates {
		p(selector)
	}
	for _, p := range tb.order {
		p(selector)
	}
	if offset := tb.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tb.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TableBGroupBy is the group-by builder for TableB entities.
type TableBGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tbb *TableBGroupBy) Aggregate(fns ...AggregateFunc) *TableBGroupBy {
	tbb.fns = append(tbb.fns, fns...)
	return tbb
}

// Scan applies the group-by query and scans the result into the given value.
func (tbb *TableBGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tbb.path(ctx)
	if err != nil {
		return err
	}
	tbb.sql = query
	return tbb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tbb *TableBGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tbb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tbb.fields) > 1 {
		return nil, errors.New("ent: TableBGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tbb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tbb *TableBGroupBy) StringsX(ctx context.Context) []string {
	v, err := tbb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tbb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tbb *TableBGroupBy) StringX(ctx context.Context) string {
	v, err := tbb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tbb.fields) > 1 {
		return nil, errors.New("ent: TableBGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tbb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tbb *TableBGroupBy) IntsX(ctx context.Context) []int {
	v, err := tbb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tbb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tbb *TableBGroupBy) IntX(ctx context.Context) int {
	v, err := tbb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tbb.fields) > 1 {
		return nil, errors.New("ent: TableBGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tbb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tbb *TableBGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tbb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tbb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tbb *TableBGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tbb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tbb.fields) > 1 {
		return nil, errors.New("ent: TableBGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tbb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tbb *TableBGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tbb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tbb *TableBGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tbb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tbb *TableBGroupBy) BoolX(ctx context.Context) bool {
	v, err := tbb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tbb *TableBGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tbb.fields {
		if !tableb.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tbb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tbb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tbb *TableBGroupBy) sqlQuery() *sql.Selector {
	selector := tbb.sql
	columns := make([]string, 0, len(tbb.fields)+len(tbb.fns))
	columns = append(columns, tbb.fields...)
	for _, fn := range tbb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tbb.fields...)
}

// TableBSelect is the builder for selecting fields of TableB entities.
type TableBSelect struct {
	*TableBQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tb *TableBSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tb.prepareQuery(ctx); err != nil {
		return err
	}
	tb.sql = tb.TableBQuery.sqlQuery(ctx)
	return tb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tb *TableBSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tb.fields) > 1 {
		return nil, errors.New("ent: TableBSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tb *TableBSelect) StringsX(ctx context.Context) []string {
	v, err := tb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tb *TableBSelect) StringX(ctx context.Context) string {
	v, err := tb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tb.fields) > 1 {
		return nil, errors.New("ent: TableBSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tb *TableBSelect) IntsX(ctx context.Context) []int {
	v, err := tb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tb *TableBSelect) IntX(ctx context.Context) int {
	v, err := tb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tb.fields) > 1 {
		return nil, errors.New("ent: TableBSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tb *TableBSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tb *TableBSelect) Float64X(ctx context.Context) float64 {
	v, err := tb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tb.fields) > 1 {
		return nil, errors.New("ent: TableBSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tb *TableBSelect) BoolsX(ctx context.Context) []bool {
	v, err := tb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tb *TableBSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tableb.Label}
	default:
		err = fmt.Errorf("ent: TableBSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tb *TableBSelect) BoolX(ctx context.Context) bool {
	v, err := tb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tb *TableBSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tb.sqlQuery().Query()
	if err := tb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tb *TableBSelect) sqlQuery() sql.Querier {
	selector := tb.sql
	selector.Select(selector.Columns(tb.fields...)...)
	return selector
}
