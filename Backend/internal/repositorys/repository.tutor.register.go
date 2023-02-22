package repositorys

import (
	"context"
	"errors"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type RepositoryTutorRegister interface {
	RegisterTutor(sr *schema.SchemaRegister) (*ent.Tutor, error)
}

type repositoryTutorRegister struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryTutorRegister(c *ent.Client) *repositoryTutorRegister {
	return &repositoryTutorRegister{client: c, ctx: context.Background()}
}

func (r repositoryTutorRegister) RegisterTutor(sr *schema.SchemaRegister) (*ent.Tutor, error) {

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
		SetRole(entUser.RoleTutor).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	newTutor, err := txc.Tutor.
		Create().
		SetUser(newUser).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		SetCitizenID(sr.CitizenID).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return newTutor, tx.Commit()	
}
