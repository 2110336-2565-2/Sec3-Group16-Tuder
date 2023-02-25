package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type PaymentHistory struct {
	ent.Schema
}

// Fields of the Todo.
func (PaymentHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		// field.String("user_id").NotEmpty().Unique(),
		// field.String("payment_id").NotEmpty().Unique(),
		field.Float("amount").Positive().Nillable().Optional(),
		field.String("type").NotEmpty(),
	}
}

func (PaymentHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type),
		edge.From("user", User.Type).
			Ref("payment_history").
			Unique().
			Required(),
		edge.From("payment", Payment.Type).
			Ref("payment_history").
			Unique().
			Required(),
	}
}
