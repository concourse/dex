// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/oauth2client"
)

// OAuth2ClientCreate is the builder for creating a OAuth2Client entity.
type OAuth2ClientCreate struct {
	config
	mutation *OAuth2ClientMutation
	hooks    []Hook
}

// SetSecret sets the "secret" field.
func (oc *OAuth2ClientCreate) SetSecret(s string) *OAuth2ClientCreate {
	oc.mutation.SetSecret(s)
	return oc
}

// SetRedirectUris sets the "redirect_uris" field.
func (oc *OAuth2ClientCreate) SetRedirectUris(s []string) *OAuth2ClientCreate {
	oc.mutation.SetRedirectUris(s)
	return oc
}

// SetTrustedPeers sets the "trusted_peers" field.
func (oc *OAuth2ClientCreate) SetTrustedPeers(s []string) *OAuth2ClientCreate {
	oc.mutation.SetTrustedPeers(s)
	return oc
}

// SetPublic sets the "public" field.
func (oc *OAuth2ClientCreate) SetPublic(b bool) *OAuth2ClientCreate {
	oc.mutation.SetPublic(b)
	return oc
}

// SetName sets the "name" field.
func (oc *OAuth2ClientCreate) SetName(s string) *OAuth2ClientCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetLogoURL sets the "logo_url" field.
func (oc *OAuth2ClientCreate) SetLogoURL(s string) *OAuth2ClientCreate {
	oc.mutation.SetLogoURL(s)
	return oc
}

// SetID sets the "id" field.
func (oc *OAuth2ClientCreate) SetID(s string) *OAuth2ClientCreate {
	oc.mutation.SetID(s)
	return oc
}

// Mutation returns the OAuth2ClientMutation object of the builder.
func (oc *OAuth2ClientCreate) Mutation() *OAuth2ClientMutation {
	return oc.mutation
}

// Save creates the OAuth2Client in the database.
func (oc *OAuth2ClientCreate) Save(ctx context.Context) (*OAuth2Client, error) {
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OAuth2ClientCreate) SaveX(ctx context.Context) *OAuth2Client {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OAuth2ClientCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OAuth2ClientCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OAuth2ClientCreate) check() error {
	if _, ok := oc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`db: missing required field "OAuth2Client.secret"`)}
	}
	if v, ok := oc.mutation.Secret(); ok {
		if err := oauth2client.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.secret": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`db: missing required field "OAuth2Client.public"`)}
	}
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "OAuth2Client.name"`)}
	}
	if v, ok := oc.mutation.Name(); ok {
		if err := oauth2client.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.name": %w`, err)}
		}
	}
	if _, ok := oc.mutation.LogoURL(); !ok {
		return &ValidationError{Name: "logo_url", err: errors.New(`db: missing required field "OAuth2Client.logo_url"`)}
	}
	if v, ok := oc.mutation.LogoURL(); ok {
		if err := oauth2client.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.logo_url": %w`, err)}
		}
	}
	if v, ok := oc.mutation.ID(); ok {
		if err := oauth2client.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`db: validator failed for field "OAuth2Client.id": %w`, err)}
		}
	}
	return nil
}

func (oc *OAuth2ClientCreate) sqlSave(ctx context.Context) (*OAuth2Client, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OAuth2Client.ID type: %T", _spec.ID.Value)
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OAuth2ClientCreate) createSpec() (*OAuth2Client, *sqlgraph.CreateSpec) {
	var (
		_node = &OAuth2Client{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(oauth2client.Table, sqlgraph.NewFieldSpec(oauth2client.FieldID, field.TypeString))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.Secret(); ok {
		_spec.SetField(oauth2client.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := oc.mutation.RedirectUris(); ok {
		_spec.SetField(oauth2client.FieldRedirectUris, field.TypeJSON, value)
		_node.RedirectUris = value
	}
	if value, ok := oc.mutation.TrustedPeers(); ok {
		_spec.SetField(oauth2client.FieldTrustedPeers, field.TypeJSON, value)
		_node.TrustedPeers = value
	}
	if value, ok := oc.mutation.Public(); ok {
		_spec.SetField(oauth2client.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(oauth2client.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oc.mutation.LogoURL(); ok {
		_spec.SetField(oauth2client.FieldLogoURL, field.TypeString, value)
		_node.LogoURL = value
	}
	return _node, _spec
}

// OAuth2ClientCreateBulk is the builder for creating many OAuth2Client entities in bulk.
type OAuth2ClientCreateBulk struct {
	config
	err      error
	builders []*OAuth2ClientCreate
}

// Save creates the OAuth2Client entities in the database.
func (ocb *OAuth2ClientCreateBulk) Save(ctx context.Context) ([]*OAuth2Client, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*OAuth2Client, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OAuth2ClientMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OAuth2ClientCreateBulk) SaveX(ctx context.Context) []*OAuth2Client {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OAuth2ClientCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OAuth2ClientCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
