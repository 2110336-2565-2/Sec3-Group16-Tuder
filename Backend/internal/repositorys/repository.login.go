package repositorys

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryLogin interface {
	LoginRepository(input *schema.SchemaLogin) (*ent.User, error)
}

type repositoryLogin struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryLogin(client *ent.Client) *repositoryLogin {
	return &repositoryLogin{client: client, ctx: context.Background()}
}

func (r *repositoryLogin) LoginRepository(input *schema.SchemaLogin) (*ent.User, error) {

	user, err := r.client.User.
		Query().
		Where(sql.FieldEQ("username", input.Username)).
		Only(r.ctx)

	if err != nil {
		return nil, errors.New("the username isn't match")
	}

	return user, nil
}
