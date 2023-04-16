package repositorys

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	appointment "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
)

type RepositoryAdminTuitionFee interface {
	GetAdminTuitionFees() ([]*ent.Appointment, error)
}

type repositoryAdminTuitionFee struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryAdminTuitionFee(c *ent.Client) *repositoryAdminTuitionFee {
	return &repositoryAdminTuitionFee{client: c, ctx: context.Background()}
}

func (r *repositoryAdminTuitionFee) GetAdminTuitionFees() ([]*ent.Appointment, error) {
	adminTuitionFees, err := r.client.Appointment.
		Query().
		WithMatch(func(q *ent.MatchQuery) {
			q.WithStudent(func(r *ent.StudentQuery) {
				r.WithUser()
			})
			q.WithCourse(func(r *ent.CourseQuery) {
				r.WithTutor(func(s *ent.TutorQuery) {
					s.WithUser()
				})
			})
		}).
		Where(appointment.StatusEQ(appointment.StatusPending)).
		Order(ent.Asc(appointment.FieldEndAt)).
		All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query appointments: %w", err)
	}

	return adminTuitionFees, nil
}
