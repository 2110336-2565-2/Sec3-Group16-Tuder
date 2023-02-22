package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Todo holds the schema definition for the Todo entity.
type ReviewTutor struct {
	ent.Schema
}

// Mixin of Tutor

func (ReviewTutor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ReviewMixin{},
	}
}

func (ReviewTutor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tutor", Tutor.Type).
			Ref("review_tutor").
			Unique().
			Required(),
	}
}
