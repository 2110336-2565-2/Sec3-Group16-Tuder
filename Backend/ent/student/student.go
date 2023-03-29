// Code generated by ent, DO NOT EDIT.

package student

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the student type in the database.
	Label = "student"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeMatch holds the string denoting the match edge name in mutations.
	EdgeMatch = "match"
	// EdgeReview holds the string denoting the review edge name in mutations.
	EdgeReview = "review"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the student in the database.
	Table = "students"
	// MatchTable is the table that holds the match relation/edge.
	MatchTable = "matches"
	// MatchInverseTable is the table name for the Match entity.
	// It exists in this package in order to avoid circular dependency with the "match" package.
	MatchInverseTable = "matches"
	// MatchColumn is the table column denoting the match relation/edge.
	MatchColumn = "student_match"
	// ReviewTable is the table that holds the review relation/edge. The primary key declared below.
	ReviewTable = "student_review"
	// ReviewInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewInverseTable = "reviews"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "students"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_student"
)

// Columns holds all SQL columns for student fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "students"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_student",
}

var (
	// ReviewPrimaryKey and ReviewColumn2 are the table columns denoting the
	// primary key for the review relation (M2M).
	ReviewPrimaryKey = []string{"student_id", "review_id"}
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
