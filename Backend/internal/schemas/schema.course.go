package schemas

import "github.com/google/uuid"

type CourseSearchResult struct {
	Course_id          uuid.UUID `json:"course_id"`
	Tutor_name         string    `json:"tutor_name"`
	Title              string    `json:"title"`
	Subject            string    `json:"subject"`
	Topic              string    `json:"topic"`
	Estimate_time      int       `json:"estimate_time"`
	Price_per_hour     int       `json:"price_per_hour"`
	Course_picture_url string    `json:"course_picture_url"`
}

type CourseSearch struct {
	Title      string  `json:"title"`
	Subject    string  `json:"subject"`
	Topic      string  `json:"topic"`
	Tutor_name string  `json:"tutor_name"`
	Days       [7]bool `json:"days"`
}
