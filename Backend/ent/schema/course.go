package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Course struct {
	ent.Schema
}

// Fields of the Todo.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.String("title").NotEmpty(),
		field.String("estimated_time").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("course_status").NotEmpty(),
		field.Int("price_per_hour").Positive(),
		field.String("level_id").NotEmpty(),
		field.String("course_picture_url").Optional().Nillable(),
	}
}

func (Course) Edges() []ent.Edge {
	return nil
}
