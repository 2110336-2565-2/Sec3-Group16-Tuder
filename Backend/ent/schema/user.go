package schema

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type User struct {
	ent.Schema
}

// Fields of the Todo.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
            Default(uuid.New).Unique().StorageKey("id"),
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.Enum("role").Values("student", "tutor", "admin"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}

