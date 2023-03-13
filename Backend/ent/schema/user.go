package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

// Todo holds the schema definition for the Todo entity.
// type UserMixin struct {
// 	mixin.Schema
// }

// Fields of the Todo.
func (User) Fields() []ent.Field {
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
		field.Enum("role").Values("student", "tutor").Nillable(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("users_username_key").Unique().Fields("username", "role"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
		edge.To("tutor", Tutor.Type).Unique(),
		edge.To("issue_report", IssueReport.Type),
		edge.To("payment", Payment.Type),
		edge.To("payment_history", PaymentHistory.Type),
	}
}
