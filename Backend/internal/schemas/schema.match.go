package schemas

import (
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

type SchemaCreateMatch struct {
	Student_id       uuid.UUID               `json:"student_id"`
	Course_id        uuid.UUID               `json:"course_id"`
	Appointment_list []SchemaAppointmentTime `json:"appointment_list"`
	Schedule         Schedule                `json:"schedule"`
}

type SchemaAppointmentTime struct {
	Begin_at time.Time `json:"begin_at"`
	End_at   time.Time `json:"end_at"`
}

type SchemaCreateAppointment struct {
	// Course_time      int                     `json:"course_time"`
	Match            *ent.Match              `json:"match"`
	Appointment_list []SchemaAppointmentTime `json:"appointment_list"`
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
