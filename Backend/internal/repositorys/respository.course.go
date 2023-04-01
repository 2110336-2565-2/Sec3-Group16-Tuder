package repositorys

import (
	"context"
	"errors"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entCourse "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryCourse interface {
	GetCourses() ([]*ent.Course, error)
	GetCourseByID(sr *schema.SchemaGetCourse) (*ent.Course, error)
	CreateCourse(sr *schema.SchemaCreateCourse) (*ent.Course, error)
	UpdateCourse(sr *schema.SchemaUpdateCourse) (*ent.Course, error)
	DeleteCourse(sr *schema.SchemaDeleteCourse) error
	SearchByTitleRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchBySucjectRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByTopicRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByTutorRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	SearchByDayRepository(sr *schema.CourseSearch) ([]*ent.Course, error)
	checkTimeAvailable(d [24]bool) bool
}

type repositoryCourse struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryCourse(c *ent.Client) *repositoryCourse {
	return &repositoryCourse{client: c, ctx: context.Background()}
}

func (r *repositoryCourse) GetCourses() ([]*ent.Course, error) {
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

func (r repositoryCourse) SearchByTitleRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
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

func (r repositoryCourse) checkTimeAvailable(d [24]bool) bool {
	for _, b := range d {
		if b {
			return true
		}
	}
	return false
}

func (r repositoryCourse) SearchByDayRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
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

func (r repositoryCourse) SearchByTutorRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
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

func (r repositoryCourse) SearchByTopicRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
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

func (r repositoryCourse) SearchBySucjectRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
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

func (r *repositoryCourse) GetCourseByID(sr *schema.SchemaGetCourse) (*ent.Course, error) {
	course, err := r.client.Course.
		Query().
		Where(entCourse.IDEQ(sr.ID)).
		WithTutor().
		Only(r.ctx)

	if err != nil {
		return nil, err
	}
	return course, nil
}

func (r *repositoryCourse) CreateCourse(sr *schema.SchemaCreateCourse) (*ent.Course, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()
	tutor, err := txc.Tutor.Query().Where(entTutor.IDEQ(sr.Tutor_id)).Only(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("getting tutor: %w", err)
	}
	// create a new course
	course, err := txc.Course.Create().
		SetCoursePictureURL(sr.Course_picture_url).
		SetSubject(sr.Subject).
		SetDescription(sr.Description).
		SetPricePerHour(sr.Price_per_hour).
		SetTutorID(tutor.ID).
		SetLevel(entCourse.Level(sr.Level)).
		SetTitle(sr.Title).
		SetTopic(sr.Topic).
		SetEstimatedTime(sr.Estimate_time).
		Save(r.ctx)
	course.Edges.Tutor = tutor

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("creating course: %w", err)
	}
	return course, tx.Commit()
}

func (r *repositoryCourse) UpdateCourse(sr *schema.SchemaUpdateCourse) (*ent.Course, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()
	tutor, err := txc.Tutor.Query().Where(entTutor.IDEQ(sr.Tutor_id)).Only(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("getting tutor: %w", err)
	}

	// update a course
	course, err := txc.Course.UpdateOneID(sr.ID).
		SetCoursePictureURL(sr.Course_picture_url).
		SetSubject(sr.Subject).
		SetDescription(sr.Description).
		SetPricePerHour(sr.Price_per_hour).
		SetTutorID(sr.Tutor_id).
		SetLevel(entCourse.Level(sr.Level)).
		SetTitle(sr.Title).
		SetTopic(sr.Topic).
		SetEstimatedTime(sr.Estimate_time).
		Save(r.ctx)
	course.Edges.Tutor = tutor
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("updating course: %w", err)
	}
	return course, tx.Commit()
}

func (r *repositoryCourse) DeleteCourse(sr *schema.SchemaDeleteCourse) error {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	// delete a course
	err = txc.Course.DeleteOneID(sr.ID).Exec(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return fmt.Errorf("deleting course: %w", err)
	}
	return tx.Commit()
}
