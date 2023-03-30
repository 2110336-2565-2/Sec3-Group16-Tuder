package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"time"
)

// Todo holds the schema definition for the Todo entity.
type Review struct {
	ent.Schema
}

// Mixin of Tutor

func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.Float32("score").Range(0.0, 5.0).Nillable().Optional(),
		field.String("review_msg").Optional().Nillable(),
		field.Time("review_time_at").Default(time.Now).Immutable(),
	}
}

func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("review").
			// Unique().
			Required(),
		edge.From("student", Student.Type).
			Ref("review").
			// Unique().
			Required(),
	}
}
