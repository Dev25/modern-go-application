// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/todoitem"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTodoItem = "TodoItem"
)

// TodoItemMutation represents an operation that mutate the TodoItems
// nodes in the graph.
type TodoItemMutation struct {
	config
	op            Op
	typ           string
	id            *int
	uid           *string
	title         *string
	completed     *bool
	order         *int
	addorder      *int
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
}

var _ ent.Mutation = (*TodoItemMutation)(nil)

// newTodoItemMutation creates new mutation for $n.Name.
func newTodoItemMutation(c config, op Op) *TodoItemMutation {
	return &TodoItemMutation{
		config:        c,
		op:            op,
		typ:           TypeTodoItem,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TodoItemMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TodoItemMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TodoItemMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUID sets the uid field.
func (m *TodoItemMutation) SetUID(s string) {
	m.uid = &s
}

// UID returns the uid value in the mutation.
func (m *TodoItemMutation) UID() (r string, exists bool) {
	v := m.uid
	if v == nil {
		return
	}
	return *v, true
}

// ResetUID reset all changes of the uid field.
func (m *TodoItemMutation) ResetUID() {
	m.uid = nil
}

// SetTitle sets the title field.
func (m *TodoItemMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the title value in the mutation.
func (m *TodoItemMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// ResetTitle reset all changes of the title field.
func (m *TodoItemMutation) ResetTitle() {
	m.title = nil
}

// SetCompleted sets the completed field.
func (m *TodoItemMutation) SetCompleted(b bool) {
	m.completed = &b
}

// Completed returns the completed value in the mutation.
func (m *TodoItemMutation) Completed() (r bool, exists bool) {
	v := m.completed
	if v == nil {
		return
	}
	return *v, true
}

// ResetCompleted reset all changes of the completed field.
func (m *TodoItemMutation) ResetCompleted() {
	m.completed = nil
}

// SetOrder sets the order field.
func (m *TodoItemMutation) SetOrder(i int) {
	m.order = &i
	m.addorder = nil
}

// Order returns the order value in the mutation.
func (m *TodoItemMutation) Order() (r int, exists bool) {
	v := m.order
	if v == nil {
		return
	}
	return *v, true
}

// AddOrder adds i to order.
func (m *TodoItemMutation) AddOrder(i int) {
	if m.addorder != nil {
		*m.addorder += i
	} else {
		m.addorder = &i
	}
}

// AddedOrder returns the value that was added to the order field in this mutation.
func (m *TodoItemMutation) AddedOrder() (r int, exists bool) {
	v := m.addorder
	if v == nil {
		return
	}
	return *v, true
}

// ResetOrder reset all changes of the order field.
func (m *TodoItemMutation) ResetOrder() {
	m.order = nil
	m.addorder = nil
}

// SetCreatedAt sets the created_at field.
func (m *TodoItemMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *TodoItemMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt reset all changes of the created_at field.
func (m *TodoItemMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *TodoItemMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *TodoItemMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdatedAt reset all changes of the updated_at field.
func (m *TodoItemMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Op returns the operation name.
func (m *TodoItemMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (TodoItem).
func (m *TodoItemMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *TodoItemMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.uid != nil {
		fields = append(fields, todoitem.FieldUID)
	}
	if m.title != nil {
		fields = append(fields, todoitem.FieldTitle)
	}
	if m.completed != nil {
		fields = append(fields, todoitem.FieldCompleted)
	}
	if m.order != nil {
		fields = append(fields, todoitem.FieldOrder)
	}
	if m.created_at != nil {
		fields = append(fields, todoitem.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, todoitem.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *TodoItemMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todoitem.FieldUID:
		return m.UID()
	case todoitem.FieldTitle:
		return m.Title()
	case todoitem.FieldCompleted:
		return m.Completed()
	case todoitem.FieldOrder:
		return m.Order()
	case todoitem.FieldCreatedAt:
		return m.CreatedAt()
	case todoitem.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TodoItemMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todoitem.FieldUID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUID(v)
		return nil
	case todoitem.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case todoitem.FieldCompleted:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCompleted(v)
		return nil
	case todoitem.FieldOrder:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOrder(v)
		return nil
	case todoitem.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case todoitem.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown TodoItem field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *TodoItemMutation) AddedFields() []string {
	var fields []string
	if m.addorder != nil {
		fields = append(fields, todoitem.FieldOrder)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *TodoItemMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case todoitem.FieldOrder:
		return m.AddedOrder()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TodoItemMutation) AddField(name string, value ent.Value) error {
	switch name {
	case todoitem.FieldOrder:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddOrder(v)
		return nil
	}
	return fmt.Errorf("unknown TodoItem numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *TodoItemMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *TodoItemMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *TodoItemMutation) ClearField(name string) error {
	return fmt.Errorf("unknown TodoItem nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *TodoItemMutation) ResetField(name string) error {
	switch name {
	case todoitem.FieldUID:
		m.ResetUID()
		return nil
	case todoitem.FieldTitle:
		m.ResetTitle()
		return nil
	case todoitem.FieldCompleted:
		m.ResetCompleted()
		return nil
	case todoitem.FieldOrder:
		m.ResetOrder()
		return nil
	case todoitem.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case todoitem.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown TodoItem field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *TodoItemMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *TodoItemMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *TodoItemMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *TodoItemMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *TodoItemMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *TodoItemMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *TodoItemMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown TodoItem unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *TodoItemMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown TodoItem edge %s", name)
}
