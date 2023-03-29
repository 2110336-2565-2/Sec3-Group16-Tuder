// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/classcancelrequest"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/google/uuid"
)

// MatchCreate is the builder for creating a Match entity.
type MatchCreate struct {
	config
	mutation *MatchMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (mc *MatchCreate) SetID(u uuid.UUID) *MatchCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MatchCreate) SetNillableID(u *uuid.UUID) *MatchCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (mc *MatchCreate) SetStudentID(id uuid.UUID) *MatchCreate {
	mc.mutation.SetStudentID(id)
	return mc
}

// SetStudent sets the "student" edge to the Student entity.
func (mc *MatchCreate) SetStudent(s *Student) *MatchCreate {
	return mc.SetStudentID(s.ID)
}

// SetCourseID sets the "course" edge to the Course entity by ID.
func (mc *MatchCreate) SetCourseID(id uuid.UUID) *MatchCreate {
	mc.mutation.SetCourseID(id)
	return mc
}

// SetCourse sets the "course" edge to the Course entity.
func (mc *MatchCreate) SetCourse(c *Course) *MatchCreate {
	return mc.SetCourseID(c.ID)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (mc *MatchCreate) SetClassID(id uuid.UUID) *MatchCreate {
	mc.mutation.SetClassID(id)
	return mc
}

// SetClass sets the "class" edge to the Class entity.
func (mc *MatchCreate) SetClass(c *Class) *MatchCreate {
	return mc.SetClassID(c.ID)
}

// SetClassCancelRequestID sets the "class_cancel_request" edge to the ClassCancelRequest entity by ID.
func (mc *MatchCreate) SetClassCancelRequestID(id uuid.UUID) *MatchCreate {
	mc.mutation.SetClassCancelRequestID(id)
	return mc
}

// SetNillableClassCancelRequestID sets the "class_cancel_request" edge to the ClassCancelRequest entity by ID if the given value is not nil.
func (mc *MatchCreate) SetNillableClassCancelRequestID(id *uuid.UUID) *MatchCreate {
	if id != nil {
		mc = mc.SetClassCancelRequestID(*id)
	}
	return mc
}

// SetClassCancelRequest sets the "class_cancel_request" edge to the ClassCancelRequest entity.
func (mc *MatchCreate) SetClassCancelRequest(c *ClassCancelRequest) *MatchCreate {
	return mc.SetClassCancelRequestID(c.ID)
}

// Mutation returns the MatchMutation object of the builder.
func (mc *MatchCreate) Mutation() *MatchMutation {
	return mc.mutation
}

// Save creates the Match in the database.
func (mc *MatchCreate) Save(ctx context.Context) (*Match, error) {
	mc.defaults()
	return withHooks[*Match, MatchMutation](ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MatchCreate) SaveX(ctx context.Context) *Match {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MatchCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MatchCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MatchCreate) defaults() {
	if _, ok := mc.mutation.ID(); !ok {
		v := match.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MatchCreate) check() error {
	if _, ok := mc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student", err: errors.New(`ent: missing required edge "Match.student"`)}
	}
	if _, ok := mc.mutation.CourseID(); !ok {
		return &ValidationError{Name: "course", err: errors.New(`ent: missing required edge "Match.course"`)}
	}
	if _, ok := mc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class", err: errors.New(`ent: missing required edge "Match.class"`)}
	}
	return nil
}

func (mc *MatchCreate) sqlSave(ctx context.Context) (*Match, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MatchCreate) createSpec() (*Match, *sqlgraph.CreateSpec) {
	var (
		_node = &Match{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(match.Table, sqlgraph.NewFieldSpec(match.FieldID, field.TypeUUID))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := mc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   match.StudentTable,
			Columns: []string{match.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.student_match = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   match.CourseTable,
			Columns: []string{match.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.course_match = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   match.ClassTable,
			Columns: []string{match.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.class_match = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ClassCancelRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   match.ClassCancelRequestTable,
			Columns: []string{match.ClassCancelRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: classcancelrequest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MatchCreateBulk is the builder for creating many Match entities in bulk.
type MatchCreateBulk struct {
	config
	builders []*MatchCreate
}

// Save creates the Match entities in the database.
func (mcb *MatchCreateBulk) Save(ctx context.Context) ([]*Match, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Match, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MatchMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MatchCreateBulk) SaveX(ctx context.Context) []*Match {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MatchCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MatchCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
