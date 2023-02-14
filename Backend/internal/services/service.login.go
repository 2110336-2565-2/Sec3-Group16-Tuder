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
	/* 
	 1 ) ใน Service จะเป็นการรับค่าจาก Controller แล้วส่งไปที่ Repository
	 2 ) ใน Service จะต้องส่ง Response กลับไปยัง Controller ในรูปแบบ Schema
	 3 ) ใน Service จะต้องเช็คค่า userLogin ให้ตรง เช่น password ต้อง hash แล้วตรงกับ password ใน database
	 4 ) ใน Service จะต้องเข้าถึง Database ผ่าน Repository เท่านั้น
	*/




	// check password
	
	user, err := s.repository.LoginRepository(userLogin)

	fmt.Println("USER LIST : ",user,"ERROR : ", err)

	return userLogin, nil
}

