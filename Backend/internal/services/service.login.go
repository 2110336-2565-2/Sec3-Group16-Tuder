package services

import (
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	// "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/models"
)

type ServiceLogin interface {
	LoginService(input *schemas.SchemaLogin) (*schemas.SchemaLogin, error)
}

type serviceLogin struct {
	repository repositorys.RepositoryLogin
}

func NewServiceLogin(repository repositorys.RepositoryLogin) *serviceLogin {
	return &serviceLogin{repository: repository}
}

func (s *serviceLogin)LoginService(userLogin *schemas.SchemaLogin) (*schemas.SchemaLogin, error) {
	
	// check password
	// if userLogin.Password == ""{
	// }
	user, err := s.repository.LoginRepository(userLogin)

	fmt.Println("USER LIST : ",user,"ERROR : ", err)

	return userLogin, nil
}

