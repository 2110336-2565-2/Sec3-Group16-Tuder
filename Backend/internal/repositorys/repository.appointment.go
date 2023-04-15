package repositorys

import (
	"context"
	"fmt"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	entAppointment "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	entCourse "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entMatch "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	entStudent "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryAppointment interface {
	GetMatchByStudentID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error)
	GetMatchByTutorID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error)
	UpdateAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error)
	GetAppointmentByMatchID(sr *schemas.SchemaGetAppointmentByMatchID) (*schemas.SchemaAppointmentsFromMatchID, error)
}

type repositoryAppointment struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryAppointment(c *ent.Client) *repositoryAppointment {
	return &repositoryAppointment{client: c, ctx: context.Background()}
}

func (r *repositoryAppointment) GetMatchByStudentID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error) {
	matches, err := r.client.Match.
		Query().
		Where(entMatch.HasStudentWith(entStudent.IDEQ(sr.ID))).
		WithAppointment(func(tq *ent.AppointmentQuery) {
			tq.Order(ent.Asc(appointment.FieldBeginAt))
		}).
		WithCourse(func(tq *ent.CourseQuery) {
			tq.WithTutor(func(tq *ent.TutorQuery) {
				tq.WithUser()
			})
		}).
		All(r.ctx)

	if err != nil {
		return nil, err
	}
	schemaAppointments := make([]*schemas.SchemaMatchesFromID, 0)
	for _, match := range matches {
		remaining := 0
		var upcoming_class time.Time
		found_upcoming := false
		for _, app := range match.Edges.Appointment {
			if app.Status.String() == "comingsoon" {
				remaining = remaining + 1
				if !found_upcoming {
					upcoming_class = app.BeginAt
					found_upcoming = true
				}
			}
		}
		tutorName := match.Edges.Course.Edges.Tutor.Edges.User.FirstName + " " + match.Edges.Course.Edges.Tutor.Edges.User.LastName
		schemaAppointments = append(schemaAppointments, &schemas.SchemaMatchesFromID{
			MatchID:       match.ID,
			CourseName:    match.Edges.Course.Title,
			TutorName:     tutorName,
			UpcomingClass: upcoming_class,
			Remaining:     remaining,
			CoursePictureURL: *match.Edges.Course.CoursePictureURL,
			Status:  	match.Status.String(),
		})
	}
	return schemaAppointments, nil
}

func (r *repositoryAppointment) GetMatchByTutorID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error) {
	matches, err := r.client.Match.
		Query().
		Where(entMatch.HasCourseWith(entCourse.HasTutorWith(entTutor.IDEQ(sr.ID)))).
		WithAppointment( func(tq *ent.AppointmentQuery) {
			tq.Order(ent.Asc(appointment.FieldBeginAt))
		}).
		WithCourse(func(tq *ent.CourseQuery) {
			tq.WithTutor(func(tq *ent.TutorQuery) {
				tq.WithUser()
			})
		}).
		All(r.ctx)

	if err != nil {
		return nil, err
	}
	schemaAppointments := make([]*schemas.SchemaMatchesFromID, 0)
	for _, match := range matches {
		remaining := 0
		var upcoming_class time.Time
		found_upcoming := false
		for _, app := range match.Edges.Appointment {
			if app.Status.String() == "comingsoon" {
				remaining = remaining + 1
				if !found_upcoming {
					upcoming_class = app.BeginAt
					found_upcoming = true
				}
			}
		}
		tutorName := match.Edges.Course.Edges.Tutor.Edges.User.FirstName + " " + match.Edges.Course.Edges.Tutor.Edges.User.LastName
		schemaAppointments = append(schemaAppointments, &schemas.SchemaMatchesFromID{
			MatchID:       match.ID,
			CourseName:    match.Edges.Course.Title,
			TutorName:     tutorName,
			UpcomingClass: upcoming_class,
			Remaining:     remaining,
			CoursePictureURL: *match.Edges.Course.CoursePictureURL,
			Status:  	match.Status.String(),
		})
	}
	return schemaAppointments, nil
}

func (r *repositoryAppointment) GetAppointmentByMatchID(sr *schemas.SchemaGetAppointmentByMatchID) (*schemas.SchemaAppointmentsFromMatchID, error) {
	appointments, err := r.client.Appointment.
		Query().
		Where(entAppointment.HasMatchWith(entMatch.IDEQ(sr.MatchID))).
		WithMatch(func(q *ent.MatchQuery) {
			q.WithCourse(func(s *ent.CourseQuery) {
				s.WithTutor(func(t *ent.TutorQuery) {
					t.WithUser()
				})
			})
		}).
		Order(ent.Asc(appointment.FieldBeginAt)).
		All(r.ctx)

	if err != nil {
		return nil, err
	}

	if len(appointments) == 0 {
		return nil, nil
	}

	schemaApps := make([]*schemas.SchemaAppointment, 0)
	for _, app := range appointments {
		schemaApps = append(schemaApps, &schemas.SchemaAppointment{
			ID:      app.ID,
			BeginAt: app.BeginAt,
			EndAt:   app.EndAt,
			Status:  app.Status.String(),
		})
	}

	app := appointments[0]

	tutorName := app.Edges.Match.Edges.Course.Edges.Tutor.Edges.User.FirstName + " " + app.Edges.Match.Edges.Course.Edges.Tutor.Edges.User.LastName
	schemaAppointment := &schemas.SchemaAppointmentsFromMatchID{
		CourseName:        app.Edges.Match.Edges.Course.Title,
		TutorName:         tutorName,
		CourseDescription: app.Edges.Match.Edges.Course.Description,
		Level:             app.Edges.Match.Edges.Course.Level.String(),
		EstimeateTime:     app.Edges.Match.Edges.Course.EstimatedTime,
		CoursePictureURL:  *app.Edges.Match.Edges.Course.CoursePictureURL,
		Appointments:      schemaApps,
	}
	fmt.Println(*schemaAppointment)
	return schemaAppointment, nil
}

func (r *repositoryAppointment) UpdateAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	txc := tx.Client()

	appointment, err := txc.Appointment.UpdateOneID(sr.ID).
		SetStatus(appointment.Status(sr.Status)).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("updating appointment status: %w", err)
	}

	// find match id from appointment
	match, err := txc.Match.Query().Where(entMatch.HasAppointmentWith(entAppointment.IDEQ(sr.ID))).Only(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("finding match id from appointment: %w", err)
	}

	// find all appointments from match
	appointments, err := txc.Appointment.Query().Where(entAppointment.HasMatchWith(entMatch.IDEQ(match.ID))).All(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("finding all appointments from match: %w", err)
	}

	state := true
	for _, app := range appointments {
		if app.Status.String() != "completed" {
			state = false
			break
		}
	}

	if state {
		// update match status
		_, err = txc.Match.UpdateOneID(match.ID).SetStatus(entMatch.Status("completed")).Save(r.ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("updating match status: %w", err)
		}
	}



	return appointment, tx.Commit()
}

func (r *repositoryAppointment) AutoUpdateAppointmentStatus() {
	appointments := r.client.Appointment.Query().AllX(r.ctx)
	now := time.Now()
	for _, app := range appointments {
		if app.BeginAt.Truncate(time.Hour).Equal(now.Truncate(time.Hour)) {
			if app.Status == "comingsoon" {
				app.Update().SetStatus("ongoing").Save(r.ctx)
			}
		}
		if app.EndAt.Truncate(time.Hour).Equal(now.Truncate(time.Hour)) {
			if app.Status == "ongoing" {
				app.Update().SetStatus("verifying").Save(r.ctx)
			}
		}
		if app.EndAt.Truncate(time.Hour).Equal(now.Add(-24 * time.Hour).Truncate(time.Hour)) {
			if app.Status == "verifying" {
				app.Update().SetStatus("pending").Save(r.ctx)
			}
		}
	}
	return
}
