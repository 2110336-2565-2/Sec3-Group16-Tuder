package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Todo.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.JSON("day_0", [24]bool{}),
		field.JSON("day_1", [24]bool{}),
		field.JSON("day_2", [24]bool{}),
		field.JSON("day_3", [24]bool{}),
		field.JSON("day_4", [24]bool{}),
		field.JSON("day_5", [24]bool{}),
		field.JSON("day_6", [24]bool{}),
	}
}

func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tutor", Tutor.Type).
			Unique(),
		edge.To("match", Match.Type).
			Unique(),
	}
}
