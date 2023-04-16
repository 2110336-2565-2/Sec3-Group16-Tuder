package repositorys

import (
	"context"
	"fmt"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/review"
	entStudent "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/google/uuid"
)

type RepositoryReview interface {
	CreateReview(review *schema.SchemaCreateReview) (*ent.Review, error)
	GetCourseWithReviewsById(courseId uuid.UUID) (*ent.Course, error)
}

type repositoryReview struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryReview(c *ent.Client) RepositoryReview {
	return &repositoryReview{
		client: c,
		ctx:    context.Background(),
	}
}

func (r repositoryReview) CreateReview(sR *schema.SchemaCreateReview) (*ent.Review, error) {

	//create translation
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start a transaction: %v", err)
	}

	// wrap client
	tcx := tx.Client()

	// begin sql query
	// 1. checking if a student has a match that mapped to the course
	student, err := tcx.Student.Query().
		Where(entStudent.IDEQ(sR.StudentId)).
		WithMatch(func(query *ent.MatchQuery) {
			query.Where(match.HasCourseWith(course.IDEQ(sR.CourseId))).
				WithCourse().
				Only(r.ctx)
		}).Only(r.ctx)
	if len(student.Edges.Match) <= 0 {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to rollback transaction: %w", err)
		}
		return nil, fmt.Errorf("the student has not enrolled in this course")
	}

	// if all criterion passed, create a review
	c := student.Edges.Match[0].Edges.Course
	review, err := tcx.Review.Create().
		SetReviewMsg(sR.ReviewMessage).
		SetScore(int8(sR.Rating)).
		SetReviewTimeAt(time.Now()).
		AddCourse(c).
		AddStudent(student).
		Save(r.ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("failed to rollback transaction: %w", err)
		}
		return nil, fmt.Errorf("failed to create review: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return review, nil
}

func (r repositoryReview) GetReviews(courseId uuid.UUID) ([]*ent.Review, error) {
	reviews, err := r.client.Review.Query().Where(
		review.HasCourseWith(course.IDEQ(courseId)),
	).All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve reviews for course %s: %v", courseId.String(), err)
	}
	return reviews, nil
}

func (r repositoryReview) GetCourseWithReviewsById(courseId uuid.UUID) (*ent.Course, error) {
	c, err := r.client.Course.Query().
		Where(course.IDEQ(courseId)).
		WithReview().
		Only(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve course with reviews for course id %s: %v", courseId.String(), err)
	}
	return c, nil
}
