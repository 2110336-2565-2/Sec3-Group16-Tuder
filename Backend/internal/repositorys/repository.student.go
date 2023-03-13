package repositorys

import (
	"context"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entStudent "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type RepositoryStudent interface {
	GetStudents() ([]*ent.Student, error)
	GetStudentByUsername(sr *schema.SchemaGetStudent) (*ent.Student, error)
	CreateStudent(sr *schema.SchemaCreateStudent) (*ent.Student, error)
	UpdateStudent(sr *schema.SchemaUpdateStudent) (*ent.Student, error)
	DeleteStudent(sr *schema.SchemaDeleteStudent) error
}

type repositoryStudent struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryStudent(c *ent.Client) RepositoryStudent {
	return &repositoryStudent{client: c, ctx: context.Background()}
}

func (rS *repositoryStudent) GetStudents() ([]*ent.Student, error) {
	students, err := rS.client.Student.Query().WithUser().All(rS.ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("\nget students successful\n")
	return students, nil
}

func (rS *repositoryStudent) GetStudentByUsername(sr *schema.SchemaGetStudent) (*ent.Student, error) {
	student, err := rS.client.Student.Query().Where(entStudent.HasUserWith(entUser.UsernameEQ(sr.Username))).WithUser().Only(rS.ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (rS *repositoryStudent) CreateStudent(sr *schema.SchemaCreateStudent) (*ent.Student, error) {
	// create a transaction
	tx, err := rS.client.Tx(rS.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()
	// create a new user
	newUser, err := txc.User.Create().
		SetUsername(sr.Username).
		SetPassword(sr.Password).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetProfilePictureURL("").
		Save(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("rolling back transaction: %v", rerr)
		}
	}
	student, err := txc.Student.Create().
		SetUser(newUser).
		SetUserID(newUser.ID).
		Save(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("rolling back transaction: %v", rerr)
		}
		return nil, fmt.Errorf("creating a new student: %w", err)
	}

	return student, tx.Commit()

}

func (rS *repositoryStudent) UpdateStudent(sr *schema.SchemaUpdateStudent) (*ent.Student, error) {
	tx, err := rS.client.Tx(rS.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	txc := tx.Client()

	user, err := txc.User.Query().
		Where(entUser.Username(sr.Username)).
		WithStudent().
		Only(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("rolling back transaction: %v", rerr)
		}
		return nil, fmt.Errorf("querying a user: %w", err)
	}

	student := user.Edges.Student

	profilePictureURL, _ := utils.GenerateProfilePictureURL(sr.ProfilePicture, sr.Username)

	user, err = txc.User.UpdateOne(user).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetProfilePictureURL(profilePictureURL).
		Save(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("rolling back transaction: %v", rerr)
		}
		return nil, fmt.Errorf("updating a user: %w", err)
	}

	student.Edges.User = user
	return student, tx.Commit()
}

func (rS *repositoryStudent) DeleteStudent(sr *schema.SchemaDeleteStudent) error {
	tx, err := rS.client.Tx(rS.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	txc := tx.Client()

	err = txc.Student.
		DeleteOneID(sr.ID).
		Exec(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("starting a transaction %w", rerr)
		}
		return err
	}

	return nil

}
