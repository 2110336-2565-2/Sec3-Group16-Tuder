package repositorys

import (
	"context"

	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entCourse "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	entUser "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
	"github.com/google/uuid"
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
	//checkTimeAvailable(d [24]bool) bool
	UpdateCourseStatus(sr *schema.SchemaUpdateCourseStatus) (*ent.Course, error)
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
		Where(entCourse.StatusEQ("open")).
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("error getting courses: %w", err)
	}

	return courses, nil
}

func (r repositoryCourse) SearchByTitleRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		Where(
			entCourse.TitleContains(sr.Title),
			entCourse.StatusEQ("open")).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("error searching courses: %w", err)
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
		Where(entCourse.StatusEQ("open")).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get courses: %w", err)
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
		Where(entCourse.StatusEQ("open"), entCourse.HasTutorWith(entTutor.HasUserWith(entUser.UsernameContains(sr.Tutor_name)))).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search courses by tutor: %w", err)
	}
	return courses, nil
}

func (r repositoryCourse) SearchByTopicRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		Where(entCourse.StatusEQ("open"), entCourse.TopicContains(sr.Topic)).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search courses by tutor: %w", err)
	}
	return courses, nil
}

func (r repositoryCourse) SearchBySucjectRepository(sr *schema.CourseSearch) ([]*ent.Course, error) {
	courses, err := r.client.Course.
		Query().
		WithTutor(func(q *ent.TutorQuery) {
			q.WithUser()
		}).
		Where(entCourse.StatusEQ("open"), entCourse.SubjectContains(sr.Subject)).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search courses by topic: %w", err)
	}
	return courses, nil
}

func (r *repositoryCourse) GetCourseByID(sr *schema.SchemaGetCourse) (*ent.Course, error) {
	course, err := r.client.Course.
		Query().
		Where(entCourse.IDEQ(sr.ID)).
		WithTutor(func(query *ent.TutorQuery) {
			query.Where().WithUser().Only(r.ctx)
		}).
		WithReview().
		Only(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get course by ID: %w", err)
	}
	return course, nil
}

func (r *repositoryCourse) CreateCourse(sr *schema.SchemaCreateCourse) (*ent.Course, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()
	tutor, err := txc.Tutor.Query().Where(entTutor.IDEQ(sr.Tutor_id)).Only(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to rollback transaction: %w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to get tutor: %w", err)
	}

	// create image url

	imgURL := fmt.Sprintf("%s/%s/%s", sr.Tutor_id, sr.Title, uuid.New())
	if sr.Course_picture != nil {
		imgURL, _ = utils.GenerateProfilePictureURL(sr.Course_picture, imgURL)
	}

	// create a new course
	course, err := txc.Course.Create().
		SetCoursePictureURL(imgURL).
		SetSubject(sr.Subject).
		SetDescription(sr.Description).
		SetPricePerHour(sr.Price_per_hour).
		SetTutorID(tutor.ID).
		SetLevel(entCourse.Level(sr.Level)).
		SetTitle(sr.Title).
		SetTopic(sr.Topic).
		SetEstimatedTime(sr.Estimate_time).
		Save(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to rollback transaction: %w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to create course: %w", err)
	}

	course.Edges.Tutor = tutor
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return course, nil
}

func (r *repositoryCourse) UpdateCourse(sr *schema.SchemaUpdateCourse) (*ent.Course, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()
	tutor, err := txc.Tutor.Query().Where(entTutor.IDEQ(sr.Tutor_id)).Only(r.ctx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to rollback transaction: %w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to update course: %w", err)
	}

	imgURL := fmt.Sprintf("%s/%s/%s", sr.Tutor_id, sr.Title, uuid.New())
	if sr.Course_picture != nil {
		imgURL, _ = utils.GenerateProfilePictureURL(sr.Course_picture, imgURL)
	}

	// update a course
	course, err := txc.Course.UpdateOneID(sr.ID).
		SetCoursePictureURL(imgURL).
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
			err = fmt.Errorf("failed to rollback transaction: %w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to create course: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return course, nil
}

func (r *repositoryCourse) DeleteCourse(sr *schema.SchemaDeleteCourse) error {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	// delete a course
	err = txc.Course.DeleteOneID(sr.ID).Exec(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to rollback transaction: %w: %v", err, rerr)
		}
		return fmt.Errorf("failed to delete course: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (r *repositoryCourse) UpdateCourseStatus(sr *schema.SchemaUpdateCourseStatus) (*ent.Course, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	status := sr.Status
	// update a course
	if status == "open" {
		status = "closed"
	} else {
		status = "open"
	}

	course, err := txc.Course.UpdateOneID(sr.ID).
		SetStatus(entCourse.Status(status)).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to rollback transaction: %w: %v", err, rerr)
		}
		return nil, fmt.Errorf("failed to update course status: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return course, nil
}
