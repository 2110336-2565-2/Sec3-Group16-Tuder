package repositorys

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryLogin interface {
	LoginRepository(input *schema.SchemaLogin) (*ent.User, error)
}

type repositoryLogin struct {
	client *ent.Client
	ctx context.Context
}

func NewRepositoryLogin(client *ent.Client) *repositoryLogin {
	return &repositoryLogin{client : client, ctx : context.Background()}
}


func (r *repositoryLogin) LoginRepository(input *schema.SchemaLogin) (*ent.User, error) {
	/*
	1 )  ต้องดึง user จาก database ที่มี username ตรงกัน
	*/

	
	user, err := r.client.User.
	Query().
	Where( sql.FieldEQ("username", input.Username) ).Only(r.ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}