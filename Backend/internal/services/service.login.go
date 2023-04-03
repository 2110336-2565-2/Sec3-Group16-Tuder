package services

import (
	"errors"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

type ServiceLogin interface {
	LoginService(l *schemas.SchemaLogin) (*schemas.SchemaLoginResponses, error)
}

type serviceLogin struct {
	repository repositorys.RepositoryLogin
}

func NewServiceLogin(l repositorys.RepositoryLogin) *serviceLogin {
	return &serviceLogin{repository: l}
}

func (s *serviceLogin) LoginService(l *schemas.SchemaLogin) (*schemas.SchemaLoginResponses, error) {

	// check password
	luser, err := s.repository.Login(l)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}

	if !utils.CheckPasswordHash(l.Password, luser.Password) {
		return nil, errors.New("the password isn't match")
	}

	// writing claims
	claims := map[string]interface{}{}

	// adding claims for user
	claims["username"] = luser.Username
	claims["userid"] = luser.ID
	claims["role"] = luser.Role

	// add claims if student
	if *luser.Role == user.RoleStudent {
		claims["studentid"] = luser.Edges.Student.ID
		// add claims if tutor
	} else if *luser.Role == user.RoleTutor {
		claims["tutorid"] = luser.Edges.Tutor.ID
	}

	token, _ := utils.GenerateLoginToken(claims, true)

	return &schemas.SchemaLoginResponses{
		Username: luser.Username,
		Token:    token,
		Role:     luser.Role.String(),
	}, nil
}
