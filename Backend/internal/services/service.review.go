package services

import (
	"fmt"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceReview interface {
	ReviewCourseService(r *schemas.SchemaCreateReview) (*schemas.SchemaCreateReviewResponse, error)
}

type serviceReview struct {
	repository repository.RepositoryReview
}

func NewServiceReview(r repository.RepositoryReview) *serviceReview {
	return &serviceReview{
		repository: r,
	}
}

func (s serviceReview) ReviewCourseService(r *schemas.SchemaCreateReview) (*schemas.SchemaCreateReviewResponse, error) {
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
		Rating:        *review.Score,
		ReviewMessage: *review.ReviewMsg,
		ReviewTime:    review.ReviewTimeAt,
	}, nil
}
