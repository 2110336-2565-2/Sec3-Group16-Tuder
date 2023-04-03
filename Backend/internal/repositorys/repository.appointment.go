package repositorys

import (
	"context"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entMatch "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	entStudent "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryAppointment interface {
	GetAppointmentByStudentID(sr *schemas.SchemaGetAppointment) ([]*schemas.SchemaAppointment, error)
}

type repositoryAppointment struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryAppointment(c *ent.Client) *repositoryAppointment {
	return &repositoryAppointment{client: c, ctx: context.Background()}
}

func (r *repositoryAppointment) GetAppointmentByStudentID(sr *schemas.SchemaGetAppointment) ([]*schemas.SchemaAppointment, error) {
	matches, err := r.client.Match.
		Query().
		Where(entMatch.HasStudentWith(entStudent.IDEQ(sr.ID))).
		WithAppointment().
		WithCourse().
		All(r.ctx)

	if err != nil {
		return nil, err
	}

	schemaAppointments := make([]*schemas.SchemaAppointment, 0)
	for _, match := range matches {
		schemaAppointments = append(schemaAppointments, &schemas.SchemaAppointment{
			ID:          match.ID,
			Appointment: match.Edges.Appointment,
			Course_name: match.Edges.Course.Title,
			Tutor_name:  match.Edges.Course.Edges.Tutor.Edges.User.Username,
		})
	}

	return schemaAppointments, nil
}
