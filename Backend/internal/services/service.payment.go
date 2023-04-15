package services

import (
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServicePayment interface {
	GetQRCodeForCoursePayment(sr *schemas.SchemaGetQRCodeForCoursePayment) (*schemas.SchemaQRCode, error)
	GetQRCodeForTuitionFree(sr *schemas.SchemaGetQRCodeForTuitionFree) (*schemas.SchemaQRCode, error)
}

type servicePayment struct {
	repository repository.RepositoryPayment
}

func NewServicePayment(repository repository.RepositoryPayment) *servicePayment {
	return &servicePayment{repository: repository}
}

// GetQRCode gets QR code of a payment
func (s *servicePayment) GetQRCodeForCoursePayment(sr *schemas.SchemaGetQRCodeForCoursePayment) (*schemas.SchemaQRCode, error) {
	qrCode, err := s.repository.GetQRCodeForCoursePayment(sr)
	if err != nil {
		return nil, err
	}
	return &schemas.SchemaQRCode{
		QrCodeUrl: qrCode,
	}, nil
}

func (s *servicePayment) GetQRCodeForTuitionFree(sr *schemas.SchemaGetQRCodeForTuitionFree) (*schemas.SchemaQRCode, error) {
	qrCode, err := s.repository.GetQRCodeForTuitionFree(sr)
	if err != nil {
		return nil, err
	}
	return &schemas.SchemaQRCode{
		QrCodeUrl: qrCode,
	}, nil
}
