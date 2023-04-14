package services

import (
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceCourse interface {
	GetCourses() ([]*schemas.CourseSearchResult, error)
	GetCourseByID(sr *schemas.SchemaGetCourse) (*schemas.SchemaCourse, error)
	CreateCourse(sr *schemas.SchemaCreateCourse) (*schemas.SchemaCourse, error)
	UpdateCourse(sr *schemas.SchemaUpdateCourse) (*schemas.SchemaCourse, error)
	DeleteCourse(sr *schemas.SchemaDeleteCourse) error
	CourseSearchService(l *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByTitle(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchBySubject(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByTopic(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByTutor(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	CourseSearchByDay(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error)
	Intercept(l [][]*schemas.CourseSearchResult) []*schemas.CourseSearchResult
	PackToSchema(course []*ent.Course) []*schemas.CourseSearchResult
	UpdateCourseStatus(sr *schemas.SchemaUpdateCourseStatus) (*schemas.SchemaUpdateCourseStatusResponse, error)
}

type serviceCourse struct {
	repository repository.RepositoryCourse
}

func NewServiceCourse(repository repository.RepositoryCourse) *serviceCourse {
	return &serviceCourse{repository: repository}
}

func (s *serviceCourse) Intercept(l [][]*schemas.CourseSearchResult) []*schemas.CourseSearchResult {
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

func (s *serviceCourse) Union(l [][]*schemas.CourseSearchResult) []*schemas.CourseSearchResult {
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

func (s *serviceCourse) PackToSchema(courses []*ent.Course) []*schemas.CourseSearchResult {
	var courseResponses []*schemas.CourseSearchResult
	for _, course := range courses {
		courseResponses = append(courseResponses, &schemas.CourseSearchResult{
			Course_id:          course.ID,
			Tutor_name:         course.Edges.Tutor.Edges.User.Username,
			Title:              course.Title,
			Subject:            course.Subject,
			Topic:              course.Topic,
			Estimate_time:      course.EstimatedTime,
			Price_per_hour:     course.PricePerHour,
			Course_picture_url: *course.CoursePictureURL,
		})
	}
	return courseResponses
}

func (s *serviceCourse) CourseSearchService(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {

	if ((((searchContent.Title == "") && (searchContent.Subject == "")) && (searchContent.Topic == "")) && (searchContent.Tutor_name == "")) && (searchContent.Days == [7]bool{false, false, false, false, false, false, false}) {
		return s.GetCourses()
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

func (s *serviceCourse) CourseSearchByTutor(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Tutor_name == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByTutorRepository(searchContent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourse) CourseSearchByTitle(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Title == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByTitleRepository(searchContent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourse) CourseSearchByDay(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Days == [7]bool{false, false, false, false, false, false, false} {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByDayRepository(searchContent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourse) CourseSearchByTopic(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Topic == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchByTopicRepository(searchContent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourse) CourseSearchBySubject(searchContent *schemas.CourseSearch) ([]*schemas.CourseSearchResult, error) {
	if searchContent.Subject == "" {
		return []*schemas.CourseSearchResult{}, nil
	}
	courses, err := s.repository.SearchBySucjectRepository(searchContent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourse) GetCourses() ([]*schemas.CourseSearchResult, error) {
	courses, err := s.repository.GetCourses()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	courseResponses := s.PackToSchema(courses)
	return courseResponses, nil
}

func (s *serviceCourse) GetCourseByID(sr *schemas.SchemaGetCourse) (*schemas.SchemaCourse, error) {
	course, err := s.repository.GetCourseByID(sr)
	if err != nil {
		return nil, err
	}
	count, rating := utils.CalcReviewStat(course.Edges.Review)
	courseResponse := &schemas.SchemaCourse{
		ID:                 course.ID,
		Title:              course.Title,
		Rating:             rating,
		ReviewCount:        count,
		Reviews:            []schemas.Review{},
		Tutor_id:           course.Edges.Tutor.ID,
		TutorFirstname:     course.Edges.Tutor.Edges.User.FirstName,
		TutorLastname:      course.Edges.Tutor.Edges.User.LastName,
		Subject:            course.Subject,
		Topic:              course.Topic,
		Level:              course.Level.String(),
		Description:        course.Description,
		Estimate_time:      course.EstimatedTime,
		Price_per_hour:     course.PricePerHour,
		Course_picture_url: *course.CoursePictureURL,
	}
	// append ent.Review to sG
	for _, r := range course.Edges.Review {
		courseResponse.Reviews = append(courseResponse.Reviews, schemas.Review{
			Score:        *r.Score,
			ReviewMsg:    *r.ReviewMsg,
			ReviewTimeAt: r.ReviewTimeAt,
		})
	}
	return courseResponse, nil
}

func (s *serviceCourse) CreateCourse(sr *schemas.SchemaCreateCourse) (*schemas.SchemaCourse, error) {
	course, err := s.repository.CreateCourse(sr)
	if err != nil {
		return nil, err
	}
	courseResponse := &schemas.SchemaCourse{
		ID:                 course.ID,
		Tutor_id:           course.Edges.Tutor.ID,
		Title:              course.Title,
		Description:        course.Description,
		Subject:            course.Subject,
		Topic:              course.Topic,
		Estimate_time:      course.EstimatedTime,
		Price_per_hour:     course.PricePerHour,
		Course_picture_url: *course.CoursePictureURL,
		Level:              course.Level.String(),
	}
	return courseResponse, nil
}

func (s *serviceCourse) UpdateCourse(sr *schemas.SchemaUpdateCourse) (*schemas.SchemaCourse, error) {
	course, err := s.repository.UpdateCourse(sr)
	if err != nil {
		return nil, err
	}
	courseResponse := &schemas.SchemaCourse{
		ID:                 course.ID,
		Tutor_id:           course.Edges.Tutor.ID,
		Title:              course.Title,
		Description:        course.Description,
		Subject:            course.Subject,
		Topic:              course.Topic,
		Estimate_time:      course.EstimatedTime,
		Price_per_hour:     course.PricePerHour,
		Course_picture_url: *course.CoursePictureURL,
		Level:              course.Level.String(),
	}
	return courseResponse, nil
}

func (s *serviceCourse) DeleteCourse(sr *schemas.SchemaDeleteCourse) error {
	err := s.repository.DeleteCourse(sr)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceCourse) UpdateCourseStatus(sr *schemas.SchemaUpdateCourseStatus) (*schemas.SchemaUpdateCourseStatusResponse, error) {
	course, err := s.repository.UpdateCourseStatus(sr)
	if err != nil {
		return nil, err
	}
	courseResponse := &schemas.SchemaUpdateCourseStatusResponse{
		ID:     course.ID,
		Title:  course.Title,
		Status: course.Status.String(),
	}
	return courseResponse, nil
}
