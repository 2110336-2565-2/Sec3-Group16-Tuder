package repositorys

import (
	"context"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryChangePassword interface {
	ChangePassword(sc *schemas.SchemaChangePassword) error
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
		Where(user.EmailEQ(sc.Email)).
		Only(r.ctx)
	if err != nil {
		return err
	}

	_, err = user.Update().
		SetPassword(sc.NewPassword).
		Save(r.ctx)
	if err != nil {
		return err
	}
	return nil
}
