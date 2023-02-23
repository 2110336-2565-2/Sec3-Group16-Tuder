package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Course struct {
	ent.Schema
}

// Fields of the Todo.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique().StorageKey("id").Immutable(),
		field.String("title").NotEmpty(),
		field.String("subject").NotEmpty(), //new attribute
		field.String("topic").NotEmpty(),   //new attribute
		field.Int("estimated_time"),
		field.String("description").NotEmpty(),
		field.Int("price_per_hour").Positive(),
		field.Enum("level").
			Values("Grade1", "Grade2", "Grade3", "Grade4", "Grade5", "Grade6", "Grade7", "Grade8", "Grade9", "Grade10", "Grade11", "Grade12").
			Optional(),
		field.String("course_picture_url").Optional().Nillable(),
	}
}

func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("review_course", ReviewCourse.Type),
		edge.To("class", Class.Type),
		edge.From("tutor", Tutor.Type).
			Ref("course").
			Unique().
			Required(),
	}
}
