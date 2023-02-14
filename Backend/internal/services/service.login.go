package services

import (
	"errors"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceLogin interface {
	LoginService(input *schemas.SchemaLogin) (*schemas.SchemaLoginResponses, error)
}

type serviceLogin struct {
	repository repositorys.RepositoryLogin
}

func NewServiceLogin(repository repositorys.RepositoryLogin) *serviceLogin {
	return &serviceLogin{repository: repository}
}

func (s *serviceLogin)LoginService(userLogin *schemas.SchemaLogin) (*schemas.SchemaLoginResponses, error) {

	// check password
	user, err := s.repository.LoginRepository(userLogin)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(userLogin.Password, user.Password){
		return nil, errors.New("the password isn't match")
	}

	token , _ := utils.GenerateToken(user.Username, true)
	return &schemas.SchemaLoginResponses{
		Username: user.Username,
		Token: token ,
	}, nil
}

