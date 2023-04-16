package schemas

import (
	"time"

	cancelrequest "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/cancelrequest"
	user "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

type SchemaCancelRequest struct {
	MatchID      uuid.UUID `json:"matchId"`
	UserID       uuid.UUID `json:"userId"`
	Title        string    `json:"title"`
	Img       []byte    `json:"img"`
	Description  string    `json:"description"`
	ReporterRole user.Role `json:"reporter_role"`
}

type SchemaCancelRequestDetail struct {
	CancelRequestID uuid.UUID            `json:"cancelRequestId"`
	MatchID         uuid.UUID            `json:"matchId"`
	Title           string               `json:"title"`
	ImgURL          string               `json:"img_url"`
	Reporter        string               `json:"reporter"`
	ReportDate      time.Time            `json:"report_date"`
	Description     string               `json:"description"`
	Status          cancelrequest.Status `json:"status"`
}

type SchemaCancelRequestApprove struct {
	MatchID 	   uuid.UUID `json:"matchId"` // for audit
	CancelRequestID uuid.UUID `json:"cancelRequestId"` // for audit
	Approve         bool      `json:"approve"`
}

type SchemaUserAcknowledge struct {
	CancelRequestID uuid.UUID `json:"cancelRequestId"`
	UserID          uuid.UUID `json:"userId"`
}
