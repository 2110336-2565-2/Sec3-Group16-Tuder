package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaAppointment struct {
	ID      uuid.UUID `json:"id"`
	BeginAt time.Time `json:"begin_at"`
	EndAt   time.Time `json:"end_at"`
	Status  string    `json:"status"`
}

type SchemaGetMatchByID struct {
	ID uuid.UUID `json:"id"`
}

type SchemaMatchesFromID struct {
	MatchID       uuid.UUID `json:"match_id"`
	CourseName    string    `json:"course_name"`
	TutorName     string    `json:"tutor_name"`
	UpcomingClass time.Time `json:"upcoming_class"`
	Remaining     int       `json:"remaining"`
	CoursePictureURL string `json:"course_picture_url"`
}

type SchemaGetAppointmentByMatchID struct {
	MatchID uuid.UUID `json:"match_id"`
}

type SchemaAppointmentsFromMatchID struct {
	CourseName        string               `json:"course_name"`
	TutorName         string               `json:"tutor_name"`
	CourseDescription string               `json:"course_description"`
	Level             string               `json:"level"`
	EstimeateTime     int                  `json:"estimate_time"`
	CoursePictureURL  string               `json:"course_picture_url"`
	Appointments      []*SchemaAppointment `json:"appointments"`
}

type SchemaUpdateAppointmentStatus struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

// type SchemaAppointment struct {
// 	ID         uuid.UUID          `json:"id"`
// 	StartDate  []*ent.Appointment `json:"appointment"`
// 	EndDate    string             `json:"course_name"`
// 	Tutor_name string             `json:"tutor_name"`
// }
