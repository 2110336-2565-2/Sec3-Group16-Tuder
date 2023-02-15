// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/google/uuid"
)

// ScheduleCreate is the builder for creating a Schedule entity.
type ScheduleCreate struct {
	config
	mutation *ScheduleMutation
	hooks    []Hook
}

// SetDay0 sets the "day_0" field.
func (sc *ScheduleCreate) SetDay0(b bool) *ScheduleCreate {
	sc.mutation.SetDay0(b)
	return sc
}

// SetNillableDay0 sets the "day_0" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDay0(b *bool) *ScheduleCreate {
	if b != nil {
		sc.SetDay0(*b)
	}
	return sc
}

// SetDay1 sets the "day_1" field.
func (sc *ScheduleCreate) SetDay1(b bool) *ScheduleCreate {
	sc.mutation.SetDay1(b)
	return sc
}

// SetNillableDay1 sets the "day_1" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDay1(b *bool) *ScheduleCreate {
	if b != nil {
		sc.SetDay1(*b)
	}
	return sc
}

// SetDay2 sets the "day_2" field.
func (sc *ScheduleCreate) SetDay2(b bool) *ScheduleCreate {
	sc.mutation.SetDay2(b)
	return sc
}

// SetNillableDay2 sets the "day_2" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDay2(b *bool) *ScheduleCreate {
	if b != nil {
		sc.SetDay2(*b)
	}
	return sc
}

// SetDay3 sets the "day_3" field.
func (sc *ScheduleCreate) SetDay3(b bool) *ScheduleCreate {
	sc.mutation.SetDay3(b)
	return sc
}

// SetNillableDay3 sets the "day_3" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDay3(b *bool) *ScheduleCreate {
	if b != nil {
		sc.SetDay3(*b)
	}
	return sc
}

// SetDay4 sets the "day_4" field.
func (sc *ScheduleCreate) SetDay4(b bool) *ScheduleCreate {
	sc.mutation.SetDay4(b)
	return sc
}

// SetNillableDay4 sets the "day_4" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDay4(b *bool) *ScheduleCreate {
	if b != nil {
		sc.SetDay4(*b)
	}
	return sc
}

// SetDay5 sets the "day_5" field.
func (sc *ScheduleCreate) SetDay5(b bool) *ScheduleCreate {
	sc.mutation.SetDay5(b)
	return sc
}

// SetNillableDay5 sets the "day_5" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDay5(b *bool) *ScheduleCreate {
	if b != nil {
		sc.SetDay5(*b)
	}
	return sc
}

// SetDay6 sets the "day_6" field.
func (sc *ScheduleCreate) SetDay6(b bool) *ScheduleCreate {
	sc.mutation.SetDay6(b)
	return sc
}

// SetNillableDay6 sets the "day_6" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDay6(b *bool) *ScheduleCreate {
	if b != nil {
		sc.SetDay6(*b)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ScheduleCreate) SetID(u uuid.UUID) *ScheduleCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableID(u *uuid.UUID) *ScheduleCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the ScheduleMutation object of the builder.
func (sc *ScheduleCreate) Mutation() *ScheduleMutation {
	return sc.mutation
}

// Save creates the Schedule in the database.
func (sc *ScheduleCreate) Save(ctx context.Context) (*Schedule, error) {
	sc.defaults()
	return withHooks[*Schedule, ScheduleMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScheduleCreate) SaveX(ctx context.Context) *Schedule {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScheduleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScheduleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ScheduleCreate) defaults() {
	if _, ok := sc.mutation.Day0(); !ok {
		v := schedule.DefaultDay0
		sc.mutation.SetDay0(v)
	}
	if _, ok := sc.mutation.Day1(); !ok {
		v := schedule.DefaultDay1
		sc.mutation.SetDay1(v)
	}
	if _, ok := sc.mutation.Day2(); !ok {
		v := schedule.DefaultDay2
		sc.mutation.SetDay2(v)
	}
	if _, ok := sc.mutation.Day3(); !ok {
		v := schedule.DefaultDay3
		sc.mutation.SetDay3(v)
	}
	if _, ok := sc.mutation.Day4(); !ok {
		v := schedule.DefaultDay4
		sc.mutation.SetDay4(v)
	}
	if _, ok := sc.mutation.Day5(); !ok {
		v := schedule.DefaultDay5
		sc.mutation.SetDay5(v)
	}
	if _, ok := sc.mutation.Day6(); !ok {
		v := schedule.DefaultDay6
		sc.mutation.SetDay6(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := schedule.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScheduleCreate) check() error {
	return nil
}

func (sc *ScheduleCreate) sqlSave(ctx context.Context) (*Schedule, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScheduleCreate) createSpec() (*Schedule, *sqlgraph.CreateSpec) {
	var (
		_node = &Schedule{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(schedule.Table, sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Day0(); ok {
		_spec.SetField(schedule.FieldDay0, field.TypeBool, value)
		_node.Day0 = &value
	}
	if value, ok := sc.mutation.Day1(); ok {
		_spec.SetField(schedule.FieldDay1, field.TypeBool, value)
		_node.Day1 = &value
	}
	if value, ok := sc.mutation.Day2(); ok {
		_spec.SetField(schedule.FieldDay2, field.TypeBool, value)
		_node.Day2 = &value
	}
	if value, ok := sc.mutation.Day3(); ok {
		_spec.SetField(schedule.FieldDay3, field.TypeBool, value)
		_node.Day3 = &value
	}
	if value, ok := sc.mutation.Day4(); ok {
		_spec.SetField(schedule.FieldDay4, field.TypeBool, value)
		_node.Day4 = &value
	}
	if value, ok := sc.mutation.Day5(); ok {
		_spec.SetField(schedule.FieldDay5, field.TypeBool, value)
		_node.Day5 = &value
	}
	if value, ok := sc.mutation.Day6(); ok {
		_spec.SetField(schedule.FieldDay6, field.TypeBool, value)
		_node.Day6 = &value
	}
	return _node, _spec
}

// ScheduleCreateBulk is the builder for creating many Schedule entities in bulk.
type ScheduleCreateBulk struct {
	config
	builders []*ScheduleCreate
}

// Save creates the Schedule entities in the database.
func (scb *ScheduleCreateBulk) Save(ctx context.Context) ([]*Schedule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Schedule, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScheduleMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScheduleCreateBulk) SaveX(ctx context.Context) []*Schedule {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScheduleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScheduleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
