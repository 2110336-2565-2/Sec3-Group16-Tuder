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
		// field.String("user_id").NotEmpty().Unique(),
		field.Bool("review_avaliable").Default(true),
		field.Time("total_hour"),
		field.Time("success_hour"),
	}
}

func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("schedule", Schedule.Type).
			Unique().
			Required(),
		edge.From("student", Student.Type).
			Ref("class").
			Unique().
			Required(),
		edge.From("course", Course.Type).
			Ref("class").
			Unique().
			Required(),
	}
}
