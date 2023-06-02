package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("班级名").Annotations(entsql.WithComments(true)),
	}
}

// Indexes of the Class.| 索引
func (Class) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

// Mixin of the Class.
func (Class) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
		mixins.SortMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("students", Student.Type),
	}
}

func (Class) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "b_class"},
	}
}
