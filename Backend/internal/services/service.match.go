package services

import (
	"fmt"
	"time"

	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceMatch interface {
	CreateMatch(sr *schemas.SchemaCreateMatch) (*schemas.SchemaMatch, error)
	GetMatchByCourseID(sr *schemas.SchemaGetMatchByCourseID) (*schemas.SchemaGetMatchByCourseIDResponse, error)
}

type serviceMatch struct {
	repository repository.RepositoryMatch
}

func NewServiceMatch(repository repository.RepositoryMatch) *serviceMatch {
	return &serviceMatch{repository: repository}
}

// CreateMatch creates a match between a student and a course
func (s *serviceMatch) CreateMatch(sr *schemas.SchemaCreateMatch) (*schemas.SchemaMatch, error) {
	match, err := s.repository.CreateMatch(sr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	match, err = s.repository.GetMatchByID(match.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rawSchedule := &schemas.SchemaRawSchedule{
		Sunday:    match.Edges.Schedule.Day0,
		Monday:    match.Edges.Schedule.Day1,
		Tuesday:   match.Edges.Schedule.Day2,
		Wednesday: match.Edges.Schedule.Day3,
		Thursday:  match.Edges.Schedule.Day4,
		Friday:    match.Edges.Schedule.Day5,
		Saturday:  match.Edges.Schedule.Day6,
	}
	return &schemas.SchemaMatch{
		ID:           match.ID,
		Student:      match.Edges.Student,
		Course:       match.Edges.Course,
		Schedule:     *rawSchedule,
		Appointments: match.Edges.Appointment,
	}, nil
}

func (s *serviceMatch) GetMatchByCourseID(sr *schemas.SchemaGetMatchByCourseID) (*schemas.SchemaGetMatchByCourseIDResponse, error) {
	matches, err := s.repository.GetMatchByCourseID(sr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var data []*schemas.IndivCourseCard
	for _, match := range matches {
		remaining := 0
		var upcoming_class time.Time
		found_upcoming := false
		for _, app := range match.Edges.Appointment {
			if app.Status.String() == "comingsoon" {
				remaining = remaining + 1
				if !found_upcoming {
					upcoming_class = app.BeginAt
					found_upcoming = true
				}
			}
		}
		data = append(data, &schemas.IndivCourseCard{
			StudentUsername:     match.Edges.Student.Edges.User.Username,
			StudentFirstname:    match.Edges.Student.Edges.User.FirstName,
			StudentLastname:     match.Edges.Student.Edges.User.LastName,
			Student_picture_url: *match.Edges.Student.Edges.User.ProfilePictureURL,
			UpcomingClass:       upcoming_class,
			Remaining:           remaining,
			Status:              match.Status.String(),
		})
	}
	response := &schemas.SchemaGetMatchByCourseIDResponse{
		CourseID:   sr.ID,
		CourseName: matches[0].Edges.Course.Title,
		DataCard:   data,
	}

	return response, nil
}
