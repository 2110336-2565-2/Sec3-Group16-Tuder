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
	SearchAll() ([]*ent.Course, error)
	checkTimeAvailable(d [24]bool) bool
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

func (r repositoryCourseSearch) checkTimeAvailable(d [24]bool) bool {
	for _, b := range d {
		if b {
			return true
		}
	}
	return false
}

func (r repositoryCourseSearch) SearchByDayRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	//implement here
	days := sr.Days
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithSchedule().WithUser()
		}).
		All(r.ctx)

	if err != nil {
		return nil, errors.New("Course not found")
	}
	var matched_courses []*ent.Course
	for _, course := range courses {
		state := true
		if days[0] {
			if r.checkTimeAvailable(course.Edges.Tutor.Edges.Schedule.Day0) == false {
				state = false
			}
		}
		if days[1] {
			if r.checkTimeAvailable(course.Edges.Tutor.Edges.Schedule.Day1) == false {
				state = false
			}
		}
		if days[2] {
			if r.checkTimeAvailable(course.Edges.Tutor.Edges.Schedule.Day2) == false {
				state = false
			}
		}
		if days[3] {
			if r.checkTimeAvailable(course.Edges.Tutor.Edges.Schedule.Day3) == false {
				state = false
			}
		}
		if days[4] {
			if r.checkTimeAvailable(course.Edges.Tutor.Edges.Schedule.Day4) == false {
				state = false
			}
		}
		if days[5] {
			if r.checkTimeAvailable(course.Edges.Tutor.Edges.Schedule.Day5) == false {
				state = false
			}
		}
		if days[6] {
			if r.checkTimeAvailable(course.Edges.Tutor.Edges.Schedule.Day6) == false {
				state = false
			}
		}
		if state {
			matched_courses = append(matched_courses, course)
		}
	}
	return matched_courses, nil
}

// ---------------------------------------------------------------------------------------------
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
	// fmt.Println(courses[0].Edges.Tutor.Edges.User.Username)

	return courses, nil
}
