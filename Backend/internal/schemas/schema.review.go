package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	Score        int8      `json:"score"`
	ReviewMsg    string    `json:"review_msg"`
	ReviewTimeAt time.Time `json:"review_time_at"`
}

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

type SchemaGetReviewsByCourseIdResponse struct {
	CourseTitle string   `json:"course_title"`
	TotalScore  float32  `json:"total_score"`
	TotalReview int      `json:"total_review"`
	Reviews     []Review `json:"reviews"`
}

type ReviewResponse struct {
	CourseTitle  string    `json:"course_title"`
	Score        int8      `json:"score"`
	ReviewMsg    string    `json:"review_msg"`
	ReviewTimeAt time.Time `json:"review_time_at"`
}
