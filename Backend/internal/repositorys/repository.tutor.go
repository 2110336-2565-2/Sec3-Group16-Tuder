package repositorys

import (
	"context"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryTutor interface {
	GetTutors() ([]*ent.Tutor, error)
	GetTutorByUsername(sr *schema.SchemaGetTutor) (*ent.Tutor, error)
	CreateTutor(sr *schema.SchemaCreateTutor) (*ent.Tutor, error)
	UpdateTutor(sr *schema.SchemaUpdateTutor) (*ent.Tutor, error)
	DeleteTutor(sr *schema.SchemaDeleteTutor) error
}

type repositoryTutor struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryTutor(c *ent.Client) *repositoryTutor {
	return &repositoryTutor{client: c, ctx: context.Background()}
}

func (r *repositoryTutor) GetTutorByUsername(sr *schema.SchemaGetTutor) (*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		Query().
		Where(entTutor.HasUserWith(entUser.UsernameEQ(sr.Username))).
		WithUser().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) GetTutors() ([]*ent.Tutor, error) {
	tutor, err := r.client.Tutor.
		Query().
		WithUser().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return tutor, nil
}

func (r *repositoryTutor) CreateTutor(sr *schema.SchemaCreateTutor) (*ent.Tutor, error) {

	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	newUser, err := txc.User.Create().
		SetUsername(sr.Username).
		SetPassword(sr.Password).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetProfilePictureURL(sr.ProfilePictureURL).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	tutor, err := txc.Tutor.
		Create().
		SetUserID(newUser.ID).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		SetCitizenID(sr.CitizenId).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	return tutor, tx.Commit()
}

func (r *repositoryTutor) UpdateTutor(sr *schema.SchemaUpdateTutor) (*ent.Tutor, error) {

	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	user, err := txc.User.Query().
		Where(entUser.UsernameEQ(sr.Username)).
		WithTutor().
		Only(r.ctx)
	if err != nil {
		// rollback
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
	}
	tutor := user.Edges.Tutor

	user, err = txc.User.
		UpdateOne(user).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetBirthDate(sr.Birthdate).
		Save(r.ctx)
	//fmt.Print(user)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
	}

	tutor, err = txc.Tutor.
		UpdateOne(tutor).
		SetDescription(sr.Description).
		SetOmiseBankToken(sr.OmiseBankToken).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	// reload tutor->user
	tutor.Edges.User, err = tutor.QueryUser().Only(r.ctx)

	return tutor, tx.Commit()
}

func (r *repositoryTutor) DeleteTutor(sr *schema.SchemaDeleteTutor) error {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	err = txc.Tutor.
		DeleteOneID(sr.ID).
		Exec(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	err = txc.User.DeleteOneID(sr.ID).Exec(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	return nil
}
