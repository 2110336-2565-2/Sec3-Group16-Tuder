package services

import (
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceIssueReport interface {
	GetIssueReports() ([]*schemas.SchemaIssueReport, error)
	CreateIssueReport(sr *schemas.SchemaCreateIssueReport) (*schemas.SchemaIssueReport, error)
	PackToSchema(issueReports []*ent.IssueReport) []*schemas.SchemaIssueReport
	UpdateIssueReport(sr *schemas.SchemaUpdateIssueReport) (*schemas.SchemaIssueReport, error)
	UpdateIssueReportStatus(sr *schemas.SchemaUpdateIssueReportStatus) (*schemas.SchemaIssueReport, error)
	DeleteIssueReport(sr *schemas.SchemaDeleteIssueReport) error
}

type serviceIssueReport struct {
	repository repository.RepositoryIssueReport
}

func NewServiceIssueReport(repository repository.RepositoryIssueReport) *serviceIssueReport {
	return &serviceIssueReport{repository: repository}
}

func (s *serviceIssueReport) PackToSchema(issueReports []*ent.IssueReport) []*schemas.SchemaIssueReport {
	var issueReportResponses []*schemas.SchemaIssueReport
	for _, issue := range issueReports {
		issueReportResponses = append(issueReportResponses, &schemas.SchemaIssueReport{
			ID:          issue.ID,
			Title:       issue.Title,
			Description: issue.Description,
			ReportDate:  issue.ReportDate,
			Contact:     issue.Contact,
			Status:      string(issue.Status),
		})
	}
	return issueReportResponses
}

func (s *serviceIssueReport) GetIssueReports() ([]*schemas.SchemaIssueReport, error) {
	issueReports, err := s.repository.GetIssueReports()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	IssueReportResponses := s.PackToSchema(issueReports)
	return IssueReportResponses, nil
}

func (s *serviceIssueReport) CreateIssueReport(sr *schemas.SchemaCreateIssueReport) (*schemas.SchemaIssueReport, error) {
	issueReport, err := s.repository.CreateIssueReport(sr)
	if err != nil {
		return nil, err
	}
	issueReportResponses := &schemas.SchemaIssueReport{
		ID:          issueReport.ID,
		Title:       issueReport.Title,
		Description: issueReport.Description,
		ReportDate:  issueReport.ReportDate,
		Contact:     issueReport.Contact,
		Status:      issueReport.Status.String(),
	}
	return issueReportResponses, nil
}

func (s *serviceIssueReport) UpdateIssueReport(sr *schemas.SchemaUpdateIssueReport) (*schemas.SchemaIssueReport, error) {
	issueReport, err := s.repository.UpdateIssueReport(sr)
	if err != nil {
		return nil, err
	}
	issueReportResponse := &schemas.SchemaIssueReport{
		ID:          issueReport.ID,
		Title:       issueReport.Title,
		Description: issueReport.Description,
		Contact:     issueReport.Contact,
		ReportDate:  issueReport.ReportDate,
		Status:      issueReport.Status.String(),
	}
	return issueReportResponse, nil
}

func (s *serviceIssueReport) UpdateIssueReportStatus(sr *schemas.SchemaUpdateIssueReportStatus) (*schemas.SchemaIssueReport, error) {
	issueReport, err := s.repository.UpdateIssueReportStatus(sr)
	if err != nil {
		return nil, err
	}
	issueReportResponse := &schemas.SchemaIssueReport{
		ID:          issueReport.ID,
		Title:       issueReport.Title,
		Description: issueReport.Description,
		Contact:     issueReport.Contact,
		ReportDate:  issueReport.ReportDate,
		Status:      issueReport.Status.String(),
	}
	return issueReportResponse, nil
}

func (s *serviceIssueReport) DeleteIssueReport(sr *schemas.SchemaDeleteIssueReport) error {
	err := s.repository.DeleteIssueReport(sr)
	if err != nil {
		return err
	}
	return nil
}
