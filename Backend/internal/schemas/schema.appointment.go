package schemas

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

type SchemaGetAppointmentByStudentID struct {
	StudentID uuid.UUID `json:"id"`
}

type SchemaAppointmentFromStudentID struct {
	ID          uuid.UUID          `json:"id"`
	Appointment []*ent.Appointment `json:"appointment"`
	Course_name string             `json:"course_name"`
	Tutor_name  string             `json:"tutor_name"`
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
