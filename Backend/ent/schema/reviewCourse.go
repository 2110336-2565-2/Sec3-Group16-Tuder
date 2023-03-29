package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"time"
)

// Todo holds the schema definition for the Todo entity.
type ReviewCourse struct {
	ent.Schema
}

// Mixin of Tutor

func (ReviewCourse) Fields() []ent.Field {
	return []ent.Field{
		field.Float32("score").Positive().Nillable().Optional(),
		field.String("review_msg").Optional().Nillable(),
		field.Time("review_time_at").Default(time.Now).Immutable(),
	}
}

func (ReviewCourse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("review_course").
			// Unique().
			Required(),
		edge.From("student", Student.Type).
			Ref("review_course").
			// Unique().
			Required(),
	}
}
