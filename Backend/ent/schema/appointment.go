package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Appointment struct {
	ent.Schema
}

// Fields of the Todo.
func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.Time("begin_at").Immutable(),
		field.Time("end_at").Immutable(),
		field.Enum("status").Values("comingsoon", "ongoing", "verifying", "pending", "completed", "postponed", "considering", "canceled").Default("comingsoon"),
		// comingsoon -> future appointment
		// ongoing -> currunt appointment
		// verifying -> appointment which waiting for student's approval
		// pending -> appointment which waitiing for payment to tutor
		// completed -> completed appointmnet
		// posponed -> appointment which student disapprove
		// considering -> appointment which have conflivt between tutor and student, waiting for admin
		// canceled -> appointment which canceled by student or tutor
	}
}

func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type).
			Unique(),
	}
}
