// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/google/uuid"
)

// ClassCreate is the builder for creating a Class entity.
type ClassCreate struct {
	config
	mutation *ClassMutation
	hooks    []Hook
}

// SetReviewAvaliable sets the "review_avaliable" field.
func (cc *ClassCreate) SetReviewAvaliable(b bool) *ClassCreate {
	cc.mutation.SetReviewAvaliable(b)
	return cc
}

// SetNillableReviewAvaliable sets the "review_avaliable" field if the given value is not nil.
func (cc *ClassCreate) SetNillableReviewAvaliable(b *bool) *ClassCreate {
	if b != nil {
		cc.SetReviewAvaliable(*b)
	}
	return cc
}

// SetTotalHour sets the "total_hour" field.
func (cc *ClassCreate) SetTotalHour(t time.Time) *ClassCreate {
	cc.mutation.SetTotalHour(t)
	return cc
}

// SetSuccessHour sets the "success_hour" field.
func (cc *ClassCreate) SetSuccessHour(t time.Time) *ClassCreate {
	cc.mutation.SetSuccessHour(t)
	return cc
}

// SetID sets the "id" field.
func (cc *ClassCreate) SetID(u uuid.UUID) *ClassCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ClassCreate) SetNillableID(u *uuid.UUID) *ClassCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddMatchIDs adds the "match" edge to the Match entity by IDs.
func (cc *ClassCreate) AddMatchIDs(ids ...int) *ClassCreate {
	cc.mutation.AddMatchIDs(ids...)
	return cc
}

// AddMatch adds the "match" edges to the Match entity.
func (cc *ClassCreate) AddMatch(m ...*Match) *ClassCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cc.AddMatchIDs(ids...)
}

// SetScheduleID sets the "schedule" edge to the Schedule entity by ID.
func (cc *ClassCreate) SetScheduleID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetScheduleID(id)
	return cc
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (cc *ClassCreate) SetSchedule(s *Schedule) *ClassCreate {
	return cc.SetScheduleID(s.ID)
}

// SetPaymentHistoryID sets the "payment_history" edge to the PaymentHistory entity by ID.
func (cc *ClassCreate) SetPaymentHistoryID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetPaymentHistoryID(id)
	return cc
}

// SetPaymentHistory sets the "payment_history" edge to the PaymentHistory entity.
func (cc *ClassCreate) SetPaymentHistory(p *PaymentHistory) *ClassCreate {
	return cc.SetPaymentHistoryID(p.ID)
}

// Mutation returns the ClassMutation object of the builder.
func (cc *ClassCreate) Mutation() *ClassMutation {
	return cc.mutation
}

// Save creates the Class in the database.
func (cc *ClassCreate) Save(ctx context.Context) (*Class, error) {
	cc.defaults()
	return withHooks[*Class, ClassMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassCreate) SaveX(ctx context.Context) *Class {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClassCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClassCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ClassCreate) defaults() {
	if _, ok := cc.mutation.ReviewAvaliable(); !ok {
		v := class.DefaultReviewAvaliable
		cc.mutation.SetReviewAvaliable(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := class.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClassCreate) check() error {
	if _, ok := cc.mutation.ReviewAvaliable(); !ok {
		return &ValidationError{Name: "review_avaliable", err: errors.New(`ent: missing required field "Class.review_avaliable"`)}
	}
	if _, ok := cc.mutation.TotalHour(); !ok {
		return &ValidationError{Name: "total_hour", err: errors.New(`ent: missing required field "Class.total_hour"`)}
	}
	if _, ok := cc.mutation.SuccessHour(); !ok {
		return &ValidationError{Name: "success_hour", err: errors.New(`ent: missing required field "Class.success_hour"`)}
	}
	if _, ok := cc.mutation.ScheduleID(); !ok {
		return &ValidationError{Name: "schedule", err: errors.New(`ent: missing required edge "Class.schedule"`)}
	}
	if _, ok := cc.mutation.PaymentHistoryID(); !ok {
		return &ValidationError{Name: "payment_history", err: errors.New(`ent: missing required edge "Class.payment_history"`)}
	}
	return nil
}

func (cc *ClassCreate) sqlSave(ctx context.Context) (*Class, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ClassCreate) createSpec() (*Class, *sqlgraph.CreateSpec) {
	var (
		_node = &Class{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(class.Table, sqlgraph.NewFieldSpec(class.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.ReviewAvaliable(); ok {
		_spec.SetField(class.FieldReviewAvaliable, field.TypeBool, value)
		_node.ReviewAvaliable = value
	}
	if value, ok := cc.mutation.TotalHour(); ok {
		_spec.SetField(class.FieldTotalHour, field.TypeTime, value)
		_node.TotalHour = value
	}
	if value, ok := cc.mutation.SuccessHour(); ok {
		_spec.SetField(class.FieldSuccessHour, field.TypeTime, value)
		_node.SuccessHour = value
	}
	if nodes := cc.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.MatchTable,
			Columns: class.MatchPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.ScheduleTable,
			Columns: []string{class.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.schedule_class = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PaymentHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.PaymentHistoryTable,
			Columns: []string{class.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: paymenthistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.payment_history_class = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClassCreateBulk is the builder for creating many Class entities in bulk.
type ClassCreateBulk struct {
	config
	builders []*ClassCreate
}

// Save creates the Class entities in the database.
func (ccb *ClassCreateBulk) Save(ctx context.Context) ([]*Class, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Class, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClassCreateBulk) SaveX(ctx context.Context) []*Class {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClassCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClassCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
