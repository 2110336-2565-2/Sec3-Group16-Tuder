package repositorys

import (
	"context"
	"errors"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
)

var (
	// DaytimeAvailable: available 8:00-12:00 , 13:00-17:00
	DaytimeAvailable = [24]bool{false, false, false, false, false, false,
		false, false, true, true, true, true,
		false, true, true, true, true, false,
		false, false, false, false, false, false}
	// NighttimeAvailable: available 18:00-21:00
	NighttimeAvailable = [24]bool{false, false, false, false, false, false,
		false, false, false, false, false, false,
		false, false, false, false, false, false,
		true, true, true, false, false, false}
	// AllDayAvailable: available 8:00-12:00 , 13:00-17:00, 18:00-21:00
	AllDayAvailable = [24]bool{false, false, false, false, false, false,
		false, false, true, true, true, true,
		false, true, true, true, true, false,
		true, true, true, false, false, false}
)

type RepositoryRegister interface {
	RegisterUser(sr *schemas.SchemaRegister) (*ent.User, error)
}

type repositoryRegister struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryRegister(c *ent.Client) *repositoryRegister {
	return &repositoryRegister{client: c, ctx: context.Background()}
}

func (r repositoryRegister) RegisterUser(sr *schemas.SchemaRegister) (*ent.User, error) {
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
	// hash password
	hashedPassword, _ := utils.HashPassword(sr.Password)

	// create user
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
		SetRole(entUser.Role(sr.Role)).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	//create a specific role
	switch sr.Role {
	case "student":
		_, err = txc.Student.
			Create().SetUserID(newUser.ID).
			Save(r.ctx)
	case "tutor":
		// create a schedule with default value: available all day, everyday
		var schedule *ent.Schedule
		schedule, err = txc.Schedule.
			Create().
			SetID(uuid.New()).
			SetDay0(AllDayAvailable).
			SetDay1(AllDayAvailable).
			SetDay2(AllDayAvailable).
			SetDay3(AllDayAvailable).
			SetDay4(AllDayAvailable).
			SetDay5(AllDayAvailable).
			SetDay6(AllDayAvailable).
			Save(r.ctx)

		_, err = txc.Tutor.
			Create().
			SetUser(newUser).
			SetDescription(sr.Description).
			SetOmiseBankToken(sr.OmiseBankToken).
			SetCitizenID(sr.CitizenID).
			SetSchedule(schedule).
			Save(r.ctx)
	default:
		err = fmt.Errorf("User Role is invalid : %s", sr.Role)
	}
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return newUser, tx.Commit()
}
