package services

import (
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceCourse interface {
	GetCourses() ([]*schemas.SchemaCourse, error)
	GetCourseByID(sr *schemas.SchemaGetCourse) (*schemas.SchemaCourse, error)
	CreateCourse(sr *schemas.SchemaCreateCourse) (*schemas.SchemaCourse, error)
	UpdateCourse(sr *schemas.SchemaUpdateCourse) (*schemas.SchemaCourse, error)
	DeleteCourse(sr *schemas.SchemaDeleteCourse) error
}

type serviceCourse struct {
	repository repository.RepositoryCourse
}

func NewServiceCourse(repository repository.RepositoryCourse) *serviceCourse {
	return &serviceCourse{repository: repository}
}

func (s *serviceCourse) GetCourses() ([]*schemas.SchemaCourse, error) {
	courses, err := s.repository.GetCourses()
	if err != nil {
		return nil, err
	}
	var courseResponses []*schemas.SchemaCourse
	for _, course := range courses {
		courseResponses = append(courseResponses, &schemas.SchemaCourse{
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
		})
	}
	return courseResponses, nil
}

func (s *serviceCourse) GetCourseByID(sr *schemas.SchemaGetCourse) (*schemas.SchemaCourse, error) {
	course, err := s.repository.GetCourseByID(sr)
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
