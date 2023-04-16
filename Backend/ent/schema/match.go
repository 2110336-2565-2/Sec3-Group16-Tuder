package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"time"
)

// Todo holds the schema definition for the Todo entity.
type Match struct {
	ent.Schema
}

// Fields of the Todo.
func (Match) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.Time("match_created_at").Default(time.Now).Immutable(),
		field.Enum("status").Values("enrolled", "completed", "cancelling", "canceled").Default("enrolled"),
		// enrolled -> student enrolled course
		// completed -> student completed course
		// cancelling -> student request to cancel course
		// canceled -> student cancel course
		

	}
}

func (Match) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("issue_report", IssueReport.Type),
		edge.From("student", Student.Type).
			Ref("match").
			Unique().
			Required(),
		edge.From("course", Course.Type).
			Ref("match").
			Unique().
			Required(),
		edge.From("appointment", Appointment.Type).
			Ref("match").
			Required(),
		edge.From("schedule", Schedule.Type).
			Ref("match").
			Unique().
			Required(),
		edge.From("cancel_request", CancelRequest.Type).
			Ref("match"),
		edge.From("payment", Payment.Type).
			Ref("match").
			Unique(),
	}
}
