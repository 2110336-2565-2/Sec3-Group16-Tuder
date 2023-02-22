package services

import (
	"errors"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"

)

type ServiceSearch interface {
	SearchService(l *schemas.SchemaSearch) (*schemas.SchemaSearchResponses, error)
	SearchByTutorName(CourseInfoByTutorName *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error)
	SearchByCourseName(CourseInfoByCourseName *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error) 
	SearchByTime(CourseInfoByTime *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error)
	SearchByDay(CourseInfoByDay *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error)
	SearchByContent(CourseInfoByContent *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error)
}

type serviceSearch struct {
	repository repositorys.RepositorySearch
}

func NewServiceSearch(l repositorys.RepositorySearch) *serviceSearch {
	return &serviceSearch{repository: l}
}

func (s *serviceSearch) SearchService(l *schemas.SchemaSearch) (*schemas.SchemaSearchResponses, error) {

	// check password
	luser, err := s.repository.LoginRepository(l)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(l.Password, luser.Password) {
		return nil, errors.New("the password isn't match")
	}

	token, _ := utils.GenerateLoginToken(luser.Username,luser.Role.String(), true)

	return &schemas.SchemaLoginResponses{
		Username: luser.Username,
		Token:    token,
		Role:     luser.Role.String(),
		}, nil
}

func (s *serviceSearch) SearchByTutorName(CourseInfoByTutorName *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error) {
	CourseInfo, err := s.repository.SearchByTutorNameRepository(CourseInfoByTutorName)
	if err != nil {
		return nil, err
	}
	return CourseInfo, nil
}

func (s *serviceSearch) SearchByCourseName(CourseInfoByCourseName *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error) {
	CourseInfo, err := s.repository.SearchByCourseNameRepository(CourseInfoByCourseName)
	if err != nil {
		return nil, err
	}
	return CourseInfo, nil
}

func (s *serviceSearch) SearchByDayAvailable(CourseInfoByDayAvailable *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error) {
	CourseInfo, err := s.repository.SearchByDayAvailableRepository(CourseInfoByDayAvailable)
	if err != nil {
		return nil, err
	}
	return CourseInfo, nil
}
func (s *serviceSearch) SearchByContent(CourseInfoByContent *schemas.SchemaUpdateTutor) ([]*ent.CourseInfo, error) {
	CourseInfo, err := s.repository.SearchByContentRepository(CourseInfoByContent)
	if err != nil {
		return nil, err
	}
	return CourseInfo, nil
}