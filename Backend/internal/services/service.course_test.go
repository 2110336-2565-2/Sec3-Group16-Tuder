package services

import (
	"errors"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	"github.com/google/uuid"
	"strings"
	"testing"
)

type Mockup_Repo struct {
	ExistedTutorID uuid.UUID
}

func (m Mockup_Repo) GetCourses() ([]*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) GetCourseByID(sr *schemas.SchemaGetCourse) (*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) UpdateCourse(sr *schemas.SchemaUpdateCourse) (*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) DeleteCourse(sr *schemas.SchemaDeleteCourse) error {
	return nil
}

func (m Mockup_Repo) SearchByTitleRepository(sr *schemas.CourseSearch) ([]*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) SearchBySucjectRepository(sr *schemas.CourseSearch) ([]*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) SearchByTopicRepository(sr *schemas.CourseSearch) ([]*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) SearchByTutorRepository(sr *schemas.CourseSearch) ([]*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) SearchByDayRepository(sr *schemas.CourseSearch) ([]*ent.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mockup_Repo) UpdateCourseStatus(sr *schemas.SchemaUpdateCourseStatus) (*ent.Course, error) {
	return nil, nil
}

func (m Mockup_Repo) CreateCourse(sr *schemas.SchemaCreateCourse) (*ent.Course, error) {
	if sr.Tutor_id != m.ExistedTutorID {
		err := errors.New("something")
		return nil, fmt.Errorf("failed to get tutor: %w", err)
	}
	url := "a string"
	course := ent.Course{
		ID:               sr.Tutor_id,
		Title:            sr.Title,
		Subject:          sr.Subject,
		Topic:            sr.Topic,
		EstimatedTime:    sr.Estimate_time,
		Description:      sr.Description,
		PricePerHour:     sr.Price_per_hour,
		Level:            course.Level(sr.Level),
		CoursePictureURL: &url,
		Status:           "",
		Edges:            ent.CourseEdges{},
	}

	course.Edges.Tutor = &ent.Tutor{
		ID:              sr.Tutor_id,
		Description:     nil,
		OmiseBankToken:  nil,
		CitizenID:       "",
		OmiseCustomerID: nil,
		Edges:           ent.TutorEdges{},
	}

	return &course, nil
}

var (
	mock_uuid, _    = uuid.Parse("a2a87ff5-eb56-4bc8-814e-47a031f5771c")
	exist_id        = mock_uuid
	non_exist_id, _ = uuid.Parse("a2a87ff5-eb56-4bc8-814e-47a031f5772d")

	mockup_repo    = Mockup_Repo{ExistedTutorID: mock_uuid}
	course_service = NewServiceCourse(mockup_repo)
)

func TestServiceCourse_CreateCourse(t *testing.T) {
	sr := schemas.SchemaCreateCourse{
		Tutor_id:       uuid.UUID{},
		Title:          "Title",
		Subject:        "Subject",
		Topic:          "Topic",
		Estimate_time:  0,
		Description:    "description",
		Price_per_hour: 0,
		Level:          "Grade1",
		Course_picture: nil,
	}
	exist_sr := sr
	exist_sr.Tutor_id = exist_id
	non_exist_sr := sr
	non_exist_sr.Tutor_id = non_exist_id
	res1, err1 := course_service.CreateCourse(&non_exist_sr)
	res2, err2 := course_service.CreateCourse(&exist_sr)

	check1, check2 := false, false

	if res1 == nil && err1 != nil {
		errMsg := err1.Error()
		if strings.HasPrefix(errMsg, "failed to get tutor:") {
			fmt.Println("pass: case 1")
			check1 = true
		}
	}

	if res2 != nil && err2 == nil {
		if res2.Tutor_id == exist_id {
			fmt.Println("pass: case 2")
			check2 = true
		}
	}

	if !check1 {
		fmt.Println("res1: ", res1)
		t.Errorf("case 1 failed\nerr1: %q", err1)
	}
	if !check2 {
		fmt.Println("res2: ", res2)
		t.Errorf("case 2 failed\nerr2: %q", err2)
	}
}
