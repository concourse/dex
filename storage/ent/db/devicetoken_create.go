// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/devicetoken"
)

// DeviceTokenCreate is the builder for creating a DeviceToken entity.
type DeviceTokenCreate struct {
	config
	mutation *DeviceTokenMutation
	hooks    []Hook
}

// SetDeviceCode sets the "device_code" field.
func (dtc *DeviceTokenCreate) SetDeviceCode(s string) *DeviceTokenCreate {
	dtc.mutation.SetDeviceCode(s)
	return dtc
}

// SetStatus sets the "status" field.
func (dtc *DeviceTokenCreate) SetStatus(s string) *DeviceTokenCreate {
	dtc.mutation.SetStatus(s)
	return dtc
}

// SetToken sets the "token" field.
func (dtc *DeviceTokenCreate) SetToken(b []byte) *DeviceTokenCreate {
	dtc.mutation.SetToken(b)
	return dtc
}

// SetExpiry sets the "expiry" field.
func (dtc *DeviceTokenCreate) SetExpiry(t time.Time) *DeviceTokenCreate {
	dtc.mutation.SetExpiry(t)
	return dtc
}

// SetLastRequest sets the "last_request" field.
func (dtc *DeviceTokenCreate) SetLastRequest(t time.Time) *DeviceTokenCreate {
	dtc.mutation.SetLastRequest(t)
	return dtc
}

// SetPollInterval sets the "poll_interval" field.
func (dtc *DeviceTokenCreate) SetPollInterval(i int) *DeviceTokenCreate {
	dtc.mutation.SetPollInterval(i)
	return dtc
}

// SetCodeChallenge sets the "code_challenge" field.
func (dtc *DeviceTokenCreate) SetCodeChallenge(s string) *DeviceTokenCreate {
	dtc.mutation.SetCodeChallenge(s)
	return dtc
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (dtc *DeviceTokenCreate) SetNillableCodeChallenge(s *string) *DeviceTokenCreate {
	if s != nil {
		dtc.SetCodeChallenge(*s)
	}
	return dtc
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (dtc *DeviceTokenCreate) SetCodeChallengeMethod(s string) *DeviceTokenCreate {
	dtc.mutation.SetCodeChallengeMethod(s)
	return dtc
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (dtc *DeviceTokenCreate) SetNillableCodeChallengeMethod(s *string) *DeviceTokenCreate {
	if s != nil {
		dtc.SetCodeChallengeMethod(*s)
	}
	return dtc
}

// Mutation returns the DeviceTokenMutation object of the builder.
func (dtc *DeviceTokenCreate) Mutation() *DeviceTokenMutation {
	return dtc.mutation
}

// Save creates the DeviceToken in the database.
func (dtc *DeviceTokenCreate) Save(ctx context.Context) (*DeviceToken, error) {
	dtc.defaults()
	return withHooks[*DeviceToken, DeviceTokenMutation](ctx, dtc.sqlSave, dtc.mutation, dtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dtc *DeviceTokenCreate) SaveX(ctx context.Context) *DeviceToken {
	v, err := dtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtc *DeviceTokenCreate) Exec(ctx context.Context) error {
	_, err := dtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtc *DeviceTokenCreate) ExecX(ctx context.Context) {
	if err := dtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dtc *DeviceTokenCreate) defaults() {
	if _, ok := dtc.mutation.CodeChallenge(); !ok {
		v := devicetoken.DefaultCodeChallenge
		dtc.mutation.SetCodeChallenge(v)
	}
	if _, ok := dtc.mutation.CodeChallengeMethod(); !ok {
		v := devicetoken.DefaultCodeChallengeMethod
		dtc.mutation.SetCodeChallengeMethod(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtc *DeviceTokenCreate) check() error {
	if _, ok := dtc.mutation.DeviceCode(); !ok {
		return &ValidationError{Name: "device_code", err: errors.New(`db: missing required field "DeviceToken.device_code"`)}
	}
	if v, ok := dtc.mutation.DeviceCode(); ok {
		if err := devicetoken.DeviceCodeValidator(v); err != nil {
			return &ValidationError{Name: "device_code", err: fmt.Errorf(`db: validator failed for field "DeviceToken.device_code": %w`, err)}
		}
	}
	if _, ok := dtc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`db: missing required field "DeviceToken.status"`)}
	}
	if v, ok := dtc.mutation.Status(); ok {
		if err := devicetoken.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "DeviceToken.status": %w`, err)}
		}
	}
	if _, ok := dtc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`db: missing required field "DeviceToken.expiry"`)}
	}
	if _, ok := dtc.mutation.LastRequest(); !ok {
		return &ValidationError{Name: "last_request", err: errors.New(`db: missing required field "DeviceToken.last_request"`)}
	}
	if _, ok := dtc.mutation.PollInterval(); !ok {
		return &ValidationError{Name: "poll_interval", err: errors.New(`db: missing required field "DeviceToken.poll_interval"`)}
	}
	if _, ok := dtc.mutation.CodeChallenge(); !ok {
		return &ValidationError{Name: "code_challenge", err: errors.New(`db: missing required field "DeviceToken.code_challenge"`)}
	}
	if _, ok := dtc.mutation.CodeChallengeMethod(); !ok {
		return &ValidationError{Name: "code_challenge_method", err: errors.New(`db: missing required field "DeviceToken.code_challenge_method"`)}
	}
	return nil
}

func (dtc *DeviceTokenCreate) sqlSave(ctx context.Context) (*DeviceToken, error) {
	if err := dtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dtc.mutation.id = &_node.ID
	dtc.mutation.done = true
	return _node, nil
}

func (dtc *DeviceTokenCreate) createSpec() (*DeviceToken, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceToken{config: dtc.config}
		_spec = sqlgraph.NewCreateSpec(devicetoken.Table, sqlgraph.NewFieldSpec(devicetoken.FieldID, field.TypeInt))
	)
	if value, ok := dtc.mutation.DeviceCode(); ok {
		_spec.SetField(devicetoken.FieldDeviceCode, field.TypeString, value)
		_node.DeviceCode = value
	}
	if value, ok := dtc.mutation.Status(); ok {
		_spec.SetField(devicetoken.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := dtc.mutation.Token(); ok {
		_spec.SetField(devicetoken.FieldToken, field.TypeBytes, value)
		_node.Token = &value
	}
	if value, ok := dtc.mutation.Expiry(); ok {
		_spec.SetField(devicetoken.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	if value, ok := dtc.mutation.LastRequest(); ok {
		_spec.SetField(devicetoken.FieldLastRequest, field.TypeTime, value)
		_node.LastRequest = value
	}
	if value, ok := dtc.mutation.PollInterval(); ok {
		_spec.SetField(devicetoken.FieldPollInterval, field.TypeInt, value)
		_node.PollInterval = value
	}
	if value, ok := dtc.mutation.CodeChallenge(); ok {
		_spec.SetField(devicetoken.FieldCodeChallenge, field.TypeString, value)
		_node.CodeChallenge = value
	}
	if value, ok := dtc.mutation.CodeChallengeMethod(); ok {
		_spec.SetField(devicetoken.FieldCodeChallengeMethod, field.TypeString, value)
		_node.CodeChallengeMethod = value
	}
	return _node, _spec
}

// DeviceTokenCreateBulk is the builder for creating many DeviceToken entities in bulk.
type DeviceTokenCreateBulk struct {
	config
	builders []*DeviceTokenCreate
}

// Save creates the DeviceToken entities in the database.
func (dtcb *DeviceTokenCreateBulk) Save(ctx context.Context) ([]*DeviceToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dtcb.builders))
	nodes := make([]*DeviceToken, len(dtcb.builders))
	mutators := make([]Mutator, len(dtcb.builders))
	for i := range dtcb.builders {
		func(i int, root context.Context) {
			builder := dtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, dtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dtcb *DeviceTokenCreateBulk) SaveX(ctx context.Context) []*DeviceToken {
	v, err := dtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtcb *DeviceTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := dtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtcb *DeviceTokenCreateBulk) ExecX(ctx context.Context) {
	if err := dtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
