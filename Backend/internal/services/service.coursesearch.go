package services

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceCourseSearch interface {
	CourseSearchService(l *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByTitle(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchBySubject(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByTopic(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByTutor(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByDay(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	SearchAllCourse() ([]*schemas.CourseSearchResult, error)
	Interception(l1 []*schemas.CourseSearchResult, l2 []*schemas.CourseSearchResult) []*schemas.CourseSearchResult
	PackToSchema(course []*ent.Course) []*schemas.CourseSearchResult
}

type serviceCourseSearch struct {
	repository repositorys.RepositoryCourseSearch
}

func NewServiceCourseSearch(l repositorys.RepositoryCourseSearch) *serviceCourseSearch {
	return &serviceCourseSearch{repository: l}
}

func (s *serviceCourseSearch) Interception(l1 []*schemas.CourseSearchResult, l2 []*schemas.CourseSearchResult) []*schemas.CourseSearchResult {
	for _, course_ref := range l2 {
		ref := course_ref.Course_id
		stattus := true
		for _, course_check := range l1 {
			if ref == course_check.Course_id {
				stattus = false
				break
			}
		}
		if stattus == true {
			l2 = append(l2, course_ref)
		}
	}
	return l2
}

func (s *serviceCourseSearch) PackToSchema(courses []*ent.Course) []*schemas.CourseSearchResult {
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
	return courseResponses
}

func (s *serviceCourseSearch) CourseSearchService(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if ((((searchContent.Title == "") && (searchContent.Subject == "")) && (searchContent.Topic == "")) && (searchContent.Tutor_name == "")) && (searchContent.Days == [7]bool{false, false, false, false, false, false, false}) {
		return s.SearchAllCourse()
	}
	searchResult, _ := s.CourseSearchByTitle(searchContent)
	subjectSearch, _ := s.CourseSearchBySubject(searchContent)
	searchResult = s.Interception(subjectSearch, searchResult)
	topicSearch, _ := s.CourseSearchByTopic(searchContent)
	searchResult = s.Interception(topicSearch, searchResult)
	tutorSearch, _ := s.CourseSearchByTutor(searchContent)
	searchResult = s.Interception(tutorSearch, searchResult)
	daySearch, _ := s.CourseSearchByDay(searchContent)
	searchResult = s.Interception(daySearch, searchResult)

	return nil, nil
}

func (s *serviceCourseSearch) CourseSearchByTutor(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Tutor_name == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByTutorRepository(searchContent)
	if err != nil {
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourseSearch) CourseSearchByTitle(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Title == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByTitleRepository(searchContent)
	if err != nil {
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourseSearch) CourseSearchByDay(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Days == [7]bool{false, false, false, false, false, false, false} {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByDayRepository(searchContent)
	if err != nil {
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}
func (s *serviceCourseSearch) CourseSearchByTopic(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Topic == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByTopicRepository(searchContent)
	if err != nil {
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourseSearch) CourseSearchBySubject(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Subject == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchBySucjectRepository(searchContent)
	if err != nil {
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourseSearch) SearchAllCourse() ([]*schemas.CourseSearchResult, error) {
	courses, err := s.repository.SearchAll()
	if err != nil {
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}
