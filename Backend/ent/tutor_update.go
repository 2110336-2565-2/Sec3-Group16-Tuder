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
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/issuereport"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewtutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// TutorUpdate is the builder for updating Tutor entities.
type TutorUpdate struct {
	config
	hooks    []Hook
	mutation *TutorMutation
}

// Where appends a list predicates to the TutorUpdate builder.
func (tu *TutorUpdate) Where(ps ...predicate.Tutor) *TutorUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TutorUpdate) SetDescription(s string) *TutorUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TutorUpdate) SetNillableDescription(s *string) *TutorUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TutorUpdate) ClearDescription() *TutorUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetOmiseBankToken sets the "omise_bank_token" field.
func (tu *TutorUpdate) SetOmiseBankToken(s string) *TutorUpdate {
	tu.mutation.SetOmiseBankToken(s)
	return tu
}

// SetNillableOmiseBankToken sets the "omise_bank_token" field if the given value is not nil.
func (tu *TutorUpdate) SetNillableOmiseBankToken(s *string) *TutorUpdate {
	if s != nil {
		tu.SetOmiseBankToken(*s)
	}
	return tu
}

// ClearOmiseBankToken clears the value of the "omise_bank_token" field.
func (tu *TutorUpdate) ClearOmiseBankToken() *TutorUpdate {
	tu.mutation.ClearOmiseBankToken()
	return tu
}

// SetCitizenID sets the "citizen_id" field.
func (tu *TutorUpdate) SetCitizenID(s string) *TutorUpdate {
	tu.mutation.SetCitizenID(s)
	return tu
}

// AddIssueReportIDs adds the "issue_report" edge to the IssueReport entity by IDs.
func (tu *TutorUpdate) AddIssueReportIDs(ids ...uuid.UUID) *TutorUpdate {
	tu.mutation.AddIssueReportIDs(ids...)
	return tu
}

// AddIssueReport adds the "issue_report" edges to the IssueReport entity.
func (tu *TutorUpdate) AddIssueReport(i ...*IssueReport) *TutorUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tu.AddIssueReportIDs(ids...)
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (tu *TutorUpdate) AddCourseIDs(ids ...uuid.UUID) *TutorUpdate {
	tu.mutation.AddCourseIDs(ids...)
	return tu
}

// AddCourse adds the "course" edges to the Course entity.
func (tu *TutorUpdate) AddCourse(c ...*Course) *TutorUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.AddCourseIDs(ids...)
}

// AddReviewTutorIDs adds the "review_tutor" edge to the ReviewTutor entity by IDs.
func (tu *TutorUpdate) AddReviewTutorIDs(ids ...int) *TutorUpdate {
	tu.mutation.AddReviewTutorIDs(ids...)
	return tu
}

// AddReviewTutor adds the "review_tutor" edges to the ReviewTutor entity.
func (tu *TutorUpdate) AddReviewTutor(r ...*ReviewTutor) *TutorUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddReviewTutorIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tu *TutorUpdate) SetUserID(id uuid.UUID) *TutorUpdate {
	tu.mutation.SetUserID(id)
	return tu
}

// SetUser sets the "user" edge to the User entity.
func (tu *TutorUpdate) SetUser(u *User) *TutorUpdate {
	return tu.SetUserID(u.ID)
}

// SetScheduleID sets the "schedule" edge to the Schedule entity by ID.
func (tu *TutorUpdate) SetScheduleID(id uuid.UUID) *TutorUpdate {
	tu.mutation.SetScheduleID(id)
	return tu
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (tu *TutorUpdate) SetSchedule(s *Schedule) *TutorUpdate {
	return tu.SetScheduleID(s.ID)
}

// Mutation returns the TutorMutation object of the builder.
func (tu *TutorUpdate) Mutation() *TutorMutation {
	return tu.mutation
}

// ClearIssueReport clears all "issue_report" edges to the IssueReport entity.
func (tu *TutorUpdate) ClearIssueReport() *TutorUpdate {
	tu.mutation.ClearIssueReport()
	return tu
}

// RemoveIssueReportIDs removes the "issue_report" edge to IssueReport entities by IDs.
func (tu *TutorUpdate) RemoveIssueReportIDs(ids ...uuid.UUID) *TutorUpdate {
	tu.mutation.RemoveIssueReportIDs(ids...)
	return tu
}

// RemoveIssueReport removes "issue_report" edges to IssueReport entities.
func (tu *TutorUpdate) RemoveIssueReport(i ...*IssueReport) *TutorUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tu.RemoveIssueReportIDs(ids...)
}

// ClearCourse clears all "course" edges to the Course entity.
func (tu *TutorUpdate) ClearCourse() *TutorUpdate {
	tu.mutation.ClearCourse()
	return tu
}

// RemoveCourseIDs removes the "course" edge to Course entities by IDs.
func (tu *TutorUpdate) RemoveCourseIDs(ids ...uuid.UUID) *TutorUpdate {
	tu.mutation.RemoveCourseIDs(ids...)
	return tu
}

// RemoveCourse removes "course" edges to Course entities.
func (tu *TutorUpdate) RemoveCourse(c ...*Course) *TutorUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.RemoveCourseIDs(ids...)
}

// ClearReviewTutor clears all "review_tutor" edges to the ReviewTutor entity.
func (tu *TutorUpdate) ClearReviewTutor() *TutorUpdate {
	tu.mutation.ClearReviewTutor()
	return tu
}

// RemoveReviewTutorIDs removes the "review_tutor" edge to ReviewTutor entities by IDs.
func (tu *TutorUpdate) RemoveReviewTutorIDs(ids ...int) *TutorUpdate {
	tu.mutation.RemoveReviewTutorIDs(ids...)
	return tu
}

// RemoveReviewTutor removes "review_tutor" edges to ReviewTutor entities.
func (tu *TutorUpdate) RemoveReviewTutor(r ...*ReviewTutor) *TutorUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveReviewTutorIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TutorUpdate) ClearUser() *TutorUpdate {
	tu.mutation.ClearUser()
	return tu
}

// ClearSchedule clears the "schedule" edge to the Schedule entity.
func (tu *TutorUpdate) ClearSchedule() *TutorUpdate {
	tu.mutation.ClearSchedule()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TutorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TutorMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TutorUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TutorUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TutorUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TutorUpdate) check() error {
	if v, ok := tu.mutation.CitizenID(); ok {
		if err := tutor.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Tutor.citizen_id": %w`, err)}
		}
	}
	if _, ok := tu.mutation.UserID(); tu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tutor.user"`)
	}
	if _, ok := tu.mutation.ScheduleID(); tu.mutation.ScheduleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tutor.schedule"`)
	}
	return nil
}

func (tu *TutorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tutor.Table, tutor.Columns, sqlgraph.NewFieldSpec(tutor.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(tutor.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(tutor.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.OmiseBankToken(); ok {
		_spec.SetField(tutor.FieldOmiseBankToken, field.TypeString, value)
	}
	if tu.mutation.OmiseBankTokenCleared() {
		_spec.ClearField(tutor.FieldOmiseBankToken, field.TypeString)
	}
	if value, ok := tu.mutation.CitizenID(); ok {
		_spec.SetField(tutor.FieldCitizenID, field.TypeString, value)
	}
	if tu.mutation.IssueReportCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedIssueReportIDs(); len(nodes) > 0 && !tu.mutation.IssueReportCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.IssueReportIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.CourseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedCourseIDs(); len(nodes) > 0 && !tu.mutation.CourseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CourseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ReviewTutorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReviewTutorIDs(); len(nodes) > 0 && !tu.mutation.ReviewTutorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReviewTutorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tutor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TutorUpdateOne is the builder for updating a single Tutor entity.
type TutorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TutorMutation
}

// SetDescription sets the "description" field.
func (tuo *TutorUpdateOne) SetDescription(s string) *TutorUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TutorUpdateOne) SetNillableDescription(s *string) *TutorUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TutorUpdateOne) ClearDescription() *TutorUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetOmiseBankToken sets the "omise_bank_token" field.
func (tuo *TutorUpdateOne) SetOmiseBankToken(s string) *TutorUpdateOne {
	tuo.mutation.SetOmiseBankToken(s)
	return tuo
}

// SetNillableOmiseBankToken sets the "omise_bank_token" field if the given value is not nil.
func (tuo *TutorUpdateOne) SetNillableOmiseBankToken(s *string) *TutorUpdateOne {
	if s != nil {
		tuo.SetOmiseBankToken(*s)
	}
	return tuo
}

// ClearOmiseBankToken clears the value of the "omise_bank_token" field.
func (tuo *TutorUpdateOne) ClearOmiseBankToken() *TutorUpdateOne {
	tuo.mutation.ClearOmiseBankToken()
	return tuo
}

// SetCitizenID sets the "citizen_id" field.
func (tuo *TutorUpdateOne) SetCitizenID(s string) *TutorUpdateOne {
	tuo.mutation.SetCitizenID(s)
	return tuo
}

// AddIssueReportIDs adds the "issue_report" edge to the IssueReport entity by IDs.
func (tuo *TutorUpdateOne) AddIssueReportIDs(ids ...uuid.UUID) *TutorUpdateOne {
	tuo.mutation.AddIssueReportIDs(ids...)
	return tuo
}

// AddIssueReport adds the "issue_report" edges to the IssueReport entity.
func (tuo *TutorUpdateOne) AddIssueReport(i ...*IssueReport) *TutorUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tuo.AddIssueReportIDs(ids...)
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (tuo *TutorUpdateOne) AddCourseIDs(ids ...uuid.UUID) *TutorUpdateOne {
	tuo.mutation.AddCourseIDs(ids...)
	return tuo
}

// AddCourse adds the "course" edges to the Course entity.
func (tuo *TutorUpdateOne) AddCourse(c ...*Course) *TutorUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.AddCourseIDs(ids...)
}

// AddReviewTutorIDs adds the "review_tutor" edge to the ReviewTutor entity by IDs.
func (tuo *TutorUpdateOne) AddReviewTutorIDs(ids ...int) *TutorUpdateOne {
	tuo.mutation.AddReviewTutorIDs(ids...)
	return tuo
}

// AddReviewTutor adds the "review_tutor" edges to the ReviewTutor entity.
func (tuo *TutorUpdateOne) AddReviewTutor(r ...*ReviewTutor) *TutorUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddReviewTutorIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tuo *TutorUpdateOne) SetUserID(id uuid.UUID) *TutorUpdateOne {
	tuo.mutation.SetUserID(id)
	return tuo
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TutorUpdateOne) SetUser(u *User) *TutorUpdateOne {
	return tuo.SetUserID(u.ID)
}

// SetScheduleID sets the "schedule" edge to the Schedule entity by ID.
func (tuo *TutorUpdateOne) SetScheduleID(id uuid.UUID) *TutorUpdateOne {
	tuo.mutation.SetScheduleID(id)
	return tuo
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (tuo *TutorUpdateOne) SetSchedule(s *Schedule) *TutorUpdateOne {
	return tuo.SetScheduleID(s.ID)
}

// Mutation returns the TutorMutation object of the builder.
func (tuo *TutorUpdateOne) Mutation() *TutorMutation {
	return tuo.mutation
}

// ClearIssueReport clears all "issue_report" edges to the IssueReport entity.
func (tuo *TutorUpdateOne) ClearIssueReport() *TutorUpdateOne {
	tuo.mutation.ClearIssueReport()
	return tuo
}

// RemoveIssueReportIDs removes the "issue_report" edge to IssueReport entities by IDs.
func (tuo *TutorUpdateOne) RemoveIssueReportIDs(ids ...uuid.UUID) *TutorUpdateOne {
	tuo.mutation.RemoveIssueReportIDs(ids...)
	return tuo
}

// RemoveIssueReport removes "issue_report" edges to IssueReport entities.
func (tuo *TutorUpdateOne) RemoveIssueReport(i ...*IssueReport) *TutorUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tuo.RemoveIssueReportIDs(ids...)
}

// ClearCourse clears all "course" edges to the Course entity.
func (tuo *TutorUpdateOne) ClearCourse() *TutorUpdateOne {
	tuo.mutation.ClearCourse()
	return tuo
}

// RemoveCourseIDs removes the "course" edge to Course entities by IDs.
func (tuo *TutorUpdateOne) RemoveCourseIDs(ids ...uuid.UUID) *TutorUpdateOne {
	tuo.mutation.RemoveCourseIDs(ids...)
	return tuo
}

// RemoveCourse removes "course" edges to Course entities.
func (tuo *TutorUpdateOne) RemoveCourse(c ...*Course) *TutorUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.RemoveCourseIDs(ids...)
}

// ClearReviewTutor clears all "review_tutor" edges to the ReviewTutor entity.
func (tuo *TutorUpdateOne) ClearReviewTutor() *TutorUpdateOne {
	tuo.mutation.ClearReviewTutor()
	return tuo
}

// RemoveReviewTutorIDs removes the "review_tutor" edge to ReviewTutor entities by IDs.
func (tuo *TutorUpdateOne) RemoveReviewTutorIDs(ids ...int) *TutorUpdateOne {
	tuo.mutation.RemoveReviewTutorIDs(ids...)
	return tuo
}

// RemoveReviewTutor removes "review_tutor" edges to ReviewTutor entities.
func (tuo *TutorUpdateOne) RemoveReviewTutor(r ...*ReviewTutor) *TutorUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveReviewTutorIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TutorUpdateOne) ClearUser() *TutorUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// ClearSchedule clears the "schedule" edge to the Schedule entity.
func (tuo *TutorUpdateOne) ClearSchedule() *TutorUpdateOne {
	tuo.mutation.ClearSchedule()
	return tuo
}

// Where appends a list predicates to the TutorUpdate builder.
func (tuo *TutorUpdateOne) Where(ps ...predicate.Tutor) *TutorUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TutorUpdateOne) Select(field string, fields ...string) *TutorUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tutor entity.
func (tuo *TutorUpdateOne) Save(ctx context.Context) (*Tutor, error) {
	return withHooks[*Tutor, TutorMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TutorUpdateOne) SaveX(ctx context.Context) *Tutor {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TutorUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TutorUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TutorUpdateOne) check() error {
	if v, ok := tuo.mutation.CitizenID(); ok {
		if err := tutor.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Tutor.citizen_id": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.UserID(); tuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tutor.user"`)
	}
	if _, ok := tuo.mutation.ScheduleID(); tuo.mutation.ScheduleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tutor.schedule"`)
	}
	return nil
}

func (tuo *TutorUpdateOne) sqlSave(ctx context.Context) (_node *Tutor, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tutor.Table, tutor.Columns, sqlgraph.NewFieldSpec(tutor.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tutor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tutor.FieldID)
		for _, f := range fields {
			if !tutor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tutor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(tutor.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(tutor.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.OmiseBankToken(); ok {
		_spec.SetField(tutor.FieldOmiseBankToken, field.TypeString, value)
	}
	if tuo.mutation.OmiseBankTokenCleared() {
		_spec.ClearField(tutor.FieldOmiseBankToken, field.TypeString)
	}
	if value, ok := tuo.mutation.CitizenID(); ok {
		_spec.SetField(tutor.FieldCitizenID, field.TypeString, value)
	}
	if tuo.mutation.IssueReportCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedIssueReportIDs(); len(nodes) > 0 && !tuo.mutation.IssueReportCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.IssueReportIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.CourseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedCourseIDs(); len(nodes) > 0 && !tuo.mutation.CourseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CourseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ReviewTutorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReviewTutorIDs(); len(nodes) > 0 && !tuo.mutation.ReviewTutorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReviewTutorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tutor{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tutor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
