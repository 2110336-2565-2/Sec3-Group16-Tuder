package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceCancelRequest interface {
	GetCancellingRequests() ([]*schemas.SchemaCancelRequestDetail, error)
	CancelRequest(sc *schemas.SchemaCancelRequest) error
	AuditRequest(sc *schemas.SchemaCancelRequestApprove) error
	// AcknowledgeCancelRequestCancellation(sc *schemas.SchemaUserAcknowledge) error
}

type serviceCancelRequest struct {
	repo repositorys.RepositoryCancelRequest
}

func NewServiceCancelRequest(repo repositorys.RepositoryCancelRequest) ServiceCancelRequest {
	return &serviceCancelRequest{
		repo: repo,
	}
}

func (s *serviceCancelRequest) GetCancellingRequests() ([]*schemas.SchemaCancelRequestDetail, error) {
	cr, err := s.repo.GetCancellingRequests()
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (s *serviceCancelRequest) CancelRequest(sc *schemas.SchemaCancelRequest) error {

	// create cancel request , wait for admin to confirm

	_, err := s.repo.CancelRequest(sc)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceCancelRequest) AuditRequest(sc *schemas.SchemaCancelRequestApprove) error {

	// change status of CancelRequest to cancelled
	err := s.repo.AuditRequest(sc)
	if err != nil {
		return err
	}
	return nil
}

// func (s *serviceCancelRequest) AcknowledgeCancelRequestCancellation(sc *schemas.SchemaUserAcknowledge) error {

// 	// change status of CancelRequest to be default
// 	err := s.repo.AcknowledgeCancelRequestCancellation(sc)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
