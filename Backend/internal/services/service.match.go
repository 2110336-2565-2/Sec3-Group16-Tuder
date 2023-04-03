package services

import (
	"fmt"

	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceMatch interface {
	CreateMatch(sr *schemas.SchemaCreateMatch) (*schemas.SchemaMatch, error)
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
	fmt.Println("create match")
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
	fmt.Println("end service")
	return &schemas.SchemaMatch{
		ID:           match.ID,
		Student:      match.Edges.Student,
		Course:       match.Edges.Course,
		Schedule:     *rawSchedule,
		Appointments: match.Edges.Appointment,
	}, nil
}
