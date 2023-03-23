package repositorys

import (
	"context"
	"errors"

	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	class "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryClass interface {
	GetCancellingClasses() ([]*schemas.SchemaCancelRequest, error)
	CancelClass(sc *schemas.SchemaCancelClass) (*ent.Class, error)
	AuditClassCancellation(sc *schemas.SchemaCancelRequestApprove) error
	AcknowledgeClassCancellation(sc *schemas.SchemaUserAcknowledge) error
}

type repositoryClass struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryClass(c *ent.Client) RepositoryClass {
	return &repositoryClass{
		client: c,
		ctx:    context.Background(),
	}
}

func (r *repositoryClass) GetCancellingClasses() ([]*schemas.SchemaCancelRequest, error) {
	classes, err := r.client.Class.
		Query().
		Where(class.StatusEQ(class.StatusCancelling)).
		WithMatch(
			func(q *ent.MatchQuery) {
				q.WithCourse(
					func(q *ent.CourseQuery) {
						q.WithTutor(
							func(q *ent.TutorQuery) {
								q.WithUser()
							},
						)
					},
				).
					WithStudent(
						func(q *ent.StudentQuery) {
							q.WithUser()
						},
					)
			},
		).
		All(r.ctx)
	if err != nil {
		return nil, err
	}

	var cancelRequests []*schemas.SchemaCancelRequest
	for _, c := range classes {
		cancelRequests = append(cancelRequests, &schemas.SchemaCancelRequest{
			ClassID:     c.ID,
			Course:      c.Edges.Match[0].Edges.Course[0].Title,
			Tutorname:   c.Edges.Match[0].Edges.Course[0].Edges.Tutor.Edges.User.FirstName + " " + c.Edges.Match[0].Edges.Course[0].Edges.Tutor.Edges.User.LastName,
			Studentname: c.Edges.Match[0].Edges.Student.Edges.User.FirstName + " " + c.Edges.Match[0].Edges.Student.Edges.User.LastName,
			Subject:     c.Edges.Match[0].Edges.Course[0].Subject,
			TotalHour:   c.TotalHour,
			SuccessHour: c.SuccessHour,
			Price:       c.Edges.Match[0].Edges.Course[0].PricePerHour,
		})
	}

	return cancelRequests, nil
}

func (r *repositoryClass) CancelClass(sc *schemas.SchemaCancelClass) (*ent.Class, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	// find class
	c, err := txc.Class.
		Query().
		Where(class.IDEQ(sc.ClassID)).
		WithMatch(
			func(q *ent.MatchQuery) {
				q.WithCourse(
					func(q *ent.CourseQuery) {
						q.WithTutor(
							func(q *ent.TutorQuery) {
								q.WithUser()
							},
						)
					},
				).
					WithStudent(
						func(q *ent.StudentQuery) {
							q.WithUser()
						},
					)
			},
		).
		Only(r.ctx)
	if err != nil {
		return nil, errors.New("class not found")
	}

	// check if class is cancelling
	if c.Status == class.StatusCancelling {
		return nil, errors.New("class cancelling is already requested")
	}

	// check if class is already cancelled
	if c.Status == class.StatusCancelled {
		return nil, errors.New("class is already cancelled")
	}

	// check if user is the owner of the class
	if sc.SubmitterRole == "tutor" {
		if c.Edges.Match[0].Edges.Course[0].Edges.Tutor.Edges.User.ID != sc.UserID {
			return nil, errors.New("user is not the owner of the class")
		}
	} else if sc.SubmitterRole == "student" {
		// check if user is the student of the class
		found := false
		for _, m := range c.Edges.Match {
			if m.Edges.Student.Edges.User.ID == sc.UserID {
				found = true
				break
			}
		}
		if !found {
			return nil, errors.New("user is not the student of the class")
		}
	} else {
		return nil, errors.New("invalid user role")
	}

	// change class state to cancelling
	c, err = txc.Class.
		UpdateOne(c).
		SetStatus(class.StatusCancelling).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return c, tx.Commit()
}

func (r *repositoryClass) AuditClassCancellation(sc *schemas.SchemaCancelRequestApprove) error {

	c, err := r.client.Class.
		Query().
		Where(class.IDEQ(sc.ClassID)).
		Only(r.ctx)

	if err != nil {
		return errors.New("class not found")
	}

	approve := sc.Approve

	if approve {
		if c.Status == class.StatusCancelled {
			return errors.New("class is already cancelled")
		}

		if c.Status != class.StatusCancelling {
			return errors.New("class is not cancelling")
		}

		_, err = r.client.Class.
			UpdateOne(c).
			SetStatus(class.StatusCancelled).
			Save(r.ctx)

		if err != nil {
			return err
		}
	} else {

		if c.Status == class.StatusRejected {
			return errors.New("class is already rejected")
		}

		if c.Status != class.StatusCancelling {
			return errors.New("class is not cancelling")
		}

		_, err = r.client.Class.
			UpdateOne(c).
			SetStatus(class.StatusRejected).
			Save(r.ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repositoryClass) AcknowledgeClassCancellation(s *schemas.SchemaUserAcknowledge) error {

	c, err := r.client.Class.
		Query().
		Where(class.IDEQ(s.ClassID)).
		WithMatch(
			func(q *ent.MatchQuery) {
				q.WithStudent(
					func(q *ent.StudentQuery) {
						q.WithUser()
					},
				)
			},
		).
		Only(r.ctx)

	if err != nil {
		return errors.New("class not found")
	}

	if c.Status != class.StatusRejected {
		return errors.New("class is not rejected")
	}

	// check if user is the student of the class
	found := false
	for _, m := range c.Edges.Match {
		if m.Edges.Student.Edges.User.ID == s.UserID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("user is not the student of the class")
	}

	_, err = r.client.Class.
		UpdateOne(c).
		SetStatus(class.StatusOngoing).
		Save(r.ctx)

	if err != nil {
		return err
	}

	return nil
}
