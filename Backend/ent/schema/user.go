package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type UserMixin struct {
	mixin.Schema
}

// Fields of the Todo.
func (UserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty().Sensitive(),
		field.String("email").NotEmpty(),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("address").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.Time("birth_date"),
		field.String("gender").NotEmpty(),
		field.String("profile_picture_URL").Optional().Nillable(),
	}
}

// func (UserMixin) MixinEdges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.To("issue_report", IssueReport.Type),
// 	}
// }

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("issue_report", IssueReport.Type),
		edge.To("payment", Payment.Type),
		edge.To("payment_history", PaymentHistory.Type),
	}
}
