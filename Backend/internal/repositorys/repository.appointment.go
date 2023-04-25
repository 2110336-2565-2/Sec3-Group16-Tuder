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
	entSchedule "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
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
			MatchID:          match.ID,
			CourseName:       match.Edges.Course.Title,
			TutorName:        tutorName,
			UpcomingClass:    upcoming_class,
			Remaining:        remaining,
			CoursePictureURL: *match.Edges.Course.CoursePictureURL,
			Status:           match.Status.String(),
		})
	}
	return schemaAppointments, nil
}

func (r *repositoryAppointment) GetMatchByTutorID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error) {
	matches, err := r.client.Match.
		Query().
		Where(entMatch.HasCourseWith(entCourse.HasTutorWith(entTutor.IDEQ(sr.ID)))).
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
			MatchID:          match.ID,
			CourseName:       match.Edges.Course.Title,
			TutorName:        tutorName,
			UpcomingClass:    upcoming_class,
			Remaining:        remaining,
			CoursePictureURL: *match.Edges.Course.CoursePictureURL,
			Status:           match.Status.String(),
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
		Status:            app.Edges.Match.Status.String(),
	}
	return schemaAppointment, nil
}

func (r *repositoryAppointment) UpdateAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	txc := tx.Client()

	appointment, err := txc.Appointment.UpdateOneID(sr.ID).
		SetStatus(appointment.Status(sr.Status)).
		Save(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update appointment status: %w", err)
	}

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to update appointment status: %w", err)
	}

	// logic for generate new appointment if posponed
	if sr.Status == "postponed" {
		match, err := txc.Match.Query().Where(entMatch.HasAppointmentWith(entAppointment.IDEQ(sr.ID))).Only(r.ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("Can't match from appointment id for creating new appointment after postpone: %w", err)
		}
		last_appointment, err1 := txc.Appointment.Query().Where(entAppointment.HasMatchWith(entMatch.IDEQ(match.ID))).Order(ent.Desc(entAppointment.FieldBeginAt)).All(r.ctx)
		if err1 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err1, rerr)
			}
			return nil, fmt.Errorf("Can't find last appointment for creating new apoointment after postponed: %w", err1)
		}
		match_schedule, err2 := txc.Schedule.Query().Where(entSchedule.HasMatchWith(entMatch.IDEQ(match.ID))).Only(r.ctx)
		if err2 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err2, rerr)
			}
			return nil, fmt.Errorf("Can't find match schedule for creating new apoointment after postponed: %w", err2)
		}
		hour := last_appointment[0].BeginAt.Hour()
		day := int(last_appointment[0].BeginAt.Weekday())
		var schedule [7][24]bool
		schedule[0] = match_schedule.Day0
		schedule[1] = match_schedule.Day1
		schedule[2] = match_schedule.Day2
		schedule[3] = match_schedule.Day3
		schedule[4] = match_schedule.Day4
		schedule[6] = match_schedule.Day5
		schedule[6] = match_schedule.Day6
		var combined_schedule []bool
		for _, arr := range schedule {
			combined_schedule = append(combined_schedule, arr[:]...)
		}
		count_hour := 1
		start := (day * 24) + hour
		for i := 1; i <= 24*7; i++ {
			idx := (start + i) % (24 * 7)
			if combined_schedule[idx] {
				break
			} else {
				count_hour = count_hour + 1
			}
		}
		next_begin_date := last_appointment[0].BeginAt.Add(time.Duration(count_hour) * time.Hour)

		_, err3 := txc.Appointment.Create().SetBeginAt(next_begin_date).SetEndAt(next_begin_date.Add(1 * time.Hour)).SetStatus("comingsoon").SetMatch(match).Save(r.ctx)
		if err3 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err3, rerr)
			}
			return nil, fmt.Errorf("Unavailable to create new appointment after postepone : %w", err3)
		}
	}

	// find match id from appointment
	match, err := txc.Match.Query().Where(entMatch.HasAppointmentWith(entAppointment.IDEQ(sr.ID))).Only(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to find match id from appointment: %w", err)
	}

	// find all appointments from match
	appointments, err := txc.Appointment.Query().Where(entAppointment.HasMatchWith(entMatch.IDEQ(match.ID))).All(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to find all appointments from match: %w", err)
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
			return nil, fmt.Errorf("failed to updating match status: %w", err)
		}

		// return time available to tutor's schedule
		var new_schedule [7][24]bool
		tutor, err1 := txc.Tutor.Query().Where(entTutor.HasCourseWith(entCourse.HasMatchWith(entMatch.IDEQ(match.ID)))).Only(r.ctx)
		if err1 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err1 = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("Tutor not found: %w", err1)
		}
		match_schedule, err2 := txc.Schedule.Query().Where(entSchedule.HasMatchWith(entMatch.IDEQ(match.ID))).Only(r.ctx)
		if err2 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err2 = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("Match not found: %w", err2)
		}
		tutor_schedule, err3 := txc.Schedule.Query().Where(entSchedule.HasTutorWith(entTutor.IDEQ(tutor.ID))).Only(r.ctx)
		if err3 != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err3 = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("Tutor's schedule not found: %w", err3)
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
				err4 = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("updating tutor's schedule: %w", err4)
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
