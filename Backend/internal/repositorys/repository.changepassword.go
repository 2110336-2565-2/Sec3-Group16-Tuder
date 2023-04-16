package repositorys

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	utils "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type RepositoryChangePassword interface {
	ChangePassword(sc *schemas.SchemaChangePassword) error
	CheckPassword(sc *schemas.SchemaCheckPassword) (*ent.User, error)
}

type repositoryChangePassword struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryChangePassword(c *ent.Client) RepositoryChangePassword {
	return &repositoryChangePassword{
		client: c,
		ctx:    context.Background(),
	}
}

func (r *repositoryChangePassword) ChangePassword(sc *schemas.SchemaChangePassword) error {
	user, err := r.client.User.
		Query().
		Where(user.UsernameEQ(sc.Username)).
		Only(r.ctx)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	// validate password

	passwordValidator := utils.PasswordValidator(user.Username, user.Email, user.FirstName, user.LastName)
	if err := passwordValidator.Validate(sc.Password); err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	hashedPassword, _ := utils.HashPassword(sc.Password)

	_, err = user.Update().
		SetPassword(hashedPassword).
		Save(r.ctx)
	if err != nil {
		return fmt.Errorf("error saving password: %w", err)
	}
	return nil
}

func (r *repositoryChangePassword) CheckPassword(sc *schemas.SchemaCheckPassword) (*ent.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.UsernameEQ(sc.Username)).
		Only(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}
