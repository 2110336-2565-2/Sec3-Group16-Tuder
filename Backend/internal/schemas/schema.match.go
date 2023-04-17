package schemas

import (
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

type SchemaCreateMatch struct {
	Student_id uuid.UUID `json:"student_id"`
	Course_id  uuid.UUID `json:"course_id"`
	// Appointment_list []SchemaAppointmentTime `json:"appointment_list"`
	Total_hour int       `json:"total_hour"`
	Schedule   Schedule  `json:"schedule"`
	Payment_id uuid.UUID `json:"payment_id"`
}

type SchemaAppointmentTime struct {
	Begin_at time.Time `json:"begin_at"`
	End_at   time.Time `json:"end_at"`
}

type SchemaCreateAppointment struct {
	// Course_time      int                     `json:"course_time"`
	// Match      *ent.Match `json:"match"`
	Total_hour int      `json:"total_hour"`
	Schedule   Schedule `json:"schedule"`
	// Appointment_list []SchemaAppointmentTime `json:"appointment_list"`
}

type SchemaCreateSchedule struct {
	Schedule Schedule `json:"schedule"`
}

type SchemaMatch struct {
	ID           uuid.UUID          `json:"id"`
	Student      *ent.Student       `json:"student"`
	Course       *ent.Course        `json:"course"`
	Schedule     SchemaRawSchedule  `json:"schedule"`
	Appointments []*ent.Appointment `json:"appointments"`
}

type SchemaUpdateTutorSchedule struct {
	Tutor_schedule_id uuid.UUID `json:"tutor_schedule_id"`
	Schedule          Schedule  `json:"schedule"`
}

type SchemaGetMatchByCourseID struct {
	ID      uuid.UUID `json:"id"`
	TutorID uuid.UUID `json:"tutor_username"`
}

type IndivCourseCard struct {
	MatchID             uuid.UUID `json:"match_id"`
	StudentUsername     string    `json:"student_username"`
	StudentFirstname    string    `json:"student_firstname"`
	StudentLastname     string    `json:"student_lastname"`
	Student_picture_url string    `json:"student_picture_url"`
	UpcomingClass       time.Time `json:"upcoming_class"`
	Remaining           int       `json:"remaining"`
	Status              string    `json:"status"`
}

type SchemaGetMatchByCourseIDResponse struct {
	CourseID   uuid.UUID          `json:"course_id"`
	CourseName string             `json:"course_name"`
	DataCard   []*IndivCourseCard `json:"data_card"`
}
