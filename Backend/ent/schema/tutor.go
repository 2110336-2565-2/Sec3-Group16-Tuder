package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Tutor struct {
	ent.Schema
}

// Mixin of Tutor

func (Tutor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserMixin{},
	}
}

// Fields of the Todo.
func (Tutor) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Nillable(),
		field.String("omise_bank_token").Optional().Nillable(),
		field.String("citizen_id").NotEmpty().Unique(),
	}
}

func (Tutor) Edges() []ent.Edge {
	return nil
}
