package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaIssueReport struct {
	ID          uuid.UUID `json:"issuereport_id"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Contact     string    `json:"contact,omitempty"`
	ReportDate  time.Time `json:"report_date,omitempty"`
	Status      string    `json:"status,omitempty"`
}

type SchemaCreateIssueReport struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Contact     string    `json:"contact,omitempty"`
	ReportDate  time.Time `json:"report_date,omitempty"`
}

type SchemaUpdateIssueReportStatus struct {
	ID     uuid.UUID `json:"issuereport_id"`
	Status string    `json:"status"`
}

type SchemaUpdateIssueReport struct {
	ID          uuid.UUID `json:"issuereport_id"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Contact     string    `json:"contact,omitempty"`
	ReportDate  time.Time `json:"report_date,omitempty"`
	Status      string    `json:"status"`
}

type SchemaDeleteIssueReport struct {
	ID uuid.UUID `json:"issuereport_id"`
}
