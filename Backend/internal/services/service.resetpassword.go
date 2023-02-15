package services

import (
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
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

func (s serviceResetPassword) ResetPasswordService(r *schemas.SchemaResetPassword) error {
	//TODO connect this function to repo
	return nil
}
