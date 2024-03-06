// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ChungKiet/cake-interview/ent/userbalance"
)

// UserBalanceCreate is the builder for creating a UserBalance entity.
type UserBalanceCreate struct {
	config
	mutation *UserBalanceMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ubc *UserBalanceCreate) SetUserID(s string) *UserBalanceCreate {
	ubc.mutation.SetUserID(s)
	return ubc
}

// SetBalance sets the "balance" field.
func (ubc *UserBalanceCreate) SetBalance(f float64) *UserBalanceCreate {
	ubc.mutation.SetBalance(f)
	return ubc
}

// SetCreatedAt sets the "created_at" field.
func (ubc *UserBalanceCreate) SetCreatedAt(t time.Time) *UserBalanceCreate {
	ubc.mutation.SetCreatedAt(t)
	return ubc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ubc *UserBalanceCreate) SetNillableCreatedAt(t *time.Time) *UserBalanceCreate {
	if t != nil {
		ubc.SetCreatedAt(*t)
	}
	return ubc
}

// SetUpdatedAt sets the "updated_at" field.
func (ubc *UserBalanceCreate) SetUpdatedAt(t time.Time) *UserBalanceCreate {
	ubc.mutation.SetUpdatedAt(t)
	return ubc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ubc *UserBalanceCreate) SetNillableUpdatedAt(t *time.Time) *UserBalanceCreate {
	if t != nil {
		ubc.SetUpdatedAt(*t)
	}
	return ubc
}

// Mutation returns the UserBalanceMutation object of the builder.
func (ubc *UserBalanceCreate) Mutation() *UserBalanceMutation {
	return ubc.mutation
}

// Save creates the UserBalance in the database.
func (ubc *UserBalanceCreate) Save(ctx context.Context) (*UserBalance, error) {
	ubc.defaults()
	return withHooks(ctx, ubc.sqlSave, ubc.mutation, ubc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ubc *UserBalanceCreate) SaveX(ctx context.Context) *UserBalance {
	v, err := ubc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubc *UserBalanceCreate) Exec(ctx context.Context) error {
	_, err := ubc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubc *UserBalanceCreate) ExecX(ctx context.Context) {
	if err := ubc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ubc *UserBalanceCreate) defaults() {
	if _, ok := ubc.mutation.CreatedAt(); !ok {
		v := userbalance.DefaultCreatedAt()
		ubc.mutation.SetCreatedAt(v)
	}
	if _, ok := ubc.mutation.UpdatedAt(); !ok {
		v := userbalance.DefaultUpdatedAt()
		ubc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ubc *UserBalanceCreate) check() error {
	if _, ok := ubc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserBalance.user_id"`)}
	}
	if _, ok := ubc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "UserBalance.balance"`)}
	}
	if _, ok := ubc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserBalance.created_at"`)}
	}
	if _, ok := ubc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserBalance.updated_at"`)}
	}
	return nil
}

func (ubc *UserBalanceCreate) sqlSave(ctx context.Context) (*UserBalance, error) {
	if err := ubc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ubc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ubc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ubc.mutation.id = &_node.ID
	ubc.mutation.done = true
	return _node, nil
}

func (ubc *UserBalanceCreate) createSpec() (*UserBalance, *sqlgraph.CreateSpec) {
	var (
		_node = &UserBalance{config: ubc.config}
		_spec = sqlgraph.NewCreateSpec(userbalance.Table, sqlgraph.NewFieldSpec(userbalance.FieldID, field.TypeInt))
	)
	if value, ok := ubc.mutation.UserID(); ok {
		_spec.SetField(userbalance.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := ubc.mutation.Balance(); ok {
		_spec.SetField(userbalance.FieldBalance, field.TypeFloat64, value)
		_node.Balance = value
	}
	if value, ok := ubc.mutation.CreatedAt(); ok {
		_spec.SetField(userbalance.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ubc.mutation.UpdatedAt(); ok {
		_spec.SetField(userbalance.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// UserBalanceCreateBulk is the builder for creating many UserBalance entities in bulk.
type UserBalanceCreateBulk struct {
	config
	err      error
	builders []*UserBalanceCreate
}

// Save creates the UserBalance entities in the database.
func (ubcb *UserBalanceCreateBulk) Save(ctx context.Context) ([]*UserBalance, error) {
	if ubcb.err != nil {
		return nil, ubcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ubcb.builders))
	nodes := make([]*UserBalance, len(ubcb.builders))
	mutators := make([]Mutator, len(ubcb.builders))
	for i := range ubcb.builders {
		func(i int, root context.Context) {
			builder := ubcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserBalanceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ubcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ubcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ubcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ubcb *UserBalanceCreateBulk) SaveX(ctx context.Context) []*UserBalance {
	v, err := ubcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubcb *UserBalanceCreateBulk) Exec(ctx context.Context) error {
	_, err := ubcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubcb *UserBalanceCreateBulk) ExecX(ctx context.Context) {
	if err := ubcb.Exec(ctx); err != nil {
		panic(err)
	}
}
