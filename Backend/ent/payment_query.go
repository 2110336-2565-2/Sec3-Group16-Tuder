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
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/payment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// PaymentQuery is the builder for querying Payment entities.
type PaymentQuery struct {
	config
	ctx                *QueryContext
	order              []OrderFunc
	inters             []Interceptor
	predicates         []predicate.Payment
	withUser           *UserQuery
	withMatch          *MatchQuery
	withAppointment    *AppointmentQuery
	withPaymentHistory *PaymentHistoryQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaymentQuery builder.
func (pq *PaymentQuery) Where(ps ...predicate.Payment) *PaymentQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PaymentQuery) Limit(limit int) *PaymentQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PaymentQuery) Offset(offset int) *PaymentQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PaymentQuery) Unique(unique bool) *PaymentQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PaymentQuery) Order(o ...OrderFunc) *PaymentQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryUser chains the current query on the "user" edge.
func (pq *PaymentQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payment.Table, payment.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payment.UserTable, payment.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMatch chains the current query on the "match" edge.
func (pq *PaymentQuery) QueryMatch() *MatchQuery {
	query := (&MatchClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payment.Table, payment.FieldID, selector),
			sqlgraph.To(match.Table, match.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, payment.MatchTable, payment.MatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAppointment chains the current query on the "appointment" edge.
func (pq *PaymentQuery) QueryAppointment() *AppointmentQuery {
	query := (&AppointmentClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payment.Table, payment.FieldID, selector),
			sqlgraph.To(appointment.Table, appointment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, payment.AppointmentTable, payment.AppointmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPaymentHistory chains the current query on the "payment_history" edge.
func (pq *PaymentQuery) QueryPaymentHistory() *PaymentHistoryQuery {
	query := (&PaymentHistoryClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payment.Table, payment.FieldID, selector),
			sqlgraph.To(paymenthistory.Table, paymenthistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, payment.PaymentHistoryTable, payment.PaymentHistoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Payment entity from the query.
// Returns a *NotFoundError when no Payment was found.
func (pq *PaymentQuery) First(ctx context.Context) (*Payment, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{payment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PaymentQuery) FirstX(ctx context.Context) *Payment {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Payment ID from the query.
// Returns a *NotFoundError when no Payment ID was found.
func (pq *PaymentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{payment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PaymentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Payment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Payment entity is found.
// Returns a *NotFoundError when no Payment entities are found.
func (pq *PaymentQuery) Only(ctx context.Context) (*Payment, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{payment.Label}
	default:
		return nil, &NotSingularError{payment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PaymentQuery) OnlyX(ctx context.Context) *Payment {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Payment ID in the query.
// Returns a *NotSingularError when more than one Payment ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PaymentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{payment.Label}
	default:
		err = &NotSingularError{payment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PaymentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Payments.
func (pq *PaymentQuery) All(ctx context.Context) ([]*Payment, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Payment, *PaymentQuery]()
	return withInterceptors[[]*Payment](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PaymentQuery) AllX(ctx context.Context) []*Payment {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Payment IDs.
func (pq *PaymentQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(payment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PaymentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PaymentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PaymentQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PaymentQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PaymentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PaymentQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaymentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PaymentQuery) Clone() *PaymentQuery {
	if pq == nil {
		return nil
	}
	return &PaymentQuery{
		config:             pq.config,
		ctx:                pq.ctx.Clone(),
		order:              append([]OrderFunc{}, pq.order...),
		inters:             append([]Interceptor{}, pq.inters...),
		predicates:         append([]predicate.Payment{}, pq.predicates...),
		withUser:           pq.withUser.Clone(),
		withMatch:          pq.withMatch.Clone(),
		withAppointment:    pq.withAppointment.Clone(),
		withPaymentHistory: pq.withPaymentHistory.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PaymentQuery) WithUser(opts ...func(*UserQuery)) *PaymentQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withUser = query
	return pq
}

// WithMatch tells the query-builder to eager-load the nodes that are connected to
// the "match" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PaymentQuery) WithMatch(opts ...func(*MatchQuery)) *PaymentQuery {
	query := (&MatchClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withMatch = query
	return pq
}

// WithAppointment tells the query-builder to eager-load the nodes that are connected to
// the "appointment" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PaymentQuery) WithAppointment(opts ...func(*AppointmentQuery)) *PaymentQuery {
	query := (&AppointmentClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withAppointment = query
	return pq
}

// WithPaymentHistory tells the query-builder to eager-load the nodes that are connected to
// the "payment_history" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PaymentQuery) WithPaymentHistory(opts ...func(*PaymentHistoryQuery)) *PaymentQuery {
	query := (&PaymentHistoryClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPaymentHistory = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		QrPictureURL string `json:"qr_picture_url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Payment.Query().
//		GroupBy(payment.FieldQrPictureURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PaymentQuery) GroupBy(field string, fields ...string) *PaymentGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PaymentGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = payment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		QrPictureURL string `json:"qr_picture_url,omitempty"`
//	}
//
//	client.Payment.Query().
//		Select(payment.FieldQrPictureURL).
//		Scan(ctx, &v)
func (pq *PaymentQuery) Select(fields ...string) *PaymentSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PaymentSelect{PaymentQuery: pq}
	sbuild.label = payment.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PaymentSelect configured with the given aggregations.
func (pq *PaymentQuery) Aggregate(fns ...AggregateFunc) *PaymentSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PaymentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !payment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PaymentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Payment, error) {
	var (
		nodes       = []*Payment{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [4]bool{
			pq.withUser != nil,
			pq.withMatch != nil,
			pq.withAppointment != nil,
			pq.withPaymentHistory != nil,
		}
	)
	if pq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, payment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Payment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Payment{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withUser; query != nil {
		if err := pq.loadUser(ctx, query, nodes, nil,
			func(n *Payment, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withMatch; query != nil {
		if err := pq.loadMatch(ctx, query, nodes, nil,
			func(n *Payment, e *Match) { n.Edges.Match = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withAppointment; query != nil {
		if err := pq.loadAppointment(ctx, query, nodes, nil,
			func(n *Payment, e *Appointment) { n.Edges.Appointment = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withPaymentHistory; query != nil {
		if err := pq.loadPaymentHistory(ctx, query, nodes,
			func(n *Payment) { n.Edges.PaymentHistory = []*PaymentHistory{} },
			func(n *Payment, e *PaymentHistory) { n.Edges.PaymentHistory = append(n.Edges.PaymentHistory, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PaymentQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Payment, init func(*Payment), assign func(*Payment, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Payment)
	for i := range nodes {
		if nodes[i].user_payment == nil {
			continue
		}
		fk := *nodes[i].user_payment
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
			return fmt.Errorf(`unexpected foreign-key "user_payment" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PaymentQuery) loadMatch(ctx context.Context, query *MatchQuery, nodes []*Payment, init func(*Payment), assign func(*Payment, *Match)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Payment)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Match(func(s *sql.Selector) {
		s.Where(sql.InValues(payment.MatchColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.payment_match
		if fk == nil {
			return fmt.Errorf(`foreign-key "payment_match" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "payment_match" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PaymentQuery) loadAppointment(ctx context.Context, query *AppointmentQuery, nodes []*Payment, init func(*Payment), assign func(*Payment, *Appointment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Payment)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.InValues(payment.AppointmentColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.payment_appointment
		if fk == nil {
			return fmt.Errorf(`foreign-key "payment_appointment" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "payment_appointment" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PaymentQuery) loadPaymentHistory(ctx context.Context, query *PaymentHistoryQuery, nodes []*Payment, init func(*Payment), assign func(*Payment, *PaymentHistory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Payment)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PaymentHistory(func(s *sql.Selector) {
		s.Where(sql.InValues(payment.PaymentHistoryColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.payment_payment_history
		if fk == nil {
			return fmt.Errorf(`foreign-key "payment_payment_history" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "payment_payment_history" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PaymentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PaymentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for i := range fields {
			if fields[i] != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PaymentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(payment.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = payment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaymentGroupBy is the group-by builder for Payment entities.
type PaymentGroupBy struct {
	selector
	build *PaymentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PaymentGroupBy) Aggregate(fns ...AggregateFunc) *PaymentGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PaymentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaymentQuery, *PaymentGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PaymentGroupBy) sqlScan(ctx context.Context, root *PaymentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PaymentSelect is the builder for selecting fields of Payment entities.
type PaymentSelect struct {
	*PaymentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PaymentSelect) Aggregate(fns ...AggregateFunc) *PaymentSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PaymentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaymentQuery, *PaymentSelect](ctx, ps.PaymentQuery, ps, ps.inters, v)
}

func (ps *PaymentSelect) sqlScan(ctx context.Context, root *PaymentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
