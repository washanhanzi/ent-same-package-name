// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/washanhanzi/ent-same-package-name/ent/predicate"
	"github.com/washanhanzi/ent-same-package-name/ent/tablea"
	"github.com/washanhanzi/ent-same-package-name/ent/tableb"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTableA = "TableA"
	TypeTableB = "TableB"
)

// TableAMutation represents an operation that mutates the TableA nodes in the graph.
type TableAMutation struct {
	config
	op            Op
	typ           string
	id            *int
	a             *entity.A
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*TableA, error)
	predicates    []predicate.TableA
}

var _ ent.Mutation = (*TableAMutation)(nil)

// tableaOption allows management of the mutation configuration using functional options.
type tableaOption func(*TableAMutation)

// newTableAMutation creates new mutation for the TableA entity.
func newTableAMutation(c config, op Op, opts ...tableaOption) *TableAMutation {
	m := &TableAMutation{
		config:        c,
		op:            op,
		typ:           TypeTableA,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTableAID sets the ID field of the mutation.
func withTableAID(id int) tableaOption {
	return func(m *TableAMutation) {
		var (
			err   error
			once  sync.Once
			value *TableA
		)
		m.oldValue = func(ctx context.Context) (*TableA, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().TableA.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTableA sets the old TableA of the mutation.
func withTableA(node *TableA) tableaOption {
	return func(m *TableAMutation) {
		m.oldValue = func(context.Context) (*TableA, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TableAMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TableAMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *TableAMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetA sets the "a" field.
func (m *TableAMutation) SetA(e entity.A) {
	m.a = &e
}

// A returns the value of the "a" field in the mutation.
func (m *TableAMutation) A() (r entity.A, exists bool) {
	v := m.a
	if v == nil {
		return
	}
	return *v, true
}

// OldA returns the old "a" field's value of the TableA entity.
// If the TableA object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TableAMutation) OldA(ctx context.Context) (v entity.A, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldA is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldA requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldA: %w", err)
	}
	return oldValue.A, nil
}

// ResetA resets all changes to the "a" field.
func (m *TableAMutation) ResetA() {
	m.a = nil
}

// Op returns the operation name.
func (m *TableAMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (TableA).
func (m *TableAMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TableAMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.a != nil {
		fields = append(fields, tablea.FieldA)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TableAMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tablea.FieldA:
		return m.A()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TableAMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tablea.FieldA:
		return m.OldA(ctx)
	}
	return nil, fmt.Errorf("unknown TableA field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TableAMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tablea.FieldA:
		v, ok := value.(entity.A)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetA(v)
		return nil
	}
	return fmt.Errorf("unknown TableA field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TableAMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TableAMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TableAMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown TableA numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TableAMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TableAMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TableAMutation) ClearField(name string) error {
	return fmt.Errorf("unknown TableA nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TableAMutation) ResetField(name string) error {
	switch name {
	case tablea.FieldA:
		m.ResetA()
		return nil
	}
	return fmt.Errorf("unknown TableA field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TableAMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TableAMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TableAMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TableAMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TableAMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TableAMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TableAMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown TableA unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TableAMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown TableA edge %s", name)
}

// TableBMutation represents an operation that mutates the TableB nodes in the graph.
type TableBMutation struct {
	config
	op            Op
	typ           string
	id            *int
	b             *entity.B
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*TableB, error)
	predicates    []predicate.TableB
}

var _ ent.Mutation = (*TableBMutation)(nil)

// tablebOption allows management of the mutation configuration using functional options.
type tablebOption func(*TableBMutation)

// newTableBMutation creates new mutation for the TableB entity.
func newTableBMutation(c config, op Op, opts ...tablebOption) *TableBMutation {
	m := &TableBMutation{
		config:        c,
		op:            op,
		typ:           TypeTableB,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTableBID sets the ID field of the mutation.
func withTableBID(id int) tablebOption {
	return func(m *TableBMutation) {
		var (
			err   error
			once  sync.Once
			value *TableB
		)
		m.oldValue = func(ctx context.Context) (*TableB, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().TableB.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTableB sets the old TableB of the mutation.
func withTableB(node *TableB) tablebOption {
	return func(m *TableBMutation) {
		m.oldValue = func(context.Context) (*TableB, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TableBMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TableBMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *TableBMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetB sets the "b" field.
func (m *TableBMutation) SetB(e entity.B) {
	m.b = &e
}

// B returns the value of the "b" field in the mutation.
func (m *TableBMutation) B() (r entity.B, exists bool) {
	v := m.b
	if v == nil {
		return
	}
	return *v, true
}

// OldB returns the old "b" field's value of the TableB entity.
// If the TableB object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TableBMutation) OldB(ctx context.Context) (v entity.B, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldB is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldB requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldB: %w", err)
	}
	return oldValue.B, nil
}

// ResetB resets all changes to the "b" field.
func (m *TableBMutation) ResetB() {
	m.b = nil
}

// Op returns the operation name.
func (m *TableBMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (TableB).
func (m *TableBMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TableBMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.b != nil {
		fields = append(fields, tableb.FieldB)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TableBMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tableb.FieldB:
		return m.B()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TableBMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tableb.FieldB:
		return m.OldB(ctx)
	}
	return nil, fmt.Errorf("unknown TableB field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TableBMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tableb.FieldB:
		v, ok := value.(entity.B)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetB(v)
		return nil
	}
	return fmt.Errorf("unknown TableB field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TableBMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TableBMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TableBMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown TableB numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TableBMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TableBMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TableBMutation) ClearField(name string) error {
	return fmt.Errorf("unknown TableB nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TableBMutation) ResetField(name string) error {
	switch name {
	case tableb.FieldB:
		m.ResetB()
		return nil
	}
	return fmt.Errorf("unknown TableB field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TableBMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TableBMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TableBMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TableBMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TableBMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TableBMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TableBMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown TableB unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TableBMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown TableB edge %s", name)
}
