package services

import (
	"errors"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockup_RepoAppointment struct {
	ExistedMatchId uuid.UUID
	ExistedStatus1 string
	ExistedStatus2 string
	ExistedStatus3 string
}

func (m mockup_RepoAppointment) GetMatchByStudentID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error) {
	return nil, nil
}

func (m mockup_RepoAppointment) GetMatchByTutorID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error) {
	return nil, nil
}

func (m mockup_RepoAppointment) GetAppointmentByMatchID(sr *schemas.SchemaGetAppointmentByMatchID) (*schemas.SchemaAppointmentsFromMatchID, error) {
	return nil, nil
}

func checkRes(app *ent.Appointment, err error) (*ent.Appointment, *schemas.SchemaResponses) {
	resp := &schemas.SchemaResponses{
		Success: true,
		Message: "Update appointment status successfully",
		Data:    app,
	}
	if err != nil {
		resp = &schemas.SchemaResponses{
			Success: false,
			Message: "failed to update appointment status",
			Data:    app,
		}
	}
	return app, resp
}

func (m mockup_RepoAppointment) UpdateAppointmentStatus(a *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error) {
	if a.ID == m.ExistedMatchId {
		if a.Status == m.ExistedStatus1 {
			return nil, nil
		}
		if a.Status == m.ExistedStatus2 {
			return nil, nil
		}
		if a.Status == m.ExistedStatus3 {
			return nil, nil
		}
	}
	return nil, errors.New("failed to update appointment status")
}

var (
	muuid, _             = uuid.NewUUID()
	rstatus              = "complete"
	mock_RepoAppointment = mockup_RepoAppointment{
		ExistedMatchId: muuid,
		ExistedStatus1: "pending",
		ExistedStatus2: "verifying",
		ExistedStatus3: "complete",
	}
	appointment_service = NewServiceAppointment(mock_RepoAppointment)

	///// bad uuid and status /////
	wuuid, _ = uuid.NewUUID()
	wstatus  = "finished"
)

// test struct
type update_results struct {
	Response *schemas.SchemaResponses
	Error    error
}

type update_testcase struct {
	Description string
	Input       *schemas.SchemaUpdateAppointmentStatus
	Expected    update_results
}

func TestMatchIDNotFound(t *testing.T) {
	Description := "TC1 MatchId not found"
	sr := &schemas.SchemaUpdateAppointmentStatus{
		ID:     wuuid,
		Status: rstatus,
	}
	app, err := appointment_service.UpdateAppointmentStatus(sr)
	app, res := checkRes(app, err)
	result := update_results{
		Response: res,
		Error:    err,
	}

	expected := update_results{
		Response: &schemas.SchemaResponses{
			Success: false,
			Message: "failed to update appointment status",
			Data:    app,
		},
		Error: errors.New("failed to update appointment status"),
	}
	assert.Equal(t, expected, result, fmt.Sprintf("Failed Test Case: %s", Description))
}

func TestStatusInvalid(t *testing.T) {
	Description := "TC2 Status Invalid"
	sr := &schemas.SchemaUpdateAppointmentStatus{
		ID:     muuid,
		Status: wstatus,
	}
	app, err := appointment_service.UpdateAppointmentStatus(sr)
	app, res := checkRes(app, err)
	result := update_results{
		Response: res,
		Error:    err,
	}
	expected := update_results{
		Response: &schemas.SchemaResponses{
			Success: false,
			Message: "failed to update appointment status",
			Data:    app,
		},
		Error: errors.New("failed to update appointment status"),
	}
	assert.Equal(t, expected, result, fmt.Sprintf("Failed Test Case: %s", Description))
}

func TestUpdateSuccess(t *testing.T) {
	Description := "TC3 Update Success"
	sr := &schemas.SchemaUpdateAppointmentStatus{
		ID:     muuid,
		Status: rstatus,
	}
	app, err := appointment_service.UpdateAppointmentStatus(sr)
	app, res := checkRes(app, err)
	result := update_results{
		Response: res,
		Error:    err,
	}
	expected := update_results{
		Response: &schemas.SchemaResponses{
			Success: true,
			Message: "Update appointment status successfully",
			Data:    app,
		},
		Error: nil,
	}
	assert.Equal(t, expected, result, fmt.Sprintf("Failed Test Case: %s", Description))
}
