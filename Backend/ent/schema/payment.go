package schema

import (
	"entgo.io/ent"
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
	}
}

func (Payment) Edges() []ent.Edge {
	return nil
}
