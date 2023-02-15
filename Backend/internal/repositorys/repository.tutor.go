package repositorys

import (
	"context"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryTutor interface {
	GetTutorsRepository() ([]*ent.Tutor, error)
	GetTutorByUsernameRepository(sr *schema.SchemaGetTutor) (*ent.Tutor, error)
	CreateTutorRepository(sr *schema.SchemaCreateTutor) (*ent.Tutor, error)
	UpdateTutorRepository(sr *schema.SchemaUpdateTutor) (*ent.Tutor, error)
	DeleteTutorRepository(sr *schema.SchemaDeleteTutor) error
}

type repositoryTutor struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryTutor(c *ent.Client) *repositoryTutor {
	return &repositoryTutor{client: c, ctx: context.Background()}
}

func (r *repositoryTutor) GetTutorByUsernameRepository(sr *schema.SchemaGetTutor) (*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		Query().
		Where(entTutor.UsernameEQ(sr.Username)).
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) GetTutorsRepository() ([]*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		Query().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) CreateTutorRepository(sr *schema.SchemaCreateTutor) (*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		Create().
		SetUsername(sr.Username).
		SetPassword(sr.Password).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetProfilePictureURL(sr.ProfilePictureURL).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		SetCitizenID(sr.CitizenId).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) UpdateTutorRepository(sr *schema.SchemaUpdateTutor) (*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		UpdateOneID(sr.ID).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetProfilePictureURL(sr.ProfilePictureURL).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) DeleteTutorRepository(sr *schema.SchemaDeleteTutor) error {
	err := r.client.Tutor.
		DeleteOneID(sr.ID).
		Exec(r.ctx)

	if err != nil {
		return err
	}
	return nil
}
