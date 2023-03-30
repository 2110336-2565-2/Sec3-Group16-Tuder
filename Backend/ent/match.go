// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/google/uuid"
)

// Match is the model entity for the Match schema.
type Match struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// MatchCreatedAt holds the value of the "match_created_at" field.
	MatchCreatedAt time.Time `json:"match_created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MatchQuery when eager-loading is set.
	Edges          MatchEdges `json:"edges"`
	course_match   *uuid.UUID
	schedule_match *uuid.UUID
	student_match  *uuid.UUID
}

// MatchEdges holds the relations/edges for other nodes in the graph.
type MatchEdges struct {
	// Student holds the value of the student edge.
	Student *Student `json:"student,omitempty"`
	// Course holds the value of the course edge.
	Course *Course `json:"course,omitempty"`
	// Appointment holds the value of the appointment edge.
	Appointment []*Appointment `json:"appointment,omitempty"`
	// Schedule holds the value of the schedule edge.
	Schedule *Schedule `json:"schedule,omitempty"`
	// CancelRequest holds the value of the cancel_request edge.
	CancelRequest []*CancelRequest `json:"cancel_request,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MatchEdges) StudentOrErr() (*Student, error) {
	if e.loadedTypes[0] {
		if e.Student == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: student.Label}
		}
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MatchEdges) CourseOrErr() (*Course, error) {
	if e.loadedTypes[1] {
		if e.Course == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: course.Label}
		}
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// AppointmentOrErr returns the Appointment value or an error if the edge
// was not loaded in eager-loading.
func (e MatchEdges) AppointmentOrErr() ([]*Appointment, error) {
	if e.loadedTypes[2] {
		return e.Appointment, nil
	}
	return nil, &NotLoadedError{edge: "appointment"}
}

// ScheduleOrErr returns the Schedule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MatchEdges) ScheduleOrErr() (*Schedule, error) {
	if e.loadedTypes[3] {
		if e.Schedule == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: schedule.Label}
		}
		return e.Schedule, nil
	}
	return nil, &NotLoadedError{edge: "schedule"}
}

// CancelRequestOrErr returns the CancelRequest value or an error if the edge
// was not loaded in eager-loading.
func (e MatchEdges) CancelRequestOrErr() ([]*CancelRequest, error) {
	if e.loadedTypes[4] {
		return e.CancelRequest, nil
	}
	return nil, &NotLoadedError{edge: "cancel_request"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Match) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case match.FieldMatchCreatedAt:
			values[i] = new(sql.NullTime)
		case match.FieldID:
			values[i] = new(uuid.UUID)
		case match.ForeignKeys[0]: // course_match
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case match.ForeignKeys[1]: // schedule_match
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case match.ForeignKeys[2]: // student_match
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Match", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Match fields.
func (m *Match) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case match.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case match.FieldMatchCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field match_created_at", values[i])
			} else if value.Valid {
				m.MatchCreatedAt = value.Time
			}
		case match.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field course_match", values[i])
			} else if value.Valid {
				m.course_match = new(uuid.UUID)
				*m.course_match = *value.S.(*uuid.UUID)
			}
		case match.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_match", values[i])
			} else if value.Valid {
				m.schedule_match = new(uuid.UUID)
				*m.schedule_match = *value.S.(*uuid.UUID)
			}
		case match.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field student_match", values[i])
			} else if value.Valid {
				m.student_match = new(uuid.UUID)
				*m.student_match = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryStudent queries the "student" edge of the Match entity.
func (m *Match) QueryStudent() *StudentQuery {
	return NewMatchClient(m.config).QueryStudent(m)
}

// QueryCourse queries the "course" edge of the Match entity.
func (m *Match) QueryCourse() *CourseQuery {
	return NewMatchClient(m.config).QueryCourse(m)
}

// QueryAppointment queries the "appointment" edge of the Match entity.
func (m *Match) QueryAppointment() *AppointmentQuery {
	return NewMatchClient(m.config).QueryAppointment(m)
}

// QuerySchedule queries the "schedule" edge of the Match entity.
func (m *Match) QuerySchedule() *ScheduleQuery {
	return NewMatchClient(m.config).QuerySchedule(m)
}

// QueryCancelRequest queries the "cancel_request" edge of the Match entity.
func (m *Match) QueryCancelRequest() *CancelRequestQuery {
	return NewMatchClient(m.config).QueryCancelRequest(m)
}

// Update returns a builder for updating this Match.
// Note that you need to call Match.Unwrap() before calling this method if this Match
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Match) Update() *MatchUpdateOne {
	return NewMatchClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Match entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Match) Unwrap() *Match {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Match is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Match) String() string {
	var builder strings.Builder
	builder.WriteString("Match(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("match_created_at=")
	builder.WriteString(m.MatchCreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Matches is a parsable slice of Match.
type Matches []*Match
