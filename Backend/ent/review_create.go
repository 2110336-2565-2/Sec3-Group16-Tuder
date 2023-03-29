// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/review"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/google/uuid"
)

// ReviewCreate is the builder for creating a Review entity.
type ReviewCreate struct {
	config
	mutation *ReviewMutation
	hooks    []Hook
}

// SetScore sets the "score" field.
func (rc *ReviewCreate) SetScore(f float32) *ReviewCreate {
	rc.mutation.SetScore(f)
	return rc
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableScore(f *float32) *ReviewCreate {
	if f != nil {
		rc.SetScore(*f)
	}
	return rc
}

// SetReviewMsg sets the "review_msg" field.
func (rc *ReviewCreate) SetReviewMsg(s string) *ReviewCreate {
	rc.mutation.SetReviewMsg(s)
	return rc
}

// SetNillableReviewMsg sets the "review_msg" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableReviewMsg(s *string) *ReviewCreate {
	if s != nil {
		rc.SetReviewMsg(*s)
	}
	return rc
}

// SetReviewTimeAt sets the "review_time_at" field.
func (rc *ReviewCreate) SetReviewTimeAt(t time.Time) *ReviewCreate {
	rc.mutation.SetReviewTimeAt(t)
	return rc
}

// SetNillableReviewTimeAt sets the "review_time_at" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableReviewTimeAt(t *time.Time) *ReviewCreate {
	if t != nil {
		rc.SetReviewTimeAt(*t)
	}
	return rc
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (rc *ReviewCreate) AddCourseIDs(ids ...uuid.UUID) *ReviewCreate {
	rc.mutation.AddCourseIDs(ids...)
	return rc
}

// AddCourse adds the "course" edges to the Course entity.
func (rc *ReviewCreate) AddCourse(c ...*Course) *ReviewCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddCourseIDs(ids...)
}

// AddStudentIDs adds the "student" edge to the Student entity by IDs.
func (rc *ReviewCreate) AddStudentIDs(ids ...uuid.UUID) *ReviewCreate {
	rc.mutation.AddStudentIDs(ids...)
	return rc
}

// AddStudent adds the "student" edges to the Student entity.
func (rc *ReviewCreate) AddStudent(s ...*Student) *ReviewCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddStudentIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (rc *ReviewCreate) Mutation() *ReviewMutation {
	return rc.mutation
}

// Save creates the Review in the database.
func (rc *ReviewCreate) Save(ctx context.Context) (*Review, error) {
	rc.defaults()
	return withHooks[*Review, ReviewMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReviewCreate) SaveX(ctx context.Context) *Review {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReviewCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReviewCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReviewCreate) defaults() {
	if _, ok := rc.mutation.ReviewTimeAt(); !ok {
		v := review.DefaultReviewTimeAt()
		rc.mutation.SetReviewTimeAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReviewCreate) check() error {
	if v, ok := rc.mutation.Score(); ok {
		if err := review.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Review.score": %w`, err)}
		}
	}
	if _, ok := rc.mutation.ReviewTimeAt(); !ok {
		return &ValidationError{Name: "review_time_at", err: errors.New(`ent: missing required field "Review.review_time_at"`)}
	}
	if len(rc.mutation.CourseIDs()) == 0 {
		return &ValidationError{Name: "course", err: errors.New(`ent: missing required edge "Review.course"`)}
	}
	if len(rc.mutation.StudentIDs()) == 0 {
		return &ValidationError{Name: "student", err: errors.New(`ent: missing required edge "Review.student"`)}
	}
	return nil
}

func (rc *ReviewCreate) sqlSave(ctx context.Context) (*Review, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReviewCreate) createSpec() (*Review, *sqlgraph.CreateSpec) {
	var (
		_node = &Review{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(review.Table, sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Score(); ok {
		_spec.SetField(review.FieldScore, field.TypeFloat32, value)
		_node.Score = &value
	}
	if value, ok := rc.mutation.ReviewMsg(); ok {
		_spec.SetField(review.FieldReviewMsg, field.TypeString, value)
		_node.ReviewMsg = &value
	}
	if value, ok := rc.mutation.ReviewTimeAt(); ok {
		_spec.SetField(review.FieldReviewTimeAt, field.TypeTime, value)
		_node.ReviewTimeAt = value
	}
	if nodes := rc.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   review.CourseTable,
			Columns: review.CoursePrimaryKey,
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
	if nodes := rc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   review.StudentTable,
			Columns: review.StudentPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReviewCreateBulk is the builder for creating many Review entities in bulk.
type ReviewCreateBulk struct {
	config
	builders []*ReviewCreate
}

// Save creates the Review entities in the database.
func (rcb *ReviewCreateBulk) Save(ctx context.Context) ([]*Review, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Review, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReviewMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReviewCreateBulk) SaveX(ctx context.Context) []*Review {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReviewCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReviewCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
