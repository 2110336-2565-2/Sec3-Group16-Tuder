// Code generated by ent, DO NOT EDIT.

package class

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the class type in the database.
	Label = "class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReviewAvaliable holds the string denoting the review_avaliable field in the database.
	FieldReviewAvaliable = "review_avaliable"
	// FieldTotalHour holds the string denoting the total_hour field in the database.
	FieldTotalHour = "total_hour"
	// FieldSuccessHour holds the string denoting the success_hour field in the database.
	FieldSuccessHour = "success_hour"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeMatch holds the string denoting the match edge name in mutations.
	EdgeMatch = "match"
	// EdgeSchedule holds the string denoting the schedule edge name in mutations.
	EdgeSchedule = "schedule"
	// EdgePaymentHistory holds the string denoting the payment_history edge name in mutations.
	EdgePaymentHistory = "payment_history"
	// Table holds the table name of the class in the database.
	Table = "classes"
	// MatchTable is the table that holds the match relation/edge. The primary key declared below.
	MatchTable = "class_match"
	// MatchInverseTable is the table name for the Match entity.
	// It exists in this package in order to avoid circular dependency with the "match" package.
	MatchInverseTable = "matches"
	// ScheduleTable is the table that holds the schedule relation/edge.
	ScheduleTable = "classes"
	// ScheduleInverseTable is the table name for the Schedule entity.
	// It exists in this package in order to avoid circular dependency with the "schedule" package.
	ScheduleInverseTable = "schedules"
	// ScheduleColumn is the table column denoting the schedule relation/edge.
	ScheduleColumn = "schedule_class"
	// PaymentHistoryTable is the table that holds the payment_history relation/edge.
	PaymentHistoryTable = "classes"
	// PaymentHistoryInverseTable is the table name for the PaymentHistory entity.
	// It exists in this package in order to avoid circular dependency with the "paymenthistory" package.
	PaymentHistoryInverseTable = "payment_histories"
	// PaymentHistoryColumn is the table column denoting the payment_history relation/edge.
	PaymentHistoryColumn = "payment_history_class"
)

// Columns holds all SQL columns for class fields.
var Columns = []string{
	FieldID,
	FieldReviewAvaliable,
	FieldTotalHour,
	FieldSuccessHour,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "classes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"payment_history_class",
	"schedule_class",
}

var (
	// MatchPrimaryKey and MatchColumn2 are the table columns denoting the
	// primary key for the match relation (M2M).
	MatchPrimaryKey = []string{"class_id", "match_id"}
)

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
	// DefaultReviewAvaliable holds the default value on creation for the "review_avaliable" field.
	DefaultReviewAvaliable bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusScheduled  Status = "scheduled"
	StatusCompleted  Status = "completed"
	StatusCancelling Status = "cancelling"
	StatusRejected   Status = "rejected"
	StatusCancelled  Status = "cancelled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusScheduled, StatusCompleted, StatusCancelling, StatusRejected, StatusCancelled:
		return nil
	default:
		return fmt.Errorf("class: invalid enum value for status field: %q", s)
	}
}
