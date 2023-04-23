package repositorys

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
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
		return nil, errors.New("the username has already been used")
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
		SetProfilePictureURL("https://se2-tuder.s3.us-west-1.amazonaws.com/ProfilePicture/Default.jpg"). // default profile picture
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("error rolling back transaction: %w", rerr)
		}
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	//create a specific role
	switch sr.Role {
	case "student":
		_, err = txc.Student.
			Create().SetUserID(newUser.ID).
			Save(r.ctx)
	case "tutor":
		// create a schedule with default value: available all day, everyday
		customerID, _ := r.CreateOmiseCustomer(&schema.SchemaCreateOmiseCustomer{Email: sr.Email, OmiseBankToken: sr.OmiseBankToken})
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
			SetOmiseCustomerID(customerID).
			SetCitizenID(sr.CitizenID).
			SetSchedule(schedule).
			Save(r.ctx)
	default:
		err = fmt.Errorf("User Role is invalid : %s", sr.Role)
	}
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("error rolling back transaction: %w", rerr)
		}
		return nil, fmt.Errorf("error creating role-specific user data: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newUser, nil
}

func (r *repositoryRegister) CreateOmiseCustomer(sr *schema.SchemaCreateOmiseCustomer) (string, error) {
	omisePublicKey := os.Getenv("OMISE_PUBLIC_KEY")
	omiseSecretKey := os.Getenv("OMISE_SECRET_KEY")

	client, err := omise.NewClient(omisePublicKey, omiseSecretKey)

	if err != nil {
		return "", fmt.Errorf("failed to create Omise client: %w", err)
	}

	// create a customer
	customer, createCustomer := &omise.Customer{}, &operations.CreateCustomer{
		Email: sr.Email,
		Card:  sr.OmiseBankToken,
	}

	if err := client.Do(customer, createCustomer); err != nil {
		return "", fmt.Errorf("failed to create Omise customer: %w", err)
	}

	return customer.ID, nil
}
