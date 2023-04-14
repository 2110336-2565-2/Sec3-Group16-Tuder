package repositorys

import (
	"context"
	"fmt"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	entMatch "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	entStudent "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryAppointment interface {
	GetAppointmentByStudentID(sr *schemas.SchemaGetAppointmentByStudentID) ([]*schemas.SchemaAppointmentFromStudentID, error)
	UpdateAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error)
}

type repositoryAppointment struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryAppointment(c *ent.Client) *repositoryAppointment {
	return &repositoryAppointment{client: c, ctx: context.Background()}
}

func (r *repositoryAppointment) GetAppointmentByStudentID(sr *schemas.SchemaGetAppointmentByStudentID) ([]*schemas.SchemaAppointmentFromStudentID, error) {
	matches, err := r.client.Match.
		Query().
		Where(entMatch.HasStudentWith(entStudent.IDEQ(sr.StudentID))).
		WithAppointment().
		WithCourse(func(tq *ent.CourseQuery) {
			tq.WithTutor(func(tq *ent.TutorQuery) {
				tq.WithUser()
			})
		}).
		All(r.ctx)

	if err != nil {
		return nil, err
	}
	schemaAppointments := make([]*schemas.SchemaAppointmentFromStudentID, 0)
	for _, match := range matches {
		schemaAppointments = append(schemaAppointments, &schemas.SchemaAppointmentFromStudentID{
			ID:          match.ID,
			Appointment: match.Edges.Appointment,
			Course_name: match.Edges.Course.Title,
			Tutor_name:  match.Edges.Course.Edges.Tutor.Edges.User.Username,
		})
	}
	return schemaAppointments, nil
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
	fmt.Println("work at repo")
	return
}

// func AutoUpdateAppointmentStatus() error {
// 	// create a transaction
// 	tx, err := r.client.Tx(r.ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("starting a transaction: %w", err)
// 	}
// 	txc := tx.Client()

// 	appointment, err := txc.Appointment.UpdateOneID(sr.ID).
// 		SetStatus(appointment.Status(sr.Status)).
// 		Save(r.ctx)

// 	if err != nil {
// 		if rerr := tx.Rollback(); rerr != nil {
// 			err = fmt.Errorf("%w: %v", err, rerr)
// 		}
// 		return nil, fmt.Errorf("updating appointment status: %w", err)
// 	}
// 	return tx.Commit()
// }
