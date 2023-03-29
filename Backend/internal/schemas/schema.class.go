package schemas

import (
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/classcancelrequest"
	user "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

type SchemaCancelClass struct {
	ClassID      uuid.UUID `json:"classId"`
	UserID       uuid.UUID `json:"userId"`
	ImgURL       string    `json:"img_url"`
	Description  string    `json:"description"`
	ReporterRole user.Role `json:"reporter_role"`
}

type SchemaCancelRequest struct {
	CancelRequestID uuid.UUID                 `json:"cancelRequestId"`
	ClassID         uuid.UUID                 `json:"classId"`
	Title           string                    `json:"title"`
	ImgURL          string                    `json:"img_url"`
	Reporter        string                    `json:"reporter"`
	ReportDate      time.Time                 `json:"report_date"`
	Description     string                    `json:"description"`
	Status          classcancelrequest.Status `json:"status"`
}

type SchemaCancelRequestApprove struct {
	CancelRequestID uuid.UUID `json:"cancelRequestId"` // for audit
	ClassID         uuid.UUID `json:"classId"`
	UserID          uuid.UUID `json:"userId"`
	Approve         bool      `json:"approve"`
}

type SchemaUserAcknowledge struct {
	UserID  uuid.UUID `json:"userId"`
	ClassID uuid.UUID `json:"classId"`
}
