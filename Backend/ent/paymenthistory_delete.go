// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
)

// PaymentHistoryDelete is the builder for deleting a PaymentHistory entity.
type PaymentHistoryDelete struct {
	config
	hooks    []Hook
	mutation *PaymentHistoryMutation
}

// Where appends a list predicates to the PaymentHistoryDelete builder.
func (phd *PaymentHistoryDelete) Where(ps ...predicate.PaymentHistory) *PaymentHistoryDelete {
	phd.mutation.Where(ps...)
	return phd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (phd *PaymentHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PaymentHistoryMutation](ctx, phd.sqlExec, phd.mutation, phd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (phd *PaymentHistoryDelete) ExecX(ctx context.Context) int {
	n, err := phd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (phd *PaymentHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(paymenthistory.Table, sqlgraph.NewFieldSpec(paymenthistory.FieldID, field.TypeUUID))
	if ps := phd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, phd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	phd.mutation.done = true
	return affected, err
}

// PaymentHistoryDeleteOne is the builder for deleting a single PaymentHistory entity.
type PaymentHistoryDeleteOne struct {
	phd *PaymentHistoryDelete
}

// Where appends a list predicates to the PaymentHistoryDelete builder.
func (phdo *PaymentHistoryDeleteOne) Where(ps ...predicate.PaymentHistory) *PaymentHistoryDeleteOne {
	phdo.phd.mutation.Where(ps...)
	return phdo
}

// Exec executes the deletion query.
func (phdo *PaymentHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := phdo.phd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{paymenthistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (phdo *PaymentHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := phdo.Exec(ctx); err != nil {
		panic(err)
	}
}
