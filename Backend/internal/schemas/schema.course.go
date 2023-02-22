package schemas

import "github.com/google/uuid"

type CourseSearchResult struct {
	Course_id          uuid.UUID `json:"Course_id"`
	Tutor_name         string    `json:"Tutor_name"`
	Tittle             string    `json:"Tittle"`
	Subject            string    `json:"Subject"`
	Topic              string    `json:"Topic"`
	Estimate_time      int       `json:"Estimate_time"`
	Price_per_hour     int       `json:"Price_per_hour"`
	Course_picture_url string    `json:"Course_picture_url"`
}

type CourseSearch struct {
	Title      string  `json:"Title"`
	Subject    string  `json:"Subject"`
	Topic      string  `json:"Topic"`
	Tutor_name string  `json:"Tutor_name"`
	Days       [7]bool `json:"Days"`
}
