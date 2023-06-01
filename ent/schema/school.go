package schema

import "entgo.io/ent"

// School holds the schema definition for the School entity.
type School struct {
	ent.Schema
}

// Fields of the School.
func (School) Fields() []ent.Field {
	return nil
}

// Edges of the School.
func (School) Edges() []ent.Edge {
	return nil
}
