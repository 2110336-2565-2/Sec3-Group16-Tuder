package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Todo holds the schema definition for the Todo entity.
type ReportMixin struct {
	mixin.Schema
}

// Fields of the Todo.
func (ReportMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Float32("score").Positive().Nillable().Optional(),
		field.String("review_msg").Optional().Nillable(),
	}
}
