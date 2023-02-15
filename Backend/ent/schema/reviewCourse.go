package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Todo holds the schema definition for the Todo entity.
type ReviewCourse struct {
	ent.Schema
}

// Mixin of Tutor

func (ReviewCourse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ReviewMixin{},
	}
}

func (ReviewCourse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("review_course").
			Unique().
			Required(),
	}
}
