// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/devicetoken"
	"github.com/dexidp/dex/storage/ent/db/predicate"
)

// DeviceTokenUpdate is the builder for updating DeviceToken entities.
type DeviceTokenUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceTokenMutation
}

// Where appends a list predicates to the DeviceTokenUpdate builder.
func (dtu *DeviceTokenUpdate) Where(ps ...predicate.DeviceToken) *DeviceTokenUpdate {
	dtu.mutation.Where(ps...)
	return dtu
}

// SetDeviceCode sets the "device_code" field.
func (dtu *DeviceTokenUpdate) SetDeviceCode(s string) *DeviceTokenUpdate {
	dtu.mutation.SetDeviceCode(s)
	return dtu
}

// SetStatus sets the "status" field.
func (dtu *DeviceTokenUpdate) SetStatus(s string) *DeviceTokenUpdate {
	dtu.mutation.SetStatus(s)
	return dtu
}

// SetToken sets the "token" field.
func (dtu *DeviceTokenUpdate) SetToken(b []byte) *DeviceTokenUpdate {
	dtu.mutation.SetToken(b)
	return dtu
}

// ClearToken clears the value of the "token" field.
func (dtu *DeviceTokenUpdate) ClearToken() *DeviceTokenUpdate {
	dtu.mutation.ClearToken()
	return dtu
}

// SetExpiry sets the "expiry" field.
func (dtu *DeviceTokenUpdate) SetExpiry(t time.Time) *DeviceTokenUpdate {
	dtu.mutation.SetExpiry(t)
	return dtu
}

// SetLastRequest sets the "last_request" field.
func (dtu *DeviceTokenUpdate) SetLastRequest(t time.Time) *DeviceTokenUpdate {
	dtu.mutation.SetLastRequest(t)
	return dtu
}

// SetPollInterval sets the "poll_interval" field.
func (dtu *DeviceTokenUpdate) SetPollInterval(i int) *DeviceTokenUpdate {
	dtu.mutation.ResetPollInterval()
	dtu.mutation.SetPollInterval(i)
	return dtu
}

// AddPollInterval adds i to the "poll_interval" field.
func (dtu *DeviceTokenUpdate) AddPollInterval(i int) *DeviceTokenUpdate {
	dtu.mutation.AddPollInterval(i)
	return dtu
}

// SetCodeChallenge sets the "code_challenge" field.
func (dtu *DeviceTokenUpdate) SetCodeChallenge(s string) *DeviceTokenUpdate {
	dtu.mutation.SetCodeChallenge(s)
	return dtu
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (dtu *DeviceTokenUpdate) SetNillableCodeChallenge(s *string) *DeviceTokenUpdate {
	if s != nil {
		dtu.SetCodeChallenge(*s)
	}
	return dtu
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (dtu *DeviceTokenUpdate) SetCodeChallengeMethod(s string) *DeviceTokenUpdate {
	dtu.mutation.SetCodeChallengeMethod(s)
	return dtu
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (dtu *DeviceTokenUpdate) SetNillableCodeChallengeMethod(s *string) *DeviceTokenUpdate {
	if s != nil {
		dtu.SetCodeChallengeMethod(*s)
	}
	return dtu
}

// Mutation returns the DeviceTokenMutation object of the builder.
func (dtu *DeviceTokenUpdate) Mutation() *DeviceTokenMutation {
	return dtu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dtu *DeviceTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dtu.sqlSave, dtu.mutation, dtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dtu *DeviceTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := dtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dtu *DeviceTokenUpdate) Exec(ctx context.Context) error {
	_, err := dtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtu *DeviceTokenUpdate) ExecX(ctx context.Context) {
	if err := dtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtu *DeviceTokenUpdate) check() error {
	if v, ok := dtu.mutation.DeviceCode(); ok {
		if err := devicetoken.DeviceCodeValidator(v); err != nil {
			return &ValidationError{Name: "device_code", err: fmt.Errorf(`db: validator failed for field "DeviceToken.device_code": %w`, err)}
		}
	}
	if v, ok := dtu.mutation.Status(); ok {
		if err := devicetoken.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "DeviceToken.status": %w`, err)}
		}
	}
	return nil
}

func (dtu *DeviceTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(devicetoken.Table, devicetoken.Columns, sqlgraph.NewFieldSpec(devicetoken.FieldID, field.TypeInt))
	if ps := dtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtu.mutation.DeviceCode(); ok {
		_spec.SetField(devicetoken.FieldDeviceCode, field.TypeString, value)
	}
	if value, ok := dtu.mutation.Status(); ok {
		_spec.SetField(devicetoken.FieldStatus, field.TypeString, value)
	}
	if value, ok := dtu.mutation.Token(); ok {
		_spec.SetField(devicetoken.FieldToken, field.TypeBytes, value)
	}
	if dtu.mutation.TokenCleared() {
		_spec.ClearField(devicetoken.FieldToken, field.TypeBytes)
	}
	if value, ok := dtu.mutation.Expiry(); ok {
		_spec.SetField(devicetoken.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := dtu.mutation.LastRequest(); ok {
		_spec.SetField(devicetoken.FieldLastRequest, field.TypeTime, value)
	}
	if value, ok := dtu.mutation.PollInterval(); ok {
		_spec.SetField(devicetoken.FieldPollInterval, field.TypeInt, value)
	}
	if value, ok := dtu.mutation.AddedPollInterval(); ok {
		_spec.AddField(devicetoken.FieldPollInterval, field.TypeInt, value)
	}
	if value, ok := dtu.mutation.CodeChallenge(); ok {
		_spec.SetField(devicetoken.FieldCodeChallenge, field.TypeString, value)
	}
	if value, ok := dtu.mutation.CodeChallengeMethod(); ok {
		_spec.SetField(devicetoken.FieldCodeChallengeMethod, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicetoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dtu.mutation.done = true
	return n, nil
}

// DeviceTokenUpdateOne is the builder for updating a single DeviceToken entity.
type DeviceTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceTokenMutation
}

// SetDeviceCode sets the "device_code" field.
func (dtuo *DeviceTokenUpdateOne) SetDeviceCode(s string) *DeviceTokenUpdateOne {
	dtuo.mutation.SetDeviceCode(s)
	return dtuo
}

// SetStatus sets the "status" field.
func (dtuo *DeviceTokenUpdateOne) SetStatus(s string) *DeviceTokenUpdateOne {
	dtuo.mutation.SetStatus(s)
	return dtuo
}

// SetToken sets the "token" field.
func (dtuo *DeviceTokenUpdateOne) SetToken(b []byte) *DeviceTokenUpdateOne {
	dtuo.mutation.SetToken(b)
	return dtuo
}

// ClearToken clears the value of the "token" field.
func (dtuo *DeviceTokenUpdateOne) ClearToken() *DeviceTokenUpdateOne {
	dtuo.mutation.ClearToken()
	return dtuo
}

// SetExpiry sets the "expiry" field.
func (dtuo *DeviceTokenUpdateOne) SetExpiry(t time.Time) *DeviceTokenUpdateOne {
	dtuo.mutation.SetExpiry(t)
	return dtuo
}

// SetLastRequest sets the "last_request" field.
func (dtuo *DeviceTokenUpdateOne) SetLastRequest(t time.Time) *DeviceTokenUpdateOne {
	dtuo.mutation.SetLastRequest(t)
	return dtuo
}

// SetPollInterval sets the "poll_interval" field.
func (dtuo *DeviceTokenUpdateOne) SetPollInterval(i int) *DeviceTokenUpdateOne {
	dtuo.mutation.ResetPollInterval()
	dtuo.mutation.SetPollInterval(i)
	return dtuo
}

// AddPollInterval adds i to the "poll_interval" field.
func (dtuo *DeviceTokenUpdateOne) AddPollInterval(i int) *DeviceTokenUpdateOne {
	dtuo.mutation.AddPollInterval(i)
	return dtuo
}

// SetCodeChallenge sets the "code_challenge" field.
func (dtuo *DeviceTokenUpdateOne) SetCodeChallenge(s string) *DeviceTokenUpdateOne {
	dtuo.mutation.SetCodeChallenge(s)
	return dtuo
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (dtuo *DeviceTokenUpdateOne) SetNillableCodeChallenge(s *string) *DeviceTokenUpdateOne {
	if s != nil {
		dtuo.SetCodeChallenge(*s)
	}
	return dtuo
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (dtuo *DeviceTokenUpdateOne) SetCodeChallengeMethod(s string) *DeviceTokenUpdateOne {
	dtuo.mutation.SetCodeChallengeMethod(s)
	return dtuo
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (dtuo *DeviceTokenUpdateOne) SetNillableCodeChallengeMethod(s *string) *DeviceTokenUpdateOne {
	if s != nil {
		dtuo.SetCodeChallengeMethod(*s)
	}
	return dtuo
}

// Mutation returns the DeviceTokenMutation object of the builder.
func (dtuo *DeviceTokenUpdateOne) Mutation() *DeviceTokenMutation {
	return dtuo.mutation
}

// Where appends a list predicates to the DeviceTokenUpdate builder.
func (dtuo *DeviceTokenUpdateOne) Where(ps ...predicate.DeviceToken) *DeviceTokenUpdateOne {
	dtuo.mutation.Where(ps...)
	return dtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dtuo *DeviceTokenUpdateOne) Select(field string, fields ...string) *DeviceTokenUpdateOne {
	dtuo.fields = append([]string{field}, fields...)
	return dtuo
}

// Save executes the query and returns the updated DeviceToken entity.
func (dtuo *DeviceTokenUpdateOne) Save(ctx context.Context) (*DeviceToken, error) {
	return withHooks(ctx, dtuo.sqlSave, dtuo.mutation, dtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dtuo *DeviceTokenUpdateOne) SaveX(ctx context.Context) *DeviceToken {
	node, err := dtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dtuo *DeviceTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := dtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtuo *DeviceTokenUpdateOne) ExecX(ctx context.Context) {
	if err := dtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtuo *DeviceTokenUpdateOne) check() error {
	if v, ok := dtuo.mutation.DeviceCode(); ok {
		if err := devicetoken.DeviceCodeValidator(v); err != nil {
			return &ValidationError{Name: "device_code", err: fmt.Errorf(`db: validator failed for field "DeviceToken.device_code": %w`, err)}
		}
	}
	if v, ok := dtuo.mutation.Status(); ok {
		if err := devicetoken.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "DeviceToken.status": %w`, err)}
		}
	}
	return nil
}

func (dtuo *DeviceTokenUpdateOne) sqlSave(ctx context.Context) (_node *DeviceToken, err error) {
	if err := dtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(devicetoken.Table, devicetoken.Columns, sqlgraph.NewFieldSpec(devicetoken.FieldID, field.TypeInt))
	id, ok := dtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "DeviceToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicetoken.FieldID)
		for _, f := range fields {
			if !devicetoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != devicetoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtuo.mutation.DeviceCode(); ok {
		_spec.SetField(devicetoken.FieldDeviceCode, field.TypeString, value)
	}
	if value, ok := dtuo.mutation.Status(); ok {
		_spec.SetField(devicetoken.FieldStatus, field.TypeString, value)
	}
	if value, ok := dtuo.mutation.Token(); ok {
		_spec.SetField(devicetoken.FieldToken, field.TypeBytes, value)
	}
	if dtuo.mutation.TokenCleared() {
		_spec.ClearField(devicetoken.FieldToken, field.TypeBytes)
	}
	if value, ok := dtuo.mutation.Expiry(); ok {
		_spec.SetField(devicetoken.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := dtuo.mutation.LastRequest(); ok {
		_spec.SetField(devicetoken.FieldLastRequest, field.TypeTime, value)
	}
	if value, ok := dtuo.mutation.PollInterval(); ok {
		_spec.SetField(devicetoken.FieldPollInterval, field.TypeInt, value)
	}
	if value, ok := dtuo.mutation.AddedPollInterval(); ok {
		_spec.AddField(devicetoken.FieldPollInterval, field.TypeInt, value)
	}
	if value, ok := dtuo.mutation.CodeChallenge(); ok {
		_spec.SetField(devicetoken.FieldCodeChallenge, field.TypeString, value)
	}
	if value, ok := dtuo.mutation.CodeChallengeMethod(); ok {
		_spec.SetField(devicetoken.FieldCodeChallengeMethod, field.TypeString, value)
	}
	_node = &DeviceToken{config: dtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicetoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dtuo.mutation.done = true
	return _node, nil
}
