// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/washanhanzi/ent-same-package-name/ent/predicate"
	"github.com/washanhanzi/ent-same-package-name/ent/tableb"
)

// TableBDelete is the builder for deleting a TableB entity.
type TableBDelete struct {
	config
	hooks    []Hook
	mutation *TableBMutation
}

// Where adds a new predicate to the TableBDelete builder.
func (tb *TableBDelete) Where(ps ...predicate.TableB) *TableBDelete {
	tb.mutation.predicates = append(tb.mutation.predicates, ps...)
	return tb
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tb *TableBDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tb.hooks) == 0 {
		affected, err = tb.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TableBMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tb.mutation = mutation
			affected, err = tb.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tb.hooks) - 1; i >= 0; i-- {
			mut = tb.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tb.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tb *TableBDelete) ExecX(ctx context.Context) int {
	n, err := tb.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tb *TableBDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tableb.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tableb.FieldID,
			},
		},
	}
	if ps := tb.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tb.driver, _spec)
}

// TableBDeleteOne is the builder for deleting a single TableB entity.
type TableBDeleteOne struct {
	tb *TableBDelete
}

// Exec executes the deletion query.
func (tbo *TableBDeleteOne) Exec(ctx context.Context) error {
	n, err := tbo.tb.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tableb.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tbo *TableBDeleteOne) ExecX(ctx context.Context) {
	tbo.tb.ExecX(ctx)
}