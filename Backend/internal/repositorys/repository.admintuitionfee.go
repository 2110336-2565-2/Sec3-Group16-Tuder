package repositorys

import (
	"context"
	"errors"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
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
		All(r.ctx)
	if err != nil {
		return nil, errors.New("Issue report not found")
	}

	return adminTuitionFees, nil
}
