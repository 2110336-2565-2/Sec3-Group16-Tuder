// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/google/uuid"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate time.Time `json:"birth_date,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// ProfilePictureURL holds the value of the "profile_picture_URL" field.
	ProfilePictureURL *string `json:"profile_picture_URL,omitempty"`
	// School holds the value of the "school" field.
	School string `json:"school,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldUsername, student.FieldPassword, student.FieldEmail, student.FieldFirstName, student.FieldLastName, student.FieldAddress, student.FieldPhone, student.FieldGender, student.FieldProfilePictureURL, student.FieldSchool:
			values[i] = new(sql.NullString)
		case student.FieldBirthDate:
			values[i] = new(sql.NullTime)
		case student.FieldID:
			values[i] = new(uuid.UUID)
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
		case student.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				s.Username = value.String
			}
		case student.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				s.Password = value.String
			}
		case student.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				s.Email = value.String
			}
		case student.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				s.FirstName = value.String
			}
		case student.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				s.LastName = value.String
			}
		case student.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				s.Address = value.String
			}
		case student.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				s.Phone = value.String
			}
		case student.FieldBirthDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				s.BirthDate = value.Time
			}
		case student.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				s.Gender = value.String
			}
		case student.FieldProfilePictureURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_picture_URL", values[i])
			} else if value.Valid {
				s.ProfilePictureURL = new(string)
				*s.ProfilePictureURL = value.String
			}
		case student.FieldSchool:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field school", values[i])
			} else if value.Valid {
				s.School = value.String
			}
		}
	}
	return nil
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
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("username=")
	builder.WriteString(s.Username)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(s.Email)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(s.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(s.LastName)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(s.Address)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(s.Phone)
	builder.WriteString(", ")
	builder.WriteString("birth_date=")
	builder.WriteString(s.BirthDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(s.Gender)
	builder.WriteString(", ")
	if v := s.ProfilePictureURL; v != nil {
		builder.WriteString("profile_picture_URL=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("school=")
	builder.WriteString(s.School)
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student
