package repositorys

import (
	"context"
	"errors"

	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	class "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryCancelClass interface {
	CancelClass(sc *schemas.SchemaCancelClass) (*ent.Class, error)
}

type repositoryCancelClass struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryCancelClass(c *ent.Client) RepositoryCancelClass {
	return &repositoryCancelClass{
		client: c,
		ctx:    context.Background(),
	}
}

func (r *repositoryCancelClass) CancelClass(sc *schemas.SchemaCancelClass) (*ent.Class, error) {
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
		Where(class.IDEQ(sc.ClassId)).
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
	if sc.Submitter_role == "tutor" {
		if c.Edges.Match[0].Edges.Course[0].Edges.Tutor.Edges.User.ID != sc.UserID {
			return nil, errors.New("user is not the owner of the class")
		}
	} else if sc.Submitter_role == "student" {
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
