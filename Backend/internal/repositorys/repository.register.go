package repositorys

import (
	"context"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryRegister interface {
	RegisterRepository(sr *schema.SchemaRegister) ([]*ent.User, error)
}

type repositoryRegister struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryRegister(c *ent.Client) *repositoryRegister {
	return &repositoryRegister{client: c, ctx: context.Background()}
}

func (r repositoryRegister) RegisterRepository(sr *schema.SchemaRegister) ([]*ent.User, error) {
	panic("implement me")
}
