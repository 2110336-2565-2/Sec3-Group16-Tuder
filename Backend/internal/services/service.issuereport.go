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
			IssueReport_id: issue.ID,
			Title:          issue.Title,
			Description:    issue.Description,
			ReportDate:     issue.ReportDate,
			Contact:        issue.Contact,
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

// func (s *serviceCourse) GetCourseByID(sr *schemas.SchemaGetCourse) (*schemas.SchemaCourse, error) {
// 	course, err := s.repository.GetCourseByID(sr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	courseResponse := &schemas.SchemaCourse{
// 		ID:                 course.ID,
// 		Tutor_id:           course.Edges.Tutor.ID,
// 		Title:              course.Title,
// 		Description:        course.Description,
// 		Subject:            course.Subject,
// 		Topic:              course.Topic,
// 		Estimate_time:      course.EstimatedTime,
// 		Price_per_hour:     course.PricePerHour,
// 		Course_picture_url: *course.CoursePictureURL,
// 		Level:              course.Level.String(),
// 	}
// 	return courseResponse, nil
// }

func (s *serviceIssueReport) CreateIssueReport(sr *schemas.SchemaCreateIssueReport) (*schemas.SchemaIssueReport, error) {
	issueReport, err := s.repository.CreateIssueReport(sr)
	if err != nil {
		return nil, err
	}
	issueReportResponses := &schemas.SchemaIssueReport{
		IssueReport_id: issueReport.ID,
		Title:          issueReport.Title,
		Description:    issueReport.Description,
		ReportDate:     issueReport.ReportDate,
		Contact:        issueReport.Contact,
	}
	return issueReportResponses, nil
}

// func (s *serviceCourse) UpdateCourse(sr *schemas.SchemaUpdateCourse) (*schemas.SchemaCourse, error) {
// 	course, err := s.repository.UpdateCourse(sr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	courseResponse := &schemas.SchemaCourse{
// 		ID:                 course.ID,
// 		Tutor_id:           course.Edges.Tutor.ID,
// 		Title:              course.Title,
// 		Description:        course.Description,
// 		Subject:            course.Subject,
// 		Topic:              course.Topic,
// 		Estimate_time:      course.EstimatedTime,
// 		Price_per_hour:     course.PricePerHour,
// 		Course_picture_url: *course.CoursePictureURL,
// 		Level:              course.Level.String(),
// 	}
// 	return courseResponse, nil
// }

// func (s *serviceCourse) DeleteCourse(sr *schemas.SchemaDeleteCourse) error {
// 	err := s.repository.DeleteCourse(sr)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
