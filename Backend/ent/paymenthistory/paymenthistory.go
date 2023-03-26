// Code generated by ent, DO NOT EDIT.

package paymenthistory

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the paymenthistory type in the database.
	Label = "payment_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePayment holds the string denoting the payment edge name in mutations.
	EdgePayment = "payment"
	// Table holds the table name of the paymenthistory in the database.
	Table = "payment_histories"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "classes"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "payment_history_class"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "payment_histories"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_payment_history"
	// PaymentTable is the table that holds the payment relation/edge.
	PaymentTable = "payment_histories"
	// PaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	PaymentInverseTable = "payments"
	// PaymentColumn is the table column denoting the payment relation/edge.
	PaymentColumn = "payment_payment_history"
)

// Columns holds all SQL columns for paymenthistory fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payment_histories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"payment_payment_history",
	"user_payment_history",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
