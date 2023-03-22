package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceCancelClass interface {
	CancelClass(sc *schemas.SchemaCancelClass) error
}

type serviceCancelClass struct {
	repo repositorys.RepositoryCancelClass
}

func NewServiceCancelClass(repo repositorys.RepositoryCancelClass) ServiceCancelClass {
	return &serviceCancelClass{
		repo: repo,
	}
}

func (s *serviceCancelClass) CancelClass(sc *schemas.SchemaCancelClass) error {

	// change status of class to cancelling , wait for admin to confirm

	_, err := s.repo.CancelClass(sc)
	if err != nil {
		return err
	}
	return nil
}
