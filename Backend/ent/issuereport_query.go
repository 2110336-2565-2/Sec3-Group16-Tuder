// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/issuereport"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/google/uuid"
)

// IssueReportQuery is the builder for querying IssueReport entities.
type IssueReportQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.IssueReport
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IssueReportQuery builder.
func (irq *IssueReportQuery) Where(ps ...predicate.IssueReport) *IssueReportQuery {
	irq.predicates = append(irq.predicates, ps...)
	return irq
}

// Limit the number of records to be returned by this query.
func (irq *IssueReportQuery) Limit(limit int) *IssueReportQuery {
	irq.ctx.Limit = &limit
	return irq
}

// Offset to start from.
func (irq *IssueReportQuery) Offset(offset int) *IssueReportQuery {
	irq.ctx.Offset = &offset
	return irq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (irq *IssueReportQuery) Unique(unique bool) *IssueReportQuery {
	irq.ctx.Unique = &unique
	return irq
}

// Order specifies how the records should be ordered.
func (irq *IssueReportQuery) Order(o ...OrderFunc) *IssueReportQuery {
	irq.order = append(irq.order, o...)
	return irq
}

// First returns the first IssueReport entity from the query.
// Returns a *NotFoundError when no IssueReport was found.
func (irq *IssueReportQuery) First(ctx context.Context) (*IssueReport, error) {
	nodes, err := irq.Limit(1).All(setContextOp(ctx, irq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{issuereport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (irq *IssueReportQuery) FirstX(ctx context.Context) *IssueReport {
	node, err := irq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IssueReport ID from the query.
// Returns a *NotFoundError when no IssueReport ID was found.
func (irq *IssueReportQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = irq.Limit(1).IDs(setContextOp(ctx, irq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{issuereport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (irq *IssueReportQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := irq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IssueReport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IssueReport entity is found.
// Returns a *NotFoundError when no IssueReport entities are found.
func (irq *IssueReportQuery) Only(ctx context.Context) (*IssueReport, error) {
	nodes, err := irq.Limit(2).All(setContextOp(ctx, irq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{issuereport.Label}
	default:
		return nil, &NotSingularError{issuereport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (irq *IssueReportQuery) OnlyX(ctx context.Context) *IssueReport {
	node, err := irq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IssueReport ID in the query.
// Returns a *NotSingularError when more than one IssueReport ID is found.
// Returns a *NotFoundError when no entities are found.
func (irq *IssueReportQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = irq.Limit(2).IDs(setContextOp(ctx, irq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{issuereport.Label}
	default:
		err = &NotSingularError{issuereport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (irq *IssueReportQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := irq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IssueReports.
func (irq *IssueReportQuery) All(ctx context.Context) ([]*IssueReport, error) {
	ctx = setContextOp(ctx, irq.ctx, "All")
	if err := irq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*IssueReport, *IssueReportQuery]()
	return withInterceptors[[]*IssueReport](ctx, irq, qr, irq.inters)
}

// AllX is like All, but panics if an error occurs.
func (irq *IssueReportQuery) AllX(ctx context.Context) []*IssueReport {
	nodes, err := irq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IssueReport IDs.
func (irq *IssueReportQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if irq.ctx.Unique == nil && irq.path != nil {
		irq.Unique(true)
	}
	ctx = setContextOp(ctx, irq.ctx, "IDs")
	if err = irq.Select(issuereport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (irq *IssueReportQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := irq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (irq *IssueReportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, irq.ctx, "Count")
	if err := irq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, irq, querierCount[*IssueReportQuery](), irq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (irq *IssueReportQuery) CountX(ctx context.Context) int {
	count, err := irq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (irq *IssueReportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, irq.ctx, "Exist")
	switch _, err := irq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (irq *IssueReportQuery) ExistX(ctx context.Context) bool {
	exist, err := irq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IssueReportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (irq *IssueReportQuery) Clone() *IssueReportQuery {
	if irq == nil {
		return nil
	}
	return &IssueReportQuery{
		config:     irq.config,
		ctx:        irq.ctx.Clone(),
		order:      append([]OrderFunc{}, irq.order...),
		inters:     append([]Interceptor{}, irq.inters...),
		predicates: append([]predicate.IssueReport{}, irq.predicates...),
		// clone intermediate query.
		sql:  irq.sql.Clone(),
		path: irq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IssueReport.Query().
//		GroupBy(issuereport.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (irq *IssueReportQuery) GroupBy(field string, fields ...string) *IssueReportGroupBy {
	irq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IssueReportGroupBy{build: irq}
	grbuild.flds = &irq.ctx.Fields
	grbuild.label = issuereport.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.IssueReport.Query().
//		Select(issuereport.FieldTitle).
//		Scan(ctx, &v)
func (irq *IssueReportQuery) Select(fields ...string) *IssueReportSelect {
	irq.ctx.Fields = append(irq.ctx.Fields, fields...)
	sbuild := &IssueReportSelect{IssueReportQuery: irq}
	sbuild.label = issuereport.Label
	sbuild.flds, sbuild.scan = &irq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IssueReportSelect configured with the given aggregations.
func (irq *IssueReportQuery) Aggregate(fns ...AggregateFunc) *IssueReportSelect {
	return irq.Select().Aggregate(fns...)
}

func (irq *IssueReportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range irq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, irq); err != nil {
				return err
			}
		}
	}
	for _, f := range irq.ctx.Fields {
		if !issuereport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if irq.path != nil {
		prev, err := irq.path(ctx)
		if err != nil {
			return err
		}
		irq.sql = prev
	}
	return nil
}

func (irq *IssueReportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IssueReport, error) {
	var (
		nodes   = []*IssueReport{}
		withFKs = irq.withFKs
		_spec   = irq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, issuereport.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*IssueReport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &IssueReport{config: irq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, irq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (irq *IssueReportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := irq.querySpec()
	_spec.Node.Columns = irq.ctx.Fields
	if len(irq.ctx.Fields) > 0 {
		_spec.Unique = irq.ctx.Unique != nil && *irq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, irq.driver, _spec)
}

func (irq *IssueReportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(issuereport.Table, issuereport.Columns, sqlgraph.NewFieldSpec(issuereport.FieldID, field.TypeUUID))
	_spec.From = irq.sql
	if unique := irq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if irq.path != nil {
		_spec.Unique = true
	}
	if fields := irq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, issuereport.FieldID)
		for i := range fields {
			if fields[i] != issuereport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := irq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := irq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := irq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := irq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (irq *IssueReportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(irq.driver.Dialect())
	t1 := builder.Table(issuereport.Table)
	columns := irq.ctx.Fields
	if len(columns) == 0 {
		columns = issuereport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if irq.sql != nil {
		selector = irq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if irq.ctx.Unique != nil && *irq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range irq.predicates {
		p(selector)
	}
	for _, p := range irq.order {
		p(selector)
	}
	if offset := irq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := irq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IssueReportGroupBy is the group-by builder for IssueReport entities.
type IssueReportGroupBy struct {
	selector
	build *IssueReportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (irgb *IssueReportGroupBy) Aggregate(fns ...AggregateFunc) *IssueReportGroupBy {
	irgb.fns = append(irgb.fns, fns...)
	return irgb
}

// Scan applies the selector query and scans the result into the given value.
func (irgb *IssueReportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, irgb.build.ctx, "GroupBy")
	if err := irgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IssueReportQuery, *IssueReportGroupBy](ctx, irgb.build, irgb, irgb.build.inters, v)
}

func (irgb *IssueReportGroupBy) sqlScan(ctx context.Context, root *IssueReportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(irgb.fns))
	for _, fn := range irgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*irgb.flds)+len(irgb.fns))
		for _, f := range *irgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*irgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := irgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IssueReportSelect is the builder for selecting fields of IssueReport entities.
type IssueReportSelect struct {
	*IssueReportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (irs *IssueReportSelect) Aggregate(fns ...AggregateFunc) *IssueReportSelect {
	irs.fns = append(irs.fns, fns...)
	return irs
}

// Scan applies the selector query and scans the result into the given value.
func (irs *IssueReportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, irs.ctx, "Select")
	if err := irs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IssueReportQuery, *IssueReportSelect](ctx, irs.IssueReportQuery, irs, irs.inters, v)
}

func (irs *IssueReportSelect) sqlScan(ctx context.Context, root *IssueReportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(irs.fns))
	for _, fn := range irs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*irs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := irs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
