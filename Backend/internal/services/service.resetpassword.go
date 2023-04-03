package services

import (
	"errors"
	"fmt"
	"os"

	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceResetPassword interface {
	SentVerificationEmailService(r *schemas.SchemaResetPassword) (err error)
	ResetPasswordService(r *schemas.SchemaNewPassword) (err error)
}

type serviceResetPassword struct {
	repository repository.RepositoryResetPassword
}

func NewServiceResetPassword(r repository.RepositoryResetPassword) *serviceResetPassword {
	return &serviceResetPassword{repository: r}
}

func (s *serviceResetPassword) SentVerificationEmailService(r *schemas.SchemaResetPassword) error {
	// Get user by email
	user, err := s.repository.FindUserByEmail(r)
	if err != nil {
		return errors.New("user not found")
	}

	token, err := utils.GenerateEmailToken(user.ID, true)

	if err != nil {
		return errors.New("internal error")
	}

	resetLink := fmt.Sprintf(os.Getenv("FRONTEND_URL")+"/reset-password/%s", token)
	body := fmt.Sprintf("Please click on the following link to reset your password: %s", resetLink)

	// Send email to user
	err = utils.SendEmail(r.Email, "Reset Your Password", body)

	if err != nil {
		return err
	}

	return nil
}

func (s *serviceResetPassword) ResetPasswordService(r *schemas.SchemaNewPassword) error {

	// identify user by id
	userId, err := utils.DecodeToken(r.Token)

	if err != nil {
		return errors.New("invalid token")
	}

	// get User by Id
	user, err := s.repository.FindUserByID(userId)

	if err != nil {
		return errors.New("user not found")
	}

	// Validate password
	passwordValidator := utils.PasswordValidator(user.Username, user.Email, user.FirstName, user.LastName)
	if err := passwordValidator.Validate(r.Password); err != nil {
		return err
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
