package repositorys

import (
	"context"
	"errors"

	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	appointment "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/cancelrequest"
	course "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	match "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	schedule "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	tutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	user "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	utils "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"

	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryCancelRequest interface {
	GetCancellingRequest(id uuid.UUID) (*schemas.SchemaCancelRequestDetail, error)
	GetCancellingRequests() ([]*schemas.SchemaCancelRequestDetail, error)
	CancelRequest(sc *schemas.SchemaCancelRequest) (*ent.CancelRequest, error)
	AuditRequest(sc *schemas.SchemaCancelRequestApprove) error
	// AcknowledgeClassCancellation(sc *schemas.SchemaUserAcknowledge) error
}

type repositoryCancelRequest struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryCancelRequest(c *ent.Client) RepositoryCancelRequest {
	return &repositoryCancelRequest{
		client: c,
		ctx:    context.Background(),
	}
}

func (r *repositoryCancelRequest) GetCancellingRequest(id uuid.UUID) (*schemas.SchemaCancelRequestDetail, error) {

	ccr, err := r.client.CancelRequest.
		Query().
		Where(cancelrequest.IDEQ(id)).
		WithMatch().
		WithUser().
		Only(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying match cancel request: %w", err)
	}

	cancelRequest := &schemas.SchemaCancelRequestDetail{
		CancelRequestID: ccr.ID,
		Title:           ccr.Title,
		MatchID:         ccr.Edges.Match.ID,
		Reporter:        ccr.Edges.User.FirstName + " " + ccr.Edges.User.LastName,
		ReportDate:      ccr.ReportDate,
		ImgURL:          ccr.ImgURL,
		Description:     ccr.Description,
		Status:          ccr.Status,
	}

	return cancelRequest, nil
}

func (r *repositoryCancelRequest) GetCancellingRequests() ([]*schemas.SchemaCancelRequestDetail, error) {

	ccr, err := r.client.CancelRequest.
		Query().
		WithMatch().
		WithUser().
		All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying all match cancel requests: %w", err)
	}

	var cancelRequests []*schemas.SchemaCancelRequestDetail
	for _, c := range ccr {
		cancelRequests = append(cancelRequests, &schemas.SchemaCancelRequestDetail{
			CancelRequestID: c.ID,
			Title:           c.Title,
			MatchID:         c.Edges.Match.ID,
			Reporter:        c.Edges.User.FirstName + " " + c.Edges.User.LastName,
			ReportDate:      c.ReportDate,
			ImgURL:          c.ImgURL,
			Description:     c.Description,
			Status:          c.Status,
		})
	}

	return cancelRequests, nil
}

func (r *repositoryCancelRequest) CancelRequest(sc *schemas.SchemaCancelRequest) (*ent.CancelRequest, error) {
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

	// identify match
	m, err := txc.Match.
		Query().
		Where(
			match.IDEQ(sc.MatchID),
		).
		WithStudent(
			func(sq *ent.StudentQuery) {
				sq.WithUser()
			},
		).
		WithCourse(
			func(sq *ent.CourseQuery) {
				sq.WithTutor(
					func(sq *ent.TutorQuery) {
						sq.WithUser()
					},
				)
			},
		).
		Only(r.ctx)
	if err != nil {
		return nil, errors.New("match not found")
	}

	// find the cancel request if already exists
	ccr, err := txc.CancelRequest.
		Query().
		Where(
			cancelrequest.HasMatchWith(
				match.IDEQ(m.ID),
			),
		).
		Only(r.ctx)
	if err == nil {
		// check if class is cancelling
		if ccr.Status.String() == cancelrequest.StatusPending.String() {
			return nil, errors.New("match cancelling is already requested")
		}

		// check if class is already cancelled
		if ccr.Status.String() == cancelrequest.StatusApproved.String() {
			return nil, errors.New("match is already cancelled")
		}
	}

	// check if student is a student of the match
	if m.Edges.Student.Edges.User.Role.String() == sc.ReporterRole.String() {
		if m.Edges.Student.Edges.User.ID != u.ID {
			return nil, errors.New("user is not a student of the match")
		}
	} else if m.Edges.Student.Edges.User.Role.String() == sc.ReporterRole.String() {
		if m.Edges.Course.Edges.Tutor.Edges.User.ID != u.ID {
			return nil, errors.New("user is not a tutor of the match")
		}
	} else {
		return nil, errors.New("user is not a student or tutor of the match")
	}

	// cancelling match status
	_, err = txc.Match.
		Update().
		Where(
			match.IDEQ(m.ID),
		).
		SetStatus(match.StatusCancelling).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	// create image url
	imgURL := fmt.Sprintf("%s/%s/%s", sc.MatchID.String(), sc.UserID.String(), uuid.New())

	// upload image to s3
	imgURL, err = utils.GenerateProfilePictureURL(sc.Img, imgURL, "ProfilePicture")

	// create new class cancel request
	ccr, err = txc.CancelRequest.
		Create().
		SetMatchID(m.ID).
		SetStatus(cancelrequest.StatusPending).
		SetImgURL(imgURL).
		SetTitle(sc.Title).
		SetDescription(sc.Description).
		SetUserID(u.ID).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return ccr, tx.Commit()
}

func (r *repositoryCancelRequest) AuditRequest(sc *schemas.SchemaCancelRequestApprove) error {

	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	// wrap the client with the transaction
	txc := tx.Client()

	ccr, err := txc.CancelRequest.
		Query().
		Where(cancelrequest.IDEQ(sc.CancelRequestID)).
		Only(r.ctx)

	if err != nil {
		return errors.New("cancel request not found")
	}

	approve := sc.Approve

	status := cancelrequest.StatusPending
	mStatus := match.StatusCancelling
	if approve {
		status = cancelrequest.StatusApproved
		appStatus := appointment.StatusCanceled
		// cancel all appointments of the match
		_, err = txc.Appointment.
			Update().
			Where(
				appointment.HasMatchWith(
					match.IDEQ(sc.MatchID),
				),
			).
			SetStatus(appStatus).
			Save(r.ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return fmt.Errorf("%w: %v", err, rerr)
			}
			return err
		}
		// return available time to tutor
		var new_schedule [7][24]bool
		Tutor, err1 := txc.Tutor.Query().Where(tutor.HasCourseWith(course.HasMatchWith(match.IDEQ(sc.MatchID)))).Only(r.ctx)
		if err1 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err1 = fmt.Errorf("%w: %v", err1, rerr)
			}
			return fmt.Errorf("Tutor not found: %w", err1)
		}
		match_schedule, err2 := txc.Schedule.Query().Where(schedule.HasMatchWith(match.IDEQ(sc.MatchID))).Only(r.ctx)
		if err2 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err2 = fmt.Errorf("%w: %v", err2, rerr)
			}
			return fmt.Errorf("Match not found: %w", err2)
		}
		tutor_schedule, err3 := txc.Schedule.Query().Where(schedule.HasTutorWith(tutor.IDEQ(Tutor.ID))).Only(r.ctx)
		if err3 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err3 = fmt.Errorf("%w: %v", err3, rerr)
			}
			return fmt.Errorf("Tutor's schedule not found: %w", err3)
		}
		for i := 0; i < 24; i++ {
			new_schedule[0][i] = match_schedule.Day0[i] || tutor_schedule.Day0[i]
		}
		for i := 0; i < 24; i++ {
			new_schedule[1][i] = match_schedule.Day1[i] || tutor_schedule.Day1[i]
		}
		for i := 0; i < 24; i++ {
			new_schedule[2][i] = match_schedule.Day2[i] || tutor_schedule.Day2[i]
		}
		for i := 0; i < 24; i++ {
			new_schedule[3][i] = match_schedule.Day3[i] || tutor_schedule.Day3[i]
		}
		for i := 0; i < 24; i++ {
			new_schedule[4][i] = match_schedule.Day4[i] || tutor_schedule.Day4[i]
		}
		for i := 0; i < 24; i++ {
			new_schedule[5][i] = match_schedule.Day5[i] || tutor_schedule.Day5[i]
		}
		for i := 0; i < 24; i++ {
			new_schedule[6][i] = match_schedule.Day6[i] || tutor_schedule.Day6[i]
		}
		_, err4 := tutor_schedule.Update().
			SetDay0(new_schedule[0]).
			SetDay1(new_schedule[1]).
			SetDay2(new_schedule[2]).
			SetDay3(new_schedule[3]).
			SetDay4(new_schedule[4]).
			SetDay5(new_schedule[5]).
			SetDay6(new_schedule[6]).
			Save(r.ctx)
		if err4 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err4 = fmt.Errorf("%w: %v", err4, rerr)
			}
			return fmt.Errorf("updating tutor's schedule: %w", err4)
		}

		mStatus = match.StatusCanceled
	} else {
		status = cancelrequest.StatusRejected
		mStatus = match.StatusEnrolled
	}

	if ccr.Status == cancelrequest.StatusApproved {
		err = errors.New("cancel request is already approved")

		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return fmt.Errorf("%w: %v", err, rerr)
			}
			return err
		}
	}
	if ccr.Status == cancelrequest.StatusRejected {
		err = errors.New("cancel request is already rejected")

		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return fmt.Errorf("%w: %v", err, rerr)
			}
			return err
		}
	}

	_, err = txc.CancelRequest.
		UpdateOne(ccr).
		SetStatus(status).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	// cancel match
	_, err = r.client.Match.
		UpdateOneID(sc.MatchID).
		SetStatus(mStatus).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	return tx.Commit()
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
