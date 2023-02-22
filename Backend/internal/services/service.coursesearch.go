package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceCourseSearch interface {
	CourseSearchService(l *schemas.CourseSearch) (*schemas.CourseSearchResponse, error)
	// CourseSearchByTutorName(CourseInfoByTutorName *schemas.SchemaUpdateTutor) ([]*ent.Course, error)
	// CourseSearchByCourseName(CourseInfoByCourseName *schemas.SchemaUpdateTutor) ([]*ent.Course, error)
	// CourseSearchByTime(CourseInfoByTime *schemas.SchemaUpdateTutor) ([]*ent.Course, error)
	// CourseSearchByDay(CourseInfoByDay *schemas.SchemaUpdateTutor) ([]*ent.Course, error)
	// CourseSearchByContent(CourseInfoByContent *schemas.SchemaUpdateTutor) ([]*ent.Course, error)
	SearchAllCourse() ([]*schemas.CourseSearchResult, error)
}

type serviceCourseSearch struct {
	repository repositorys.RepositoryCourseSearch
}

func NewServiceCourseSearch(l repositorys.RepositoryCourseSearch) *serviceCourseSearch {
	return &serviceCourseSearch{repository: l}
}

func (s *serviceCourseSearch) CourseSearchService(l *schemas.CourseSearch) (*schemas.CourseSearchResponse, error) {

	// check password
	// luser, err := s.repository.LoginRepository(l)
	// if err != nil {
	// 	return nil, err
	// }

	// if !utils.CheckPasswordHash(l.Password, luser.Password) {
	// 	return nil, errors.New("the password isn't match")
	// }

	// token, _ := utils.GenerateLoginToken(luser.Username,luser.Role.String(), true)

	// return &schemas.SchemaLoginResponses{
	// 	Username: luser.Username,
	// 	Token:    token,
	// 	Role:     luser.Role.String(),
	// 	}, nil
	return nil, nil
}

func (s *serviceCourseSearch) CourseSearchByTutorName(CourseInfoByTutorName *schemas.SchemaUpdateTutor) ([]*ent.Course, error) {
	// CourseInfo, err := s.repository.SearchByTutorNameRepository(CourseInfoByTutorName)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (s *serviceCourseSearch) CourseSearchByCourseName(CourseInfoByCourseName *schemas.SchemaUpdateTutor) ([]*ent.Course, error) {
	// CourseInfo, err := s.repository.SearchByCourseNameRepository(CourseInfoByCourseName)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (s *serviceCourseSearch) CourseSearchByDayAvailable(CourseInfoByDayAvailable *schemas.SchemaUpdateTutor) ([]*ent.Course, error) {
	// CourseInfo, err := s.repository.SearchByDayAvailableRepository(CourseInfoByDayAvailable)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
func (s *serviceCourseSearch) CourseSearchByContent(CourseInfoByContent *schemas.SchemaUpdateTutor) ([]*ent.Course, error) {
	// CourseInfo, err := s.repository.SearchByContentRepository(CourseInfoByContent)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (s *serviceCourseSearch) SearchAllCourse() ([]*schemas.CourseSearchResult, error) {
	courses, err := s.repository.SearchAll()
	if err != nil {
		return nil, err
	}
	var courseResponses []*schemas.CourseSearchResult
	for _, course := range courses {
		

		courseResponses = append(courseResponses, &schemas.CourseSearchResult{
			Course_id:          course.ID,
			Tutor_name:         course.Edges.Tutor.Edges.User.Username,
			Tittle:             course.Title,
			Subject:            course.Subject,
			Topic:              course.Topic,
			Estimate_time:      course.EstimatedTime,
			Price_per_hour:     course.PricePerHour,
			Course_picture_url: *course.CoursePictureURL,
		})
	}

	return courseResponses, nil
}
