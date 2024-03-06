// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ChungKiet/cake-interview/ent/uservoucher"
)

// UserVoucherCreate is the builder for creating a UserVoucher entity.
type UserVoucherCreate struct {
	config
	mutation *UserVoucherMutation
	hooks    []Hook
}

// SetVoucherCampaignID sets the "voucher_campaign_id" field.
func (uvc *UserVoucherCreate) SetVoucherCampaignID(i int64) *UserVoucherCreate {
	uvc.mutation.SetVoucherCampaignID(i)
	return uvc
}

// SetUserID sets the "user_id" field.
func (uvc *UserVoucherCreate) SetUserID(s string) *UserVoucherCreate {
	uvc.mutation.SetUserID(s)
	return uvc
}

// SetCreatedAt sets the "created_at" field.
func (uvc *UserVoucherCreate) SetCreatedAt(t time.Time) *UserVoucherCreate {
	uvc.mutation.SetCreatedAt(t)
	return uvc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uvc *UserVoucherCreate) SetNillableCreatedAt(t *time.Time) *UserVoucherCreate {
	if t != nil {
		uvc.SetCreatedAt(*t)
	}
	return uvc
}

// SetUpdatedAt sets the "updated_at" field.
func (uvc *UserVoucherCreate) SetUpdatedAt(t time.Time) *UserVoucherCreate {
	uvc.mutation.SetUpdatedAt(t)
	return uvc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uvc *UserVoucherCreate) SetNillableUpdatedAt(t *time.Time) *UserVoucherCreate {
	if t != nil {
		uvc.SetUpdatedAt(*t)
	}
	return uvc
}

// Mutation returns the UserVoucherMutation object of the builder.
func (uvc *UserVoucherCreate) Mutation() *UserVoucherMutation {
	return uvc.mutation
}

// Save creates the UserVoucher in the database.
func (uvc *UserVoucherCreate) Save(ctx context.Context) (*UserVoucher, error) {
	uvc.defaults()
	return withHooks(ctx, uvc.sqlSave, uvc.mutation, uvc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uvc *UserVoucherCreate) SaveX(ctx context.Context) *UserVoucher {
	v, err := uvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvc *UserVoucherCreate) Exec(ctx context.Context) error {
	_, err := uvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvc *UserVoucherCreate) ExecX(ctx context.Context) {
	if err := uvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uvc *UserVoucherCreate) defaults() {
	if _, ok := uvc.mutation.CreatedAt(); !ok {
		v := uservoucher.DefaultCreatedAt()
		uvc.mutation.SetCreatedAt(v)
	}
	if _, ok := uvc.mutation.UpdatedAt(); !ok {
		v := uservoucher.DefaultUpdatedAt()
		uvc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvc *UserVoucherCreate) check() error {
	if _, ok := uvc.mutation.VoucherCampaignID(); !ok {
		return &ValidationError{Name: "voucher_campaign_id", err: errors.New(`ent: missing required field "UserVoucher.voucher_campaign_id"`)}
	}
	if _, ok := uvc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserVoucher.user_id"`)}
	}
	if _, ok := uvc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserVoucher.created_at"`)}
	}
	if _, ok := uvc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserVoucher.updated_at"`)}
	}
	return nil
}

func (uvc *UserVoucherCreate) sqlSave(ctx context.Context) (*UserVoucher, error) {
	if err := uvc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uvc.mutation.id = &_node.ID
	uvc.mutation.done = true
	return _node, nil
}

func (uvc *UserVoucherCreate) createSpec() (*UserVoucher, *sqlgraph.CreateSpec) {
	var (
		_node = &UserVoucher{config: uvc.config}
		_spec = sqlgraph.NewCreateSpec(uservoucher.Table, sqlgraph.NewFieldSpec(uservoucher.FieldID, field.TypeInt))
	)
	if value, ok := uvc.mutation.VoucherCampaignID(); ok {
		_spec.SetField(uservoucher.FieldVoucherCampaignID, field.TypeInt64, value)
		_node.VoucherCampaignID = value
	}
	if value, ok := uvc.mutation.UserID(); ok {
		_spec.SetField(uservoucher.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := uvc.mutation.CreatedAt(); ok {
		_spec.SetField(uservoucher.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uvc.mutation.UpdatedAt(); ok {
		_spec.SetField(uservoucher.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// UserVoucherCreateBulk is the builder for creating many UserVoucher entities in bulk.
type UserVoucherCreateBulk struct {
	config
	err      error
	builders []*UserVoucherCreate
}

// Save creates the UserVoucher entities in the database.
func (uvcb *UserVoucherCreateBulk) Save(ctx context.Context) ([]*UserVoucher, error) {
	if uvcb.err != nil {
		return nil, uvcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uvcb.builders))
	nodes := make([]*UserVoucher, len(uvcb.builders))
	mutators := make([]Mutator, len(uvcb.builders))
	for i := range uvcb.builders {
		func(i int, root context.Context) {
			builder := uvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserVoucherMutation)
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
					_, err = mutators[i+1].Mutate(root, uvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uvcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uvcb *UserVoucherCreateBulk) SaveX(ctx context.Context) []*UserVoucher {
	v, err := uvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvcb *UserVoucherCreateBulk) Exec(ctx context.Context) error {
	_, err := uvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvcb *UserVoucherCreateBulk) ExecX(ctx context.Context) {
	if err := uvcb.Exec(ctx); err != nil {
		panic(err)
	}
}
