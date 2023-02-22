package services

import (
	"errors"

	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceResetPassword interface {
	ResetPasswordService(r *schemas.SchemaResetPassword) (err error)
}

type serviceResetPassword struct {
	repository repository.RepositoryResetPassword
}

func NewServiceResetPassword(r repository.RepositoryResetPassword) *serviceResetPassword {
	return &serviceResetPassword{repository: r}
}

func (s *serviceResetPassword) ResetPasswordService(r *schemas.SchemaResetPassword) error {
	// Get user by email
	user, err := s.repository.FindUserByEmail(r)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(r.Password)
	if err != nil {
		return err
	}

	// Update user's password
	_, err = s.repository.UpdateUserPassword(user, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
