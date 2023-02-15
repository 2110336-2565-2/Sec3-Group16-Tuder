package repositorys

import (
	"context"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryResetPassword interface {
	ResetPasswordRepository(sr *schema.SchemaResetPassword) ([]*ent.User, error)
}

type repositoryResetPassword struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryResetPassword(c *ent.Client) *repositoryResetPassword {
	return &repositoryResetPassword{client: c, ctx: context.Background()}
}

func (r repositoryResetPassword) ResetPasswordRepository(sr *schema.SchemaResetPassword) ([]*ent.User, error) {
	panic("implement me")
}
