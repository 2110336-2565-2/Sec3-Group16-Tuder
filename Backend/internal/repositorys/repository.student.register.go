package repositorys

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
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

	_, err := r.client.Student.
		Query().Where(sql.FieldEQ("username", sr.Username)).Only(r.ctx)

	if err == nil {
		return nil, errors.New("the username has been used")
	}
	hashedPassword, _ := utils.HashPassword(sr.Password)

	newUser, err := r.client.User.Create().
		SetUsername(sr.Username).
		SetPassword(hashedPassword).
		SetEmail(sr.Email).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetAddress(sr.Address).
		SetPhone(sr.Phone).
		SetBirthDate(sr.Birthdate).
		SetGender(sr.Gender).
		SetRole(user.RoleStudent).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	
	newStudent, err := r.client.Student.
		Create().SetUserID( newUser.ID).
		Save(r.ctx)
		
	if err != nil {
		return nil, err
	}

	return newStudent, nil
}
