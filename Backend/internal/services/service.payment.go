package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServicePayment interface {
	GetQRCodeForCoursePayment(sr *schemas.SchemaGetQRCodeForCoursePayment) (*schemas.SchemaQRCode, error)
	GetQRCodeForTuitionFree(sr *schemas.SchemaGetQRCodeForTuitionFree) (*schemas.SchemaQRCode, error)
	WebhookChargeHandler(sr *schemas.SchemaChargeWebhook) (*ent.Payment, error)
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
	return qrCode, nil
}

func (s *servicePayment) GetQRCodeForTuitionFree(sr *schemas.SchemaGetQRCodeForTuitionFree) (*schemas.SchemaQRCode, error) {
	qrCode, err := s.repository.GetQRCodeForTuitionFree(sr)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}

func (s *servicePayment) WebhookChargeHandler(sr *schemas.SchemaChargeWebhook) (*ent.Payment, error) {
	payment, err := s.repository.WebhookChargeHandler(sr)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
