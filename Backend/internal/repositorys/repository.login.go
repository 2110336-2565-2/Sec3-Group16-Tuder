package repositorys

import (
	"context"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	// models "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/models"
)

type RepositoryLogin interface {
	LoginRepository(input *schema.SchemaLogin) ([]*ent.User, error)
}

type repositoryLogin struct {
	client *ent.Client
	ctx context.Context
}

func NewRepositoryLogin(client *ent.Client) *repositoryLogin {
	return &repositoryLogin{client : client, ctx : context.Background()}
}


func (r *repositoryLogin) LoginRepository(input *schema.SchemaLogin) ([]*ent.User, error) {

	user, err := r.client.User.Query().All(r.ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}