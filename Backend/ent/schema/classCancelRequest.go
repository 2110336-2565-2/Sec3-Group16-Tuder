package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	time "time"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type ClassCancelRequest struct {
	ent.Schema
}

// Fields of the Todo.
func (ClassCancelRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.String("title").NotEmpty(),
		field.Time("report_date").Default(func() time.Time {
			return time.Now()
		}),
		field.String("img_url"),
		field.String("description").NotEmpty(),
		field.Enum("status").Values("pending", "approved", "rejected"),
	}
}

func (ClassCancelRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("match", Match.Type).
			Ref("class_cancel_request").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("class_cancel_request").
			Unique().
			Required(),
	}
}
