package schema

import (
	"entgo.io/ent"
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
		// field.String("user_id").NotEmpty().Unique(),
		field.Bool("review_avaliable").Default(true),
		field.String("total_hour").NotEmpty(),
		field.String("success_hour").NotEmpty(),
	}
}

func (Class) Edges() []ent.Edge {
	return nil
}
