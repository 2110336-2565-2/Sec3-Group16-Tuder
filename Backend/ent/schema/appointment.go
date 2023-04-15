package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Appointment struct {
	ent.Schema
}

// Fields of the Todo.
func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.Time("begin_at").Immutable(),
		field.Time("end_at").Immutable(),
		field.Enum("status").Values("ongoing", "completed", "cancelling", "rejected", "cancelled"),
	}
}

func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type).
			Unique(),
		edge.From("payment", Payment.Type).
			Ref("appointment").
			Unique(),
	}
}
