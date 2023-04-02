package schemas

import (
	"time"

	entAppointment "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	"github.com/google/uuid"
)

type SchemaGetAppointment struct {
	ID uuid.UUID `json:"id"`
}

type SchemaAppointment struct {
	ID          uuid.UUID             `json:"id"`
	Begin_at    time.Time             `json:"begin_at"`
	End_at      time.Time             `json:"end_at"`
	Status      entAppointment.Status `json:"status"`
	Course_name string                `json:"course_name"`
	Tutor_name  string                `json:"tutor_name"`
}
