package schema

import (
	"entgo.io/ent"
)

// Todo holds the schema definition for the Todo entity.
type ReviewCourse struct {
	ent.Schema
}

// Mixin of Tutor

func (ReviewCourse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ReportMixin{},
	}
}

func (ReviewCourse) Edges() []ent.Edge {
	return nil
}
