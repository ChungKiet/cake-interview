// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ChungKiet/cake-interview/ent/predicate"
	"github.com/ChungKiet/cake-interview/ent/userbalance"
)

// UserBalanceDelete is the builder for deleting a UserBalance entity.
type UserBalanceDelete struct {
	config
	hooks    []Hook
	mutation *UserBalanceMutation
}

// Where appends a list predicates to the UserBalanceDelete builder.
func (ubd *UserBalanceDelete) Where(ps ...predicate.UserBalance) *UserBalanceDelete {
	ubd.mutation.Where(ps...)
	return ubd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ubd *UserBalanceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ubd.sqlExec, ubd.mutation, ubd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ubd *UserBalanceDelete) ExecX(ctx context.Context) int {
	n, err := ubd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ubd *UserBalanceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userbalance.Table, sqlgraph.NewFieldSpec(userbalance.FieldID, field.TypeInt))
	if ps := ubd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ubd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ubd.mutation.done = true
	return affected, err
}

// UserBalanceDeleteOne is the builder for deleting a single UserBalance entity.
type UserBalanceDeleteOne struct {
	ubd *UserBalanceDelete
}

// Where appends a list predicates to the UserBalanceDelete builder.
func (ubdo *UserBalanceDeleteOne) Where(ps ...predicate.UserBalance) *UserBalanceDeleteOne {
	ubdo.ubd.mutation.Where(ps...)
	return ubdo
}

// Exec executes the deletion query.
func (ubdo *UserBalanceDeleteOne) Exec(ctx context.Context) error {
	n, err := ubdo.ubd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userbalance.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ubdo *UserBalanceDeleteOne) ExecX(ctx context.Context) {
	if err := ubdo.Exec(ctx); err != nil {
		panic(err)
	}
}
