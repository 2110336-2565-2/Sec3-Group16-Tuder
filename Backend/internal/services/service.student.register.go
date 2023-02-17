package services

import (
	"errors"

	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceStudentRegister interface {
	RegisterStudentService(r *schemas.SchemaRegister) (*schemas.SchemaRegisterResponse, error)
}

type serviceStudentRegister struct {
	repository repository.RepositoryStudentRegister
}

func NewServiceStudentRegister(r repository.RepositoryStudentRegister) *serviceStudentRegister {
	return &serviceStudentRegister{repository: r}
}

func (s serviceStudentRegister) RegisterStudentService(r *schemas.SchemaRegister) (*schemas.SchemaRegisterResponse, error) {
	//TODO connect this function to repo
	if r.Password != r.ConfirmPassword {
		return nil, errors.New("the password isn't match")
	}

	passwordValidator := utils.PasswordValidator(r.Username, r.Email, r.Firstname, r.Lastname)
	if err := passwordValidator.Validate(r.Password); err != nil {
		return nil, err
	}

	_, err := s.repository.RegisterStudentRepository(r)
	if err != nil {
		return nil, err
	}
	token, _ := utils.GenerateToken(r.Username, true)
	return &schemas.SchemaRegisterResponse{
		Success: true,
		Message: "Register Success",
		Token:   token,
		Error:   nil,
	}, nil
}
