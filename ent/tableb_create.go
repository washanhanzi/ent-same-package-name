// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/washanhanzi/ent-same-package-name/b/entity"
	"github.com/washanhanzi/ent-same-package-name/ent/tableb"
)

// TableBCreate is the builder for creating a TableB entity.
type TableBCreate struct {
	config
	mutation *TableBMutation
	hooks    []Hook
}

// SetB sets the "b" field.
func (tb *TableBCreate) SetB(e entity.B) *TableBCreate {
	tb.mutation.SetB(e)
	return tb
}

// Mutation returns the TableBMutation object of the builder.
func (tb *TableBCreate) Mutation() *TableBMutation {
	return tb.mutation
}

// Save creates the TableB in the database.
func (tb *TableBCreate) Save(ctx context.Context) (*TableB, error) {
	var (
		err  error
		node *TableB
	)
	if len(tb.hooks) == 0 {
		if err = tb.check(); err != nil {
			return nil, err
		}
		node, err = tb.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TableBMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tb.check(); err != nil {
				return nil, err
			}
			tb.mutation = mutation
			node, err = tb.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tb.hooks) - 1; i >= 0; i-- {
			mut = tb.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tb.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tb *TableBCreate) SaveX(ctx context.Context) *TableB {
	v, err := tb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (tb *TableBCreate) check() error {
	if _, ok := tb.mutation.B(); !ok {
		return &ValidationError{Name: "b", err: errors.New("ent: missing required field \"b\"")}
	}
	return nil
}

func (tb *TableBCreate) sqlSave(ctx context.Context) (*TableB, error) {
	_node, _spec := tb.createSpec()
	if err := sqlgraph.CreateNode(ctx, tb.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tb *TableBCreate) createSpec() (*TableB, *sqlgraph.CreateSpec) {
	var (
		_node = &TableB{config: tb.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tableb.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tableb.FieldID,
			},
		}
	)
	if value, ok := tb.mutation.B(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: tableb.FieldB,
		})
		_node.B = value
	}
	return _node, _spec
}

// TableBCreateBulk is the builder for creating many TableB entities in bulk.
type TableBCreateBulk struct {
	config
	builders []*TableBCreate
}

// Save creates the TableB entities in the database.
func (tbb *TableBCreateBulk) Save(ctx context.Context) ([]*TableB, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tbb.builders))
	nodes := make([]*TableB, len(tbb.builders))
	mutators := make([]Mutator, len(tbb.builders))
	for i := range tbb.builders {
		func(i int, root context.Context) {
			builder := tbb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TableBMutation)
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
					_, err = mutators[i+1].Mutate(root, tbb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tbb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tbb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tbb *TableBCreateBulk) SaveX(ctx context.Context) []*TableB {
	v, err := tbb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
