package repositorys

import (
	"context"
	"errors"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type RepositorySearch interface {
	SearchRepository(sr *schema.SchemaSearch) ([]*ent.Search, error)
}

type repositorySearch  struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryCourseSearch (c *ent.Client) *repositorySearch  {
	return &repositorySearch {client: c, ctx: context.Background()}
}

func (r repositorySearch ) SearchByTutorNameRepository (sr *schema.SchemaSearch) ([]*ent.CourseInfo, error) {
	Tutor, err := r.client.Course.
		Query().
		Where(entUser.UsernameEQ(([a-z]+) + sr.Username + ([a-z]+) ) && entUser.RoleEQ("Tutor") ).
		Only(r.ctx)

	if err != nil {
		return nil, errors.New("the TutorName isn't match")
	}


	return user, nil
}


