package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Todo holds the schema definition for the Todo entity.
type Student struct {
	ent.Schema
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserMixin{},
	}
}

// Fields of the Todo.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		// field.String("school").NotEmpty(),
	}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("issue_report", IssueReport.Type),
		edge.To("course", Course.Type).
			Unique(),
		edge.To("class", Class.Type).
			Unique(),
	}
}
