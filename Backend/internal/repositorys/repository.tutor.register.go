package repositorys

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type RepositoryTutorRegister interface {
	RegisterTutorRepository(sr *schema.SchemaRegister) (*ent.Tutor, error)
}

type repositoryTutorRegister struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryTutorRegister(c *ent.Client) *repositoryTutorRegister {
	return &repositoryTutorRegister{client: c, ctx: context.Background()}
}

func (r repositoryTutorRegister) RegisterTutorRepository(sr *schema.SchemaRegister) (*ent.Tutor, error) {
	_, err := r.client.Tutor.
		Query().Where(sql.FieldEQ("username", sr.Username)).Only(r.ctx)

	if err == nil {
		return nil, errors.New("The Username has been used")
	}
	hashedPassword, _ := utils.HashPassword(sr.Password)

	newTutor, err := r.client.Tutor.
		Create().
		SetUsername(sr.Username).
		SetPassword(hashedPassword).
		SetEmail(sr.Email).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetAddress(sr.Address).
		SetPhone(sr.Phone).
		SetBirthDate(sr.Birthdate).
		SetGender(sr.Gender).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		SetCitizenID(sr.CitizenID).
		Save(r.ctx)

	return newTutor, nil
}
