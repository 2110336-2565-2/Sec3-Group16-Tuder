package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceAppointment interface {
	GetAppointmentByStudentID(sr *schemas.SchemaGetAppointmentByStudentID) ([]*schemas.SchemaAppointmentFromStudentID, error)
	UpdaAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error)
}

type serviceAppointment struct {
	repository repository.RepositoryAppointment
}

func NewServiceAppointment(repository repository.RepositoryAppointment) *serviceAppointment {
	return &serviceAppointment{repository: repository}
}

// GetAppointmentByStudentID gets all appointments of a student
func (s *serviceAppointment) GetAppointmentByStudentID(sr *schemas.SchemaGetAppointmentByStudentID) ([]*schemas.SchemaAppointmentFromStudentID, error) {
	appointments, err := s.repository.GetAppointmentByStudentID(sr)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s *serviceAppointment) UpdaAppointmentStatus(sr *schemas.SchemaUpdateAppointmentStatus) (*ent.Appointment, error) {
	appointment, err := s.repository.UpdateAppointmentStatus(sr)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}
