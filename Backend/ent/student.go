// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// Student is the model entity for the Student schema.
type Student struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges        StudentEdges `json:"edges"`
	user_student *uuid.UUID
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// IssueReport holds the value of the issue_report edge.
	IssueReport []*IssueReport `json:"issue_report,omitempty"`
	// Course holds the value of the course edge.
	Course *Course `json:"course,omitempty"`
	// Class holds the value of the class edge.
	Class []*Class `json:"class,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// IssueReportOrErr returns the IssueReport value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) IssueReportOrErr() ([]*IssueReport, error) {
	if e.loadedTypes[0] {
		return e.IssueReport, nil
	}
	return nil, &NotLoadedError{edge: "issue_report"}
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) CourseOrErr() (*Course, error) {
	if e.loadedTypes[1] {
		if e.Course == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: course.Label}
		}
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) ClassOrErr() ([]*Class, error) {
	if e.loadedTypes[2] {
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			values[i] = new(uuid.UUID)
		case student.ForeignKeys[0]: // user_student
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Student", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case student.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_student", values[i])
			} else if value.Valid {
				s.user_student = new(uuid.UUID)
				*s.user_student = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryIssueReport queries the "issue_report" edge of the Student entity.
func (s *Student) QueryIssueReport() *IssueReportQuery {
	return NewStudentClient(s.config).QueryIssueReport(s)
}

// QueryCourse queries the "course" edge of the Student entity.
func (s *Student) QueryCourse() *CourseQuery {
	return NewStudentClient(s.config).QueryCourse(s)
}

// QueryClass queries the "class" edge of the Student entity.
func (s *Student) QueryClass() *ClassQuery {
	return NewStudentClient(s.config).QueryClass(s)
}

// QueryUser queries the "user" edge of the Student entity.
func (s *Student) QueryUser() *UserQuery {
	return NewStudentClient(s.config).QueryUser(s)
}

// Update returns a builder for updating this Student.
// Note that you need to call Student.Unwrap() before calling this method if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return NewStudentClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Student entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student
