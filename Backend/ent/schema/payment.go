package schema

import (
	"time"

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
		field.String("qr_picture_url").Optional().Nillable(),
		field.Enum("payment_status").Values("pending", "successful", "failed", "processing"),
		field.Int("amount").Positive(),
		field.String("currency").NotEmpty(),
		field.String("charge_id").Unique().NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			StorageKey("created_at").
			StructTag(`json:"created_at,omitempty"`),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			StorageKey("updated_at").
			StructTag(`json:"updated_at,omitempty"`),
	}
}

func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("payment").
			Unique().
			Required(),
		edge.To("match", Match.Type).Unique(),
		edge.To("appointment", Appointment.Type).Unique(),
		edge.To("payment_history", PaymentHistory.Type),
	}
}
