package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"time"
)

type CancelRequest struct {
	ent.Schema
}

func (CancelRequest) Fields() []ent.Field {
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

func (CancelRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type).
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("cancel_request").
			Unique().
			Required(),
	}
}
