package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaIssueReport struct {
	IssueReport_id uuid.UUID `json:"issuereport_id"`
	Title          string    `json:"title,omitempty"`
	Description    string    `json:"description,omitempty"`
	Contact        string    `json:"contact,omitempty"`
	ReportDate     time.Time `json:"report_date,omitempty"`
}

type SchemaCreateIssueReport struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Contact     string    `json:"contact,omitempty"`
	ReportDate  time.Time `json:"report_date,omitempty"`
}
