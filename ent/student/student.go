// Code generated by ent, DO NOT EDIT.

package student

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the student type in the database.
	Label = "student"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIDCard holds the string denoting the id_card field in the database.
	FieldIDCard = "id_card"
	// FieldClassID holds the string denoting the class_id field in the database.
	FieldClassID = "class_id"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// Table holds the table name of the student in the database.
	Table = "b_student"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "b_student"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "b_class"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_id"
)

// Columns holds all SQL columns for student fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSort,
	FieldStatus,
	FieldName,
	FieldIDCard,
	FieldClassID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort uint32
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)

// OrderOption defines the ordering options for the Student queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIDCard orders the results by the id_card field.
func ByIDCard(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIDCard, opts...).ToFunc()
}

// ByClassID orders the results by the class_id field.
func ByClassID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClassID, opts...).ToFunc()
}

// ByClassField orders the results by class field.
func ByClassField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassStep(), sql.OrderByField(field, opts...))
	}
}
func newClassStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
	)
}
