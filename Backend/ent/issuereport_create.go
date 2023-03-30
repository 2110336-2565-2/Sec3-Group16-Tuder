// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/issuereport"
	"github.com/google/uuid"
)

// IssueReportCreate is the builder for creating a IssueReport entity.
type IssueReportCreate struct {
	config
	mutation *IssueReportMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (irc *IssueReportCreate) SetTitle(s string) *IssueReportCreate {
	irc.mutation.SetTitle(s)
	return irc
}

// SetDescription sets the "description" field.
func (irc *IssueReportCreate) SetDescription(s string) *IssueReportCreate {
	irc.mutation.SetDescription(s)
	return irc
}

// SetContact sets the "contact" field.
func (irc *IssueReportCreate) SetContact(s string) *IssueReportCreate {
	irc.mutation.SetContact(s)
	return irc
}

// SetNillableContact sets the "contact" field if the given value is not nil.
func (irc *IssueReportCreate) SetNillableContact(s *string) *IssueReportCreate {
	if s != nil {
		irc.SetContact(*s)
	}
	return irc
}

// SetReportDate sets the "report_date" field.
func (irc *IssueReportCreate) SetReportDate(t time.Time) *IssueReportCreate {
	irc.mutation.SetReportDate(t)
	return irc
}

// SetStatus sets the "status" field.
func (irc *IssueReportCreate) SetStatus(i issuereport.Status) *IssueReportCreate {
	irc.mutation.SetStatus(i)
	return irc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (irc *IssueReportCreate) SetNillableStatus(i *issuereport.Status) *IssueReportCreate {
	if i != nil {
		irc.SetStatus(*i)
	}
	return irc
}

// SetID sets the "id" field.
func (irc *IssueReportCreate) SetID(u uuid.UUID) *IssueReportCreate {
	irc.mutation.SetID(u)
	return irc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (irc *IssueReportCreate) SetNillableID(u *uuid.UUID) *IssueReportCreate {
	if u != nil {
		irc.SetID(*u)
	}
	return irc
}

// Mutation returns the IssueReportMutation object of the builder.
func (irc *IssueReportCreate) Mutation() *IssueReportMutation {
	return irc.mutation
}

// Save creates the IssueReport in the database.
func (irc *IssueReportCreate) Save(ctx context.Context) (*IssueReport, error) {
	irc.defaults()
	return withHooks[*IssueReport, IssueReportMutation](ctx, irc.sqlSave, irc.mutation, irc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (irc *IssueReportCreate) SaveX(ctx context.Context) *IssueReport {
	v, err := irc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (irc *IssueReportCreate) Exec(ctx context.Context) error {
	_, err := irc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (irc *IssueReportCreate) ExecX(ctx context.Context) {
	if err := irc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (irc *IssueReportCreate) defaults() {
	if _, ok := irc.mutation.Contact(); !ok {
		v := issuereport.DefaultContact
		irc.mutation.SetContact(v)
	}
	if _, ok := irc.mutation.Status(); !ok {
		v := issuereport.DefaultStatus
		irc.mutation.SetStatus(v)
	}
	if _, ok := irc.mutation.ID(); !ok {
		v := issuereport.DefaultID()
		irc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (irc *IssueReportCreate) check() error {
	if _, ok := irc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "IssueReport.title"`)}
	}
	if v, ok := irc.mutation.Title(); ok {
		if err := issuereport.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "IssueReport.title": %w`, err)}
		}
	}
	if _, ok := irc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "IssueReport.description"`)}
	}
	if v, ok := irc.mutation.Description(); ok {
		if err := issuereport.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "IssueReport.description": %w`, err)}
		}
	}
	if _, ok := irc.mutation.Contact(); !ok {
		return &ValidationError{Name: "contact", err: errors.New(`ent: missing required field "IssueReport.contact"`)}
	}
	if _, ok := irc.mutation.ReportDate(); !ok {
		return &ValidationError{Name: "report_date", err: errors.New(`ent: missing required field "IssueReport.report_date"`)}
	}
	if _, ok := irc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "IssueReport.status"`)}
	}
	if v, ok := irc.mutation.Status(); ok {
		if err := issuereport.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "IssueReport.status": %w`, err)}
		}
	}
	return nil
}

func (irc *IssueReportCreate) sqlSave(ctx context.Context) (*IssueReport, error) {
	if err := irc.check(); err != nil {
		return nil, err
	}
	_node, _spec := irc.createSpec()
	if err := sqlgraph.CreateNode(ctx, irc.driver, _spec); err != nil {
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
	irc.mutation.id = &_node.ID
	irc.mutation.done = true
	return _node, nil
}

func (irc *IssueReportCreate) createSpec() (*IssueReport, *sqlgraph.CreateSpec) {
	var (
		_node = &IssueReport{config: irc.config}
		_spec = sqlgraph.NewCreateSpec(issuereport.Table, sqlgraph.NewFieldSpec(issuereport.FieldID, field.TypeUUID))
	)
	if id, ok := irc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := irc.mutation.Title(); ok {
		_spec.SetField(issuereport.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := irc.mutation.Description(); ok {
		_spec.SetField(issuereport.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := irc.mutation.Contact(); ok {
		_spec.SetField(issuereport.FieldContact, field.TypeString, value)
		_node.Contact = value
	}
	if value, ok := irc.mutation.ReportDate(); ok {
		_spec.SetField(issuereport.FieldReportDate, field.TypeTime, value)
		_node.ReportDate = value
	}
	if value, ok := irc.mutation.Status(); ok {
		_spec.SetField(issuereport.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	return _node, _spec
}

// IssueReportCreateBulk is the builder for creating many IssueReport entities in bulk.
type IssueReportCreateBulk struct {
	config
	builders []*IssueReportCreate
}

// Save creates the IssueReport entities in the database.
func (ircb *IssueReportCreateBulk) Save(ctx context.Context) ([]*IssueReport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ircb.builders))
	nodes := make([]*IssueReport, len(ircb.builders))
	mutators := make([]Mutator, len(ircb.builders))
	for i := range ircb.builders {
		func(i int, root context.Context) {
			builder := ircb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IssueReportMutation)
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
					_, err = mutators[i+1].Mutate(root, ircb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ircb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ircb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ircb *IssueReportCreateBulk) SaveX(ctx context.Context) []*IssueReport {
	v, err := ircb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ircb *IssueReportCreateBulk) Exec(ctx context.Context) error {
	_, err := ircb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ircb *IssueReportCreateBulk) ExecX(ctx context.Context) {
	if err := ircb.Exec(ctx); err != nil {
		panic(err)
	}
}
