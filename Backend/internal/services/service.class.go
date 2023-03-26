package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceClass interface {
	GetCancellingClasses() ([]*schemas.SchemaCancelRequest, error)
	CancelClass(sc *schemas.SchemaCancelClass) error
	AuditClassCancellation(sc *schemas.SchemaCancelRequestApprove) error
	AcknowledgeClassCancellation(sc *schemas.SchemaUserAcknowledge) error
}

type serviceClass struct {
	repo repositorys.RepositoryClass
}

func NewServiceClass(repo repositorys.RepositoryClass) ServiceClass {
	return &serviceClass{
		repo: repo,
	}
}

func (s *serviceClass) GetCancellingClasses() ([]*schemas.SchemaCancelRequest, error) {
	classes, err := s.repo.GetCancellingClasses()
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *serviceClass) CancelClass(sc *schemas.SchemaCancelClass) error {

	// change status of class to cancelling , wait for admin to confirm

	_, err := s.repo.CancelClass(sc)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceClass) AuditClassCancellation(sc *schemas.SchemaCancelRequestApprove) error {

	// change status of class to cancelled
	err := s.repo.AuditClassCancellation(sc)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceClass) AcknowledgeClassCancellation(sc *schemas.SchemaUserAcknowledge) error {
	
	// change status of class to be default
	err := s.repo.AcknowledgeClassCancellation(sc)
	if err != nil {
		return err
	}
	return nil
}
