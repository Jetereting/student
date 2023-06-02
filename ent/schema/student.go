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

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("姓名").Annotations(entsql.WithComments(true)),
		field.String("id_card").Comment("身份证号").Annotations(entsql.WithComments(true)),
		field.Uint64("class_id").Optional().Comment("班级代码").Annotations(entsql.WithComments(true)),
	}
}

// Indexes of the Student.| 索引
func (Student) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("id_card").Unique(),
		index.Fields("class_id"),
	}
}

// Mixin of the Student.
func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
		mixins.SortMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("students").Unique().Comment("班级").Field("class_id").Annotations(entsql.WithComments(true)),
	}
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "b_student"},
	}
}
