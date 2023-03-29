package repositorys

import (
	"context"
	"errors"

	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	class "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	cancelrequest "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/classcancelrequest"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	match "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	user "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryClass interface {
	GetCancellingClasses() ([]*schemas.SchemaCancelRequest, error)
	CancelClass(sc *schemas.SchemaCancelClass) (*ent.ClassCancelRequest, error)
	AuditClassCancellation(sc *schemas.SchemaCancelRequestApprove) error
	// AcknowledgeClassCancellation(sc *schemas.SchemaUserAcknowledge) error
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

	ccr, err := r.client.ClassCancelRequest.
		Query().
		WithUser().
		All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying all class cancel requests: %w", err)
	}

	var cancelRequests []*schemas.SchemaCancelRequest
	for _, c := range ccr {
		cancelRequests = append(cancelRequests, &schemas.SchemaCancelRequest{
			CancelRequestID: c.ID,
			ClassID:         c.ID,
			Reporter:        c.Edges.User.FirstName + " " + c.Edges.User.LastName,
			ReportDate:      c.ReportDate,
			ImgURL:          c.ImgURL,
			Description:     c.Description,
			Status:          c.Status,
		})
	}

	return cancelRequests, nil
}

func (r *repositoryClass) CancelClass(sc *schemas.SchemaCancelClass) (*ent.ClassCancelRequest, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	// identify user
	u, err := txc.User.
		Query().
		Where(
			user.IDEQ(sc.UserID),
		).
		Only(r.ctx)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// check if user is a reporter
	if u.Role.String() != sc.ReporterRole.String() {
		return nil, errors.New("user is not a reporter")
	}

	var m *ent.Match

	if u.Role.String() == user.RoleTutor.String() {

		m, err = txc.Match.
			Query().
			Where(
				match.And(
					match.HasClassWith( // find the class
						class.IDEQ(sc.ClassID),
					),
					match.HasCourseWith( // find the tutor from course
						course.HasTutorWith( // find the tutor
							tutor.HasUserWith(
								user.IDEQ(sc.UserID),
							),
						),
					),
				),
			).
			Only(r.ctx)
		if err != nil {
			return nil, errors.New(err.Error())
			//return nil, errors.New("matching not found, tutor doesn't have this class")
		}

	} else if u.Role.String() == user.RoleStudent.String() {
		// find the matching

		m, err = txc.Match.
			Query().
			Where(
				match.HasClassWith( // find the class
					class.IDEQ(sc.ClassID),
				),
				match.HasStudentWith( // find the student
					student.HasUserWith( // find the user
						user.IDEQ(sc.UserID),
					),
				),
			).
			Only(r.ctx)
		if err != nil {
			return nil, errors.New("matching not found, student doesn't have this class")
		}

	}

	// find the cancel request if already exists
	ccr, err := txc.ClassCancelRequest.
		Query().
		Where(
			cancelrequest.HasMatchWith(match.IDEQ(m.ID)),
		).
		Only(r.ctx)
	if err == nil {
		return nil, errors.New("cancel request not found")
	}

	// check if class is cancelling
	if ccr.Status == cancelrequest.StatusPending {
		return nil, errors.New("class cancelling is already requested")
	}

	// check if class is already cancelled
	if ccr.Status == cancelrequest.StatusApproved {
		return nil, errors.New("class is already cancelled")
	}

	// create new class cancel request
	ccr, err = txc.ClassCancelRequest.
		Create().
		SetStatus(cancelrequest.StatusPending).
		SetMatchID(m.ID).
		SetDescription(sc.Description).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return ccr, tx.Commit()
}

func (r *repositoryClass) AuditClassCancellation(sc *schemas.SchemaCancelRequestApprove) error {

	ccr, err := r.client.ClassCancelRequest.
		Query().
		Where(cancelrequest.IDEQ(sc.CancelRequestID)).
		Only(r.ctx)

	if err != nil {
		return errors.New("cancel request not found")
	}

	approve := sc.Approve
	if ccr.Status == cancelrequest.StatusApproved {
		return errors.New("cancel request is already approved")
	}

	if ccr.Status != cancelrequest.StatusRejected {
		return errors.New("cancel request is already rejected")
	}

	status := cancelrequest.StatusPending
	if approve {
		status = cancelrequest.StatusApproved
	} else {
		status = cancelrequest.StatusRejected
	}

	_, err = r.client.ClassCancelRequest.
		UpdateOne(ccr).
		SetStatus(status).
		Save(r.ctx)

	if err != nil {
		return err
	}

	return nil
}

// func (r *repositoryClass) AcknowledgeClassCancellation(s *schemas.SchemaUserAcknowledge) error {

// 	c, err := r.client.Class.
// 		Query().
// 		Where(class.IDEQ(s.ClassID)).
// 		WithMatch(
// 			func(q *ent.MatchQuery) {
// 				q.WithStudent(
// 					func(q *ent.StudentQuery) {
// 						q.WithUser()
// 					},
// 				)
// 			},
// 		).
// 		Only(r.ctx)

// 	if err != nil {
// 		return errors.New("class not found")
// 	}

// 	if c.Status != class.StatusRejected {
// 		return errors.New("class is not rejected")
// 	}

// 	// check if user is the student of the class
// 	found := false
// 	for _, m := range c.Edges.Match {
// 		if m.Edges.Student.Edges.User.ID == s.UserID {
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		return errors.New("user is not the student of the class")
// 	}

// 	_, err = r.client.Class.
// 		UpdateOne(c).
// 		SetStatus(class.StatusOngoing).
// 		Save(r.ctx)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
