// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/issuereport"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (sc *StudentCreate) SetID(u uuid.UUID) *StudentCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StudentCreate) SetNillableID(u *uuid.UUID) *StudentCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddIssueReportIDs adds the "issue_report" edge to the IssueReport entity by IDs.
func (sc *StudentCreate) AddIssueReportIDs(ids ...uuid.UUID) *StudentCreate {
	sc.mutation.AddIssueReportIDs(ids...)
	return sc
}

// AddIssueReport adds the "issue_report" edges to the IssueReport entity.
func (sc *StudentCreate) AddIssueReport(i ...*IssueReport) *StudentCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return sc.AddIssueReportIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (sc *StudentCreate) AddClasIDs(ids ...uuid.UUID) *StudentCreate {
	sc.mutation.AddClasIDs(ids...)
	return sc
}

// AddClass adds the "class" edges to the Class entity.
func (sc *StudentCreate) AddClass(c ...*Class) *StudentCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddClasIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *StudentCreate) SetUserID(id uuid.UUID) *StudentCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *StudentCreate) SetUser(u *User) *StudentCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	sc.defaults()
	return withHooks[*Student, StudentMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StudentCreate) defaults() {
	if _, ok := sc.mutation.ID(); !ok {
		v := student.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudentCreate) check() error {
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Student.user"`)}
	}
	return nil
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
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

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		_node = &Student{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(student.Table, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := sc.mutation.IssueReportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.IssueReportTable,
			Columns: []string{student.IssueReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: issuereport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClassTable,
			Columns: []string{student.ClassColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_student = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudentCreateBulk is the builder for creating many Student entities in bulk.
type StudentCreateBulk struct {
	config
	builders []*StudentCreate
}

// Save creates the Student entities in the database.
func (scb *StudentCreateBulk) Save(ctx context.Context) ([]*Student, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Student, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentMutation)
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
func (scb *StudentCreateBulk) SaveX(ctx context.Context) []*Student {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
