// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/review"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/google/uuid"
)

// CourseUpdate is the builder for updating Course entities.
type CourseUpdate struct {
	config
	hooks    []Hook
	mutation *CourseMutation
}

// Where appends a list predicates to the CourseUpdate builder.
func (cu *CourseUpdate) Where(ps ...predicate.Course) *CourseUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *CourseUpdate) SetTitle(s string) *CourseUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetSubject sets the "subject" field.
func (cu *CourseUpdate) SetSubject(s string) *CourseUpdate {
	cu.mutation.SetSubject(s)
	return cu
}

// SetTopic sets the "topic" field.
func (cu *CourseUpdate) SetTopic(s string) *CourseUpdate {
	cu.mutation.SetTopic(s)
	return cu
}

// SetEstimatedTime sets the "estimated_time" field.
func (cu *CourseUpdate) SetEstimatedTime(i int) *CourseUpdate {
	cu.mutation.ResetEstimatedTime()
	cu.mutation.SetEstimatedTime(i)
	return cu
}

// AddEstimatedTime adds i to the "estimated_time" field.
func (cu *CourseUpdate) AddEstimatedTime(i int) *CourseUpdate {
	cu.mutation.AddEstimatedTime(i)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CourseUpdate) SetDescription(s string) *CourseUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetPricePerHour sets the "price_per_hour" field.
func (cu *CourseUpdate) SetPricePerHour(i int) *CourseUpdate {
	cu.mutation.ResetPricePerHour()
	cu.mutation.SetPricePerHour(i)
	return cu
}

// AddPricePerHour adds i to the "price_per_hour" field.
func (cu *CourseUpdate) AddPricePerHour(i int) *CourseUpdate {
	cu.mutation.AddPricePerHour(i)
	return cu
}

// SetLevel sets the "level" field.
func (cu *CourseUpdate) SetLevel(c course.Level) *CourseUpdate {
	cu.mutation.SetLevel(c)
	return cu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableLevel(c *course.Level) *CourseUpdate {
	if c != nil {
		cu.SetLevel(*c)
	}
	return cu
}

// ClearLevel clears the value of the "level" field.
func (cu *CourseUpdate) ClearLevel() *CourseUpdate {
	cu.mutation.ClearLevel()
	return cu
}

// SetCoursePictureURL sets the "course_picture_url" field.
func (cu *CourseUpdate) SetCoursePictureURL(s string) *CourseUpdate {
	cu.mutation.SetCoursePictureURL(s)
	return cu
}

// SetNillableCoursePictureURL sets the "course_picture_url" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableCoursePictureURL(s *string) *CourseUpdate {
	if s != nil {
		cu.SetCoursePictureURL(*s)
	}
	return cu
}

// ClearCoursePictureURL clears the value of the "course_picture_url" field.
func (cu *CourseUpdate) ClearCoursePictureURL() *CourseUpdate {
	cu.mutation.ClearCoursePictureURL()
	return cu
}

// AddReviewIDs adds the "review" edge to the Review entity by IDs.
func (cu *CourseUpdate) AddReviewIDs(ids ...int) *CourseUpdate {
	cu.mutation.AddReviewIDs(ids...)
	return cu
}

// AddReview adds the "review" edges to the Review entity.
func (cu *CourseUpdate) AddReview(r ...*Review) *CourseUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddReviewIDs(ids...)
}

// AddMatchIDs adds the "match" edge to the Match entity by IDs.
func (cu *CourseUpdate) AddMatchIDs(ids ...uuid.UUID) *CourseUpdate {
	cu.mutation.AddMatchIDs(ids...)
	return cu
}

// AddMatch adds the "match" edges to the Match entity.
func (cu *CourseUpdate) AddMatch(m ...*Match) *CourseUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMatchIDs(ids...)
}

// SetTutorID sets the "tutor" edge to the Tutor entity by ID.
func (cu *CourseUpdate) SetTutorID(id uuid.UUID) *CourseUpdate {
	cu.mutation.SetTutorID(id)
	return cu
}

// SetTutor sets the "tutor" edge to the Tutor entity.
func (cu *CourseUpdate) SetTutor(t *Tutor) *CourseUpdate {
	return cu.SetTutorID(t.ID)
}

// Mutation returns the CourseMutation object of the builder.
func (cu *CourseUpdate) Mutation() *CourseMutation {
	return cu.mutation
}

// ClearReview clears all "review" edges to the Review entity.
func (cu *CourseUpdate) ClearReview() *CourseUpdate {
	cu.mutation.ClearReview()
	return cu
}

// RemoveReviewIDs removes the "review" edge to Review entities by IDs.
func (cu *CourseUpdate) RemoveReviewIDs(ids ...int) *CourseUpdate {
	cu.mutation.RemoveReviewIDs(ids...)
	return cu
}

// RemoveReview removes "review" edges to Review entities.
func (cu *CourseUpdate) RemoveReview(r ...*Review) *CourseUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveReviewIDs(ids...)
}

// ClearMatch clears all "match" edges to the Match entity.
func (cu *CourseUpdate) ClearMatch() *CourseUpdate {
	cu.mutation.ClearMatch()
	return cu
}

// RemoveMatchIDs removes the "match" edge to Match entities by IDs.
func (cu *CourseUpdate) RemoveMatchIDs(ids ...uuid.UUID) *CourseUpdate {
	cu.mutation.RemoveMatchIDs(ids...)
	return cu
}

// RemoveMatch removes "match" edges to Match entities.
func (cu *CourseUpdate) RemoveMatch(m ...*Match) *CourseUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMatchIDs(ids...)
}

// ClearTutor clears the "tutor" edge to the Tutor entity.
func (cu *CourseUpdate) ClearTutor() *CourseUpdate {
	cu.mutation.ClearTutor()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CourseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CourseMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CourseUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CourseUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CourseUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CourseUpdate) check() error {
	if v, ok := cu.mutation.Title(); ok {
		if err := course.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Course.title": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Subject(); ok {
		if err := course.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "Course.subject": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Topic(); ok {
		if err := course.TopicValidator(v); err != nil {
			return &ValidationError{Name: "topic", err: fmt.Errorf(`ent: validator failed for field "Course.topic": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Description(); ok {
		if err := course.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Course.description": %w`, err)}
		}
	}
	if v, ok := cu.mutation.PricePerHour(); ok {
		if err := course.PricePerHourValidator(v); err != nil {
			return &ValidationError{Name: "price_per_hour", err: fmt.Errorf(`ent: validator failed for field "Course.price_per_hour": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Level(); ok {
		if err := course.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Course.level": %w`, err)}
		}
	}
	if _, ok := cu.mutation.TutorID(); cu.mutation.TutorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.tutor"`)
	}
	return nil
}

func (cu *CourseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(course.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.Subject(); ok {
		_spec.SetField(course.FieldSubject, field.TypeString, value)
	}
	if value, ok := cu.mutation.Topic(); ok {
		_spec.SetField(course.FieldTopic, field.TypeString, value)
	}
	if value, ok := cu.mutation.EstimatedTime(); ok {
		_spec.SetField(course.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedEstimatedTime(); ok {
		_spec.AddField(course.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(course.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.PricePerHour(); ok {
		_spec.SetField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedPricePerHour(); ok {
		_spec.AddField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Level(); ok {
		_spec.SetField(course.FieldLevel, field.TypeEnum, value)
	}
	if cu.mutation.LevelCleared() {
		_spec.ClearField(course.FieldLevel, field.TypeEnum)
	}
	if value, ok := cu.mutation.CoursePictureURL(); ok {
		_spec.SetField(course.FieldCoursePictureURL, field.TypeString, value)
	}
	if cu.mutation.CoursePictureURLCleared() {
		_spec.ClearField(course.FieldCoursePictureURL, field.TypeString)
	}
	if cu.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.ReviewTable,
			Columns: course.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedReviewIDs(); len(nodes) > 0 && !cu.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.ReviewTable,
			Columns: course.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ReviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.ReviewTable,
			Columns: course.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.MatchTable,
			Columns: course.MatchPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: match.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMatchIDs(); len(nodes) > 0 && !cu.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.MatchTable,
			Columns: course.MatchPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.MatchTable,
			Columns: course.MatchPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TutorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TutorTable,
			Columns: []string{course.TutorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tutor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TutorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TutorTable,
			Columns: []string{course.TutorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tutor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CourseUpdateOne is the builder for updating a single Course entity.
type CourseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseMutation
}

// SetTitle sets the "title" field.
func (cuo *CourseUpdateOne) SetTitle(s string) *CourseUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetSubject sets the "subject" field.
func (cuo *CourseUpdateOne) SetSubject(s string) *CourseUpdateOne {
	cuo.mutation.SetSubject(s)
	return cuo
}

// SetTopic sets the "topic" field.
func (cuo *CourseUpdateOne) SetTopic(s string) *CourseUpdateOne {
	cuo.mutation.SetTopic(s)
	return cuo
}

// SetEstimatedTime sets the "estimated_time" field.
func (cuo *CourseUpdateOne) SetEstimatedTime(i int) *CourseUpdateOne {
	cuo.mutation.ResetEstimatedTime()
	cuo.mutation.SetEstimatedTime(i)
	return cuo
}

// AddEstimatedTime adds i to the "estimated_time" field.
func (cuo *CourseUpdateOne) AddEstimatedTime(i int) *CourseUpdateOne {
	cuo.mutation.AddEstimatedTime(i)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CourseUpdateOne) SetDescription(s string) *CourseUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetPricePerHour sets the "price_per_hour" field.
func (cuo *CourseUpdateOne) SetPricePerHour(i int) *CourseUpdateOne {
	cuo.mutation.ResetPricePerHour()
	cuo.mutation.SetPricePerHour(i)
	return cuo
}

// AddPricePerHour adds i to the "price_per_hour" field.
func (cuo *CourseUpdateOne) AddPricePerHour(i int) *CourseUpdateOne {
	cuo.mutation.AddPricePerHour(i)
	return cuo
}

// SetLevel sets the "level" field.
func (cuo *CourseUpdateOne) SetLevel(c course.Level) *CourseUpdateOne {
	cuo.mutation.SetLevel(c)
	return cuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableLevel(c *course.Level) *CourseUpdateOne {
	if c != nil {
		cuo.SetLevel(*c)
	}
	return cuo
}

// ClearLevel clears the value of the "level" field.
func (cuo *CourseUpdateOne) ClearLevel() *CourseUpdateOne {
	cuo.mutation.ClearLevel()
	return cuo
}

// SetCoursePictureURL sets the "course_picture_url" field.
func (cuo *CourseUpdateOne) SetCoursePictureURL(s string) *CourseUpdateOne {
	cuo.mutation.SetCoursePictureURL(s)
	return cuo
}

// SetNillableCoursePictureURL sets the "course_picture_url" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableCoursePictureURL(s *string) *CourseUpdateOne {
	if s != nil {
		cuo.SetCoursePictureURL(*s)
	}
	return cuo
}

// ClearCoursePictureURL clears the value of the "course_picture_url" field.
func (cuo *CourseUpdateOne) ClearCoursePictureURL() *CourseUpdateOne {
	cuo.mutation.ClearCoursePictureURL()
	return cuo
}

// AddReviewIDs adds the "review" edge to the Review entity by IDs.
func (cuo *CourseUpdateOne) AddReviewIDs(ids ...int) *CourseUpdateOne {
	cuo.mutation.AddReviewIDs(ids...)
	return cuo
}

// AddReview adds the "review" edges to the Review entity.
func (cuo *CourseUpdateOne) AddReview(r ...*Review) *CourseUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddReviewIDs(ids...)
}

// AddMatchIDs adds the "match" edge to the Match entity by IDs.
func (cuo *CourseUpdateOne) AddMatchIDs(ids ...uuid.UUID) *CourseUpdateOne {
	cuo.mutation.AddMatchIDs(ids...)
	return cuo
}

// AddMatch adds the "match" edges to the Match entity.
func (cuo *CourseUpdateOne) AddMatch(m ...*Match) *CourseUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMatchIDs(ids...)
}

// SetTutorID sets the "tutor" edge to the Tutor entity by ID.
func (cuo *CourseUpdateOne) SetTutorID(id uuid.UUID) *CourseUpdateOne {
	cuo.mutation.SetTutorID(id)
	return cuo
}

// SetTutor sets the "tutor" edge to the Tutor entity.
func (cuo *CourseUpdateOne) SetTutor(t *Tutor) *CourseUpdateOne {
	return cuo.SetTutorID(t.ID)
}

// Mutation returns the CourseMutation object of the builder.
func (cuo *CourseUpdateOne) Mutation() *CourseMutation {
	return cuo.mutation
}

// ClearReview clears all "review" edges to the Review entity.
func (cuo *CourseUpdateOne) ClearReview() *CourseUpdateOne {
	cuo.mutation.ClearReview()
	return cuo
}

// RemoveReviewIDs removes the "review" edge to Review entities by IDs.
func (cuo *CourseUpdateOne) RemoveReviewIDs(ids ...int) *CourseUpdateOne {
	cuo.mutation.RemoveReviewIDs(ids...)
	return cuo
}

// RemoveReview removes "review" edges to Review entities.
func (cuo *CourseUpdateOne) RemoveReview(r ...*Review) *CourseUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveReviewIDs(ids...)
}

// ClearMatch clears all "match" edges to the Match entity.
func (cuo *CourseUpdateOne) ClearMatch() *CourseUpdateOne {
	cuo.mutation.ClearMatch()
	return cuo
}

// RemoveMatchIDs removes the "match" edge to Match entities by IDs.
func (cuo *CourseUpdateOne) RemoveMatchIDs(ids ...uuid.UUID) *CourseUpdateOne {
	cuo.mutation.RemoveMatchIDs(ids...)
	return cuo
}

// RemoveMatch removes "match" edges to Match entities.
func (cuo *CourseUpdateOne) RemoveMatch(m ...*Match) *CourseUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMatchIDs(ids...)
}

// ClearTutor clears the "tutor" edge to the Tutor entity.
func (cuo *CourseUpdateOne) ClearTutor() *CourseUpdateOne {
	cuo.mutation.ClearTutor()
	return cuo
}

// Where appends a list predicates to the CourseUpdate builder.
func (cuo *CourseUpdateOne) Where(ps ...predicate.Course) *CourseUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CourseUpdateOne) Select(field string, fields ...string) *CourseUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Course entity.
func (cuo *CourseUpdateOne) Save(ctx context.Context) (*Course, error) {
	return withHooks[*Course, CourseMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CourseUpdateOne) SaveX(ctx context.Context) *Course {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CourseUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CourseUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CourseUpdateOne) check() error {
	if v, ok := cuo.mutation.Title(); ok {
		if err := course.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Course.title": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Subject(); ok {
		if err := course.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "Course.subject": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Topic(); ok {
		if err := course.TopicValidator(v); err != nil {
			return &ValidationError{Name: "topic", err: fmt.Errorf(`ent: validator failed for field "Course.topic": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Description(); ok {
		if err := course.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Course.description": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.PricePerHour(); ok {
		if err := course.PricePerHourValidator(v); err != nil {
			return &ValidationError{Name: "price_per_hour", err: fmt.Errorf(`ent: validator failed for field "Course.price_per_hour": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Level(); ok {
		if err := course.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Course.level": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.TutorID(); cuo.mutation.TutorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.tutor"`)
	}
	return nil
}

func (cuo *CourseUpdateOne) sqlSave(ctx context.Context) (_node *Course, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Course.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, course.FieldID)
		for _, f := range fields {
			if !course.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != course.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(course.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Subject(); ok {
		_spec.SetField(course.FieldSubject, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Topic(); ok {
		_spec.SetField(course.FieldTopic, field.TypeString, value)
	}
	if value, ok := cuo.mutation.EstimatedTime(); ok {
		_spec.SetField(course.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedEstimatedTime(); ok {
		_spec.AddField(course.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(course.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.PricePerHour(); ok {
		_spec.SetField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedPricePerHour(); ok {
		_spec.AddField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Level(); ok {
		_spec.SetField(course.FieldLevel, field.TypeEnum, value)
	}
	if cuo.mutation.LevelCleared() {
		_spec.ClearField(course.FieldLevel, field.TypeEnum)
	}
	if value, ok := cuo.mutation.CoursePictureURL(); ok {
		_spec.SetField(course.FieldCoursePictureURL, field.TypeString, value)
	}
	if cuo.mutation.CoursePictureURLCleared() {
		_spec.ClearField(course.FieldCoursePictureURL, field.TypeString)
	}
	if cuo.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.ReviewTable,
			Columns: course.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedReviewIDs(); len(nodes) > 0 && !cuo.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.ReviewTable,
			Columns: course.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ReviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.ReviewTable,
			Columns: course.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.MatchTable,
			Columns: course.MatchPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: match.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMatchIDs(); len(nodes) > 0 && !cuo.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.MatchTable,
			Columns: course.MatchPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   course.MatchTable,
			Columns: course.MatchPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TutorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TutorTable,
			Columns: []string{course.TutorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tutor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TutorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TutorTable,
			Columns: []string{course.TutorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tutor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Course{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
