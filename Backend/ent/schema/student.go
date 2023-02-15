package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Student struct {
	ent.Schema
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserMixin{},
	}
}

// Fields of the Todo.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("school").NotEmpty(),
	}
}

func (Student) Edges() []ent.Edge {
	return nil
}
