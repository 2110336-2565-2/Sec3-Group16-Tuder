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
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/classcancelrequest"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/payment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetAddress sets the "address" field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.mutation.SetAddress(s)
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetBirthDate sets the "birth_date" field.
func (uu *UserUpdate) SetBirthDate(t time.Time) *UserUpdate {
	uu.mutation.SetBirthDate(t)
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(s string) *UserUpdate {
	uu.mutation.SetGender(s)
	return uu
}

// SetProfilePictureURL sets the "profile_picture_URL" field.
func (uu *UserUpdate) SetProfilePictureURL(s string) *UserUpdate {
	uu.mutation.SetProfilePictureURL(s)
	return uu
}

// SetNillableProfilePictureURL sets the "profile_picture_URL" field if the given value is not nil.
func (uu *UserUpdate) SetNillableProfilePictureURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetProfilePictureURL(*s)
	}
	return uu
}

// ClearProfilePictureURL clears the value of the "profile_picture_URL" field.
func (uu *UserUpdate) ClearProfilePictureURL() *UserUpdate {
	uu.mutation.ClearProfilePictureURL()
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(u user.Role) *UserUpdate {
	uu.mutation.SetRole(u)
	return uu
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (uu *UserUpdate) SetStudentID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetStudentID(id)
	return uu
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableStudentID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetStudentID(*id)
	}
	return uu
}

// SetStudent sets the "student" edge to the Student entity.
func (uu *UserUpdate) SetStudent(s *Student) *UserUpdate {
	return uu.SetStudentID(s.ID)
}

// SetTutorID sets the "tutor" edge to the Tutor entity by ID.
func (uu *UserUpdate) SetTutorID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetTutorID(id)
	return uu
}

// SetNillableTutorID sets the "tutor" edge to the Tutor entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableTutorID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetTutorID(*id)
	}
	return uu
}

// SetTutor sets the "tutor" edge to the Tutor entity.
func (uu *UserUpdate) SetTutor(t *Tutor) *UserUpdate {
	return uu.SetTutorID(t.ID)
}

// AddPaymentIDs adds the "payment" edge to the Payment entity by IDs.
func (uu *UserUpdate) AddPaymentIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddPaymentIDs(ids...)
	return uu
}

// AddPayment adds the "payment" edges to the Payment entity.
func (uu *UserUpdate) AddPayment(p ...*Payment) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPaymentIDs(ids...)
}

// AddPaymentHistoryIDs adds the "payment_history" edge to the PaymentHistory entity by IDs.
func (uu *UserUpdate) AddPaymentHistoryIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddPaymentHistoryIDs(ids...)
	return uu
}

// AddPaymentHistory adds the "payment_history" edges to the PaymentHistory entity.
func (uu *UserUpdate) AddPaymentHistory(p ...*PaymentHistory) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPaymentHistoryIDs(ids...)
}

// AddClassCancelRequestIDs adds the "class_cancel_request" edge to the ClassCancelRequest entity by IDs.
func (uu *UserUpdate) AddClassCancelRequestIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddClassCancelRequestIDs(ids...)
	return uu
}

// AddClassCancelRequest adds the "class_cancel_request" edges to the ClassCancelRequest entity.
func (uu *UserUpdate) AddClassCancelRequest(c ...*ClassCancelRequest) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddClassCancelRequestIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (uu *UserUpdate) ClearStudent() *UserUpdate {
	uu.mutation.ClearStudent()
	return uu
}

// ClearTutor clears the "tutor" edge to the Tutor entity.
func (uu *UserUpdate) ClearTutor() *UserUpdate {
	uu.mutation.ClearTutor()
	return uu
}

// ClearPayment clears all "payment" edges to the Payment entity.
func (uu *UserUpdate) ClearPayment() *UserUpdate {
	uu.mutation.ClearPayment()
	return uu
}

// RemovePaymentIDs removes the "payment" edge to Payment entities by IDs.
func (uu *UserUpdate) RemovePaymentIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemovePaymentIDs(ids...)
	return uu
}

// RemovePayment removes "payment" edges to Payment entities.
func (uu *UserUpdate) RemovePayment(p ...*Payment) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePaymentIDs(ids...)
}

// ClearPaymentHistory clears all "payment_history" edges to the PaymentHistory entity.
func (uu *UserUpdate) ClearPaymentHistory() *UserUpdate {
	uu.mutation.ClearPaymentHistory()
	return uu
}

// RemovePaymentHistoryIDs removes the "payment_history" edge to PaymentHistory entities by IDs.
func (uu *UserUpdate) RemovePaymentHistoryIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemovePaymentHistoryIDs(ids...)
	return uu
}

// RemovePaymentHistory removes "payment_history" edges to PaymentHistory entities.
func (uu *UserUpdate) RemovePaymentHistory(p ...*PaymentHistory) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePaymentHistoryIDs(ids...)
}

// ClearClassCancelRequest clears all "class_cancel_request" edges to the ClassCancelRequest entity.
func (uu *UserUpdate) ClearClassCancelRequest() *UserUpdate {
	uu.mutation.ClearClassCancelRequest()
	return uu
}

// RemoveClassCancelRequestIDs removes the "class_cancel_request" edge to ClassCancelRequest entities by IDs.
func (uu *UserUpdate) RemoveClassCancelRequestIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveClassCancelRequestIDs(ids...)
	return uu
}

// RemoveClassCancelRequest removes "class_cancel_request" edges to ClassCancelRequest entities.
func (uu *UserUpdate) RemoveClassCancelRequest(c ...*ClassCancelRequest) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveClassCancelRequestIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Address(); ok {
		if err := user.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "User.address": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uu.mutation.BirthDate(); ok {
		_spec.SetField(user.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeString, value)
	}
	if value, ok := uu.mutation.ProfilePictureURL(); ok {
		_spec.SetField(user.FieldProfilePictureURL, field.TypeString, value)
	}
	if uu.mutation.ProfilePictureURLCleared() {
		_spec.ClearField(user.FieldProfilePictureURL, field.TypeString)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if uu.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.StudentTable,
			Columns: []string{user.StudentColumn},
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
	if nodes := uu.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.StudentTable,
			Columns: []string{user.StudentColumn},
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
	if uu.mutation.TutorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TutorTable,
			Columns: []string{user.TutorColumn},
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
	if nodes := uu.mutation.TutorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TutorTable,
			Columns: []string{user.TutorColumn},
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
	if uu.mutation.PaymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentTable,
			Columns: []string{user.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: payment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPaymentIDs(); len(nodes) > 0 && !uu.mutation.PaymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentTable,
			Columns: []string{user.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentTable,
			Columns: []string{user.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentHistoryTable,
			Columns: []string{user.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: paymenthistory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPaymentHistoryIDs(); len(nodes) > 0 && !uu.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentHistoryTable,
			Columns: []string{user.PaymentHistoryColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PaymentHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentHistoryTable,
			Columns: []string{user.PaymentHistoryColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ClassCancelRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClassCancelRequestTable,
			Columns: []string{user.ClassCancelRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: classcancelrequest.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedClassCancelRequestIDs(); len(nodes) > 0 && !uu.mutation.ClassCancelRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClassCancelRequestTable,
			Columns: []string{user.ClassCancelRequestColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ClassCancelRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClassCancelRequestTable,
			Columns: []string{user.ClassCancelRequestColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetAddress sets the "address" field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.mutation.SetAddress(s)
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetBirthDate sets the "birth_date" field.
func (uuo *UserUpdateOne) SetBirthDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetBirthDate(t)
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(s string) *UserUpdateOne {
	uuo.mutation.SetGender(s)
	return uuo
}

// SetProfilePictureURL sets the "profile_picture_URL" field.
func (uuo *UserUpdateOne) SetProfilePictureURL(s string) *UserUpdateOne {
	uuo.mutation.SetProfilePictureURL(s)
	return uuo
}

// SetNillableProfilePictureURL sets the "profile_picture_URL" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProfilePictureURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetProfilePictureURL(*s)
	}
	return uuo
}

// ClearProfilePictureURL clears the value of the "profile_picture_URL" field.
func (uuo *UserUpdateOne) ClearProfilePictureURL() *UserUpdateOne {
	uuo.mutation.ClearProfilePictureURL()
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(u user.Role) *UserUpdateOne {
	uuo.mutation.SetRole(u)
	return uuo
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (uuo *UserUpdateOne) SetStudentID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetStudentID(id)
	return uuo
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStudentID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetStudentID(*id)
	}
	return uuo
}

// SetStudent sets the "student" edge to the Student entity.
func (uuo *UserUpdateOne) SetStudent(s *Student) *UserUpdateOne {
	return uuo.SetStudentID(s.ID)
}

// SetTutorID sets the "tutor" edge to the Tutor entity by ID.
func (uuo *UserUpdateOne) SetTutorID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetTutorID(id)
	return uuo
}

// SetNillableTutorID sets the "tutor" edge to the Tutor entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTutorID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetTutorID(*id)
	}
	return uuo
}

// SetTutor sets the "tutor" edge to the Tutor entity.
func (uuo *UserUpdateOne) SetTutor(t *Tutor) *UserUpdateOne {
	return uuo.SetTutorID(t.ID)
}

// AddPaymentIDs adds the "payment" edge to the Payment entity by IDs.
func (uuo *UserUpdateOne) AddPaymentIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddPaymentIDs(ids...)
	return uuo
}

// AddPayment adds the "payment" edges to the Payment entity.
func (uuo *UserUpdateOne) AddPayment(p ...*Payment) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPaymentIDs(ids...)
}

// AddPaymentHistoryIDs adds the "payment_history" edge to the PaymentHistory entity by IDs.
func (uuo *UserUpdateOne) AddPaymentHistoryIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddPaymentHistoryIDs(ids...)
	return uuo
}

// AddPaymentHistory adds the "payment_history" edges to the PaymentHistory entity.
func (uuo *UserUpdateOne) AddPaymentHistory(p ...*PaymentHistory) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPaymentHistoryIDs(ids...)
}

// AddClassCancelRequestIDs adds the "class_cancel_request" edge to the ClassCancelRequest entity by IDs.
func (uuo *UserUpdateOne) AddClassCancelRequestIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddClassCancelRequestIDs(ids...)
	return uuo
}

// AddClassCancelRequest adds the "class_cancel_request" edges to the ClassCancelRequest entity.
func (uuo *UserUpdateOne) AddClassCancelRequest(c ...*ClassCancelRequest) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddClassCancelRequestIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (uuo *UserUpdateOne) ClearStudent() *UserUpdateOne {
	uuo.mutation.ClearStudent()
	return uuo
}

// ClearTutor clears the "tutor" edge to the Tutor entity.
func (uuo *UserUpdateOne) ClearTutor() *UserUpdateOne {
	uuo.mutation.ClearTutor()
	return uuo
}

// ClearPayment clears all "payment" edges to the Payment entity.
func (uuo *UserUpdateOne) ClearPayment() *UserUpdateOne {
	uuo.mutation.ClearPayment()
	return uuo
}

// RemovePaymentIDs removes the "payment" edge to Payment entities by IDs.
func (uuo *UserUpdateOne) RemovePaymentIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemovePaymentIDs(ids...)
	return uuo
}

// RemovePayment removes "payment" edges to Payment entities.
func (uuo *UserUpdateOne) RemovePayment(p ...*Payment) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePaymentIDs(ids...)
}

// ClearPaymentHistory clears all "payment_history" edges to the PaymentHistory entity.
func (uuo *UserUpdateOne) ClearPaymentHistory() *UserUpdateOne {
	uuo.mutation.ClearPaymentHistory()
	return uuo
}

// RemovePaymentHistoryIDs removes the "payment_history" edge to PaymentHistory entities by IDs.
func (uuo *UserUpdateOne) RemovePaymentHistoryIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemovePaymentHistoryIDs(ids...)
	return uuo
}

// RemovePaymentHistory removes "payment_history" edges to PaymentHistory entities.
func (uuo *UserUpdateOne) RemovePaymentHistory(p ...*PaymentHistory) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePaymentHistoryIDs(ids...)
}

// ClearClassCancelRequest clears all "class_cancel_request" edges to the ClassCancelRequest entity.
func (uuo *UserUpdateOne) ClearClassCancelRequest() *UserUpdateOne {
	uuo.mutation.ClearClassCancelRequest()
	return uuo
}

// RemoveClassCancelRequestIDs removes the "class_cancel_request" edge to ClassCancelRequest entities by IDs.
func (uuo *UserUpdateOne) RemoveClassCancelRequestIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveClassCancelRequestIDs(ids...)
	return uuo
}

// RemoveClassCancelRequest removes "class_cancel_request" edges to ClassCancelRequest entities.
func (uuo *UserUpdateOne) RemoveClassCancelRequest(c ...*ClassCancelRequest) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveClassCancelRequestIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Address(); ok {
		if err := user.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "User.address": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uuo.mutation.BirthDate(); ok {
		_spec.SetField(user.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeString, value)
	}
	if value, ok := uuo.mutation.ProfilePictureURL(); ok {
		_spec.SetField(user.FieldProfilePictureURL, field.TypeString, value)
	}
	if uuo.mutation.ProfilePictureURLCleared() {
		_spec.ClearField(user.FieldProfilePictureURL, field.TypeString)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if uuo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.StudentTable,
			Columns: []string{user.StudentColumn},
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
	if nodes := uuo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.StudentTable,
			Columns: []string{user.StudentColumn},
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
	if uuo.mutation.TutorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TutorTable,
			Columns: []string{user.TutorColumn},
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
	if nodes := uuo.mutation.TutorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TutorTable,
			Columns: []string{user.TutorColumn},
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
	if uuo.mutation.PaymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentTable,
			Columns: []string{user.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: payment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPaymentIDs(); len(nodes) > 0 && !uuo.mutation.PaymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentTable,
			Columns: []string{user.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentTable,
			Columns: []string{user.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentHistoryTable,
			Columns: []string{user.PaymentHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: paymenthistory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPaymentHistoryIDs(); len(nodes) > 0 && !uuo.mutation.PaymentHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentHistoryTable,
			Columns: []string{user.PaymentHistoryColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PaymentHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PaymentHistoryTable,
			Columns: []string{user.PaymentHistoryColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ClassCancelRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClassCancelRequestTable,
			Columns: []string{user.ClassCancelRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: classcancelrequest.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedClassCancelRequestIDs(); len(nodes) > 0 && !uuo.mutation.ClassCancelRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClassCancelRequestTable,
			Columns: []string{user.ClassCancelRequestColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ClassCancelRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClassCancelRequestTable,
			Columns: []string{user.ClassCancelRequestColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
