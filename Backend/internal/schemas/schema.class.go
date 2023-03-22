package schemas

import (
	user "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

type SchemaCancelClass struct {
	ClassID       uuid.UUID `json:"classId"`
	UserID        uuid.UUID `json:"userId"`
	SubmitterRole user.Role `json:"submitter_role"`
}

type SchemaCancelRequest struct {
	ClassID     uuid.UUID `json:"classId"`
	Course      string    `json:"course"`
	Tutorname   string    `json:"tutor_name"`
	Studentname string    `json:"student_name"`
	Subject     string    `json:"subject"`
	TotalHour   int       `json:"total_hours"`
	SuccessHour int       `json:"success_hours"`
	Price       int       `json:"price"`
}

type SchemaCancelRequestApprove struct {
	ClassID uuid.UUID `json:"classId"`
	UserID  uuid.UUID `json:"userId"`
	Approve bool      `json:"approve"`
}

type SchemaUserAcknowledge struct {
	UserID  uuid.UUID `json:"userId"`
	ClassID uuid.UUID `json:"classId"`
}
