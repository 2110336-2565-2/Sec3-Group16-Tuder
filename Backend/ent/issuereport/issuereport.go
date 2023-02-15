// Code generated by ent, DO NOT EDIT.

package issuereport

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the issuereport type in the database.
	Label = "issue_report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldReportDate holds the string denoting the report_date field in the database.
	FieldReportDate = "report_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the issuereport in the database.
	Table = "issue_reports"
)

// Columns holds all SQL columns for issuereport fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldReportDate,
	FieldStatus,
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// ReportDateValidator is a validator for the "report_date" field. It is called by the builders before save.
	ReportDateValidator func(string) error
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
