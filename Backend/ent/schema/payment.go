package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Payment struct {
	ent.Schema
}

// Fields of the Todo.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		// field.String("user_id").NotEmpty().Unique(),
		field.String("qr_picture_url").Optional().Nillable(),
		field.String("payment_status").NotEmpty(),
		field.String("card").NotEmpty(),
		field.Int("amount").Positive(),
	}
}

func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("payment").
			Unique().
			Required(),
		edge.To("payment_history", PaymentHistory.Type),
	}
}
