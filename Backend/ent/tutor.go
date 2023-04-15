// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// Tutor is the model entity for the Tutor schema.
type Tutor struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// OmiseBankToken holds the value of the "omise_bank_token" field.
	OmiseBankToken *string `json:"omise_bank_token,omitempty"`
	// CitizenID holds the value of the "citizen_id" field.
	CitizenID string `json:"citizen_id,omitempty"`
	// OmiseCustomerID holds the value of the "omise_customer_id" field.
	OmiseCustomerID *string `json:"omise_customer_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TutorQuery when eager-loading is set.
	Edges          TutorEdges `json:"edges"`
	schedule_tutor *uuid.UUID
	user_tutor     *uuid.UUID
}

// TutorEdges holds the relations/edges for other nodes in the graph.
type TutorEdges struct {
	// IssueReport holds the value of the issue_report edge.
	IssueReport []*IssueReport `json:"issue_report,omitempty"`
	// Course holds the value of the course edge.
	Course []*Course `json:"course,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Schedule holds the value of the schedule edge.
	Schedule *Schedule `json:"schedule,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// IssueReportOrErr returns the IssueReport value or an error if the edge
// was not loaded in eager-loading.
func (e TutorEdges) IssueReportOrErr() ([]*IssueReport, error) {
	if e.loadedTypes[0] {
		return e.IssueReport, nil
	}
	return nil, &NotLoadedError{edge: "issue_report"}
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading.
func (e TutorEdges) CourseOrErr() ([]*Course, error) {
	if e.loadedTypes[1] {
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TutorEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ScheduleOrErr returns the Schedule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TutorEdges) ScheduleOrErr() (*Schedule, error) {
	if e.loadedTypes[3] {
		if e.Schedule == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: schedule.Label}
		}
		return e.Schedule, nil
	}
	return nil, &NotLoadedError{edge: "schedule"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tutor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tutor.FieldDescription, tutor.FieldOmiseBankToken, tutor.FieldCitizenID, tutor.FieldOmiseCustomerID:
			values[i] = new(sql.NullString)
		case tutor.FieldID:
			values[i] = new(uuid.UUID)
		case tutor.ForeignKeys[0]: // schedule_tutor
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case tutor.ForeignKeys[1]: // user_tutor
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tutor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tutor fields.
func (t *Tutor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tutor.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tutor.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = new(string)
				*t.Description = value.String
			}
		case tutor.FieldOmiseBankToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field omise_bank_token", values[i])
			} else if value.Valid {
				t.OmiseBankToken = new(string)
				*t.OmiseBankToken = value.String
			}
		case tutor.FieldCitizenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field citizen_id", values[i])
			} else if value.Valid {
				t.CitizenID = value.String
			}
		case tutor.FieldOmiseCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field omise_customer_id", values[i])
			} else if value.Valid {
				t.OmiseCustomerID = new(string)
				*t.OmiseCustomerID = value.String
			}
		case tutor.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_tutor", values[i])
			} else if value.Valid {
				t.schedule_tutor = new(uuid.UUID)
				*t.schedule_tutor = *value.S.(*uuid.UUID)
			}
		case tutor.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_tutor", values[i])
			} else if value.Valid {
				t.user_tutor = new(uuid.UUID)
				*t.user_tutor = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryIssueReport queries the "issue_report" edge of the Tutor entity.
func (t *Tutor) QueryIssueReport() *IssueReportQuery {
	return NewTutorClient(t.config).QueryIssueReport(t)
}

// QueryCourse queries the "course" edge of the Tutor entity.
func (t *Tutor) QueryCourse() *CourseQuery {
	return NewTutorClient(t.config).QueryCourse(t)
}

// QueryUser queries the "user" edge of the Tutor entity.
func (t *Tutor) QueryUser() *UserQuery {
	return NewTutorClient(t.config).QueryUser(t)
}

// QuerySchedule queries the "schedule" edge of the Tutor entity.
func (t *Tutor) QuerySchedule() *ScheduleQuery {
	return NewTutorClient(t.config).QuerySchedule(t)
}

// Update returns a builder for updating this Tutor.
// Note that you need to call Tutor.Unwrap() before calling this method if this Tutor
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tutor) Update() *TutorUpdateOne {
	return NewTutorClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tutor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tutor) Unwrap() *Tutor {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tutor is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tutor) String() string {
	var builder strings.Builder
	builder.WriteString("Tutor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	if v := t.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := t.OmiseBankToken; v != nil {
		builder.WriteString("omise_bank_token=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("citizen_id=")
	builder.WriteString(t.CitizenID)
	builder.WriteString(", ")
	if v := t.OmiseCustomerID; v != nil {
		builder.WriteString("omise_customer_id=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Tutors is a parsable slice of Tutor.
type Tutors []*Tutor
