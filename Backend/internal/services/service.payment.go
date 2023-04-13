package services

import (
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServicePayment interface {
	GetQRCode(sr *schemas.SchemaGetQRCode) (*schemas.SchemaQRCode, error)
}

type servicePayment struct {
	repository repository.RepositoryPayment
}

func NewServicePayment(repository repository.RepositoryPayment) *servicePayment {
	return &servicePayment{repository: repository}
}

// GetQRCode gets QR code of a payment
func (s *servicePayment) GetQRCode(sr *schemas.SchemaGetQRCode) (*schemas.SchemaQRCode, error) {
	qrCode, err := s.repository.GetQRCode(sr)
	if err != nil {
		return nil, err
	}
	return &schemas.SchemaQRCode{
		QrCodeUrl: qrCode,
	}, nil
}
