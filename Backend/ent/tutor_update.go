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
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
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

// SetUsername sets the "username" field.
func (tu *TutorUpdate) SetUsername(s string) *TutorUpdate {
	tu.mutation.SetUsername(s)
	return tu
}

// SetPassword sets the "password" field.
func (tu *TutorUpdate) SetPassword(s string) *TutorUpdate {
	tu.mutation.SetPassword(s)
	return tu
}

// SetEmail sets the "email" field.
func (tu *TutorUpdate) SetEmail(s string) *TutorUpdate {
	tu.mutation.SetEmail(s)
	return tu
}

// SetFirstName sets the "first_name" field.
func (tu *TutorUpdate) SetFirstName(s string) *TutorUpdate {
	tu.mutation.SetFirstName(s)
	return tu
}

// SetLastName sets the "last_name" field.
func (tu *TutorUpdate) SetLastName(s string) *TutorUpdate {
	tu.mutation.SetLastName(s)
	return tu
}

// SetAddress sets the "address" field.
func (tu *TutorUpdate) SetAddress(s string) *TutorUpdate {
	tu.mutation.SetAddress(s)
	return tu
}

// SetPhone sets the "phone" field.
func (tu *TutorUpdate) SetPhone(s string) *TutorUpdate {
	tu.mutation.SetPhone(s)
	return tu
}

// SetBirthDate sets the "birth_date" field.
func (tu *TutorUpdate) SetBirthDate(t time.Time) *TutorUpdate {
	tu.mutation.SetBirthDate(t)
	return tu
}

// SetGender sets the "gender" field.
func (tu *TutorUpdate) SetGender(s string) *TutorUpdate {
	tu.mutation.SetGender(s)
	return tu
}

// SetProfilePictureURL sets the "profile_picture_URL" field.
func (tu *TutorUpdate) SetProfilePictureURL(s string) *TutorUpdate {
	tu.mutation.SetProfilePictureURL(s)
	return tu
}

// SetNillableProfilePictureURL sets the "profile_picture_URL" field if the given value is not nil.
func (tu *TutorUpdate) SetNillableProfilePictureURL(s *string) *TutorUpdate {
	if s != nil {
		tu.SetProfilePictureURL(*s)
	}
	return tu
}

// ClearProfilePictureURL clears the value of the "profile_picture_URL" field.
func (tu *TutorUpdate) ClearProfilePictureURL() *TutorUpdate {
	tu.mutation.ClearProfilePictureURL()
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

// Mutation returns the TutorMutation object of the builder.
func (tu *TutorUpdate) Mutation() *TutorMutation {
	return tu.mutation
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
	if v, ok := tu.mutation.Username(); ok {
		if err := tutor.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Tutor.username": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Password(); ok {
		if err := tutor.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Tutor.password": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Email(); ok {
		if err := tutor.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Tutor.email": %w`, err)}
		}
	}
	if v, ok := tu.mutation.FirstName(); ok {
		if err := tutor.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Tutor.first_name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.LastName(); ok {
		if err := tutor.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Tutor.last_name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Address(); ok {
		if err := tutor.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Tutor.address": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Phone(); ok {
		if err := tutor.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Tutor.phone": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Gender(); ok {
		if err := tutor.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Tutor.gender": %w`, err)}
		}
	}
	if v, ok := tu.mutation.CitizenID(); ok {
		if err := tutor.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Tutor.citizen_id": %w`, err)}
		}
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
	if value, ok := tu.mutation.Username(); ok {
		_spec.SetField(tutor.FieldUsername, field.TypeString, value)
	}
	if value, ok := tu.mutation.Password(); ok {
		_spec.SetField(tutor.FieldPassword, field.TypeString, value)
	}
	if value, ok := tu.mutation.Email(); ok {
		_spec.SetField(tutor.FieldEmail, field.TypeString, value)
	}
	if value, ok := tu.mutation.FirstName(); ok {
		_spec.SetField(tutor.FieldFirstName, field.TypeString, value)
	}
	if value, ok := tu.mutation.LastName(); ok {
		_spec.SetField(tutor.FieldLastName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Address(); ok {
		_spec.SetField(tutor.FieldAddress, field.TypeString, value)
	}
	if value, ok := tu.mutation.Phone(); ok {
		_spec.SetField(tutor.FieldPhone, field.TypeString, value)
	}
	if value, ok := tu.mutation.BirthDate(); ok {
		_spec.SetField(tutor.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Gender(); ok {
		_spec.SetField(tutor.FieldGender, field.TypeString, value)
	}
	if value, ok := tu.mutation.ProfilePictureURL(); ok {
		_spec.SetField(tutor.FieldProfilePictureURL, field.TypeString, value)
	}
	if tu.mutation.ProfilePictureURLCleared() {
		_spec.ClearField(tutor.FieldProfilePictureURL, field.TypeString)
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

// SetUsername sets the "username" field.
func (tuo *TutorUpdateOne) SetUsername(s string) *TutorUpdateOne {
	tuo.mutation.SetUsername(s)
	return tuo
}

// SetPassword sets the "password" field.
func (tuo *TutorUpdateOne) SetPassword(s string) *TutorUpdateOne {
	tuo.mutation.SetPassword(s)
	return tuo
}

// SetEmail sets the "email" field.
func (tuo *TutorUpdateOne) SetEmail(s string) *TutorUpdateOne {
	tuo.mutation.SetEmail(s)
	return tuo
}

// SetFirstName sets the "first_name" field.
func (tuo *TutorUpdateOne) SetFirstName(s string) *TutorUpdateOne {
	tuo.mutation.SetFirstName(s)
	return tuo
}

// SetLastName sets the "last_name" field.
func (tuo *TutorUpdateOne) SetLastName(s string) *TutorUpdateOne {
	tuo.mutation.SetLastName(s)
	return tuo
}

// SetAddress sets the "address" field.
func (tuo *TutorUpdateOne) SetAddress(s string) *TutorUpdateOne {
	tuo.mutation.SetAddress(s)
	return tuo
}

// SetPhone sets the "phone" field.
func (tuo *TutorUpdateOne) SetPhone(s string) *TutorUpdateOne {
	tuo.mutation.SetPhone(s)
	return tuo
}

// SetBirthDate sets the "birth_date" field.
func (tuo *TutorUpdateOne) SetBirthDate(t time.Time) *TutorUpdateOne {
	tuo.mutation.SetBirthDate(t)
	return tuo
}

// SetGender sets the "gender" field.
func (tuo *TutorUpdateOne) SetGender(s string) *TutorUpdateOne {
	tuo.mutation.SetGender(s)
	return tuo
}

// SetProfilePictureURL sets the "profile_picture_URL" field.
func (tuo *TutorUpdateOne) SetProfilePictureURL(s string) *TutorUpdateOne {
	tuo.mutation.SetProfilePictureURL(s)
	return tuo
}

// SetNillableProfilePictureURL sets the "profile_picture_URL" field if the given value is not nil.
func (tuo *TutorUpdateOne) SetNillableProfilePictureURL(s *string) *TutorUpdateOne {
	if s != nil {
		tuo.SetProfilePictureURL(*s)
	}
	return tuo
}

// ClearProfilePictureURL clears the value of the "profile_picture_URL" field.
func (tuo *TutorUpdateOne) ClearProfilePictureURL() *TutorUpdateOne {
	tuo.mutation.ClearProfilePictureURL()
	return tuo
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

// Mutation returns the TutorMutation object of the builder.
func (tuo *TutorUpdateOne) Mutation() *TutorMutation {
	return tuo.mutation
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
	if v, ok := tuo.mutation.Username(); ok {
		if err := tutor.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Tutor.username": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Password(); ok {
		if err := tutor.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Tutor.password": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Email(); ok {
		if err := tutor.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Tutor.email": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.FirstName(); ok {
		if err := tutor.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Tutor.first_name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.LastName(); ok {
		if err := tutor.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Tutor.last_name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Address(); ok {
		if err := tutor.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Tutor.address": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Phone(); ok {
		if err := tutor.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Tutor.phone": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Gender(); ok {
		if err := tutor.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Tutor.gender": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.CitizenID(); ok {
		if err := tutor.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Tutor.citizen_id": %w`, err)}
		}
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
	if value, ok := tuo.mutation.Username(); ok {
		_spec.SetField(tutor.FieldUsername, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Password(); ok {
		_spec.SetField(tutor.FieldPassword, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Email(); ok {
		_spec.SetField(tutor.FieldEmail, field.TypeString, value)
	}
	if value, ok := tuo.mutation.FirstName(); ok {
		_spec.SetField(tutor.FieldFirstName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.LastName(); ok {
		_spec.SetField(tutor.FieldLastName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Address(); ok {
		_spec.SetField(tutor.FieldAddress, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Phone(); ok {
		_spec.SetField(tutor.FieldPhone, field.TypeString, value)
	}
	if value, ok := tuo.mutation.BirthDate(); ok {
		_spec.SetField(tutor.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Gender(); ok {
		_spec.SetField(tutor.FieldGender, field.TypeString, value)
	}
	if value, ok := tuo.mutation.ProfilePictureURL(); ok {
		_spec.SetField(tutor.FieldProfilePictureURL, field.TypeString, value)
	}
	if tuo.mutation.ProfilePictureURLCleared() {
		_spec.ClearField(tutor.FieldProfilePictureURL, field.TypeString)
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
