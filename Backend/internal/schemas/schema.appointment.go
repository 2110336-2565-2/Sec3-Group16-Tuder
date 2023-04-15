package schemas

import (
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

type SchemaGetMatchByID struct {
	ID uuid.UUID `json:"id"`
}

type SchemaMatchesFromID struct {
	MatchID       uuid.UUID `json:"match_id"`
	CourseName    string    `json:"course_name"`
	TutorName     string    `json:"tutor_name"`
	UpcomingClass time.Time `json:"upcoming_class"`
	Remaining     int       `json:"remaining"`
}

type SchemaGetAppointmentByMatchID struct {
	MatchID uuid.UUID `json:"match_id"`
}

type SchemaAppointmentsFromMatchID struct {
	CourseName        string             `json:"course_name"`
	TutorName         string             `json:"tutor_name"`
	CourseDescription string             `json:"course_description"`
	Appointments      []*ent.Appointment `json:"appointments"`
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
