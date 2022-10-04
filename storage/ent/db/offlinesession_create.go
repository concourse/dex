// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/offlinesession"
)

// OfflineSessionCreate is the builder for creating a OfflineSession entity.
type OfflineSessionCreate struct {
	config
	mutation *OfflineSessionMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (osc *OfflineSessionCreate) SetUserID(s string) *OfflineSessionCreate {
	osc.mutation.SetUserID(s)
	return osc
}

// SetConnID sets the "conn_id" field.
func (osc *OfflineSessionCreate) SetConnID(s string) *OfflineSessionCreate {
	osc.mutation.SetConnID(s)
	return osc
}

// SetRefresh sets the "refresh" field.
func (osc *OfflineSessionCreate) SetRefresh(b []byte) *OfflineSessionCreate {
	osc.mutation.SetRefresh(b)
	return osc
}

// SetConnectorData sets the "connector_data" field.
func (osc *OfflineSessionCreate) SetConnectorData(b []byte) *OfflineSessionCreate {
	osc.mutation.SetConnectorData(b)
	return osc
}

// SetID sets the "id" field.
func (osc *OfflineSessionCreate) SetID(s string) *OfflineSessionCreate {
	osc.mutation.SetID(s)
	return osc
}

// Mutation returns the OfflineSessionMutation object of the builder.
func (osc *OfflineSessionCreate) Mutation() *OfflineSessionMutation {
	return osc.mutation
}

// Save creates the OfflineSession in the database.
func (osc *OfflineSessionCreate) Save(ctx context.Context) (*OfflineSession, error) {
	var (
		err  error
		node *OfflineSession
	)
	if len(osc.hooks) == 0 {
		if err = osc.check(); err != nil {
			return nil, err
		}
		node, err = osc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfflineSessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = osc.check(); err != nil {
				return nil, err
			}
			osc.mutation = mutation
			if node, err = osc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(osc.hooks) - 1; i >= 0; i-- {
			if osc.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = osc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, osc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OfflineSession)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OfflineSessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (osc *OfflineSessionCreate) SaveX(ctx context.Context) *OfflineSession {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (osc *OfflineSessionCreate) Exec(ctx context.Context) error {
	_, err := osc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osc *OfflineSessionCreate) ExecX(ctx context.Context) {
	if err := osc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osc *OfflineSessionCreate) check() error {
	if _, ok := osc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`db: missing required field "OfflineSession.user_id"`)}
	}
	if v, ok := osc.mutation.UserID(); ok {
		if err := offlinesession.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`db: validator failed for field "OfflineSession.user_id": %w`, err)}
		}
	}
	if _, ok := osc.mutation.ConnID(); !ok {
		return &ValidationError{Name: "conn_id", err: errors.New(`db: missing required field "OfflineSession.conn_id"`)}
	}
	if v, ok := osc.mutation.ConnID(); ok {
		if err := offlinesession.ConnIDValidator(v); err != nil {
			return &ValidationError{Name: "conn_id", err: fmt.Errorf(`db: validator failed for field "OfflineSession.conn_id": %w`, err)}
		}
	}
	if _, ok := osc.mutation.Refresh(); !ok {
		return &ValidationError{Name: "refresh", err: errors.New(`db: missing required field "OfflineSession.refresh"`)}
	}
	if v, ok := osc.mutation.ID(); ok {
		if err := offlinesession.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`db: validator failed for field "OfflineSession.id": %w`, err)}
		}
	}
	return nil
}

func (osc *OfflineSessionCreate) sqlSave(ctx context.Context) (*OfflineSession, error) {
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OfflineSession.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (osc *OfflineSessionCreate) createSpec() (*OfflineSession, *sqlgraph.CreateSpec) {
	var (
		_node = &OfflineSession{config: osc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: offlinesession.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: offlinesession.FieldID,
			},
		}
	)
	if id, ok := osc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := osc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: offlinesession.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := osc.mutation.ConnID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: offlinesession.FieldConnID,
		})
		_node.ConnID = value
	}
	if value, ok := osc.mutation.Refresh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: offlinesession.FieldRefresh,
		})
		_node.Refresh = value
	}
	if value, ok := osc.mutation.ConnectorData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: offlinesession.FieldConnectorData,
		})
		_node.ConnectorData = &value
	}
	return _node, _spec
}

// OfflineSessionCreateBulk is the builder for creating many OfflineSession entities in bulk.
type OfflineSessionCreateBulk struct {
	config
	builders []*OfflineSessionCreate
}

// Save creates the OfflineSession entities in the database.
func (oscb *OfflineSessionCreateBulk) Save(ctx context.Context) ([]*OfflineSession, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OfflineSession, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OfflineSessionMutation)
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
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OfflineSessionCreateBulk) SaveX(ctx context.Context) []*OfflineSession {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oscb *OfflineSessionCreateBulk) Exec(ctx context.Context) error {
	_, err := oscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscb *OfflineSessionCreateBulk) ExecX(ctx context.Context) {
	if err := oscb.Exec(ctx); err != nil {
		panic(err)
	}
}
