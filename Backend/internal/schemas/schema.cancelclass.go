package schemas

import (
	user "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

type SchemaCancelClass struct {
	ClassId        uuid.UUID `json:"classId"`
	UserID         uuid.UUID `json:"userId"`
	Submitter_role user.Role `json:"submitter_role"`
}

type SchemaCancelRequest struct {
	ClassId     string `json:"classId"`
	Course      string `json:"course"`
	Tutorname   string `json:"tutorname"`
	Subject     string `json:"subject"`
	Time        string `json:"time"`
	Price       string `json:"price"`
	UserID      string `json:"userId"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type SchemaCancelRequestApprove struct {
	ClassId string `json:"classId"`
	UserID  string `json:"userId"`
}
