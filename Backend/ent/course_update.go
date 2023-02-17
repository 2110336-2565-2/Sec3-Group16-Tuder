// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewcourse"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
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

// SetEstimatedTime sets the "estimated_time" field.
func (cu *CourseUpdate) SetEstimatedTime(t time.Time) *CourseUpdate {
	cu.mutation.SetEstimatedTime(t)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CourseUpdate) SetDescription(s string) *CourseUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetCourseStatus sets the "course_status" field.
func (cu *CourseUpdate) SetCourseStatus(s string) *CourseUpdate {
	cu.mutation.SetCourseStatus(s)
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

// SetLevelID sets the "level_id" field.
func (cu *CourseUpdate) SetLevelID(s string) *CourseUpdate {
	cu.mutation.SetLevelID(s)
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

// AddReviewCourseIDs adds the "review_course" edge to the ReviewCourse entity by IDs.
func (cu *CourseUpdate) AddReviewCourseIDs(ids ...int) *CourseUpdate {
	cu.mutation.AddReviewCourseIDs(ids...)
	return cu
}

// AddReviewCourse adds the "review_course" edges to the ReviewCourse entity.
func (cu *CourseUpdate) AddReviewCourse(r ...*ReviewCourse) *CourseUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddReviewCourseIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (cu *CourseUpdate) AddClasIDs(ids ...uuid.UUID) *CourseUpdate {
	cu.mutation.AddClasIDs(ids...)
	return cu
}

// AddClass adds the "class" edges to the Class entity.
func (cu *CourseUpdate) AddClass(c ...*Class) *CourseUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddClasIDs(ids...)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (cu *CourseUpdate) SetStudentID(id uuid.UUID) *CourseUpdate {
	cu.mutation.SetStudentID(id)
	return cu
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (cu *CourseUpdate) SetNillableStudentID(id *uuid.UUID) *CourseUpdate {
	if id != nil {
		cu = cu.SetStudentID(*id)
	}
	return cu
}

// SetStudent sets the "student" edge to the Student entity.
func (cu *CourseUpdate) SetStudent(s *Student) *CourseUpdate {
	return cu.SetStudentID(s.ID)
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

// ClearReviewCourse clears all "review_course" edges to the ReviewCourse entity.
func (cu *CourseUpdate) ClearReviewCourse() *CourseUpdate {
	cu.mutation.ClearReviewCourse()
	return cu
}

// RemoveReviewCourseIDs removes the "review_course" edge to ReviewCourse entities by IDs.
func (cu *CourseUpdate) RemoveReviewCourseIDs(ids ...int) *CourseUpdate {
	cu.mutation.RemoveReviewCourseIDs(ids...)
	return cu
}

// RemoveReviewCourse removes "review_course" edges to ReviewCourse entities.
func (cu *CourseUpdate) RemoveReviewCourse(r ...*ReviewCourse) *CourseUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveReviewCourseIDs(ids...)
}

// ClearClass clears all "class" edges to the Class entity.
func (cu *CourseUpdate) ClearClass() *CourseUpdate {
	cu.mutation.ClearClass()
	return cu
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (cu *CourseUpdate) RemoveClasIDs(ids ...uuid.UUID) *CourseUpdate {
	cu.mutation.RemoveClasIDs(ids...)
	return cu
}

// RemoveClass removes "class" edges to Class entities.
func (cu *CourseUpdate) RemoveClass(c ...*Class) *CourseUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveClasIDs(ids...)
}

// ClearStudent clears the "student" edge to the Student entity.
func (cu *CourseUpdate) ClearStudent() *CourseUpdate {
	cu.mutation.ClearStudent()
	return cu
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
	if v, ok := cu.mutation.Description(); ok {
		if err := course.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Course.description": %w`, err)}
		}
	}
	if v, ok := cu.mutation.CourseStatus(); ok {
		if err := course.CourseStatusValidator(v); err != nil {
			return &ValidationError{Name: "course_status", err: fmt.Errorf(`ent: validator failed for field "Course.course_status": %w`, err)}
		}
	}
	if v, ok := cu.mutation.PricePerHour(); ok {
		if err := course.PricePerHourValidator(v); err != nil {
			return &ValidationError{Name: "price_per_hour", err: fmt.Errorf(`ent: validator failed for field "Course.price_per_hour": %w`, err)}
		}
	}
	if v, ok := cu.mutation.LevelID(); ok {
		if err := course.LevelIDValidator(v); err != nil {
			return &ValidationError{Name: "level_id", err: fmt.Errorf(`ent: validator failed for field "Course.level_id": %w`, err)}
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
	if value, ok := cu.mutation.EstimatedTime(); ok {
		_spec.SetField(course.FieldEstimatedTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(course.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.CourseStatus(); ok {
		_spec.SetField(course.FieldCourseStatus, field.TypeString, value)
	}
	if value, ok := cu.mutation.PricePerHour(); ok {
		_spec.SetField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedPricePerHour(); ok {
		_spec.AddField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cu.mutation.LevelID(); ok {
		_spec.SetField(course.FieldLevelID, field.TypeString, value)
	}
	if value, ok := cu.mutation.CoursePictureURL(); ok {
		_spec.SetField(course.FieldCoursePictureURL, field.TypeString, value)
	}
	if cu.mutation.CoursePictureURLCleared() {
		_spec.ClearField(course.FieldCoursePictureURL, field.TypeString)
	}
	if cu.mutation.ReviewCourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ReviewCourseTable,
			Columns: []string{course.ReviewCourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reviewcourse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedReviewCourseIDs(); len(nodes) > 0 && !cu.mutation.ReviewCourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ReviewCourseTable,
			Columns: []string{course.ReviewCourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reviewcourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ReviewCourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ReviewCourseTable,
			Columns: []string{course.ReviewCourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reviewcourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ClassTable,
			Columns: []string{course.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClassIDs(); len(nodes) > 0 && !cu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ClassTable,
			Columns: []string{course.ClassColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ClassTable,
			Columns: []string{course.ClassColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.StudentTable,
			Columns: []string{course.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.StudentTable,
			Columns: []string{course.StudentColumn},
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

// SetEstimatedTime sets the "estimated_time" field.
func (cuo *CourseUpdateOne) SetEstimatedTime(t time.Time) *CourseUpdateOne {
	cuo.mutation.SetEstimatedTime(t)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CourseUpdateOne) SetDescription(s string) *CourseUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetCourseStatus sets the "course_status" field.
func (cuo *CourseUpdateOne) SetCourseStatus(s string) *CourseUpdateOne {
	cuo.mutation.SetCourseStatus(s)
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

// SetLevelID sets the "level_id" field.
func (cuo *CourseUpdateOne) SetLevelID(s string) *CourseUpdateOne {
	cuo.mutation.SetLevelID(s)
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

// AddReviewCourseIDs adds the "review_course" edge to the ReviewCourse entity by IDs.
func (cuo *CourseUpdateOne) AddReviewCourseIDs(ids ...int) *CourseUpdateOne {
	cuo.mutation.AddReviewCourseIDs(ids...)
	return cuo
}

// AddReviewCourse adds the "review_course" edges to the ReviewCourse entity.
func (cuo *CourseUpdateOne) AddReviewCourse(r ...*ReviewCourse) *CourseUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddReviewCourseIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (cuo *CourseUpdateOne) AddClasIDs(ids ...uuid.UUID) *CourseUpdateOne {
	cuo.mutation.AddClasIDs(ids...)
	return cuo
}

// AddClass adds the "class" edges to the Class entity.
func (cuo *CourseUpdateOne) AddClass(c ...*Class) *CourseUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddClasIDs(ids...)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (cuo *CourseUpdateOne) SetStudentID(id uuid.UUID) *CourseUpdateOne {
	cuo.mutation.SetStudentID(id)
	return cuo
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableStudentID(id *uuid.UUID) *CourseUpdateOne {
	if id != nil {
		cuo = cuo.SetStudentID(*id)
	}
	return cuo
}

// SetStudent sets the "student" edge to the Student entity.
func (cuo *CourseUpdateOne) SetStudent(s *Student) *CourseUpdateOne {
	return cuo.SetStudentID(s.ID)
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

// ClearReviewCourse clears all "review_course" edges to the ReviewCourse entity.
func (cuo *CourseUpdateOne) ClearReviewCourse() *CourseUpdateOne {
	cuo.mutation.ClearReviewCourse()
	return cuo
}

// RemoveReviewCourseIDs removes the "review_course" edge to ReviewCourse entities by IDs.
func (cuo *CourseUpdateOne) RemoveReviewCourseIDs(ids ...int) *CourseUpdateOne {
	cuo.mutation.RemoveReviewCourseIDs(ids...)
	return cuo
}

// RemoveReviewCourse removes "review_course" edges to ReviewCourse entities.
func (cuo *CourseUpdateOne) RemoveReviewCourse(r ...*ReviewCourse) *CourseUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveReviewCourseIDs(ids...)
}

// ClearClass clears all "class" edges to the Class entity.
func (cuo *CourseUpdateOne) ClearClass() *CourseUpdateOne {
	cuo.mutation.ClearClass()
	return cuo
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (cuo *CourseUpdateOne) RemoveClasIDs(ids ...uuid.UUID) *CourseUpdateOne {
	cuo.mutation.RemoveClasIDs(ids...)
	return cuo
}

// RemoveClass removes "class" edges to Class entities.
func (cuo *CourseUpdateOne) RemoveClass(c ...*Class) *CourseUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveClasIDs(ids...)
}

// ClearStudent clears the "student" edge to the Student entity.
func (cuo *CourseUpdateOne) ClearStudent() *CourseUpdateOne {
	cuo.mutation.ClearStudent()
	return cuo
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
	if v, ok := cuo.mutation.Description(); ok {
		if err := course.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Course.description": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.CourseStatus(); ok {
		if err := course.CourseStatusValidator(v); err != nil {
			return &ValidationError{Name: "course_status", err: fmt.Errorf(`ent: validator failed for field "Course.course_status": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.PricePerHour(); ok {
		if err := course.PricePerHourValidator(v); err != nil {
			return &ValidationError{Name: "price_per_hour", err: fmt.Errorf(`ent: validator failed for field "Course.price_per_hour": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.LevelID(); ok {
		if err := course.LevelIDValidator(v); err != nil {
			return &ValidationError{Name: "level_id", err: fmt.Errorf(`ent: validator failed for field "Course.level_id": %w`, err)}
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
	if value, ok := cuo.mutation.EstimatedTime(); ok {
		_spec.SetField(course.FieldEstimatedTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(course.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CourseStatus(); ok {
		_spec.SetField(course.FieldCourseStatus, field.TypeString, value)
	}
	if value, ok := cuo.mutation.PricePerHour(); ok {
		_spec.SetField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedPricePerHour(); ok {
		_spec.AddField(course.FieldPricePerHour, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.LevelID(); ok {
		_spec.SetField(course.FieldLevelID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CoursePictureURL(); ok {
		_spec.SetField(course.FieldCoursePictureURL, field.TypeString, value)
	}
	if cuo.mutation.CoursePictureURLCleared() {
		_spec.ClearField(course.FieldCoursePictureURL, field.TypeString)
	}
	if cuo.mutation.ReviewCourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ReviewCourseTable,
			Columns: []string{course.ReviewCourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reviewcourse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedReviewCourseIDs(); len(nodes) > 0 && !cuo.mutation.ReviewCourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ReviewCourseTable,
			Columns: []string{course.ReviewCourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reviewcourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ReviewCourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ReviewCourseTable,
			Columns: []string{course.ReviewCourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reviewcourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ClassTable,
			Columns: []string{course.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClassIDs(); len(nodes) > 0 && !cuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ClassTable,
			Columns: []string{course.ClassColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   course.ClassTable,
			Columns: []string{course.ClassColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.StudentTable,
			Columns: []string{course.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.StudentTable,
			Columns: []string{course.StudentColumn},
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
