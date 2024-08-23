// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-use/ent/diary"
	"kratos-use/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DiaryDelete is the builder for deleting a Diary entity.
type DiaryDelete struct {
	config
	hooks    []Hook
	mutation *DiaryMutation
}

// Where appends a list predicates to the DiaryDelete builder.
func (dd *DiaryDelete) Where(ps ...predicate.Diary) *DiaryDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DiaryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DiaryDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DiaryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(diary.Table, sqlgraph.NewFieldSpec(diary.FieldID, field.TypeString))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DiaryDeleteOne is the builder for deleting a single Diary entity.
type DiaryDeleteOne struct {
	dd *DiaryDelete
}

// Where appends a list predicates to the DiaryDelete builder.
func (ddo *DiaryDeleteOne) Where(ps ...predicate.Diary) *DiaryDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DiaryDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{diary.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DiaryDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}
