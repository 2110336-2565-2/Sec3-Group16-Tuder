package repositorys

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	entCourse "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entTutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryCourse interface {
	GetCourses() ([]*ent.Course, error)
	GetCourseByID(sr *schema.SchemaGetCourse) (*ent.Course, error)
	CreateCourse(sr *schema.SchemaCreateCourse) (*ent.Course, error)
	UpdateCourse(sr *schema.SchemaUpdateCourse) (*ent.Course, error)
	DeleteCourse(sr *schema.SchemaDeleteCourse) error
}

type repositoryCourse struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryCourse(c *ent.Client) *repositoryCourse {
	return &repositoryCourse{client: c, ctx: context.Background()}
}

func (r *repositoryCourse) GetCourses() ([]*ent.Course, error) {
	course, err := r.client.Course.
		Query().
		WithTutor().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return course, nil
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
