package services

import (
	"errors"
	"fmt"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceRegister interface {
	RegisterService(r *schemas.SchemaRegister) (*schemas.SchemaRegisterResponse, error)
}

type serviceRegister struct {
	repository repository.RepositoryRegister
}

func NewServiceRegister(r repository.RepositoryRegister) *serviceRegister {
	return &serviceRegister{repository: r}
}

func (s serviceRegister) RegisterService(r *schemas.SchemaRegister) (*schemas.SchemaRegisterResponse, error) {
	if r.Password != r.ConfirmPassword {
		return nil, errors.New("the password isn't match")
	}

	passwordValidator := utils.PasswordValidator(r.Username, r.Email, r.Firstname, r.Lastname)
	if err := passwordValidator.Validate(r.Password); err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err := s.repository.RegisterUser(r)
	if err != nil {
		fmt.Println(err)
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
