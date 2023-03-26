package services

import (
	"errors"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	schemas "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	utils "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceChangePassword interface {
	ChangePassword(sc *schemas.SchemaChangePassword) error
	CheckPassword(sc *schemas.SchemaCheckPassword) error
}

type serviceChangePassword struct {
	repo repositorys.RepositoryChangePassword
}

func NewServiceChangePassword(repo repositorys.RepositoryChangePassword) ServiceChangePassword {
	return &serviceChangePassword{
		repo: repo,
	}
}

func (s *serviceChangePassword) ChangePassword(sc *schemas.SchemaChangePassword) error {
	if sc.Password != sc.ConfirmPassword {
		return errors.New("password and confirm password are not the same")
	}

	err := s.repo.ChangePassword(sc)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceChangePassword) CheckPassword(sc *schemas.SchemaCheckPassword) error {
	user, err := s.repo.CheckPassword(sc)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(sc.Password, user.Password) {
		return errors.New("wrong password")
	}

	return nil
}
