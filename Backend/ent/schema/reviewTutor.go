package schema

import (
	"entgo.io/ent"
)

// Todo holds the schema definition for the Todo entity.
type ReportTutor struct {
	ent.Schema
}

// Mixin of Tutor

func (ReportTutor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ReportMixin{},
	}
}

func (ReportTutor) Edges() []ent.Edge {
	return nil
}
