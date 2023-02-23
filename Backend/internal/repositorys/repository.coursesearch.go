package repositorys

import (
	"context"
	"errors"

	//"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	//"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	entCourse "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryCourseSearch interface {
	SearchByTitleRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchBySucjectRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByTopicRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByTutorRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByDayRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchAll(sr *schema.CourseSearch) ([]*ent.Course, error)
}

type repositoryCourseSearch struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryCourseSearch(c *ent.Client) *repositoryCourseSearch {
	return &repositoryCourseSearch{client: c, ctx: context.Background()}
}

//	func (r repositoryCourseSearch) CourseSearchRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
//		return nil, nil
//	}
//
// ---------------------------------------------------------------------------------------------
func (r repositoryCourseSearch) SearchByTitleRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		Where(entCourse.TitleContains(sr.Title)).
		All(r.ctx)

	if err != nil {
		return nil, errors.New("Course not found")
	}
	return courses, nil
}

// ---------------------------------------------------------------------------------------------
func (r repositoryCourseSearch) SearchBySucjectRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		Where(entCourse.SubjectContains(sr.Subject)).
		All(r.ctx)

	if err != nil {
		return nil, errors.New("Course not found")
	}
	return courses, nil
}

//---------------------------------------------------------------------------------------------

func (r repositoryCourseSearch) SearchByTopicRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		Where(entCourse.TopicContains(sr.Topic)).
		All(r.ctx)

	if err != nil {
		return nil, errors.New("Course not found")
	}
	return courses, nil
}

// ---------------------------------------------------------------------------------------------
func (r repositoryCourseSearch) SearchByTutorRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		Where(entCourse.HasTutorWith(entTutor.HasUserWith(entUser.UsernameContains(sr.Tutor_name)))).
		All(r.ctx)

	if err != nil {
		return nil, errors.New("Course not found")
	}
	return courses, nil
}

// ---------------------------------------------------------------------------------------------
func (r repositoryCourseSearch) SearchByDayRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	//implement here
	return nil, nil
}

// ---------------------------------------------------------------------------------------------
func (r repositoryCourseSearch) SearchAll(sr *schema.CourseSearch) ([]*ent.Course, error) {

	return nil, nil
}
