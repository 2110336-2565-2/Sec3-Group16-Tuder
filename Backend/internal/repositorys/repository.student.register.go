package repositorys

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryStudentRegister interface {
	RegisterStudentRepository(sr *schema.SchemaRegister) (*ent.Student, error)
}

type repositoryStudentRegister struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryStudentRegister(c *ent.Client) *repositoryStudentRegister {
	return &repositoryStudentRegister{client: c, ctx: context.Background()}
}

func (r repositoryStudentRegister) RegisterStudentRepository(sr *schema.SchemaRegister) (*ent.Student, error) {
	// panic("implement me")
	_, err := r.client.Student.
		Query().Where(sql.FieldEQ("username", sr.Username)).Only(r.ctx)

	if err == nil {
		return nil, errors.New("the username has been used")
	}

	newStudent, err := r.client.Student.
		Create().
		SetUsername(sr.Username).
		SetPassword(sr.Password).
		SetEmail(sr.Email).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetAddress(sr.Address).
		SetPhone(sr.Phone).
		SetBirthDate(sr.Birthdate).
		SetGender(sr.Gender).
		Save(r.ctx)

	return newStudent, nil
}
