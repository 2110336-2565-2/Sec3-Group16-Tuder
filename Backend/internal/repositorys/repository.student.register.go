package repositorys

import (
	"context"
	"errors"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entUser"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
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

	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	
	// wrap the client with the transaction
	txc := tx.Client()


	_, err = txc.User.
		Query().
		Where(entUser.UsernameEQ(sr.Username)).
		Only(r.ctx)


	if err == nil {
		return nil, errors.New("the username has been used")
	}
	hashedPassword, _ := utils.HashPassword(sr.Password)

	newUser, err := txc.User.Create().
		SetUsername(sr.Username).
		SetPassword(hashedPassword).
		SetEmail(sr.Email).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetAddress(sr.Address).
		SetPhone(sr.Phone).
		SetBirthDate(sr.Birthdate).
		SetGender(sr.Gender).
		SetRole(entUser.RoleStudent).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	
	newStudent, err := txc.Student.
		Create().SetUserID( newUser.ID).
		Save(r.ctx)
		
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return newStudent, tx.Commit()
}
