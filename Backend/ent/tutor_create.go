// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/issuereport"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewtutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// TutorCreate is the builder for creating a Tutor entity.
type TutorCreate struct {
	config
	mutation *TutorMutation
	hooks    []Hook
}

// SetDescription sets the "description" field.
func (tc *TutorCreate) SetDescription(s string) *TutorCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TutorCreate) SetNillableDescription(s *string) *TutorCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetOmiseBankToken sets the "omise_bank_token" field.
func (tc *TutorCreate) SetOmiseBankToken(s string) *TutorCreate {
	tc.mutation.SetOmiseBankToken(s)
	return tc
}

// SetNillableOmiseBankToken sets the "omise_bank_token" field if the given value is not nil.
func (tc *TutorCreate) SetNillableOmiseBankToken(s *string) *TutorCreate {
	if s != nil {
		tc.SetOmiseBankToken(*s)
	}
	return tc
}

// SetCitizenID sets the "citizen_id" field.
func (tc *TutorCreate) SetCitizenID(s string) *TutorCreate {
	tc.mutation.SetCitizenID(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TutorCreate) SetID(u uuid.UUID) *TutorCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TutorCreate) SetNillableID(u *uuid.UUID) *TutorCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// AddIssueReportIDs adds the "issue_report" edge to the IssueReport entity by IDs.
func (tc *TutorCreate) AddIssueReportIDs(ids ...uuid.UUID) *TutorCreate {
	tc.mutation.AddIssueReportIDs(ids...)
	return tc
}

// AddIssueReport adds the "issue_report" edges to the IssueReport entity.
func (tc *TutorCreate) AddIssueReport(i ...*IssueReport) *TutorCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tc.AddIssueReportIDs(ids...)
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (tc *TutorCreate) AddCourseIDs(ids ...uuid.UUID) *TutorCreate {
	tc.mutation.AddCourseIDs(ids...)
	return tc
}

// AddCourse adds the "course" edges to the Course entity.
func (tc *TutorCreate) AddCourse(c ...*Course) *TutorCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddCourseIDs(ids...)
}

// AddReviewTutorIDs adds the "review_tutor" edge to the ReviewTutor entity by IDs.
func (tc *TutorCreate) AddReviewTutorIDs(ids ...int) *TutorCreate {
	tc.mutation.AddReviewTutorIDs(ids...)
	return tc
}

// AddReviewTutor adds the "review_tutor" edges to the ReviewTutor entity.
func (tc *TutorCreate) AddReviewTutor(r ...*ReviewTutor) *TutorCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddReviewTutorIDs(ids...)
}

// AddScheduleIDs adds the "schedule" edge to the Schedule entity by IDs.
func (tc *TutorCreate) AddScheduleIDs(ids ...uuid.UUID) *TutorCreate {
	tc.mutation.AddScheduleIDs(ids...)
	return tc
}

// AddSchedule adds the "schedule" edges to the Schedule entity.
func (tc *TutorCreate) AddSchedule(s ...*Schedule) *TutorCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddScheduleIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tc *TutorCreate) SetUserID(id uuid.UUID) *TutorCreate {
	tc.mutation.SetUserID(id)
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TutorCreate) SetUser(u *User) *TutorCreate {
	return tc.SetUserID(u.ID)
}

// Mutation returns the TutorMutation object of the builder.
func (tc *TutorCreate) Mutation() *TutorMutation {
	return tc.mutation
}

// Save creates the Tutor in the database.
func (tc *TutorCreate) Save(ctx context.Context) (*Tutor, error) {
	tc.defaults()
	return withHooks[*Tutor, TutorMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TutorCreate) SaveX(ctx context.Context) *Tutor {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TutorCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TutorCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TutorCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := tutor.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TutorCreate) check() error {
	if _, ok := tc.mutation.CitizenID(); !ok {
		return &ValidationError{Name: "citizen_id", err: errors.New(`ent: missing required field "Tutor.citizen_id"`)}
	}
	if v, ok := tc.mutation.CitizenID(); ok {
		if err := tutor.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Tutor.citizen_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Tutor.user"`)}
	}
	return nil
}

func (tc *TutorCreate) sqlSave(ctx context.Context) (*Tutor, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TutorCreate) createSpec() (*Tutor, *sqlgraph.CreateSpec) {
	var (
		_node = &Tutor{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tutor.Table, sqlgraph.NewFieldSpec(tutor.FieldID, field.TypeUUID))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(tutor.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := tc.mutation.OmiseBankToken(); ok {
		_spec.SetField(tutor.FieldOmiseBankToken, field.TypeString, value)
		_node.OmiseBankToken = &value
	}
	if value, ok := tc.mutation.CitizenID(); ok {
		_spec.SetField(tutor.FieldCitizenID, field.TypeString, value)
		_node.CitizenID = value
	}
	if nodes := tc.mutation.IssueReportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tutor.IssueReportTable,
			Columns: []string{tutor.IssueReportColumn},
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
	if nodes := tc.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tutor.CourseTable,
			Columns: []string{tutor.CourseColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ReviewTutorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tutor.ReviewTutorTable,
			Columns: []string{tutor.ReviewTutorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reviewtutor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tutor.ScheduleTable,
			Columns: []string{tutor.ScheduleColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tutor.UserTable,
			Columns: []string{tutor.UserColumn},
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
		_node.user_tutor = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TutorCreateBulk is the builder for creating many Tutor entities in bulk.
type TutorCreateBulk struct {
	config
	builders []*TutorCreate
}

// Save creates the Tutor entities in the database.
func (tcb *TutorCreateBulk) Save(ctx context.Context) ([]*Tutor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tutor, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TutorMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TutorCreateBulk) SaveX(ctx context.Context) []*Tutor {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TutorCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TutorCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
