// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewcourse"
	"github.com/google/uuid"
)

// ReviewCourse is the model entity for the ReviewCourse schema.
type ReviewCourse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Score holds the value of the "score" field.
	Score *float32 `json:"score,omitempty"`
	// ReviewMsg holds the value of the "review_msg" field.
	ReviewMsg *string `json:"review_msg,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewCourseQuery when eager-loading is set.
	Edges                ReviewCourseEdges `json:"edges"`
	course_review_course *uuid.UUID
}

// ReviewCourseEdges holds the relations/edges for other nodes in the graph.
type ReviewCourseEdges struct {
	// Course holds the value of the course edge.
	Course *Course `json:"course,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewCourseEdges) CourseOrErr() (*Course, error) {
	if e.loadedTypes[0] {
		if e.Course == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: course.Label}
		}
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReviewCourse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reviewcourse.FieldScore:
			values[i] = new(sql.NullFloat64)
		case reviewcourse.FieldID:
			values[i] = new(sql.NullInt64)
		case reviewcourse.FieldReviewMsg:
			values[i] = new(sql.NullString)
		case reviewcourse.ForeignKeys[0]: // course_review_course
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ReviewCourse", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReviewCourse fields.
func (rc *ReviewCourse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reviewcourse.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rc.ID = int(value.Int64)
		case reviewcourse.FieldScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				rc.Score = new(float32)
				*rc.Score = float32(value.Float64)
			}
		case reviewcourse.FieldReviewMsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field review_msg", values[i])
			} else if value.Valid {
				rc.ReviewMsg = new(string)
				*rc.ReviewMsg = value.String
			}
		case reviewcourse.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field course_review_course", values[i])
			} else if value.Valid {
				rc.course_review_course = new(uuid.UUID)
				*rc.course_review_course = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCourse queries the "course" edge of the ReviewCourse entity.
func (rc *ReviewCourse) QueryCourse() *CourseQuery {
	return NewReviewCourseClient(rc.config).QueryCourse(rc)
}

// Update returns a builder for updating this ReviewCourse.
// Note that you need to call ReviewCourse.Unwrap() before calling this method if this ReviewCourse
// was returned from a transaction, and the transaction was committed or rolled back.
func (rc *ReviewCourse) Update() *ReviewCourseUpdateOne {
	return NewReviewCourseClient(rc.config).UpdateOne(rc)
}

// Unwrap unwraps the ReviewCourse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rc *ReviewCourse) Unwrap() *ReviewCourse {
	_tx, ok := rc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReviewCourse is not a transactional entity")
	}
	rc.config.driver = _tx.drv
	return rc
}

// String implements the fmt.Stringer.
func (rc *ReviewCourse) String() string {
	var builder strings.Builder
	builder.WriteString("ReviewCourse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rc.ID))
	if v := rc.Score; v != nil {
		builder.WriteString("score=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := rc.ReviewMsg; v != nil {
		builder.WriteString("review_msg=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// ReviewCourses is a parsable slice of ReviewCourse.
type ReviewCourses []*ReviewCourse
