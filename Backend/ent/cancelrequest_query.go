// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/cancelrequest"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// CancelRequestQuery is the builder for querying CancelRequest entities.
type CancelRequestQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.CancelRequest
	withMatch  *MatchQuery
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CancelRequestQuery builder.
func (crq *CancelRequestQuery) Where(ps ...predicate.CancelRequest) *CancelRequestQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit the number of records to be returned by this query.
func (crq *CancelRequestQuery) Limit(limit int) *CancelRequestQuery {
	crq.ctx.Limit = &limit
	return crq
}

// Offset to start from.
func (crq *CancelRequestQuery) Offset(offset int) *CancelRequestQuery {
	crq.ctx.Offset = &offset
	return crq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crq *CancelRequestQuery) Unique(unique bool) *CancelRequestQuery {
	crq.ctx.Unique = &unique
	return crq
}

// Order specifies how the records should be ordered.
func (crq *CancelRequestQuery) Order(o ...OrderFunc) *CancelRequestQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// QueryMatch chains the current query on the "match" edge.
func (crq *CancelRequestQuery) QueryMatch() *MatchQuery {
	query := (&MatchClient{config: crq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cancelrequest.Table, cancelrequest.FieldID, selector),
			sqlgraph.To(match.Table, match.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cancelrequest.MatchTable, cancelrequest.MatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (crq *CancelRequestQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: crq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cancelrequest.Table, cancelrequest.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cancelrequest.UserTable, cancelrequest.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CancelRequest entity from the query.
// Returns a *NotFoundError when no CancelRequest was found.
func (crq *CancelRequestQuery) First(ctx context.Context) (*CancelRequest, error) {
	nodes, err := crq.Limit(1).All(setContextOp(ctx, crq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cancelrequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *CancelRequestQuery) FirstX(ctx context.Context) *CancelRequest {
	node, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CancelRequest ID from the query.
// Returns a *NotFoundError when no CancelRequest ID was found.
func (crq *CancelRequestQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = crq.Limit(1).IDs(setContextOp(ctx, crq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cancelrequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crq *CancelRequestQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CancelRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CancelRequest entity is found.
// Returns a *NotFoundError when no CancelRequest entities are found.
func (crq *CancelRequestQuery) Only(ctx context.Context) (*CancelRequest, error) {
	nodes, err := crq.Limit(2).All(setContextOp(ctx, crq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cancelrequest.Label}
	default:
		return nil, &NotSingularError{cancelrequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *CancelRequestQuery) OnlyX(ctx context.Context) *CancelRequest {
	node, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CancelRequest ID in the query.
// Returns a *NotSingularError when more than one CancelRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (crq *CancelRequestQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = crq.Limit(2).IDs(setContextOp(ctx, crq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cancelrequest.Label}
	default:
		err = &NotSingularError{cancelrequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *CancelRequestQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CancelRequests.
func (crq *CancelRequestQuery) All(ctx context.Context) ([]*CancelRequest, error) {
	ctx = setContextOp(ctx, crq.ctx, "All")
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CancelRequest, *CancelRequestQuery]()
	return withInterceptors[[]*CancelRequest](ctx, crq, qr, crq.inters)
}

// AllX is like All, but panics if an error occurs.
func (crq *CancelRequestQuery) AllX(ctx context.Context) []*CancelRequest {
	nodes, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CancelRequest IDs.
func (crq *CancelRequestQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if crq.ctx.Unique == nil && crq.path != nil {
		crq.Unique(true)
	}
	ctx = setContextOp(ctx, crq.ctx, "IDs")
	if err = crq.Select(cancelrequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *CancelRequestQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *CancelRequestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, crq.ctx, "Count")
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, crq, querierCount[*CancelRequestQuery](), crq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (crq *CancelRequestQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *CancelRequestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, crq.ctx, "Exist")
	switch _, err := crq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *CancelRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CancelRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *CancelRequestQuery) Clone() *CancelRequestQuery {
	if crq == nil {
		return nil
	}
	return &CancelRequestQuery{
		config:     crq.config,
		ctx:        crq.ctx.Clone(),
		order:      append([]OrderFunc{}, crq.order...),
		inters:     append([]Interceptor{}, crq.inters...),
		predicates: append([]predicate.CancelRequest{}, crq.predicates...),
		withMatch:  crq.withMatch.Clone(),
		withUser:   crq.withUser.Clone(),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

// WithMatch tells the query-builder to eager-load the nodes that are connected to
// the "match" edge. The optional arguments are used to configure the query builder of the edge.
func (crq *CancelRequestQuery) WithMatch(opts ...func(*MatchQuery)) *CancelRequestQuery {
	query := (&MatchClient{config: crq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crq.withMatch = query
	return crq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (crq *CancelRequestQuery) WithUser(opts ...func(*UserQuery)) *CancelRequestQuery {
	query := (&UserClient{config: crq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crq.withUser = query
	return crq
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
//	client.CancelRequest.Query().
//		GroupBy(cancelrequest.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crq *CancelRequestQuery) GroupBy(field string, fields ...string) *CancelRequestGroupBy {
	crq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CancelRequestGroupBy{build: crq}
	grbuild.flds = &crq.ctx.Fields
	grbuild.label = cancelrequest.Label
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
//	client.CancelRequest.Query().
//		Select(cancelrequest.FieldTitle).
//		Scan(ctx, &v)
func (crq *CancelRequestQuery) Select(fields ...string) *CancelRequestSelect {
	crq.ctx.Fields = append(crq.ctx.Fields, fields...)
	sbuild := &CancelRequestSelect{CancelRequestQuery: crq}
	sbuild.label = cancelrequest.Label
	sbuild.flds, sbuild.scan = &crq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CancelRequestSelect configured with the given aggregations.
func (crq *CancelRequestQuery) Aggregate(fns ...AggregateFunc) *CancelRequestSelect {
	return crq.Select().Aggregate(fns...)
}

func (crq *CancelRequestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range crq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, crq); err != nil {
				return err
			}
		}
	}
	for _, f := range crq.ctx.Fields {
		if !cancelrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *CancelRequestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CancelRequest, error) {
	var (
		nodes       = []*CancelRequest{}
		withFKs     = crq.withFKs
		_spec       = crq.querySpec()
		loadedTypes = [2]bool{
			crq.withMatch != nil,
			crq.withUser != nil,
		}
	)
	if crq.withMatch != nil || crq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cancelrequest.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CancelRequest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CancelRequest{config: crq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := crq.withMatch; query != nil {
		if err := crq.loadMatch(ctx, query, nodes, nil,
			func(n *CancelRequest, e *Match) { n.Edges.Match = e }); err != nil {
			return nil, err
		}
	}
	if query := crq.withUser; query != nil {
		if err := crq.loadUser(ctx, query, nodes, nil,
			func(n *CancelRequest, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (crq *CancelRequestQuery) loadMatch(ctx context.Context, query *MatchQuery, nodes []*CancelRequest, init func(*CancelRequest), assign func(*CancelRequest, *Match)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CancelRequest)
	for i := range nodes {
		if nodes[i].cancel_request_match == nil {
			continue
		}
		fk := *nodes[i].cancel_request_match
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(match.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cancel_request_match" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (crq *CancelRequestQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*CancelRequest, init func(*CancelRequest), assign func(*CancelRequest, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CancelRequest)
	for i := range nodes {
		if nodes[i].user_cancel_request == nil {
			continue
		}
		fk := *nodes[i].user_cancel_request
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_cancel_request" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (crq *CancelRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	_spec.Node.Columns = crq.ctx.Fields
	if len(crq.ctx.Fields) > 0 {
		_spec.Unique = crq.ctx.Unique != nil && *crq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *CancelRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cancelrequest.Table, cancelrequest.Columns, sqlgraph.NewFieldSpec(cancelrequest.FieldID, field.TypeUUID))
	_spec.From = crq.sql
	if unique := crq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if crq.path != nil {
		_spec.Unique = true
	}
	if fields := crq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cancelrequest.FieldID)
		for i := range fields {
			if fields[i] != cancelrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *CancelRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(cancelrequest.Table)
	columns := crq.ctx.Fields
	if len(columns) == 0 {
		columns = cancelrequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crq.ctx.Unique != nil && *crq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CancelRequestGroupBy is the group-by builder for CancelRequest entities.
type CancelRequestGroupBy struct {
	selector
	build *CancelRequestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *CancelRequestGroupBy) Aggregate(fns ...AggregateFunc) *CancelRequestGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the selector query and scans the result into the given value.
func (crgb *CancelRequestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crgb.build.ctx, "GroupBy")
	if err := crgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CancelRequestQuery, *CancelRequestGroupBy](ctx, crgb.build, crgb, crgb.build.inters, v)
}

func (crgb *CancelRequestGroupBy) sqlScan(ctx context.Context, root *CancelRequestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(crgb.fns))
	for _, fn := range crgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*crgb.flds)+len(crgb.fns))
		for _, f := range *crgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*crgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CancelRequestSelect is the builder for selecting fields of CancelRequest entities.
type CancelRequestSelect struct {
	*CancelRequestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (crs *CancelRequestSelect) Aggregate(fns ...AggregateFunc) *CancelRequestSelect {
	crs.fns = append(crs.fns, fns...)
	return crs
}

// Scan applies the selector query and scans the result into the given value.
func (crs *CancelRequestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crs.ctx, "Select")
	if err := crs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CancelRequestQuery, *CancelRequestSelect](ctx, crs.CancelRequestQuery, crs, crs.inters, v)
}

func (crs *CancelRequestSelect) sqlScan(ctx context.Context, root *CancelRequestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(crs.fns))
	for _, fn := range crs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*crs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
