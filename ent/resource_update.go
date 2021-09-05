// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lemon-mint/open-backend/ent/group"
	"github.com/lemon-mint/open-backend/ent/predicate"
	"github.com/lemon-mint/open-backend/ent/resource"
)

// ResourceUpdate is the builder for updating Resource entities.
type ResourceUpdate struct {
	config
	hooks    []Hook
	mutation *ResourceMutation
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ru *ResourceUpdate) Where(ps ...predicate.Resource) *ResourceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *ResourceUpdate) SetName(s string) *ResourceUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetGroup sets the "group" field.
func (ru *ResourceUpdate) SetGroup(s []string) *ResourceUpdate {
	ru.mutation.SetGroup(s)
	return ru
}

// SetOthers sets the "others" field.
func (ru *ResourceUpdate) SetOthers(s []string) *ResourceUpdate {
	ru.mutation.SetOthers(s)
	return ru
}

// SetOwnerID sets the "owner" edge to the Group entity by ID.
func (ru *ResourceUpdate) SetOwnerID(id int) *ResourceUpdate {
	ru.mutation.SetOwnerID(id)
	return ru
}

// SetNillableOwnerID sets the "owner" edge to the Group entity by ID if the given value is not nil.
func (ru *ResourceUpdate) SetNillableOwnerID(id *int) *ResourceUpdate {
	if id != nil {
		ru = ru.SetOwnerID(*id)
	}
	return ru
}

// SetOwner sets the "owner" edge to the Group entity.
func (ru *ResourceUpdate) SetOwner(g *Group) *ResourceUpdate {
	return ru.SetOwnerID(g.ID)
}

// Mutation returns the ResourceMutation object of the builder.
func (ru *ResourceUpdate) Mutation() *ResourceMutation {
	return ru.mutation
}

// ClearOwner clears the "owner" edge to the Group entity.
func (ru *ResourceUpdate) ClearOwner() *ResourceUpdate {
	ru.mutation.ClearOwner()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResourceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResourceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResourceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resource.Table,
			Columns: resource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: resource.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldName,
		})
	}
	if value, ok := ru.mutation.Group(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: resource.FieldGroup,
		})
	}
	if value, ok := ru.mutation.Others(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: resource.FieldOthers,
		})
	}
	if ru.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ResourceUpdateOne is the builder for updating a single Resource entity.
type ResourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResourceMutation
}

// SetName sets the "name" field.
func (ruo *ResourceUpdateOne) SetName(s string) *ResourceUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetGroup sets the "group" field.
func (ruo *ResourceUpdateOne) SetGroup(s []string) *ResourceUpdateOne {
	ruo.mutation.SetGroup(s)
	return ruo
}

// SetOthers sets the "others" field.
func (ruo *ResourceUpdateOne) SetOthers(s []string) *ResourceUpdateOne {
	ruo.mutation.SetOthers(s)
	return ruo
}

// SetOwnerID sets the "owner" edge to the Group entity by ID.
func (ruo *ResourceUpdateOne) SetOwnerID(id int) *ResourceUpdateOne {
	ruo.mutation.SetOwnerID(id)
	return ruo
}

// SetNillableOwnerID sets the "owner" edge to the Group entity by ID if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableOwnerID(id *int) *ResourceUpdateOne {
	if id != nil {
		ruo = ruo.SetOwnerID(*id)
	}
	return ruo
}

// SetOwner sets the "owner" edge to the Group entity.
func (ruo *ResourceUpdateOne) SetOwner(g *Group) *ResourceUpdateOne {
	return ruo.SetOwnerID(g.ID)
}

// Mutation returns the ResourceMutation object of the builder.
func (ruo *ResourceUpdateOne) Mutation() *ResourceMutation {
	return ruo.mutation
}

// ClearOwner clears the "owner" edge to the Group entity.
func (ruo *ResourceUpdateOne) ClearOwner() *ResourceUpdateOne {
	ruo.mutation.ClearOwner()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResourceUpdateOne) Select(field string, fields ...string) *ResourceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Resource entity.
func (ruo *ResourceUpdateOne) Save(ctx context.Context) (*Resource, error) {
	var (
		err  error
		node *Resource
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResourceUpdateOne) SaveX(ctx context.Context) *Resource {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResourceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ResourceUpdateOne) sqlSave(ctx context.Context) (_node *Resource, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resource.Table,
			Columns: resource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: resource.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Resource.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resource.FieldID)
		for _, f := range fields {
			if !resource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != resource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldName,
		})
	}
	if value, ok := ruo.mutation.Group(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: resource.FieldGroup,
		})
	}
	if value, ok := ruo.mutation.Others(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: resource.FieldOthers,
		})
	}
	if ruo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Resource{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}