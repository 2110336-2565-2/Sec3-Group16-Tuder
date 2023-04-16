package schemas

import (
	"github.com/google/uuid"
)

type CourseSearchResult struct {
	Course_id          uuid.UUID `json:"course_id"`
	Tutor_name         string    `json:"tutor_name"`
	Title              string    `json:"title"`
	Subject            string    `json:"subject"`
	Topic              string    `json:"topic"`
	Estimate_time      int       `json:"estimate_time"`
	Price_per_hour     int       `json:"price_per_hour"`
	Course_picture_url string    `json:"course_picture_url"`
	CourseStatus       string    `json:"course_status"`
}

type CourseSearch struct {
	Title      string  `json:"title"`
	Subject    string  `json:"subject"`
	Topic      string  `json:"topic"`
	Tutor_name string  `json:"tutor_name"`
	Days       [7]bool `json:"days"`
}

type SchemaCourse struct {
	ID                 uuid.UUID `json:"id"`
	Title              string    `json:"title"`
	Rating             float32
	ReviewCount        int
	Reviews            []Review
	Tutor_id           uuid.UUID `json:"tutor_id"`
	TutorUsername      string
	TutorFirstname     string
	TutorLastname      string
	Subject            string `json:"subject"`
	Topic              string `json:"topic"`
	Level              string `json:"level"`
	Description        string `json:"description"`
	Estimate_time      int    `json:"estimate_time"`
	Price_per_hour     int    `json:"price_per_hour"`
	Course_picture_url string `json:"course_picture_url"`
	CourseStatus       string `json:"course_status"`
}

type SchemaGetCourse struct {
	ID uuid.UUID `json:"id"`
}

type SchemaCreateCourse struct {
	Tutor_id       uuid.UUID `json:"tutor_id"`
	Title          string    `json:"title"`
	Subject        string    `json:"subject"`
	Topic          string    `json:"topic"`
	Estimate_time  int       `json:"estimate_time"`
	Description    string    `json:"description"`
	Price_per_hour int       `json:"price_per_hour"`
	Level          string    `json:"level"`
	Course_picture []byte    `json:"course_picture"`
}

type SchemaUpdateCourse struct {
	ID             uuid.UUID `json:"id"`
	Tutor_id       uuid.UUID `json:"tutor_id"`
	Title          string    `json:"title"`
	Subject        string    `json:"subject"`
	Topic          string    `json:"topic"`
	Estimate_time  int       `json:"estimate_time"`
	Description    string    `json:"description"`
	Price_per_hour int       `json:"price_per_hour"`
	Level          string    `json:"level"`
	Course_picture []byte    `json:"course_picture"`
}

type SchemaDeleteCourse struct {
	ID uuid.UUID `json:"id"`
}

type CourseResponse struct {
	ID                 uuid.UUID `json:"id"`
	Title              string    `json:"title"`
	Subject            string    `json:"subject"`
	Topic              string    `json:"topic"`
	EstimatedTime      int       `json:"estimated_time"`
	Level              string    `json:"level"`
	Course_picture_url string    `json:"course_picture_url"`
	Status             string    `json:"status"`
	NumOfClass         int       `json:"num_of_class"`
}

type CourseStatus struct {
	Status string `json:"status"`
}

type SchemaUpdateCourseStatus struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type SchemaUpdateCourseStatusResponse struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Status string    `json:"status"`
}
