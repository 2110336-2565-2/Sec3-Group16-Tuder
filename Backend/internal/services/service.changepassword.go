package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceChangePassword interface {
	ChangePassword(user *schemas.SchemaChangePassword) error
}

type serviceChangePassword struct {
	repo repositorys.RepositoryChangePassword
}

func NewServiceChangePassword(repo repositorys.RepositoryChangePassword) ServiceChangePassword {
	return &serviceChangePassword{
		repo: repo,
	}
}

func (s *serviceChangePassword) ChangePassword(user *schemas.SchemaChangePassword) error {
	err := s.repo.ChangePassword(user)
	if err != nil {
		return err
	}
	return nil
}
