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
	Intercept(l [][]*schemas.CourseSearchResult) []*schemas.CourseSearchResult
	PackToSchema(course []*ent.Course) []*schemas.CourseSearchResult
}

type serviceCourseSearch struct {
	repository repositorys.RepositoryCourseSearch
}

func NewServiceCourseSearch(l repositorys.RepositoryCourseSearch) *serviceCourseSearch {
	return &serviceCourseSearch{repository: l}
}

func (s *serviceCourseSearch) Intercept(l [][]*schemas.CourseSearchResult) []*schemas.CourseSearchResult {
	result := []*schemas.CourseSearchResult{}
	start := 0
	for idx, r := range l {
		if len(r) > 0 {
			result = r
			start = idx
			break
		}
	}
	for _, lx := range l[start:] {
		tmp := []*schemas.CourseSearchResult{}
		if len(lx) > 0 {
			for _, course_ref := range lx {
				id_ref := course_ref.Course_id
				for _, course_check := range result {
					if id_ref == course_check.Course_id {
						tmp = append(tmp, course_ref)
						break
					}
				}
			}
			result = tmp
		}
	}
	return result
}

func (s *serviceCourseSearch) Union(l [][]*schemas.CourseSearchResult) []*schemas.CourseSearchResult {
	result := []*schemas.CourseSearchResult{}
	for _, lx := range l {
		for _, course_ref := range lx {
			id_ref := course_ref.Course_id
			status := true
			for _, course_check := range result {
				if id_ref == course_check.Course_id {
					status = false
					break
				}
			}
			if status == true {
				result = append(result, course_ref)
			}
		}
	}
	return result
}

func (s *serviceCourseSearch) PackToSchema(courses []*ent.Course) []*schemas.CourseSearchResult {
	var courseResponses []*schemas.CourseSearchResult
	for _, course := range courses {
		courseResponses = append(courseResponses, &schemas.CourseSearchResult{
			Course_id:          course.ID,
			Tutor_name:         course.Edges.Tutor.Edges.User.Username,
			Title:             course.Title,
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

	titleResult, _ := s.CourseSearchByTitle(searchContent)
	subjectSearch, _ := s.CourseSearchBySubject(searchContent)
	topicSearch, _ := s.CourseSearchByTopic(searchContent)
	tutorSearch, _ := s.CourseSearchByTutor(searchContent)
	daySearch, _ := s.CourseSearchByDay(searchContent)

	searchResult := s.Intercept([][]*schemas.CourseSearchResult{titleResult, subjectSearch, topicSearch, tutorSearch, daySearch})
	if len(searchResult) == 0 {
		searchResult = s.Union([][]*schemas.CourseSearchResult{titleResult, subjectSearch, topicSearch, tutorSearch, daySearch})
	}

	// fmt.Println(daySearch)

	return searchResult, nil
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

// ---------------------------------------------------------------------------------------------
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

// ---------------------------------------------------------------------------------------------
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

// ---------------------------------------------------------------------------------------------
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

// ---------------------------------------------------------------------------------------------
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

// ---------------------------------------------------------------------------------------------
func (s *serviceCourseSearch) SearchAllCourse() ([]*schemas.CourseSearchResult, error) {
	courses, err := s.repository.SearchAll()
	if err != nil {
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

//---------------------------------------------------------------------------------------------
