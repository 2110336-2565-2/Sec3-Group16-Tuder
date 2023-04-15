package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaAdminTuitionFee struct {
	AppointmentID      uuid.UUID `json:"appointmentID"`
	AppointmentBeginAt time.Time `json:"appointmentBeginAt"`
	AppointmentEndAt   time.Time `json:"appointmentEndAt"`
	MatchID            uuid.UUID `json:"matchId"`
	StudentID          uuid.UUID `json:"studentID"`
	TutorID            uuid.UUID `json:"tutorID"`
	Student_name       string    `json:"student_name"`
	Tutor_name         string    `json:"tutor_name"`
	Price_per_hour     int       `json:"price_per_hour"`
	Title              string    `json:"title"`
	Subject            string    `json:"subject"`
	Topic              string    `json:"topic"`
	Img                []byte    `json:"img"`

	//Course_picture_url string    `json:"course_picture_url"`
}
