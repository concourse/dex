// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/oauth2client"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// OAuth2ClientUpdate is the builder for updating OAuth2Client entities.
type OAuth2ClientUpdate struct {
	config
	hooks    []Hook
	mutation *OAuth2ClientMutation
}

// Where appends a list predicates to the OAuth2ClientUpdate builder.
func (ou *OAuth2ClientUpdate) Where(ps ...predicate.OAuth2Client) *OAuth2ClientUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetSecret sets the "secret" field.
func (ou *OAuth2ClientUpdate) SetSecret(s string) *OAuth2ClientUpdate {
	ou.mutation.SetSecret(s)
	return ou
}

// SetRedirectUris sets the "redirect_uris" field.
func (ou *OAuth2ClientUpdate) SetRedirectUris(s []string) *OAuth2ClientUpdate {
	ou.mutation.SetRedirectUris(s)
	return ou
}

// AppendRedirectUris appends s to the "redirect_uris" field.
func (ou *OAuth2ClientUpdate) AppendRedirectUris(s []string) *OAuth2ClientUpdate {
	ou.mutation.AppendRedirectUris(s)
	return ou
}

// ClearRedirectUris clears the value of the "redirect_uris" field.
func (ou *OAuth2ClientUpdate) ClearRedirectUris() *OAuth2ClientUpdate {
	ou.mutation.ClearRedirectUris()
	return ou
}

// SetTrustedPeers sets the "trusted_peers" field.
func (ou *OAuth2ClientUpdate) SetTrustedPeers(s []string) *OAuth2ClientUpdate {
	ou.mutation.SetTrustedPeers(s)
	return ou
}

// AppendTrustedPeers appends s to the "trusted_peers" field.
func (ou *OAuth2ClientUpdate) AppendTrustedPeers(s []string) *OAuth2ClientUpdate {
	ou.mutation.AppendTrustedPeers(s)
	return ou
}

// ClearTrustedPeers clears the value of the "trusted_peers" field.
func (ou *OAuth2ClientUpdate) ClearTrustedPeers() *OAuth2ClientUpdate {
	ou.mutation.ClearTrustedPeers()
	return ou
}

// SetPublic sets the "public" field.
func (ou *OAuth2ClientUpdate) SetPublic(b bool) *OAuth2ClientUpdate {
	ou.mutation.SetPublic(b)
	return ou
}

// SetName sets the "name" field.
func (ou *OAuth2ClientUpdate) SetName(s string) *OAuth2ClientUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetLogoURL sets the "logo_url" field.
func (ou *OAuth2ClientUpdate) SetLogoURL(s string) *OAuth2ClientUpdate {
	ou.mutation.SetLogoURL(s)
	return ou
}

// Mutation returns the OAuth2ClientMutation object of the builder.
func (ou *OAuth2ClientUpdate) Mutation() *OAuth2ClientMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OAuth2ClientUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OAuth2ClientMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OAuth2ClientUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OAuth2ClientUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OAuth2ClientUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OAuth2ClientUpdate) check() error {
	if v, ok := ou.mutation.Secret(); ok {
		if err := oauth2client.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.secret": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Name(); ok {
		if err := oauth2client.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.LogoURL(); ok {
		if err := oauth2client.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.logo_url": %w`, err)}
		}
	}
	return nil
}

func (ou *OAuth2ClientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauth2client.Table, oauth2client.Columns, sqlgraph.NewFieldSpec(oauth2client.FieldID, field.TypeString))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Secret(); ok {
		_spec.SetField(oauth2client.FieldSecret, field.TypeString, value)
	}
	if value, ok := ou.mutation.RedirectUris(); ok {
		_spec.SetField(oauth2client.FieldRedirectUris, field.TypeJSON, value)
	}
	if value, ok := ou.mutation.AppendedRedirectUris(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oauth2client.FieldRedirectUris, value)
		})
	}
	if ou.mutation.RedirectUrisCleared() {
		_spec.ClearField(oauth2client.FieldRedirectUris, field.TypeJSON)
	}
	if value, ok := ou.mutation.TrustedPeers(); ok {
		_spec.SetField(oauth2client.FieldTrustedPeers, field.TypeJSON, value)
	}
	if value, ok := ou.mutation.AppendedTrustedPeers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oauth2client.FieldTrustedPeers, value)
		})
	}
	if ou.mutation.TrustedPeersCleared() {
		_spec.ClearField(oauth2client.FieldTrustedPeers, field.TypeJSON)
	}
	if value, ok := ou.mutation.Public(); ok {
		_spec.SetField(oauth2client.FieldPublic, field.TypeBool, value)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(oauth2client.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.LogoURL(); ok {
		_spec.SetField(oauth2client.FieldLogoURL, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth2client.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OAuth2ClientUpdateOne is the builder for updating a single OAuth2Client entity.
type OAuth2ClientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OAuth2ClientMutation
}

// SetSecret sets the "secret" field.
func (ouo *OAuth2ClientUpdateOne) SetSecret(s string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetSecret(s)
	return ouo
}

// SetRedirectUris sets the "redirect_uris" field.
func (ouo *OAuth2ClientUpdateOne) SetRedirectUris(s []string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetRedirectUris(s)
	return ouo
}

// AppendRedirectUris appends s to the "redirect_uris" field.
func (ouo *OAuth2ClientUpdateOne) AppendRedirectUris(s []string) *OAuth2ClientUpdateOne {
	ouo.mutation.AppendRedirectUris(s)
	return ouo
}

// ClearRedirectUris clears the value of the "redirect_uris" field.
func (ouo *OAuth2ClientUpdateOne) ClearRedirectUris() *OAuth2ClientUpdateOne {
	ouo.mutation.ClearRedirectUris()
	return ouo
}

// SetTrustedPeers sets the "trusted_peers" field.
func (ouo *OAuth2ClientUpdateOne) SetTrustedPeers(s []string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetTrustedPeers(s)
	return ouo
}

// AppendTrustedPeers appends s to the "trusted_peers" field.
func (ouo *OAuth2ClientUpdateOne) AppendTrustedPeers(s []string) *OAuth2ClientUpdateOne {
	ouo.mutation.AppendTrustedPeers(s)
	return ouo
}

// ClearTrustedPeers clears the value of the "trusted_peers" field.
func (ouo *OAuth2ClientUpdateOne) ClearTrustedPeers() *OAuth2ClientUpdateOne {
	ouo.mutation.ClearTrustedPeers()
	return ouo
}

// SetPublic sets the "public" field.
func (ouo *OAuth2ClientUpdateOne) SetPublic(b bool) *OAuth2ClientUpdateOne {
	ouo.mutation.SetPublic(b)
	return ouo
}

// SetName sets the "name" field.
func (ouo *OAuth2ClientUpdateOne) SetName(s string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetLogoURL sets the "logo_url" field.
func (ouo *OAuth2ClientUpdateOne) SetLogoURL(s string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetLogoURL(s)
	return ouo
}

// Mutation returns the OAuth2ClientMutation object of the builder.
func (ouo *OAuth2ClientUpdateOne) Mutation() *OAuth2ClientMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OAuth2ClientUpdate builder.
func (ouo *OAuth2ClientUpdateOne) Where(ps ...predicate.OAuth2Client) *OAuth2ClientUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OAuth2ClientUpdateOne) Select(field string, fields ...string) *OAuth2ClientUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated OAuth2Client entity.
func (ouo *OAuth2ClientUpdateOne) Save(ctx context.Context) (*OAuth2Client, error) {
	return withHooks[*OAuth2Client, OAuth2ClientMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OAuth2ClientUpdateOne) SaveX(ctx context.Context) *OAuth2Client {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OAuth2ClientUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OAuth2ClientUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OAuth2ClientUpdateOne) check() error {
	if v, ok := ouo.mutation.Secret(); ok {
		if err := oauth2client.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.secret": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Name(); ok {
		if err := oauth2client.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.LogoURL(); ok {
		if err := oauth2client.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.logo_url": %w`, err)}
		}
	}
	return nil
}

func (ouo *OAuth2ClientUpdateOne) sqlSave(ctx context.Context) (_node *OAuth2Client, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauth2client.Table, oauth2client.Columns, sqlgraph.NewFieldSpec(oauth2client.FieldID, field.TypeString))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "OAuth2Client.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauth2client.FieldID)
		for _, f := range fields {
			if !oauth2client.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != oauth2client.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Secret(); ok {
		_spec.SetField(oauth2client.FieldSecret, field.TypeString, value)
	}
	if value, ok := ouo.mutation.RedirectUris(); ok {
		_spec.SetField(oauth2client.FieldRedirectUris, field.TypeJSON, value)
	}
	if value, ok := ouo.mutation.AppendedRedirectUris(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oauth2client.FieldRedirectUris, value)
		})
	}
	if ouo.mutation.RedirectUrisCleared() {
		_spec.ClearField(oauth2client.FieldRedirectUris, field.TypeJSON)
	}
	if value, ok := ouo.mutation.TrustedPeers(); ok {
		_spec.SetField(oauth2client.FieldTrustedPeers, field.TypeJSON, value)
	}
	if value, ok := ouo.mutation.AppendedTrustedPeers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oauth2client.FieldTrustedPeers, value)
		})
	}
	if ouo.mutation.TrustedPeersCleared() {
		_spec.ClearField(oauth2client.FieldTrustedPeers, field.TypeJSON)
	}
	if value, ok := ouo.mutation.Public(); ok {
		_spec.SetField(oauth2client.FieldPublic, field.TypeBool, value)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(oauth2client.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.LogoURL(); ok {
		_spec.SetField(oauth2client.FieldLogoURL, field.TypeString, value)
	}
	_node = &OAuth2Client{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth2client.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
