// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/google/uuid"
)

// Course is the model entity for the Course schema.
type Course struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Topic holds the value of the "topic" field.
	Topic string `json:"topic,omitempty"`
	// EstimatedTime holds the value of the "estimated_time" field.
	EstimatedTime int `json:"estimated_time,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// PricePerHour holds the value of the "price_per_hour" field.
	PricePerHour int `json:"price_per_hour,omitempty"`
	// Level holds the value of the "level" field.
	Level course.Level `json:"level,omitempty"`
	// CoursePictureURL holds the value of the "course_picture_url" field.
	CoursePictureURL *string `json:"course_picture_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseQuery when eager-loading is set.
	Edges        CourseEdges `json:"edges"`
	tutor_course *uuid.UUID
}

// CourseEdges holds the relations/edges for other nodes in the graph.
type CourseEdges struct {
	// ReviewCourse holds the value of the review_course edge.
	ReviewCourse []*ReviewCourse `json:"review_course,omitempty"`
	// Match holds the value of the match edge.
	Match []*Match `json:"match,omitempty"`
	// Tutor holds the value of the tutor edge.
	Tutor *Tutor `json:"tutor,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ReviewCourseOrErr returns the ReviewCourse value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) ReviewCourseOrErr() ([]*ReviewCourse, error) {
	if e.loadedTypes[0] {
		return e.ReviewCourse, nil
	}
	return nil, &NotLoadedError{edge: "review_course"}
}

// MatchOrErr returns the Match value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) MatchOrErr() ([]*Match, error) {
	if e.loadedTypes[1] {
		return e.Match, nil
	}
	return nil, &NotLoadedError{edge: "match"}
}

// TutorOrErr returns the Tutor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CourseEdges) TutorOrErr() (*Tutor, error) {
	if e.loadedTypes[2] {
		if e.Tutor == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tutor.Label}
		}
		return e.Tutor, nil
	}
	return nil, &NotLoadedError{edge: "tutor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Course) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case course.FieldEstimatedTime, course.FieldPricePerHour:
			values[i] = new(sql.NullInt64)
		case course.FieldTitle, course.FieldSubject, course.FieldTopic, course.FieldDescription, course.FieldLevel, course.FieldCoursePictureURL:
			values[i] = new(sql.NullString)
		case course.FieldID:
			values[i] = new(uuid.UUID)
		case course.ForeignKeys[0]: // tutor_course
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Course", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Course fields.
func (c *Course) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case course.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case course.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case course.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				c.Subject = value.String
			}
		case course.FieldTopic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field topic", values[i])
			} else if value.Valid {
				c.Topic = value.String
			}
		case course.FieldEstimatedTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field estimated_time", values[i])
			} else if value.Valid {
				c.EstimatedTime = int(value.Int64)
			}
		case course.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case course.FieldPricePerHour:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price_per_hour", values[i])
			} else if value.Valid {
				c.PricePerHour = int(value.Int64)
			}
		case course.FieldLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				c.Level = course.Level(value.String)
			}
		case course.FieldCoursePictureURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field course_picture_url", values[i])
			} else if value.Valid {
				c.CoursePictureURL = new(string)
				*c.CoursePictureURL = value.String
			}
		case course.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field tutor_course", values[i])
			} else if value.Valid {
				c.tutor_course = new(uuid.UUID)
				*c.tutor_course = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryReviewCourse queries the "review_course" edge of the Course entity.
func (c *Course) QueryReviewCourse() *ReviewCourseQuery {
	return NewCourseClient(c.config).QueryReviewCourse(c)
}

// QueryMatch queries the "match" edge of the Course entity.
func (c *Course) QueryMatch() *MatchQuery {
	return NewCourseClient(c.config).QueryMatch(c)
}

// QueryTutor queries the "tutor" edge of the Course entity.
func (c *Course) QueryTutor() *TutorQuery {
	return NewCourseClient(c.config).QueryTutor(c)
}

// Update returns a builder for updating this Course.
// Note that you need to call Course.Unwrap() before calling this method if this Course
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Course) Update() *CourseUpdateOne {
	return NewCourseClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Course entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Course) Unwrap() *Course {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Course is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Course) String() string {
	var builder strings.Builder
	builder.WriteString("Course(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(c.Subject)
	builder.WriteString(", ")
	builder.WriteString("topic=")
	builder.WriteString(c.Topic)
	builder.WriteString(", ")
	builder.WriteString("estimated_time=")
	builder.WriteString(fmt.Sprintf("%v", c.EstimatedTime))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("price_per_hour=")
	builder.WriteString(fmt.Sprintf("%v", c.PricePerHour))
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", c.Level))
	builder.WriteString(", ")
	if v := c.CoursePictureURL; v != nil {
		builder.WriteString("course_picture_url=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Courses is a parsable slice of Course.
type Courses []*Course
