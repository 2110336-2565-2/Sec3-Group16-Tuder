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
		return nil, fmt.Errorf("failed to get students: %v", err)
	}
	return students, nil
}

func (rS *repositoryStudent) GetStudentByUsername(sr *schema.SchemaGetStudent) (*ent.Student, error) {
	student, err := rS.client.Student.Query().Where(entStudent.HasUserWith(entUser.UsernameEQ(sr.Username))).WithUser().Only(rS.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get student by username: %v", err)
	}
	return student, nil
}

func (rS *repositoryStudent) CreateStudent(sr *schema.SchemaCreateStudent) (*ent.Student, error) {
	// create a transaction
	tx, err := rS.client.Tx(rS.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start a transaction: %v", err)
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
			return nil, fmt.Errorf("failed to rollback transaction: %v", rerr)
		}
		return nil, fmt.Errorf("failed to create new user: %v", err)
	}
	student, err := txc.Student.Create().
		SetUser(newUser).
		SetUserID(newUser.ID).
		Save(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("failed to rollback transaction: %v", rerr)
		}
		return nil, fmt.Errorf("failed to create new student: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return student, nil
}

func (rS *repositoryStudent) UpdateStudent(sr *schema.SchemaUpdateStudent) (*ent.Student, error) {
	tx, err := rS.client.Tx(rS.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	txc := tx.Client()

	user, err := txc.User.Query().
		Where(entUser.Username(sr.Username)).
		WithStudent().
		Only(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("failed to rollback transaction: %v", rerr)
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	student := user.Edges.Student

	profilePictureURL := *user.ProfilePictureURL
	if sr.ProfilePicture != nil {
		profilePictureURL, _ = utils.GenerateProfilePictureURL(sr.ProfilePicture, sr.Username, "ProfilePicture")

	}

	user, err = txc.User.UpdateOne(user).
		SetFirstName(sr.Firstname).
		SetLastName(sr.Lastname).
		SetEmail(sr.Email).
		SetPhone(sr.Phone).
		SetAddress(sr.Address).
		SetBirthDate(sr.Birthdate).
		SetGender(sr.Gender).
		SetProfilePictureURL(profilePictureURL).
		Save(rS.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return nil, fmt.Errorf("failed to rollback transaction: %v", rerr)
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	student.Edges.User = user
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return student, nil
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
			return fmt.Errorf("failed to rollback transaction: %v", rerr)
		}
		return fmt.Errorf("failed to delete student: %w", err)
	}
	return nil

}
