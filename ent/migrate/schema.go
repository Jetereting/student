// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SchoolsColumns holds the columns for the "schools" table.
	SchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SchoolsTable holds the schema information for the "schools" table.
	SchoolsTable = &schema.Table{
		Name:       "schools",
		Columns:    SchoolsColumns,
		PrimaryKey: []*schema.Column{SchoolsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SchoolsTable,
	}
)

func init() {
}