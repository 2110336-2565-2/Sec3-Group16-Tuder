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
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/cancelrequest"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// CancelRequestUpdate is the builder for updating CancelRequest entities.
type CancelRequestUpdate struct {
	config
	hooks    []Hook
	mutation *CancelRequestMutation
}

// Where appends a list predicates to the CancelRequestUpdate builder.
func (cru *CancelRequestUpdate) Where(ps ...predicate.CancelRequest) *CancelRequestUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetTitle sets the "title" field.
func (cru *CancelRequestUpdate) SetTitle(s string) *CancelRequestUpdate {
	cru.mutation.SetTitle(s)
	return cru
}

// SetReportDate sets the "report_date" field.
func (cru *CancelRequestUpdate) SetReportDate(t time.Time) *CancelRequestUpdate {
	cru.mutation.SetReportDate(t)
	return cru
}

// SetNillableReportDate sets the "report_date" field if the given value is not nil.
func (cru *CancelRequestUpdate) SetNillableReportDate(t *time.Time) *CancelRequestUpdate {
	if t != nil {
		cru.SetReportDate(*t)
	}
	return cru
}

// SetImgURL sets the "img_url" field.
func (cru *CancelRequestUpdate) SetImgURL(s string) *CancelRequestUpdate {
	cru.mutation.SetImgURL(s)
	return cru
}

// SetDescription sets the "description" field.
func (cru *CancelRequestUpdate) SetDescription(s string) *CancelRequestUpdate {
	cru.mutation.SetDescription(s)
	return cru
}

// SetStatus sets the "status" field.
func (cru *CancelRequestUpdate) SetStatus(c cancelrequest.Status) *CancelRequestUpdate {
	cru.mutation.SetStatus(c)
	return cru
}

// SetMatchID sets the "match" edge to the Match entity by ID.
func (cru *CancelRequestUpdate) SetMatchID(id uuid.UUID) *CancelRequestUpdate {
	cru.mutation.SetMatchID(id)
	return cru
}

// SetMatch sets the "match" edge to the Match entity.
func (cru *CancelRequestUpdate) SetMatch(m *Match) *CancelRequestUpdate {
	return cru.SetMatchID(m.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cru *CancelRequestUpdate) SetUserID(id uuid.UUID) *CancelRequestUpdate {
	cru.mutation.SetUserID(id)
	return cru
}

// SetUser sets the "user" edge to the User entity.
func (cru *CancelRequestUpdate) SetUser(u *User) *CancelRequestUpdate {
	return cru.SetUserID(u.ID)
}

// Mutation returns the CancelRequestMutation object of the builder.
func (cru *CancelRequestUpdate) Mutation() *CancelRequestMutation {
	return cru.mutation
}

// ClearMatch clears the "match" edge to the Match entity.
func (cru *CancelRequestUpdate) ClearMatch() *CancelRequestUpdate {
	cru.mutation.ClearMatch()
	return cru
}

// ClearUser clears the "user" edge to the User entity.
func (cru *CancelRequestUpdate) ClearUser() *CancelRequestUpdate {
	cru.mutation.ClearUser()
	return cru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CancelRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CancelRequestMutation](ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CancelRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CancelRequestUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CancelRequestUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *CancelRequestUpdate) check() error {
	if v, ok := cru.mutation.Title(); ok {
		if err := cancelrequest.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CancelRequest.title": %w`, err)}
		}
	}
	if v, ok := cru.mutation.Description(); ok {
		if err := cancelrequest.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "CancelRequest.description": %w`, err)}
		}
	}
	if v, ok := cru.mutation.Status(); ok {
		if err := cancelrequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "CancelRequest.status": %w`, err)}
		}
	}
	if _, ok := cru.mutation.MatchID(); cru.mutation.MatchCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CancelRequest.match"`)
	}
	if _, ok := cru.mutation.UserID(); cru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CancelRequest.user"`)
	}
	return nil
}

func (cru *CancelRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cancelrequest.Table, cancelrequest.Columns, sqlgraph.NewFieldSpec(cancelrequest.FieldID, field.TypeUUID))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.Title(); ok {
		_spec.SetField(cancelrequest.FieldTitle, field.TypeString, value)
	}
	if value, ok := cru.mutation.ReportDate(); ok {
		_spec.SetField(cancelrequest.FieldReportDate, field.TypeTime, value)
	}
	if value, ok := cru.mutation.ImgURL(); ok {
		_spec.SetField(cancelrequest.FieldImgURL, field.TypeString, value)
	}
	if value, ok := cru.mutation.Description(); ok {
		_spec.SetField(cancelrequest.FieldDescription, field.TypeString, value)
	}
	if value, ok := cru.mutation.Status(); ok {
		_spec.SetField(cancelrequest.FieldStatus, field.TypeEnum, value)
	}
	if cru.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cancelrequest.MatchTable,
			Columns: []string{cancelrequest.MatchColumn},
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
	if nodes := cru.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cancelrequest.MatchTable,
			Columns: []string{cancelrequest.MatchColumn},
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
	if cru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cancelrequest.UserTable,
			Columns: []string{cancelrequest.UserColumn},
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
	if nodes := cru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cancelrequest.UserTable,
			Columns: []string{cancelrequest.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cancelrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// CancelRequestUpdateOne is the builder for updating a single CancelRequest entity.
type CancelRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CancelRequestMutation
}

// SetTitle sets the "title" field.
func (cruo *CancelRequestUpdateOne) SetTitle(s string) *CancelRequestUpdateOne {
	cruo.mutation.SetTitle(s)
	return cruo
}

// SetReportDate sets the "report_date" field.
func (cruo *CancelRequestUpdateOne) SetReportDate(t time.Time) *CancelRequestUpdateOne {
	cruo.mutation.SetReportDate(t)
	return cruo
}

// SetNillableReportDate sets the "report_date" field if the given value is not nil.
func (cruo *CancelRequestUpdateOne) SetNillableReportDate(t *time.Time) *CancelRequestUpdateOne {
	if t != nil {
		cruo.SetReportDate(*t)
	}
	return cruo
}

// SetImgURL sets the "img_url" field.
func (cruo *CancelRequestUpdateOne) SetImgURL(s string) *CancelRequestUpdateOne {
	cruo.mutation.SetImgURL(s)
	return cruo
}

// SetDescription sets the "description" field.
func (cruo *CancelRequestUpdateOne) SetDescription(s string) *CancelRequestUpdateOne {
	cruo.mutation.SetDescription(s)
	return cruo
}

// SetStatus sets the "status" field.
func (cruo *CancelRequestUpdateOne) SetStatus(c cancelrequest.Status) *CancelRequestUpdateOne {
	cruo.mutation.SetStatus(c)
	return cruo
}

// SetMatchID sets the "match" edge to the Match entity by ID.
func (cruo *CancelRequestUpdateOne) SetMatchID(id uuid.UUID) *CancelRequestUpdateOne {
	cruo.mutation.SetMatchID(id)
	return cruo
}

// SetMatch sets the "match" edge to the Match entity.
func (cruo *CancelRequestUpdateOne) SetMatch(m *Match) *CancelRequestUpdateOne {
	return cruo.SetMatchID(m.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cruo *CancelRequestUpdateOne) SetUserID(id uuid.UUID) *CancelRequestUpdateOne {
	cruo.mutation.SetUserID(id)
	return cruo
}

// SetUser sets the "user" edge to the User entity.
func (cruo *CancelRequestUpdateOne) SetUser(u *User) *CancelRequestUpdateOne {
	return cruo.SetUserID(u.ID)
}

// Mutation returns the CancelRequestMutation object of the builder.
func (cruo *CancelRequestUpdateOne) Mutation() *CancelRequestMutation {
	return cruo.mutation
}

// ClearMatch clears the "match" edge to the Match entity.
func (cruo *CancelRequestUpdateOne) ClearMatch() *CancelRequestUpdateOne {
	cruo.mutation.ClearMatch()
	return cruo
}

// ClearUser clears the "user" edge to the User entity.
func (cruo *CancelRequestUpdateOne) ClearUser() *CancelRequestUpdateOne {
	cruo.mutation.ClearUser()
	return cruo
}

// Where appends a list predicates to the CancelRequestUpdate builder.
func (cruo *CancelRequestUpdateOne) Where(ps ...predicate.CancelRequest) *CancelRequestUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CancelRequestUpdateOne) Select(field string, fields ...string) *CancelRequestUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CancelRequest entity.
func (cruo *CancelRequestUpdateOne) Save(ctx context.Context) (*CancelRequest, error) {
	return withHooks[*CancelRequest, CancelRequestMutation](ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CancelRequestUpdateOne) SaveX(ctx context.Context) *CancelRequest {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CancelRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CancelRequestUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *CancelRequestUpdateOne) check() error {
	if v, ok := cruo.mutation.Title(); ok {
		if err := cancelrequest.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CancelRequest.title": %w`, err)}
		}
	}
	if v, ok := cruo.mutation.Description(); ok {
		if err := cancelrequest.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "CancelRequest.description": %w`, err)}
		}
	}
	if v, ok := cruo.mutation.Status(); ok {
		if err := cancelrequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "CancelRequest.status": %w`, err)}
		}
	}
	if _, ok := cruo.mutation.MatchID(); cruo.mutation.MatchCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CancelRequest.match"`)
	}
	if _, ok := cruo.mutation.UserID(); cruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CancelRequest.user"`)
	}
	return nil
}

func (cruo *CancelRequestUpdateOne) sqlSave(ctx context.Context) (_node *CancelRequest, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cancelrequest.Table, cancelrequest.Columns, sqlgraph.NewFieldSpec(cancelrequest.FieldID, field.TypeUUID))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CancelRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cancelrequest.FieldID)
		for _, f := range fields {
			if !cancelrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cancelrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.Title(); ok {
		_spec.SetField(cancelrequest.FieldTitle, field.TypeString, value)
	}
	if value, ok := cruo.mutation.ReportDate(); ok {
		_spec.SetField(cancelrequest.FieldReportDate, field.TypeTime, value)
	}
	if value, ok := cruo.mutation.ImgURL(); ok {
		_spec.SetField(cancelrequest.FieldImgURL, field.TypeString, value)
	}
	if value, ok := cruo.mutation.Description(); ok {
		_spec.SetField(cancelrequest.FieldDescription, field.TypeString, value)
	}
	if value, ok := cruo.mutation.Status(); ok {
		_spec.SetField(cancelrequest.FieldStatus, field.TypeEnum, value)
	}
	if cruo.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cancelrequest.MatchTable,
			Columns: []string{cancelrequest.MatchColumn},
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
	if nodes := cruo.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cancelrequest.MatchTable,
			Columns: []string{cancelrequest.MatchColumn},
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
	if cruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cancelrequest.UserTable,
			Columns: []string{cancelrequest.UserColumn},
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
	if nodes := cruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cancelrequest.UserTable,
			Columns: []string{cancelrequest.UserColumn},
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
	_node = &CancelRequest{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cancelrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}
