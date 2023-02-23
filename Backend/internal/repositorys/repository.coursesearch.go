package repositorys

import (
	"context"
	"errors"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryCourseSearch interface {
	SearchByTitleRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchBySucjectRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByTopicRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByTutorRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByDayRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchAll() ([]*ent.Course, error)
}

type repositoryCourseSearch struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryCourseSearch(c *ent.Client) *repositoryCourseSearch {
	return &repositoryCourseSearch{client: c, ctx: context.Background()}
}

// func (r repositoryCourseSearch) CourseSearchRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
// 	return nil, nil
// }

func (r repositoryCourseSearch) SearchByTitleRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	//implement here
	return nil, nil
}

func (r repositoryCourseSearch) SearchBySucjectRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	//implement here
	return nil, nil
}

func (r repositoryCourseSearch) SearchByTopicRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	//implement here
	return nil, nil
}

func (r repositoryCourseSearch) SearchByTutorRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	// Tutor, err := r.client.Course.
	// 	Query().
	// 	Where(entUser.UsernameEQ(([a-z]+) + sr.Username + ([a-z]+) ) && entUser.RoleEQ("Tutor") ).
	// 	Only(r.ctx)

	// if err != nil {
	// 	return nil, errors.New("the TutorName isn't match")
	// }

	return nil, nil
}

func (r repositoryCourseSearch) SearchByDayRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	//implement here
	return nil, nil
}

func (r repositoryCourseSearch) SearchAll() ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		All(r.ctx)

	if err != nil {
		return nil, errors.New("Course not found")
	}

	return courses, nil
}
