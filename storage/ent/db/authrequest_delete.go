// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/authrequest"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// AuthRequestDelete is the builder for deleting a AuthRequest entity.
type AuthRequestDelete struct {
	config
	hooks    []Hook
	mutation *AuthRequestMutation
}

// Where appends a list predicates to the AuthRequestDelete builder.
func (ard *AuthRequestDelete) Where(ps ...predicate.AuthRequest) *AuthRequestDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AuthRequestDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ard.hooks) == 0 {
		affected, err = ard.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ard.mutation = mutation
			affected, err = ard.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ard.hooks) - 1; i >= 0; i-- {
			if ard.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = ard.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ard.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AuthRequestDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AuthRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authrequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authrequest.FieldID,
			},
		},
	}
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AuthRequestDeleteOne is the builder for deleting a single AuthRequest entity.
type AuthRequestDeleteOne struct {
	ard *AuthRequestDelete
}

// Exec executes the deletion query.
func (ardo *AuthRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authrequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AuthRequestDeleteOne) ExecX(ctx context.Context) {
	ardo.ard.ExecX(ctx)
}
