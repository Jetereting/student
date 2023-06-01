// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"school/ent/school"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// School is the model entity for the School schema.
type School struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*School) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case school.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the School fields.
func (s *School) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case school.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the School.
// This includes values selected through modifiers, order, etc.
func (s *School) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this School.
// Note that you need to call School.Unwrap() before calling this method if this School
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *School) Update() *SchoolUpdateOne {
	return NewSchoolClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the School entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *School) Unwrap() *School {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: School is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *School) String() string {
	var builder strings.Builder
	builder.WriteString("School(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Schools is a parsable slice of School.
type Schools []*School