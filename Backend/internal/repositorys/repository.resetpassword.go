package repositorys

import (
	"context"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"

	"github.com/google/uuid"
)

type RepositoryResetPassword interface {
	FindUserByEmail(sr *schema.SchemaResetPassword) (*ent.User, error)
	FindUserByID(id uuid.UUID) (*ent.User, error)
	UpdateUserPassword(user *ent.User, password string) (*ent.User, error)
}

type repositoryResetPassword struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryResetPassword(c *ent.Client) *repositoryResetPassword {
	return &repositoryResetPassword{
		client: c,
		ctx:    context.Background(),
	}
}

func (r *repositoryResetPassword) FindUserByEmail(sr *schema.SchemaResetPassword) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.EmailEQ(sr.Email)).
		Only(r.ctx)
}

func (r *repositoryResetPassword) FindUserByID(id uuid.UUID) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(r.ctx)
}

func (r *repositoryResetPassword) UpdateUserPassword(user *ent.User, password string) (*ent.User, error) {
	return user.Update().
		SetPassword(password).
		Save(r.ctx)
}
