package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Todo holds the schema definition for the Todo entity.
type Match struct {
	ent.Schema
}

// func (Match) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		UserMixin{},
// 	}
// }

// // Fields of the Todo.
// func (Match) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.String("school").NotEmpty(),
// 	}
// }

func (Match) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("issue_report", IssueReport.Type),
		edge.From("student", Student.Type).
			Ref("match").
			Required(),
		edge.From("course", Course.Type).
			Ref("match").
			Required(),
		edge.From("class", Class.Type).
			Ref("match").
			Required(),
	}
}
