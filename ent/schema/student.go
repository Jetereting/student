package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("name").Comment("姓名"),
		field.Strings("id_card").Comment("身份证号"),
	}
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
		mixins.SortMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "b_student"},
	}
}
