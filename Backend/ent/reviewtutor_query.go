// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewtutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/google/uuid"
)

// ReviewTutorQuery is the builder for querying ReviewTutor entities.
type ReviewTutorQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.ReviewTutor
	withTutor   *TutorQuery
	withStudent *StudentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReviewTutorQuery builder.
func (rtq *ReviewTutorQuery) Where(ps ...predicate.ReviewTutor) *ReviewTutorQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit the number of records to be returned by this query.
func (rtq *ReviewTutorQuery) Limit(limit int) *ReviewTutorQuery {
	rtq.ctx.Limit = &limit
	return rtq
}

// Offset to start from.
func (rtq *ReviewTutorQuery) Offset(offset int) *ReviewTutorQuery {
	rtq.ctx.Offset = &offset
	return rtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rtq *ReviewTutorQuery) Unique(unique bool) *ReviewTutorQuery {
	rtq.ctx.Unique = &unique
	return rtq
}

// Order specifies how the records should be ordered.
func (rtq *ReviewTutorQuery) Order(o ...OrderFunc) *ReviewTutorQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// QueryTutor chains the current query on the "tutor" edge.
func (rtq *ReviewTutorQuery) QueryTutor() *TutorQuery {
	query := (&TutorClient{config: rtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reviewtutor.Table, reviewtutor.FieldID, selector),
			sqlgraph.To(tutor.Table, tutor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, reviewtutor.TutorTable, reviewtutor.TutorPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudent chains the current query on the "student" edge.
func (rtq *ReviewTutorQuery) QueryStudent() *StudentQuery {
	query := (&StudentClient{config: rtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reviewtutor.Table, reviewtutor.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, reviewtutor.StudentTable, reviewtutor.StudentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ReviewTutor entity from the query.
// Returns a *NotFoundError when no ReviewTutor was found.
func (rtq *ReviewTutorQuery) First(ctx context.Context) (*ReviewTutor, error) {
	nodes, err := rtq.Limit(1).All(setContextOp(ctx, rtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reviewtutor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *ReviewTutorQuery) FirstX(ctx context.Context) *ReviewTutor {
	node, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ReviewTutor ID from the query.
// Returns a *NotFoundError when no ReviewTutor ID was found.
func (rtq *ReviewTutorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(1).IDs(setContextOp(ctx, rtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reviewtutor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rtq *ReviewTutorQuery) FirstIDX(ctx context.Context) int {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ReviewTutor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ReviewTutor entity is found.
// Returns a *NotFoundError when no ReviewTutor entities are found.
func (rtq *ReviewTutorQuery) Only(ctx context.Context) (*ReviewTutor, error) {
	nodes, err := rtq.Limit(2).All(setContextOp(ctx, rtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reviewtutor.Label}
	default:
		return nil, &NotSingularError{reviewtutor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *ReviewTutorQuery) OnlyX(ctx context.Context) *ReviewTutor {
	node, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ReviewTutor ID in the query.
// Returns a *NotSingularError when more than one ReviewTutor ID is found.
// Returns a *NotFoundError when no entities are found.
func (rtq *ReviewTutorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(2).IDs(setContextOp(ctx, rtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reviewtutor.Label}
	default:
		err = &NotSingularError{reviewtutor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *ReviewTutorQuery) OnlyIDX(ctx context.Context) int {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReviewTutors.
func (rtq *ReviewTutorQuery) All(ctx context.Context) ([]*ReviewTutor, error) {
	ctx = setContextOp(ctx, rtq.ctx, "All")
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ReviewTutor, *ReviewTutorQuery]()
	return withInterceptors[[]*ReviewTutor](ctx, rtq, qr, rtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rtq *ReviewTutorQuery) AllX(ctx context.Context) []*ReviewTutor {
	nodes, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ReviewTutor IDs.
func (rtq *ReviewTutorQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rtq.ctx.Unique == nil && rtq.path != nil {
		rtq.Unique(true)
	}
	ctx = setContextOp(ctx, rtq.ctx, "IDs")
	if err = rtq.Select(reviewtutor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *ReviewTutorQuery) IDsX(ctx context.Context) []int {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *ReviewTutorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Count")
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rtq, querierCount[*ReviewTutorQuery](), rtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *ReviewTutorQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *ReviewTutorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Exist")
	switch _, err := rtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *ReviewTutorQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReviewTutorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *ReviewTutorQuery) Clone() *ReviewTutorQuery {
	if rtq == nil {
		return nil
	}
	return &ReviewTutorQuery{
		config:      rtq.config,
		ctx:         rtq.ctx.Clone(),
		order:       append([]OrderFunc{}, rtq.order...),
		inters:      append([]Interceptor{}, rtq.inters...),
		predicates:  append([]predicate.ReviewTutor{}, rtq.predicates...),
		withTutor:   rtq.withTutor.Clone(),
		withStudent: rtq.withStudent.Clone(),
		// clone intermediate query.
		sql:  rtq.sql.Clone(),
		path: rtq.path,
	}
}

// WithTutor tells the query-builder to eager-load the nodes that are connected to
// the "tutor" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *ReviewTutorQuery) WithTutor(opts ...func(*TutorQuery)) *ReviewTutorQuery {
	query := (&TutorClient{config: rtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rtq.withTutor = query
	return rtq
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *ReviewTutorQuery) WithStudent(opts ...func(*StudentQuery)) *ReviewTutorQuery {
	query := (&StudentClient{config: rtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rtq.withStudent = query
	return rtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Score float32 `json:"score,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ReviewTutor.Query().
//		GroupBy(reviewtutor.FieldScore).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rtq *ReviewTutorQuery) GroupBy(field string, fields ...string) *ReviewTutorGroupBy {
	rtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReviewTutorGroupBy{build: rtq}
	grbuild.flds = &rtq.ctx.Fields
	grbuild.label = reviewtutor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Score float32 `json:"score,omitempty"`
//	}
//
//	client.ReviewTutor.Query().
//		Select(reviewtutor.FieldScore).
//		Scan(ctx, &v)
func (rtq *ReviewTutorQuery) Select(fields ...string) *ReviewTutorSelect {
	rtq.ctx.Fields = append(rtq.ctx.Fields, fields...)
	sbuild := &ReviewTutorSelect{ReviewTutorQuery: rtq}
	sbuild.label = reviewtutor.Label
	sbuild.flds, sbuild.scan = &rtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReviewTutorSelect configured with the given aggregations.
func (rtq *ReviewTutorQuery) Aggregate(fns ...AggregateFunc) *ReviewTutorSelect {
	return rtq.Select().Aggregate(fns...)
}

func (rtq *ReviewTutorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rtq); err != nil {
				return err
			}
		}
	}
	for _, f := range rtq.ctx.Fields {
		if !reviewtutor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *ReviewTutorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ReviewTutor, error) {
	var (
		nodes       = []*ReviewTutor{}
		_spec       = rtq.querySpec()
		loadedTypes = [2]bool{
			rtq.withTutor != nil,
			rtq.withStudent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ReviewTutor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ReviewTutor{config: rtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rtq.withTutor; query != nil {
		if err := rtq.loadTutor(ctx, query, nodes,
			func(n *ReviewTutor) { n.Edges.Tutor = []*Tutor{} },
			func(n *ReviewTutor, e *Tutor) { n.Edges.Tutor = append(n.Edges.Tutor, e) }); err != nil {
			return nil, err
		}
	}
	if query := rtq.withStudent; query != nil {
		if err := rtq.loadStudent(ctx, query, nodes,
			func(n *ReviewTutor) { n.Edges.Student = []*Student{} },
			func(n *ReviewTutor, e *Student) { n.Edges.Student = append(n.Edges.Student, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rtq *ReviewTutorQuery) loadTutor(ctx context.Context, query *TutorQuery, nodes []*ReviewTutor, init func(*ReviewTutor), assign func(*ReviewTutor, *Tutor)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ReviewTutor)
	nids := make(map[uuid.UUID]map[*ReviewTutor]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(reviewtutor.TutorTable)
		s.Join(joinT).On(s.C(tutor.FieldID), joinT.C(reviewtutor.TutorPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(reviewtutor.TutorPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(reviewtutor.TutorPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*ReviewTutor]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tutor](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tutor" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rtq *ReviewTutorQuery) loadStudent(ctx context.Context, query *StudentQuery, nodes []*ReviewTutor, init func(*ReviewTutor), assign func(*ReviewTutor, *Student)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ReviewTutor)
	nids := make(map[uuid.UUID]map[*ReviewTutor]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(reviewtutor.StudentTable)
		s.Join(joinT).On(s.C(student.FieldID), joinT.C(reviewtutor.StudentPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(reviewtutor.StudentPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(reviewtutor.StudentPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*ReviewTutor]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Student](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "student" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (rtq *ReviewTutorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	_spec.Node.Columns = rtq.ctx.Fields
	if len(rtq.ctx.Fields) > 0 {
		_spec.Unique = rtq.ctx.Unique != nil && *rtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *ReviewTutorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(reviewtutor.Table, reviewtutor.Columns, sqlgraph.NewFieldSpec(reviewtutor.FieldID, field.TypeInt))
	_spec.From = rtq.sql
	if unique := rtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rtq.path != nil {
		_spec.Unique = true
	}
	if fields := rtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reviewtutor.FieldID)
		for i := range fields {
			if fields[i] != reviewtutor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *ReviewTutorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(reviewtutor.Table)
	columns := rtq.ctx.Fields
	if len(columns) == 0 {
		columns = reviewtutor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rtq.ctx.Unique != nil && *rtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReviewTutorGroupBy is the group-by builder for ReviewTutor entities.
type ReviewTutorGroupBy struct {
	selector
	build *ReviewTutorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *ReviewTutorGroupBy) Aggregate(fns ...AggregateFunc) *ReviewTutorGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the selector query and scans the result into the given value.
func (rtgb *ReviewTutorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rtgb.build.ctx, "GroupBy")
	if err := rtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReviewTutorQuery, *ReviewTutorGroupBy](ctx, rtgb.build, rtgb, rtgb.build.inters, v)
}

func (rtgb *ReviewTutorGroupBy) sqlScan(ctx context.Context, root *ReviewTutorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rtgb.fns))
	for _, fn := range rtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rtgb.flds)+len(rtgb.fns))
		for _, f := range *rtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReviewTutorSelect is the builder for selecting fields of ReviewTutor entities.
type ReviewTutorSelect struct {
	*ReviewTutorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rts *ReviewTutorSelect) Aggregate(fns ...AggregateFunc) *ReviewTutorSelect {
	rts.fns = append(rts.fns, fns...)
	return rts
}

// Scan applies the selector query and scans the result into the given value.
func (rts *ReviewTutorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rts.ctx, "Select")
	if err := rts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReviewTutorQuery, *ReviewTutorSelect](ctx, rts.ReviewTutorQuery, rts, rts.inters, v)
}

func (rts *ReviewTutorSelect) sqlScan(ctx context.Context, root *ReviewTutorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rts.fns))
	for _, fn := range rts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
