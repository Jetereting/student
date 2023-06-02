// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BClassColumns holds the columns for the "b_class" table.
	BClassColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "班级名"},
	}
	// BClassTable holds the schema information for the "b_class" table.
	BClassTable = &schema.Table{
		Name:       "b_class",
		Columns:    BClassColumns,
		PrimaryKey: []*schema.Column{BClassColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "class_name",
				Unique:  false,
				Columns: []*schema.Column{BClassColumns[5]},
			},
		},
	}
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
	// BStudentColumns holds the columns for the "b_student" table.
	BStudentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "姓名", Default: ""},
		{Name: "id_card", Type: field.TypeString, Comment: "身份证号"},
		{Name: "class_id", Type: field.TypeUint64, Nullable: true, Comment: "班级代码"},
	}
	// BStudentTable holds the schema information for the "b_student" table.
	BStudentTable = &schema.Table{
		Name:       "b_student",
		Columns:    BStudentColumns,
		PrimaryKey: []*schema.Column{BStudentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "b_student_b_class_students",
				Columns:    []*schema.Column{BStudentColumns[7]},
				RefColumns: []*schema.Column{BClassColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "student_name",
				Unique:  false,
				Columns: []*schema.Column{BStudentColumns[5]},
			},
			{
				Name:    "student_id_card",
				Unique:  true,
				Columns: []*schema.Column{BStudentColumns[6]},
			},
			{
				Name:    "student_class_id",
				Unique:  false,
				Columns: []*schema.Column{BStudentColumns[7]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BClassTable,
		SchoolsTable,
		BStudentTable,
	}
)

func init() {
	BClassTable.Annotation = &entsql.Annotation{
		Table: "b_class",
	}
	BStudentTable.ForeignKeys[0].RefTable = BClassTable
	BStudentTable.Annotation = &entsql.Annotation{
		Table: "b_student",
	}
}
