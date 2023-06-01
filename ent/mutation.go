// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"school/ent/predicate"
	"school/ent/student"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSchool  = "School"
	TypeStudent = "Student"
)

// SchoolMutation represents an operation that mutates the School nodes in the graph.
type SchoolMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*School, error)
	predicates    []predicate.School
}

var _ ent.Mutation = (*SchoolMutation)(nil)

// schoolOption allows management of the mutation configuration using functional options.
type schoolOption func(*SchoolMutation)

// newSchoolMutation creates new mutation for the School entity.
func newSchoolMutation(c config, op Op, opts ...schoolOption) *SchoolMutation {
	m := &SchoolMutation{
		config:        c,
		op:            op,
		typ:           TypeSchool,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSchoolID sets the ID field of the mutation.
func withSchoolID(id int) schoolOption {
	return func(m *SchoolMutation) {
		var (
			err   error
			once  sync.Once
			value *School
		)
		m.oldValue = func(ctx context.Context) (*School, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().School.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSchool sets the old School of the mutation.
func withSchool(node *School) schoolOption {
	return func(m *SchoolMutation) {
		m.oldValue = func(context.Context) (*School, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SchoolMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SchoolMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SchoolMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SchoolMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().School.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// Where appends a list predicates to the SchoolMutation builder.
func (m *SchoolMutation) Where(ps ...predicate.School) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SchoolMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SchoolMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.School, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SchoolMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SchoolMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (School).
func (m *SchoolMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SchoolMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SchoolMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SchoolMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown School field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SchoolMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown School field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SchoolMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SchoolMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SchoolMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown School numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SchoolMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SchoolMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SchoolMutation) ClearField(name string) error {
	return fmt.Errorf("unknown School nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SchoolMutation) ResetField(name string) error {
	return fmt.Errorf("unknown School field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SchoolMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SchoolMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SchoolMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SchoolMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SchoolMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SchoolMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SchoolMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown School unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SchoolMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown School edge %s", name)
}

// StudentMutation represents an operation that mutates the Student nodes in the graph.
type StudentMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	created_at    *time.Time
	updated_at    *time.Time
	sort          *uint32
	addsort       *int32
	status        *uint8
	addstatus     *int8
	name          *[]string
	appendname    []string
	id_card       *[]string
	appendid_card []string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Student, error)
	predicates    []predicate.Student
}

var _ ent.Mutation = (*StudentMutation)(nil)

// studentOption allows management of the mutation configuration using functional options.
type studentOption func(*StudentMutation)

// newStudentMutation creates new mutation for the Student entity.
func newStudentMutation(c config, op Op, opts ...studentOption) *StudentMutation {
	m := &StudentMutation{
		config:        c,
		op:            op,
		typ:           TypeStudent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStudentID sets the ID field of the mutation.
func withStudentID(id uint64) studentOption {
	return func(m *StudentMutation) {
		var (
			err   error
			once  sync.Once
			value *Student
		)
		m.oldValue = func(ctx context.Context) (*Student, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Student.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStudent sets the old Student of the mutation.
func withStudent(node *Student) studentOption {
	return func(m *StudentMutation) {
		m.oldValue = func(context.Context) (*Student, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StudentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StudentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Student entities.
func (m *StudentMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StudentMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StudentMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Student.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *StudentMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *StudentMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *StudentMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *StudentMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *StudentMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *StudentMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetSort sets the "sort" field.
func (m *StudentMutation) SetSort(u uint32) {
	m.sort = &u
	m.addsort = nil
}

// Sort returns the value of the "sort" field in the mutation.
func (m *StudentMutation) Sort() (r uint32, exists bool) {
	v := m.sort
	if v == nil {
		return
	}
	return *v, true
}

// OldSort returns the old "sort" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldSort(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSort is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSort requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSort: %w", err)
	}
	return oldValue.Sort, nil
}

// AddSort adds u to the "sort" field.
func (m *StudentMutation) AddSort(u int32) {
	if m.addsort != nil {
		*m.addsort += u
	} else {
		m.addsort = &u
	}
}

// AddedSort returns the value that was added to the "sort" field in this mutation.
func (m *StudentMutation) AddedSort() (r int32, exists bool) {
	v := m.addsort
	if v == nil {
		return
	}
	return *v, true
}

// ResetSort resets all changes to the "sort" field.
func (m *StudentMutation) ResetSort() {
	m.sort = nil
	m.addsort = nil
}

// SetStatus sets the "status" field.
func (m *StudentMutation) SetStatus(u uint8) {
	m.status = &u
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *StudentMutation) Status() (r uint8, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldStatus(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds u to the "status" field.
func (m *StudentMutation) AddStatus(u int8) {
	if m.addstatus != nil {
		*m.addstatus += u
	} else {
		m.addstatus = &u
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *StudentMutation) AddedStatus() (r int8, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ClearStatus clears the value of the "status" field.
func (m *StudentMutation) ClearStatus() {
	m.status = nil
	m.addstatus = nil
	m.clearedFields[student.FieldStatus] = struct{}{}
}

// StatusCleared returns if the "status" field was cleared in this mutation.
func (m *StudentMutation) StatusCleared() bool {
	_, ok := m.clearedFields[student.FieldStatus]
	return ok
}

// ResetStatus resets all changes to the "status" field.
func (m *StudentMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
	delete(m.clearedFields, student.FieldStatus)
}

// SetName sets the "name" field.
func (m *StudentMutation) SetName(s []string) {
	m.name = &s
	m.appendname = nil
}

// Name returns the value of the "name" field in the mutation.
func (m *StudentMutation) Name() (r []string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldName(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// AppendName adds s to the "name" field.
func (m *StudentMutation) AppendName(s []string) {
	m.appendname = append(m.appendname, s...)
}

// AppendedName returns the list of values that were appended to the "name" field in this mutation.
func (m *StudentMutation) AppendedName() ([]string, bool) {
	if len(m.appendname) == 0 {
		return nil, false
	}
	return m.appendname, true
}

// ResetName resets all changes to the "name" field.
func (m *StudentMutation) ResetName() {
	m.name = nil
	m.appendname = nil
}

// SetIDCard sets the "id_card" field.
func (m *StudentMutation) SetIDCard(s []string) {
	m.id_card = &s
	m.appendid_card = nil
}

// IDCard returns the value of the "id_card" field in the mutation.
func (m *StudentMutation) IDCard() (r []string, exists bool) {
	v := m.id_card
	if v == nil {
		return
	}
	return *v, true
}

// OldIDCard returns the old "id_card" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldIDCard(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIDCard is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIDCard requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIDCard: %w", err)
	}
	return oldValue.IDCard, nil
}

// AppendIDCard adds s to the "id_card" field.
func (m *StudentMutation) AppendIDCard(s []string) {
	m.appendid_card = append(m.appendid_card, s...)
}

// AppendedIDCard returns the list of values that were appended to the "id_card" field in this mutation.
func (m *StudentMutation) AppendedIDCard() ([]string, bool) {
	if len(m.appendid_card) == 0 {
		return nil, false
	}
	return m.appendid_card, true
}

// ResetIDCard resets all changes to the "id_card" field.
func (m *StudentMutation) ResetIDCard() {
	m.id_card = nil
	m.appendid_card = nil
}

// Where appends a list predicates to the StudentMutation builder.
func (m *StudentMutation) Where(ps ...predicate.Student) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StudentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StudentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Student, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StudentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StudentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Student).
func (m *StudentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StudentMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, student.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, student.FieldUpdatedAt)
	}
	if m.sort != nil {
		fields = append(fields, student.FieldSort)
	}
	if m.status != nil {
		fields = append(fields, student.FieldStatus)
	}
	if m.name != nil {
		fields = append(fields, student.FieldName)
	}
	if m.id_card != nil {
		fields = append(fields, student.FieldIDCard)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StudentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case student.FieldCreatedAt:
		return m.CreatedAt()
	case student.FieldUpdatedAt:
		return m.UpdatedAt()
	case student.FieldSort:
		return m.Sort()
	case student.FieldStatus:
		return m.Status()
	case student.FieldName:
		return m.Name()
	case student.FieldIDCard:
		return m.IDCard()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StudentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case student.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case student.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case student.FieldSort:
		return m.OldSort(ctx)
	case student.FieldStatus:
		return m.OldStatus(ctx)
	case student.FieldName:
		return m.OldName(ctx)
	case student.FieldIDCard:
		return m.OldIDCard(ctx)
	}
	return nil, fmt.Errorf("unknown Student field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case student.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case student.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case student.FieldSort:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSort(v)
		return nil
	case student.FieldStatus:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case student.FieldName:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case student.FieldIDCard:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIDCard(v)
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StudentMutation) AddedFields() []string {
	var fields []string
	if m.addsort != nil {
		fields = append(fields, student.FieldSort)
	}
	if m.addstatus != nil {
		fields = append(fields, student.FieldStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StudentMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case student.FieldSort:
		return m.AddedSort()
	case student.FieldStatus:
		return m.AddedStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) AddField(name string, value ent.Value) error {
	switch name {
	case student.FieldSort:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSort(v)
		return nil
	case student.FieldStatus:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Student numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StudentMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(student.FieldStatus) {
		fields = append(fields, student.FieldStatus)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StudentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StudentMutation) ClearField(name string) error {
	switch name {
	case student.FieldStatus:
		m.ClearStatus()
		return nil
	}
	return fmt.Errorf("unknown Student nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StudentMutation) ResetField(name string) error {
	switch name {
	case student.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case student.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case student.FieldSort:
		m.ResetSort()
		return nil
	case student.FieldStatus:
		m.ResetStatus()
		return nil
	case student.FieldName:
		m.ResetName()
		return nil
	case student.FieldIDCard:
		m.ResetIDCard()
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StudentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StudentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StudentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StudentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StudentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StudentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StudentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Student unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StudentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Student edge %s", name)
}
