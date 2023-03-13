// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewtutor"
)

// ReviewTutor is the model entity for the ReviewTutor schema.
type ReviewTutor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Score holds the value of the "score" field.
	Score *float32 `json:"score,omitempty"`
	// ReviewMsg holds the value of the "review_msg" field.
	ReviewMsg *string `json:"review_msg,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewTutorQuery when eager-loading is set.
	Edges ReviewTutorEdges `json:"edges"`
}

// ReviewTutorEdges holds the relations/edges for other nodes in the graph.
type ReviewTutorEdges struct {
	// Tutor holds the value of the tutor edge.
	Tutor []*Tutor `json:"tutor,omitempty"`
	// Student holds the value of the student edge.
	Student []*Student `json:"student,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TutorOrErr returns the Tutor value or an error if the edge
// was not loaded in eager-loading.
func (e ReviewTutorEdges) TutorOrErr() ([]*Tutor, error) {
	if e.loadedTypes[0] {
		return e.Tutor, nil
	}
	return nil, &NotLoadedError{edge: "tutor"}
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading.
func (e ReviewTutorEdges) StudentOrErr() ([]*Student, error) {
	if e.loadedTypes[1] {
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReviewTutor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reviewtutor.FieldScore:
			values[i] = new(sql.NullFloat64)
		case reviewtutor.FieldID:
			values[i] = new(sql.NullInt64)
		case reviewtutor.FieldReviewMsg:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ReviewTutor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReviewTutor fields.
func (rt *ReviewTutor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reviewtutor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rt.ID = int(value.Int64)
		case reviewtutor.FieldScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				rt.Score = new(float32)
				*rt.Score = float32(value.Float64)
			}
		case reviewtutor.FieldReviewMsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field review_msg", values[i])
			} else if value.Valid {
				rt.ReviewMsg = new(string)
				*rt.ReviewMsg = value.String
			}
		}
	}
	return nil
}

// QueryTutor queries the "tutor" edge of the ReviewTutor entity.
func (rt *ReviewTutor) QueryTutor() *TutorQuery {
	return NewReviewTutorClient(rt.config).QueryTutor(rt)
}

// QueryStudent queries the "student" edge of the ReviewTutor entity.
func (rt *ReviewTutor) QueryStudent() *StudentQuery {
	return NewReviewTutorClient(rt.config).QueryStudent(rt)
}

// Update returns a builder for updating this ReviewTutor.
// Note that you need to call ReviewTutor.Unwrap() before calling this method if this ReviewTutor
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *ReviewTutor) Update() *ReviewTutorUpdateOne {
	return NewReviewTutorClient(rt.config).UpdateOne(rt)
}

// Unwrap unwraps the ReviewTutor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rt *ReviewTutor) Unwrap() *ReviewTutor {
	_tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReviewTutor is not a transactional entity")
	}
	rt.config.driver = _tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *ReviewTutor) String() string {
	var builder strings.Builder
	builder.WriteString("ReviewTutor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rt.ID))
	if v := rt.Score; v != nil {
		builder.WriteString("score=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := rt.ReviewMsg; v != nil {
		builder.WriteString("review_msg=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// ReviewTutors is a parsable slice of ReviewTutor.
type ReviewTutors []*ReviewTutor
