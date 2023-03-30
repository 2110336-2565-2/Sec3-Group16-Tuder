// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldBirthDate holds the string denoting the birth_date field in the database.
	FieldBirthDate = "birth_date"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldProfilePictureURL holds the string denoting the profile_picture_url field in the database.
	FieldProfilePictureURL = "profile_picture_url"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// EdgeTutor holds the string denoting the tutor edge name in mutations.
	EdgeTutor = "tutor"
	// EdgePayment holds the string denoting the payment edge name in mutations.
	EdgePayment = "payment"
	// EdgePaymentHistory holds the string denoting the payment_history edge name in mutations.
	EdgePaymentHistory = "payment_history"
	// EdgeCancelRequest holds the string denoting the cancel_request edge name in mutations.
	EdgeCancelRequest = "cancel_request"
	// Table holds the table name of the user in the database.
	Table = "users"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "students"
	// StudentInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentInverseTable = "students"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "user_student"
	// TutorTable is the table that holds the tutor relation/edge.
	TutorTable = "tutors"
	// TutorInverseTable is the table name for the Tutor entity.
	// It exists in this package in order to avoid circular dependency with the "tutor" package.
	TutorInverseTable = "tutors"
	// TutorColumn is the table column denoting the tutor relation/edge.
	TutorColumn = "user_tutor"
	// PaymentTable is the table that holds the payment relation/edge.
	PaymentTable = "payments"
	// PaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	PaymentInverseTable = "payments"
	// PaymentColumn is the table column denoting the payment relation/edge.
	PaymentColumn = "user_payment"
	// PaymentHistoryTable is the table that holds the payment_history relation/edge.
	PaymentHistoryTable = "payment_histories"
	// PaymentHistoryInverseTable is the table name for the PaymentHistory entity.
	// It exists in this package in order to avoid circular dependency with the "paymenthistory" package.
	PaymentHistoryInverseTable = "payment_histories"
	// PaymentHistoryColumn is the table column denoting the payment_history relation/edge.
	PaymentHistoryColumn = "user_payment_history"
	// CancelRequestTable is the table that holds the cancel_request relation/edge.
	CancelRequestTable = "cancel_requests"
	// CancelRequestInverseTable is the table name for the CancelRequest entity.
	// It exists in this package in order to avoid circular dependency with the "cancelrequest" package.
	CancelRequestInverseTable = "cancel_requests"
	// CancelRequestColumn is the table column denoting the cancel_request relation/edge.
	CancelRequestColumn = "user_cancel_request"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldEmail,
	FieldFirstName,
	FieldLastName,
	FieldAddress,
	FieldPhone,
	FieldBirthDate,
	FieldGender,
	FieldProfilePictureURL,
	FieldRole,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	GenderValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// Role values.
const (
	RoleStudent Role = "student"
	RoleTutor   Role = "tutor"
	RoleAdmin   Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleStudent, RoleTutor, RoleAdmin:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}
