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
	ClassID       uuid.UUID `json:"classId"`
	Course        string    `json:"course"`
	Tutorname     string    `json:"tutorname"`
	Studentname   string    `json:"studentname"`
	Subject       string    `json:"subject"`
	Time          int       `json:"time"`
	Price         int       `json:"price"`
}

type SchemaCancelRequestApprove struct {
	ClassID string `json:"classId"`
	UserID  string `json:"userId"`
}
