// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/payment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetQrPictureURL sets the "qr_picture_url" field.
func (pu *PaymentUpdate) SetQrPictureURL(s string) *PaymentUpdate {
	pu.mutation.SetQrPictureURL(s)
	return pu
}

// SetNillableQrPictureURL sets the "qr_picture_url" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableQrPictureURL(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetQrPictureURL(*s)
	}
	return pu
}

// ClearQrPictureURL clears the value of the "qr_picture_url" field.
func (pu *PaymentUpdate) ClearQrPictureURL() *PaymentUpdate {
	pu.mutation.ClearQrPictureURL()
	return pu
}

// SetPaymentStatus sets the "payment_status" field.
func (pu *PaymentUpdate) SetPaymentStatus(s string) *PaymentUpdate {
	pu.mutation.SetPaymentStatus(s)
	return pu
}

// SetCard sets the "card" field.
func (pu *PaymentUpdate) SetCard(s string) *PaymentUpdate {
	pu.mutation.SetCard(s)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(i int) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(i)
	return pu
}

// AddAmount adds i to the "amount" field.
func (pu *PaymentUpdate) AddAmount(i int) *PaymentUpdate {
	pu.mutation.AddAmount(i)
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PaymentUpdate) SetUserID(id uuid.UUID) *PaymentUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PaymentUpdate) SetUser(u *User) *PaymentUpdate {
	return pu.SetUserID(u.ID)
}

// AddPaymentHistoryIDs adds the "payment_history" edge to the PaymentHistory entity by IDs.
func (pu *PaymentUpdate) AddPaymentHistoryIDs(ids ...uuid.UUID) *PaymentUpdate {
	pu.mutation.AddPaymentHistoryIDs(ids...)
	return pu
}

// AddPaymentHistory adds the "payment_history" edges to the PaymentHistory entity.
func (pu *PaymentUpdate) AddPaymentHistory(p ...*PaymentHistory) *PaymentUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPaymentHistoryIDs(ids...)
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PaymentUpdate) ClearUser() *PaymentUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearPaymentHistory clears all "payment_history" edges to the PaymentHistory entity.
func (pu *PaymentUpdate) ClearPaymentHistory() *PaymentUpdate {
	pu.mutation.ClearPaymentHistory()
	return pu
}

// RemovePaymentHistoryIDs removes the "payment_history" edge to PaymentHistory entities by IDs.
func (pu *PaymentUpdate) RemovePaymentHistoryIDs(ids ...uuid.UUID) *PaymentUpdate {
	pu.mutation.RemovePaymentHistoryIDs(ids...)
	return pu
}

// RemovePaymentHistory removes "payment_history" edges to PaymentHistory entities.
func (pu *PaymentUpdate) RemovePaymentHistory(p ...*PaymentHistory) *PaymentUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePaymentHistoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PaymentMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.PaymentStatus(); ok {
		if err := payment.PaymentStatusValidator(v); err != nil {
			return &ValidationError{Name: "payment_status", err: fmt.Errorf(`ent: validator failed for field "Payment.payment_status": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Card(); ok {
		if err := payment.CardValidator(v); err != nil {
			return &ValidationError{Name: "card", err: fmt.Errorf(`ent: validator failed for field "Payment.card": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Payment.user"`)
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.QrPictureURL(); ok {
		_spec.SetField(payment.FieldQrPictureURL, field.TypeString, value)
	}
	if pu.mutation.QrPictureURLCleared() {
		_spec.ClearField(payment.FieldQrPictureURL, field.TypeString)
	}
	if value, ok := pu.mutation.PaymentStatus(); ok {
		_spec.SetField(payment.FieldPaymentStatus, field.TypeString, value)
	}
	if value, ok := pu.mutation.Card(); ok {
		_spec.SetField(payment.FieldCard, field.TypeString, value)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeInt, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.PaymentHistoryTable,
			Columns: []string{payment.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPaymentHistoryIDs(); len(nodes) > 0 && !pu.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.PaymentHistoryTable,
			Columns: []string{payment.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PaymentHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.PaymentHistoryTable,
			Columns: []string{payment.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetQrPictureURL sets the "qr_picture_url" field.
func (puo *PaymentUpdateOne) SetQrPictureURL(s string) *PaymentUpdateOne {
	puo.mutation.SetQrPictureURL(s)
	return puo
}

// SetNillableQrPictureURL sets the "qr_picture_url" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableQrPictureURL(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetQrPictureURL(*s)
	}
	return puo
}

// ClearQrPictureURL clears the value of the "qr_picture_url" field.
func (puo *PaymentUpdateOne) ClearQrPictureURL() *PaymentUpdateOne {
	puo.mutation.ClearQrPictureURL()
	return puo
}

// SetPaymentStatus sets the "payment_status" field.
func (puo *PaymentUpdateOne) SetPaymentStatus(s string) *PaymentUpdateOne {
	puo.mutation.SetPaymentStatus(s)
	return puo
}

// SetCard sets the "card" field.
func (puo *PaymentUpdateOne) SetCard(s string) *PaymentUpdateOne {
	puo.mutation.SetCard(s)
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(i int) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(i)
	return puo
}

// AddAmount adds i to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(i int) *PaymentUpdateOne {
	puo.mutation.AddAmount(i)
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PaymentUpdateOne) SetUserID(id uuid.UUID) *PaymentUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PaymentUpdateOne) SetUser(u *User) *PaymentUpdateOne {
	return puo.SetUserID(u.ID)
}

// AddPaymentHistoryIDs adds the "payment_history" edge to the PaymentHistory entity by IDs.
func (puo *PaymentUpdateOne) AddPaymentHistoryIDs(ids ...uuid.UUID) *PaymentUpdateOne {
	puo.mutation.AddPaymentHistoryIDs(ids...)
	return puo
}

// AddPaymentHistory adds the "payment_history" edges to the PaymentHistory entity.
func (puo *PaymentUpdateOne) AddPaymentHistory(p ...*PaymentHistory) *PaymentUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPaymentHistoryIDs(ids...)
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PaymentUpdateOne) ClearUser() *PaymentUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearPaymentHistory clears all "payment_history" edges to the PaymentHistory entity.
func (puo *PaymentUpdateOne) ClearPaymentHistory() *PaymentUpdateOne {
	puo.mutation.ClearPaymentHistory()
	return puo
}

// RemovePaymentHistoryIDs removes the "payment_history" edge to PaymentHistory entities by IDs.
func (puo *PaymentUpdateOne) RemovePaymentHistoryIDs(ids ...uuid.UUID) *PaymentUpdateOne {
	puo.mutation.RemovePaymentHistoryIDs(ids...)
	return puo
}

// RemovePaymentHistory removes "payment_history" edges to PaymentHistory entities.
func (puo *PaymentUpdateOne) RemovePaymentHistory(p ...*PaymentHistory) *PaymentUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePaymentHistoryIDs(ids...)
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	return withHooks[*Payment, PaymentMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.PaymentStatus(); ok {
		if err := payment.PaymentStatusValidator(v); err != nil {
			return &ValidationError{Name: "payment_status", err: fmt.Errorf(`ent: validator failed for field "Payment.payment_status": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Card(); ok {
		if err := payment.CardValidator(v); err != nil {
			return &ValidationError{Name: "card", err: fmt.Errorf(`ent: validator failed for field "Payment.card": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Payment.user"`)
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.QrPictureURL(); ok {
		_spec.SetField(payment.FieldQrPictureURL, field.TypeString, value)
	}
	if puo.mutation.QrPictureURLCleared() {
		_spec.ClearField(payment.FieldQrPictureURL, field.TypeString)
	}
	if value, ok := puo.mutation.PaymentStatus(); ok {
		_spec.SetField(payment.FieldPaymentStatus, field.TypeString, value)
	}
	if value, ok := puo.mutation.Card(); ok {
		_spec.SetField(payment.FieldCard, field.TypeString, value)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeInt, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.PaymentHistoryTable,
			Columns: []string{payment.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPaymentHistoryIDs(); len(nodes) > 0 && !puo.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.PaymentHistoryTable,
			Columns: []string{payment.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PaymentHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.PaymentHistoryTable,
			Columns: []string{payment.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
