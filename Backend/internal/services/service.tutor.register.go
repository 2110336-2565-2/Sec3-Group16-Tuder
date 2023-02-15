package services

import (
	"errors"

	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceTutorRegister interface {
	RegisterTutorService(r *schemas.SchemaRegister) (err error)
}

type serviceTutorRegister struct {
	repository repository.RepositoryTutorRegister
}

func NewServiceTutorRegister(r repository.RepositoryTutorRegister) *serviceTutorRegister {
	return &serviceTutorRegister{repository: r}
}

func (s serviceTutorRegister) RegisterTutorService(r *schemas.SchemaRegister) error {
	//TODO connect this function to repo
	if r.Password != r.ConfirmPassword {
		return errors.New("the password isn't match")
	}
	passwordValidator := utils.PasswordValidator(r.Username, r.Email, r.Firstname, r.Lastname)
	err := passwordValidator.Validate(r.Password)
	return err
}
