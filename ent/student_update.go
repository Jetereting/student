// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"school/ent/predicate"
	"school/ent/student"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StudentUpdate) SetUpdatedAt(t time.Time) *StudentUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetSort sets the "sort" field.
func (su *StudentUpdate) SetSort(u uint32) *StudentUpdate {
	su.mutation.ResetSort()
	su.mutation.SetSort(u)
	return su
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (su *StudentUpdate) SetNillableSort(u *uint32) *StudentUpdate {
	if u != nil {
		su.SetSort(*u)
	}
	return su
}

// AddSort adds u to the "sort" field.
func (su *StudentUpdate) AddSort(u int32) *StudentUpdate {
	su.mutation.AddSort(u)
	return su
}

// SetStatus sets the "status" field.
func (su *StudentUpdate) SetStatus(u uint8) *StudentUpdate {
	su.mutation.ResetStatus()
	su.mutation.SetStatus(u)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *StudentUpdate) SetNillableStatus(u *uint8) *StudentUpdate {
	if u != nil {
		su.SetStatus(*u)
	}
	return su
}

// AddStatus adds u to the "status" field.
func (su *StudentUpdate) AddStatus(u int8) *StudentUpdate {
	su.mutation.AddStatus(u)
	return su
}

// ClearStatus clears the value of the "status" field.
func (su *StudentUpdate) ClearStatus() *StudentUpdate {
	su.mutation.ClearStatus()
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s []string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// AppendName appends s to the "name" field.
func (su *StudentUpdate) AppendName(s []string) *StudentUpdate {
	su.mutation.AppendName(s)
	return su
}

// SetIDCard sets the "id_card" field.
func (su *StudentUpdate) SetIDCard(s []string) *StudentUpdate {
	su.mutation.SetIDCard(s)
	return su
}

// AppendIDCard appends s to the "id_card" field.
func (su *StudentUpdate) AppendIDCard(s []string) *StudentUpdate {
	su.mutation.AppendIDCard(s)
	return su
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StudentUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Sort(); ok {
		_spec.SetField(student.FieldSort, field.TypeUint32, value)
	}
	if value, ok := su.mutation.AddedSort(); ok {
		_spec.AddField(student.FieldSort, field.TypeUint32, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(student.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AddedStatus(); ok {
		_spec.AddField(student.FieldStatus, field.TypeUint8, value)
	}
	if su.mutation.StatusCleared() {
		_spec.ClearField(student.FieldStatus, field.TypeUint8)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeJSON, value)
	}
	if value, ok := su.mutation.AppendedName(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, student.FieldName, value)
		})
	}
	if value, ok := su.mutation.IDCard(); ok {
		_spec.SetField(student.FieldIDCard, field.TypeJSON, value)
	}
	if value, ok := su.mutation.AppendedIDCard(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, student.FieldIDCard, value)
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StudentUpdateOne) SetUpdatedAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetSort sets the "sort" field.
func (suo *StudentUpdateOne) SetSort(u uint32) *StudentUpdateOne {
	suo.mutation.ResetSort()
	suo.mutation.SetSort(u)
	return suo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableSort(u *uint32) *StudentUpdateOne {
	if u != nil {
		suo.SetSort(*u)
	}
	return suo
}

// AddSort adds u to the "sort" field.
func (suo *StudentUpdateOne) AddSort(u int32) *StudentUpdateOne {
	suo.mutation.AddSort(u)
	return suo
}

// SetStatus sets the "status" field.
func (suo *StudentUpdateOne) SetStatus(u uint8) *StudentUpdateOne {
	suo.mutation.ResetStatus()
	suo.mutation.SetStatus(u)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableStatus(u *uint8) *StudentUpdateOne {
	if u != nil {
		suo.SetStatus(*u)
	}
	return suo
}

// AddStatus adds u to the "status" field.
func (suo *StudentUpdateOne) AddStatus(u int8) *StudentUpdateOne {
	suo.mutation.AddStatus(u)
	return suo
}

// ClearStatus clears the value of the "status" field.
func (suo *StudentUpdateOne) ClearStatus() *StudentUpdateOne {
	suo.mutation.ClearStatus()
	return suo
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s []string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// AppendName appends s to the "name" field.
func (suo *StudentUpdateOne) AppendName(s []string) *StudentUpdateOne {
	suo.mutation.AppendName(s)
	return suo
}

// SetIDCard sets the "id_card" field.
func (suo *StudentUpdateOne) SetIDCard(s []string) *StudentUpdateOne {
	suo.mutation.SetIDCard(s)
	return suo
}

// AppendIDCard appends s to the "id_card" field.
func (suo *StudentUpdateOne) AppendIDCard(s []string) *StudentUpdateOne {
	suo.mutation.AppendIDCard(s)
	return suo
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StudentUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Sort(); ok {
		_spec.SetField(student.FieldSort, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.AddedSort(); ok {
		_spec.AddField(student.FieldSort, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(student.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AddedStatus(); ok {
		_spec.AddField(student.FieldStatus, field.TypeUint8, value)
	}
	if suo.mutation.StatusCleared() {
		_spec.ClearField(student.FieldStatus, field.TypeUint8)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.AppendedName(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, student.FieldName, value)
		})
	}
	if value, ok := suo.mutation.IDCard(); ok {
		_spec.SetField(student.FieldIDCard, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.AppendedIDCard(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, student.FieldIDCard, value)
		})
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}