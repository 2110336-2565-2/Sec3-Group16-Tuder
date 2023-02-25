package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type IssueReport struct {
	ent.Schema
}

// Fields of the Todo.
func (IssueReport) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		// field.String("user_id").NotEmpty().Unique(),
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
		field.Time("report_date"),
		field.String("status").NotEmpty(),
	}
}

func (IssueReport) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("student", Student.Type).
		// 	Ref("issue_report").
		// 	Unique().
		// 	Required(),
		// edge.From("tutor", Tutor.Type).
		// 	Ref("issue_report").
		// 	Unique().
		// 	Required(),

		edge.From("user", User.Type).
			Ref("issue_report").
			Unique().
			Required(),
	}
}
