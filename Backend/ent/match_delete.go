// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
)

// MatchDelete is the builder for deleting a Match entity.
type MatchDelete struct {
	config
	hooks    []Hook
	mutation *MatchMutation
}

// Where appends a list predicates to the MatchDelete builder.
func (md *MatchDelete) Where(ps ...predicate.Match) *MatchDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MatchDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MatchMutation](ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MatchDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MatchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(match.Table, sqlgraph.NewFieldSpec(match.FieldID, field.TypeInt))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MatchDeleteOne is the builder for deleting a single Match entity.
type MatchDeleteOne struct {
	md *MatchDelete
}

// Where appends a list predicates to the MatchDelete builder.
func (mdo *MatchDeleteOne) Where(ps ...predicate.Match) *MatchDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MatchDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{match.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MatchDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
