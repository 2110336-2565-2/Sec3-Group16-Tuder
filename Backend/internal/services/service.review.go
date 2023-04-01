package services

import (
	"fmt"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/google/uuid"
)

type ServiceReview interface {
	ReviewService(r *schemas.SchemaCreateReview) (*schemas.SchemaCreateReviewResponse, error)
	GetRating(courseId string) (*schemas.SchemaGetReviewsByCourseIdResponse, error)
}

type serviceReview struct {
	repository repository.RepositoryReview
}

func NewServiceReview(r repository.RepositoryReview) *serviceReview {
	return &serviceReview{
		repository: r,
	}
}

func (s serviceReview) ReviewService(r *schemas.SchemaCreateReview) (*schemas.SchemaCreateReviewResponse, error) {
	fmt.Println(r)
	// 1. check if the rating is valid (between 0, 5)
	if r.Rating < 0 || r.Rating > 5 {
		return nil, fmt.Errorf("invalid Rating")
	}
	review, err := s.repository.CreateReview(r)
	if err != nil {
		return nil, err
	}
	return &schemas.SchemaCreateReviewResponse{
		CourseId:      r.CourseId,
		Rating:        float32(*review.Score),
		ReviewMessage: *review.ReviewMsg,
		ReviewTime:    review.ReviewTimeAt,
	}, nil
}

func (s serviceReview) GetRating(courseId string) (*schemas.SchemaGetReviewsByCourseIdResponse, error) {
	// parse uuid from string
	id, err := uuid.Parse(courseId)
	if err != nil {
		return nil, err
	}
	// get course with reviews
	c, err := s.repository.GetCourseWithReviewsById(id)
	if err != nil {
		return nil, err
	}
	// calculate stats
	var rating float32
	var count int
	var accRating float32 = 0
	for _, r := range c.Edges.Review {
		accRating += *r.Score
	}
	// check if a course have never been rate yed
	if len(c.Edges.Review) <= 0 {
		count = 0
		rating = 0.0
		// prevent zero division
	} else if accRating == 0.0 {
		count = len(c.Edges.Review)
		rating = 0.0
		// start counting
	} else {
		count = len(c.Edges.Review)
		rating = accRating / float32(count)
	}

	// assign values to fields
	sG := &schemas.SchemaGetReviewsByCourseIdResponse{
		CourseTitle: c.Title,
		TotalScore:  rating,
		TotalReview: count,
		Reviews:     []schemas.Review{},
	}
	// append ent.Review to sG
	for _, r := range c.Edges.Review {
		sG.Reviews = append(sG.Reviews, schemas.Review{
			Score:        *r.Score,
			ReviewMsg:    *r.ReviewMsg,
			ReviewTimeAt: r.ReviewTimeAt,
		})
	}
	return sG, err
}
