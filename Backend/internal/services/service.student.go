package services

import (

	//"fmt"
	// "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schema"

	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceStudent interface {
	GetStudentByUsername(student *schemas.SchemaGetStudent) (*schemas.SchemaStudent, error) // Get student by username
	GetStudents() ([]*schemas.SchemaStudent, error)                                         // Get all students
	CreateStudent(student *schemas.SchemaCreateStudent) (*schemas.SchemaStudent, error)     // Create student
	UpdateStudent(student *schemas.SchemaUpdateStudent) (*schemas.SchemaStudent, error)     // Update student
	DeleteStudent(student *schemas.SchemaDeleteStudent) error                               // Delete student

}

type serviceStudent struct {
	studentRepository repository.RepositoryStudent
}

func NewServiceStudent(r repository.RepositoryStudent) ServiceStudent {
	return &serviceStudent{studentRepository: r}
}

func (sS *serviceStudent) GetStudentByUsername(studentGet *schemas.SchemaGetStudent) (*schemas.SchemaStudent, error) {
	student, err := sS.studentRepository.GetStudentByUsername(studentGet)
	if err != nil {

		return nil, err
	}
	return &schemas.SchemaStudent{
		ID:                student.ID,
		Username:          student.Edges.User.Username,
		Firstname:         student.Edges.User.FirstName,
		Lastname:          student.Edges.User.LastName,
		Email:             student.Edges.User.Email,
		Phone:             student.Edges.User.Phone,
		Address:           student.Edges.User.Address,
		Birthdate:         student.Edges.User.BirthDate,
		Gender:            student.Edges.User.Gender,
		ProfilePictureURL: student.Edges.User.ProfilePictureURL,
	}, nil
}

func (sS *serviceStudent) GetStudents() ([]*schemas.SchemaStudent, error) {
	students, err := sS.studentRepository.GetStudents()
	if err != nil {
		return nil, err
	}
	var studentResponses []*schemas.SchemaStudent
	for _, student := range students {
		studentResponses = append(studentResponses, &schemas.SchemaStudent{
			ID:                student.ID,
			Username:          student.Edges.User.Username,
			Firstname:         student.Edges.User.FirstName,
			Lastname:          student.Edges.User.LastName,
			Email:             student.Edges.User.Email,
			Phone:             student.Edges.User.Phone,
			Address:           student.Edges.User.Address,
			Birthdate:         student.Edges.User.BirthDate,
			Gender:            student.Edges.User.Gender,
			ProfilePictureURL: student.Edges.User.ProfilePictureURL,
		})
	}

	return studentResponses, nil
}

func (sS *serviceStudent) CreateStudent(studentCreate *schemas.SchemaCreateStudent) (*schemas.SchemaStudent, error) {
	student, err := sS.studentRepository.CreateStudent(studentCreate)
	if err != nil {
		return nil, err
	}

	return &schemas.SchemaStudent{
		ID:                student.ID,
		Username:          student.Edges.User.Username,
		Firstname:         student.Edges.User.FirstName,
		Lastname:          student.Edges.User.LastName,
		Email:             student.Edges.User.Email,
		Phone:             student.Edges.User.Phone,
		Address:           student.Edges.User.Address,
		Birthdate:         student.Edges.User.BirthDate,
		Gender:            student.Edges.User.Gender,
		ProfilePictureURL: student.Edges.User.ProfilePictureURL,
	}, nil
}

func (sS *serviceStudent) UpdateStudent(studentUpdate *schemas.SchemaUpdateStudent) (*schemas.SchemaStudent, error) {
	student, err := sS.studentRepository.UpdateStudent(studentUpdate)
	if err != nil {
		return nil, err
	}
	return &schemas.SchemaStudent{
		ID:                student.ID,
		Username:          student.Edges.User.Username,
		Firstname:         student.Edges.User.FirstName,
		Lastname:          student.Edges.User.LastName,
		Email:             student.Edges.User.Email,
		Phone:             student.Edges.User.Phone,
		Address:           student.Edges.User.Address,
		Birthdate:         student.Edges.User.BirthDate,
		Gender:            student.Edges.User.Gender,
		ProfilePictureURL: student.Edges.User.ProfilePictureURL,
	}, nil
}

func (sS *serviceStudent) DeleteStudent(studentDelete *schemas.SchemaDeleteStudent) error {
	err := sS.studentRepository.DeleteStudent(studentDelete)
	if err != nil {
		return err
	}
	return nil
}
