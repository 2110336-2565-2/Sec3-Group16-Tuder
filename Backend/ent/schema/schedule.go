package schema

import (
	"entgo.io/ent"
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
		field.Bool("day_0").Optional().Nillable().Default(false),
		field.Bool("day_1").Optional().Nillable().Default(false),
		field.Bool("day_2").Optional().Nillable().Default(false),
		field.Bool("day_3").Optional().Nillable().Default(false),
		field.Bool("day_4").Optional().Nillable().Default(false),
		field.Bool("day_5").Optional().Nillable().Default(false),
		field.Bool("day_6").Optional().Nillable().Default(false),
	}
}

func (Schedule) Edges() []ent.Edge {
	return nil
}
