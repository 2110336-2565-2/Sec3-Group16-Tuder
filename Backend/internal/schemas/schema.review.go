package schemas

import (
	"github.com/google/uuid"
	"time"
)

type SchemaCreateReview struct {
	CourseId      uuid.UUID `json:"course_id"`
	Rating        int       `json:"rating"`
	ReviewMessage string    `json:"review_message"`
	StudentId     uuid.UUID `json:"student_id"`
}

type SchemaCreateReviewResponse struct {
	CourseId      uuid.UUID `json:"course_id,omitempty"`
	Rating        float32   `json:"rating,omitempty"`
	ReviewMessage string    `json:"review_message,omitempty"`
	ReviewTime    time.Time `json:"review_time,omitempty"`
}
