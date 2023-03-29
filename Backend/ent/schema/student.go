package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Student struct {
	ent.Schema
}

// func (Student) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		UserMixin{},
// 	}
// }

// Fields of the Todo.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		// field.String("school").NotEmpty(),
	}
}

func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("issue_report", IssueReport.Type),
		edge.To("match", Match.Type),
		// edge.To("class", Class.Type),
		edge.To("review", Review.Type),
		edge.From("user", User.Type).
			Ref("student").
			Unique().
			Required(),
	}
}
