package repositorys

import (
	"context"
	"errors"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryLogin interface {
	Login(l *schema.SchemaLogin) (*ent.User, error)
}

type repositoryLogin struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryLogin(c *ent.Client) *repositoryLogin {
	return &repositoryLogin{client: c, ctx: context.Background()}
}

func (r *repositoryLogin) Login(l *schema.SchemaLogin) (*ent.User, error) {

	user, err := r.client.User.
		Query().
		Where(entUser.UsernameEQ(l.Username)).
		Only(r.ctx)

	if err != nil {
		return nil, errors.New("the username isn't match")
	}

	return user, nil
}
