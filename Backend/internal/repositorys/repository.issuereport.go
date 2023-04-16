package repositorys

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/issuereport"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryIssueReport interface {
	CreateIssueReport(sr *schema.SchemaCreateIssueReport) (*ent.IssueReport, error)
	GetIssueReports() ([]*ent.IssueReport, error)
	UpdateIssueReport(sr *schema.SchemaUpdateIssueReport) (*ent.IssueReport, error)
	UpdateIssueReportStatus(sr *schema.SchemaUpdateIssueReportStatus) (*ent.IssueReport, error)
	DeleteIssueReport(sr *schema.SchemaDeleteIssueReport) error
}

type repositoryIssueReport struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryIssueReport(c *ent.Client) *repositoryIssueReport {
	return &repositoryIssueReport{client: c, ctx: context.Background()}
}

func (r *repositoryIssueReport) GetIssueReports() ([]*ent.IssueReport, error) {
	issueReports, err := r.client.IssueReport.
		Query().
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve issue reports: %w", err)
	}

	return issueReports, nil
}

func (r *repositoryIssueReport) CreateIssueReport(sr *schema.SchemaCreateIssueReport) (*ent.IssueReport, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start a transaction: %w", err)
	}
	txc := tx.Client()
	contact := "No contact"
	if sr.Contact != "" {
		contact = sr.Contact
	}
	issueReport, err := txc.IssueReport.Create().
		SetTitle(sr.Title).
		SetDescription(sr.Description).
		SetContact(contact).
		SetReportDate(sr.ReportDate).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to create issue report: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return issueReport, nil
}

func (r *repositoryIssueReport) UpdateIssueReport(sr *schema.SchemaUpdateIssueReport) (*ent.IssueReport, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start a transaction: %w", err)
	}
	txc := tx.Client()
	contact := "No contact"
	if sr.Contact != "" {
		contact = sr.Contact
	}

	issueReport, err := txc.IssueReport.UpdateOneID(sr.ID).
		SetTitle(sr.Title).
		SetDescription(sr.Description).
		SetContact(contact).
		SetReportDate(sr.ReportDate).
		SetStatus(issuereport.Status(sr.Status)).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to update issue report: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return issueReport, nil
}

func (r *repositoryIssueReport) UpdateIssueReportStatus(sr *schema.SchemaUpdateIssueReportStatus) (*ent.IssueReport, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start a transaction: %w", err)
	}
	txc := tx.Client()

	issueReport, err := txc.IssueReport.UpdateOneID(sr.ID).
		SetStatus(issuereport.Status(sr.Status)).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to update issue report status: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return issueReport, nil
}

func (r *repositoryIssueReport) DeleteIssueReport(sr *schema.SchemaDeleteIssueReport) error {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return fmt.Errorf("failed to start a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	// delete a course
	err = txc.IssueReport.DeleteOneID(sr.ID).Exec(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return fmt.Errorf("failed to delete Issue report: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
