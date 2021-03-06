// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lemon-mint/open-backend/ent/group"
	"github.com/lemon-mint/open-backend/ent/resource"
)

// ResourceCreate is the builder for creating a Resource entity.
type ResourceCreate struct {
	config
	mutation *ResourceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *ResourceCreate) SetName(s string) *ResourceCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetAcls sets the "acls" field.
func (rc *ResourceCreate) SetAcls(s []string) *ResourceCreate {
	rc.mutation.SetAcls(s)
	return rc
}

// SetDefault sets the "default" field.
func (rc *ResourceCreate) SetDefault(s []string) *ResourceCreate {
	rc.mutation.SetDefault(s)
	return rc
}

// SetID sets the "id" field.
func (rc *ResourceCreate) SetID(u uuid.UUID) *ResourceCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetOwnerID sets the "owner" edge to the Group entity by ID.
func (rc *ResourceCreate) SetOwnerID(id int) *ResourceCreate {
	rc.mutation.SetOwnerID(id)
	return rc
}

// SetNillableOwnerID sets the "owner" edge to the Group entity by ID if the given value is not nil.
func (rc *ResourceCreate) SetNillableOwnerID(id *int) *ResourceCreate {
	if id != nil {
		rc = rc.SetOwnerID(*id)
	}
	return rc
}

// SetOwner sets the "owner" edge to the Group entity.
func (rc *ResourceCreate) SetOwner(g *Group) *ResourceCreate {
	return rc.SetOwnerID(g.ID)
}

// Mutation returns the ResourceMutation object of the builder.
func (rc *ResourceCreate) Mutation() *ResourceMutation {
	return rc.mutation
}

// Save creates the Resource in the database.
func (rc *ResourceCreate) Save(ctx context.Context) (*Resource, error) {
	var (
		err  error
		node *Resource
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResourceCreate) SaveX(ctx context.Context) *Resource {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResourceCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResourceCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResourceCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := rc.mutation.Acls(); !ok {
		return &ValidationError{Name: "acls", err: errors.New(`ent: missing required field "acls"`)}
	}
	if _, ok := rc.mutation.Default(); !ok {
		return &ValidationError{Name: "default", err: errors.New(`ent: missing required field "default"`)}
	}
	return nil
}

func (rc *ResourceCreate) sqlSave(ctx context.Context) (*Resource, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (rc *ResourceCreate) createSpec() (*Resource, *sqlgraph.CreateSpec) {
	var (
		_node = &Resource{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: resource.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Acls(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: resource.FieldAcls,
		})
		_node.Acls = value
	}
	if value, ok := rc.mutation.Default(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: resource.FieldDefault,
		})
		_node.Default = value
	}
	if nodes := rc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resource.OwnerTable,
			Columns: []string{resource.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_resources = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResourceCreateBulk is the builder for creating many Resource entities in bulk.
type ResourceCreateBulk struct {
	config
	builders []*ResourceCreate
}

// Save creates the Resource entities in the database.
func (rcb *ResourceCreateBulk) Save(ctx context.Context) ([]*Resource, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Resource, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResourceCreateBulk) SaveX(ctx context.Context) []*Resource {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResourceCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResourceCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
