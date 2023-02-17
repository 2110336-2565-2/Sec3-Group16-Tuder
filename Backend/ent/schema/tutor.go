package schema

import (
	"github.com/google/uuid"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Tutor struct {
	ent.Schema
}

// Mixin of Tutor

// func (Tutor) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		UserMixin{},
// 	}
// }

// Fields of the Todo.
func (Tutor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.String("description").Optional().Nillable(),
		field.String("omise_bank_token").Optional().Nillable(),
		field.String("citizen_id").NotEmpty().Unique(),
	}
}

func (Tutor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("issue_report", IssueReport.Type),
		edge.To("course", Course.Type),
		edge.To("review_tutor", ReviewTutor.Type),
		edge.To("schedule", Schedule.Type),
		edge.From("user", User.Type).
			Ref("tutor").
			Unique().
			Required(),
	}
}
