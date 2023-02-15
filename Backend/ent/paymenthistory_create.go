// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/google/uuid"
)

// PaymentHistoryCreate is the builder for creating a PaymentHistory entity.
type PaymentHistoryCreate struct {
	config
	mutation *PaymentHistoryMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (phc *PaymentHistoryCreate) SetAmount(f float64) *PaymentHistoryCreate {
	phc.mutation.SetAmount(f)
	return phc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (phc *PaymentHistoryCreate) SetNillableAmount(f *float64) *PaymentHistoryCreate {
	if f != nil {
		phc.SetAmount(*f)
	}
	return phc
}

// SetType sets the "type" field.
func (phc *PaymentHistoryCreate) SetType(s string) *PaymentHistoryCreate {
	phc.mutation.SetType(s)
	return phc
}

// SetID sets the "id" field.
func (phc *PaymentHistoryCreate) SetID(u uuid.UUID) *PaymentHistoryCreate {
	phc.mutation.SetID(u)
	return phc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (phc *PaymentHistoryCreate) SetNillableID(u *uuid.UUID) *PaymentHistoryCreate {
	if u != nil {
		phc.SetID(*u)
	}
	return phc
}

// Mutation returns the PaymentHistoryMutation object of the builder.
func (phc *PaymentHistoryCreate) Mutation() *PaymentHistoryMutation {
	return phc.mutation
}

// Save creates the PaymentHistory in the database.
func (phc *PaymentHistoryCreate) Save(ctx context.Context) (*PaymentHistory, error) {
	phc.defaults()
	return withHooks[*PaymentHistory, PaymentHistoryMutation](ctx, phc.sqlSave, phc.mutation, phc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (phc *PaymentHistoryCreate) SaveX(ctx context.Context) *PaymentHistory {
	v, err := phc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (phc *PaymentHistoryCreate) Exec(ctx context.Context) error {
	_, err := phc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (phc *PaymentHistoryCreate) ExecX(ctx context.Context) {
	if err := phc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (phc *PaymentHistoryCreate) defaults() {
	if _, ok := phc.mutation.ID(); !ok {
		v := paymenthistory.DefaultID()
		phc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (phc *PaymentHistoryCreate) check() error {
	if v, ok := phc.mutation.Amount(); ok {
		if err := paymenthistory.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "PaymentHistory.amount": %w`, err)}
		}
	}
	if _, ok := phc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "PaymentHistory.type"`)}
	}
	if v, ok := phc.mutation.GetType(); ok {
		if err := paymenthistory.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PaymentHistory.type": %w`, err)}
		}
	}
	return nil
}

func (phc *PaymentHistoryCreate) sqlSave(ctx context.Context) (*PaymentHistory, error) {
	if err := phc.check(); err != nil {
		return nil, err
	}
	_node, _spec := phc.createSpec()
	if err := sqlgraph.CreateNode(ctx, phc.driver, _spec); err != nil {
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
	phc.mutation.id = &_node.ID
	phc.mutation.done = true
	return _node, nil
}

func (phc *PaymentHistoryCreate) createSpec() (*PaymentHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &PaymentHistory{config: phc.config}
		_spec = sqlgraph.NewCreateSpec(paymenthistory.Table, sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID))
	)
	if id, ok := phc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := phc.mutation.Amount(); ok {
		_spec.SetField(paymenthistory.FieldAmount, field.TypeFloat64, value)
		_node.Amount = &value
	}
	if value, ok := phc.mutation.GetType(); ok {
		_spec.SetField(paymenthistory.FieldType, field.TypeString, value)
		_node.Type = value
	}
	return _node, _spec
}

// PaymentHistoryCreateBulk is the builder for creating many PaymentHistory entities in bulk.
type PaymentHistoryCreateBulk struct {
	config
	builders []*PaymentHistoryCreate
}

// Save creates the PaymentHistory entities in the database.
func (phcb *PaymentHistoryCreateBulk) Save(ctx context.Context) ([]*PaymentHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(phcb.builders))
	nodes := make([]*PaymentHistory, len(phcb.builders))
	mutators := make([]Mutator, len(phcb.builders))
	for i := range phcb.builders {
		func(i int, root context.Context) {
			builder := phcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, phcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, phcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, phcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (phcb *PaymentHistoryCreateBulk) SaveX(ctx context.Context) []*PaymentHistory {
	v, err := phcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (phcb *PaymentHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := phcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (phcb *PaymentHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := phcb.Exec(ctx); err != nil {
		panic(err)
	}
}
