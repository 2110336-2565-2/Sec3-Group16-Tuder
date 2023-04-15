package services

import (
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceAdminTuitionFee interface {
	GetAdminTuitionFees() ([]*schemas.SchemaAdminTuitionFee, error)
	PackToSchema(adminTuitionFees []*ent.Appointment) []*schemas.SchemaAdminTuitionFee
}

type serviceAdminTuitionFee struct {
	repository repository.RepositoryAdminTuitionFee
}

func NewServiceAdminTuitionFee(repository repository.RepositoryAdminTuitionFee) *serviceAdminTuitionFee {
	return &serviceAdminTuitionFee{repository: repository}
}

func (s *serviceAdminTuitionFee) PackToSchema(adminTuitionFees []*ent.Appointment) []*schemas.SchemaAdminTuitionFee {
	var adminTuitionFeeResponses []*schemas.SchemaAdminTuitionFee
	for _, Appointment := range adminTuitionFees {
		adminTuitionFeeResponses = append(adminTuitionFeeResponses, &schemas.SchemaAdminTuitionFee{
			AppointmentID:      Appointment.ID,
			AppointmentBeginAt: Appointment.BeginAt,
			AppointmentEndAt:   Appointment.EndAt,
			MatchID:            Appointment.Edges.Match.ID,
			StudentID:          Appointment.Edges.Match.Edges.Student.ID,
			TutorID:            Appointment.Edges.Match.Edges.Course.Edges.Tutor.ID,
			Student_name:       Appointment.Edges.Match.Edges.Student.Edges.User.Username,
			Tutor_name:         Appointment.Edges.Match.Edges.Course.Edges.Tutor.Edges.User.Username,
			Price_per_hour:     Appointment.Edges.Match.Edges.Course.PricePerHour,
			Title:              Appointment.Edges.Match.Edges.Course.Title,
			Subject:            Appointment.Edges.Match.Edges.Course.Subject,
			Topic:              Appointment.Edges.Match.Edges.Course.Topic,
			Img:                []byte(*Appointment.Edges.Match.Edges.Course.CoursePictureURL),
		})
	}
	return adminTuitionFeeResponses
}

func (s *serviceAdminTuitionFee) GetAdminTuitionFees() ([]*schemas.SchemaAdminTuitionFee, error) {
	adminTuitionFees, err := s.repository.GetAdminTuitionFees()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	AdminTuitionFeeResponses := s.PackToSchema(adminTuitionFees)
	return AdminTuitionFeeResponses, nil
}
