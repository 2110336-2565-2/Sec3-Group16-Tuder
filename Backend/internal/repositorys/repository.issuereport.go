package repositorys

import (
	"context"
	"errors"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryIssueReport interface {
	CreateIssueReport(sr *schema.SchemaCreateIssueReport) (*ent.IssueReport, error)
	GetIssueReports() ([]*ent.IssueReport, error)
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
		return nil, errors.New("Issue report not found")
	}

	return issueReports, nil
}

func (r *repositoryIssueReport) CreateIssueReport(sr *schema.SchemaCreateIssueReport) (*ent.IssueReport, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()
	// issueReport, err := txc.IssueReport.Query().Where(entIssueReport.IDEQ(sr.IssueReport_id)).Only(r.ctx)

	// if err != nil {
	// 	if rerr := tx.Rollback(); rerr != nil {
	// 		err = fmt.Errorf("%w: %v", err, rerr)
	// 	}
	// 	return nil, fmt.Errorf("getting tutor: %w", err)
	// }
	// create a new course
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
		return nil, fmt.Errorf("creating course: %w", err)
	}
	return issueReport, tx.Commit()
}

// func (r *repositoryCourse) UpdateCourse(sr *schema.SchemaUpdateCourse) (*ent.Course, error) {
// 	// create a transaction
// 	tx, err := r.client.Tx(r.ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("starting a transaction: %w", err)
// 	}
// 	// wrap the client with the transaction
// 	txc := tx.Client()
// 	tutor, err := txc.Tutor.Query().Where(entTutor.IDEQ(sr.Tutor_id)).Only(r.ctx)

// 	if err != nil {
// 		if rerr := tx.Rollback(); rerr != nil {
// 			err = fmt.Errorf("%w: %v", err, rerr)
// 		}
// 		return nil, fmt.Errorf("getting tutor: %w", err)
// 	}

// 	// update a course
// 	course, err := txc.Course.UpdateOneID(sr.ID).
// 		SetCoursePictureURL(sr.Course_picture_url).
// 		SetSubject(sr.Subject).
// 		SetDescription(sr.Description).
// 		SetPricePerHour(sr.Price_per_hour).
// 		SetTutorID(sr.Tutor_id).
// 		SetLevel(entCourse.Level(sr.Level)).
// 		SetTitle(sr.Title).
// 		SetTopic(sr.Topic).
// 		SetEstimatedTime(sr.Estimate_time).
// 		Save(r.ctx)
// 	course.Edges.Tutor = tutor
// 	if err != nil {
// 		if rerr := tx.Rollback(); rerr != nil {
// 			err = fmt.Errorf("%w: %v", err, rerr)
// 		}
// 		return nil, fmt.Errorf("updating course: %w", err)
// 	}
// 	return course, tx.Commit()
// }

// func (r *repositoryCourse) DeleteCourse(sr *schema.SchemaDeleteCourse) error {
// 	// create a transaction
// 	tx, err := r.client.Tx(r.ctx)
// 	if err != nil {
// 		return fmt.Errorf("starting a transaction: %w", err)
// 	}
// 	// wrap the client with the transaction
// 	txc := tx.Client()

// 	// delete a course
// 	err = txc.Course.DeleteOneID(sr.ID).Exec(r.ctx)
// 	if err != nil {
// 		if rerr := tx.Rollback(); rerr != nil {
// 			err = fmt.Errorf("%w: %v", err, rerr)
// 		}
// 		return fmt.Errorf("deleting course: %w", err)
// 	}
// 	return tx.Commit()
// }
