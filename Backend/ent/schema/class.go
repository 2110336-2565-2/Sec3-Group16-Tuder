package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Class struct {
	ent.Schema
}

// Fields of the Todo.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.Bool("review_avaliable").Default(true),
		field.Int("total_hour"),
		field.Int("success_hour"),
		field.Enum("status").Values("ongoing", "completed", "cancelled"),
	}
}

func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type),
		edge.From("schedule", Schedule.Type).
			Ref("class").
			Unique().
			Required(),
		edge.From("payment_history", PaymentHistory.Type).
			Ref("class").
			Unique().
			Required(),
	}
}
