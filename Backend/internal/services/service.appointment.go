package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceAppointment interface {
	GetMatchByID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error)
	UpdateAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error)
	GetAppointmentByMatchID(sr *schemas.SchemaGetAppointmentByMatchID) (*schemas.SchemaAppointmentsFromMatchID, error)
}

type serviceAppointment struct {
	repository repository.RepositoryAppointment
}

func NewServiceAppointment(repository repository.RepositoryAppointment) *serviceAppointment {
	return &serviceAppointment{repository: repository}
}

// GetAppointmentByStudentID gets all appointments of a student
func (s *serviceAppointment) GetMatchByID(sr *schemas.SchemaGetMatchByID) ([]*schemas.SchemaMatchesFromID, error) {
	appointments, err := s.repository.GetMatchByStudentID(sr)
	if len(appointments) == 0 {
		appointments, err = s.repository.GetMatchByTutorID(sr)
	}
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s *serviceAppointment) GetAppointmentByMatchID(sr *schemas.SchemaGetAppointmentByMatchID) (*schemas.SchemaAppointmentsFromMatchID, error) {
	appointments, err := s.repository.GetAppointmentByMatchID(sr)

	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (s *serviceAppointment) UpdateAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error) {
	appointment, err := s.repository.UpdateAppointmentStatus(sr)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}
