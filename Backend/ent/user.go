// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
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
	// Role holds the value of the "role" field.
	Role *user.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Student holds the value of the student edge.
	Student *Student `json:"student,omitempty"`
	// Tutor holds the value of the tutor edge.
	Tutor *Tutor `json:"tutor,omitempty"`
	// IssueReport holds the value of the issue_report edge.
	IssueReport []*IssueReport `json:"issue_report,omitempty"`
	// Payment holds the value of the payment edge.
	Payment []*Payment `json:"payment,omitempty"`
	// PaymentHistory holds the value of the payment_history edge.
	PaymentHistory []*PaymentHistory `json:"payment_history,omitempty"`
	// ClassCancelRequest holds the value of the class_cancel_request edge.
	ClassCancelRequest []*ClassCancelRequest `json:"class_cancel_request,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) StudentOrErr() (*Student, error) {
	if e.loadedTypes[0] {
		if e.Student == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: student.Label}
		}
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// TutorOrErr returns the Tutor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TutorOrErr() (*Tutor, error) {
	if e.loadedTypes[1] {
		if e.Tutor == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tutor.Label}
		}
		return e.Tutor, nil
	}
	return nil, &NotLoadedError{edge: "tutor"}
}

// IssueReportOrErr returns the IssueReport value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IssueReportOrErr() ([]*IssueReport, error) {
	if e.loadedTypes[2] {
		return e.IssueReport, nil
	}
	return nil, &NotLoadedError{edge: "issue_report"}
}

// PaymentOrErr returns the Payment value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PaymentOrErr() ([]*Payment, error) {
	if e.loadedTypes[3] {
		return e.Payment, nil
	}
	return nil, &NotLoadedError{edge: "payment"}
}

// PaymentHistoryOrErr returns the PaymentHistory value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PaymentHistoryOrErr() ([]*PaymentHistory, error) {
	if e.loadedTypes[4] {
		return e.PaymentHistory, nil
	}
	return nil, &NotLoadedError{edge: "payment_history"}
}

// ClassCancelRequestOrErr returns the ClassCancelRequest value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ClassCancelRequestOrErr() ([]*ClassCancelRequest, error) {
	if e.loadedTypes[5] {
		return e.ClassCancelRequest, nil
	}
	return nil, &NotLoadedError{edge: "class_cancel_request"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldUsername, user.FieldPassword, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldAddress, user.FieldPhone, user.FieldGender, user.FieldProfilePictureURL, user.FieldRole:
			values[i] = new(sql.NullString)
		case user.FieldBirthDate:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				u.Address = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldBirthDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				u.BirthDate = value.Time
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = value.String
			}
		case user.FieldProfilePictureURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_picture_URL", values[i])
			} else if value.Valid {
				u.ProfilePictureURL = new(string)
				*u.ProfilePictureURL = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = new(user.Role)
				*u.Role = user.Role(value.String)
			}
		}
	}
	return nil
}

// QueryStudent queries the "student" edge of the User entity.
func (u *User) QueryStudent() *StudentQuery {
	return NewUserClient(u.config).QueryStudent(u)
}

// QueryTutor queries the "tutor" edge of the User entity.
func (u *User) QueryTutor() *TutorQuery {
	return NewUserClient(u.config).QueryTutor(u)
}

// QueryIssueReport queries the "issue_report" edge of the User entity.
func (u *User) QueryIssueReport() *IssueReportQuery {
	return NewUserClient(u.config).QueryIssueReport(u)
}

// QueryPayment queries the "payment" edge of the User entity.
func (u *User) QueryPayment() *PaymentQuery {
	return NewUserClient(u.config).QueryPayment(u)
}

// QueryPaymentHistory queries the "payment_history" edge of the User entity.
func (u *User) QueryPaymentHistory() *PaymentHistoryQuery {
	return NewUserClient(u.config).QueryPaymentHistory(u)
}

// QueryClassCancelRequest queries the "class_cancel_request" edge of the User entity.
func (u *User) QueryClassCancelRequest() *ClassCancelRequestQuery {
	return NewUserClient(u.config).QueryClassCancelRequest(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(u.Address)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("birth_date=")
	builder.WriteString(u.BirthDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(u.Gender)
	builder.WriteString(", ")
	if v := u.ProfilePictureURL; v != nil {
		builder.WriteString("profile_picture_URL=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Role; v != nil {
		builder.WriteString("role=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
