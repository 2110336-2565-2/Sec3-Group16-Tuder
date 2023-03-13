package services

import (
	"fmt"
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceTutor interface {
	GetTutorByUsername(tutorGet *schemas.SchemaGetTutor) (*schemas.SchemaTutor, error)
	GetTutors() ([]*schemas.SchemaTutor, error)
	CreateTutor(tutorCreate *schemas.SchemaCreateTutor) (*schemas.SchemaTutor, error)
	UpdateTutor(tutorUpdate *schemas.SchemaUpdateTutor) (*schemas.SchemaTutor, error)
	DeleteTutor(tutorDelete *schemas.SchemaDeleteTutor) error
	UpdateTutorSchedule(scheduleUpdate *schemas.SchemaUpdateSchedule) (*schemas.SchemaRawSchedule, error)
	GetTutorSchedule(scheduleRequest *schemas.SchemaGetSchedule) (*schemas.SchemaRawSchedule, error)
}

type serviceTutor struct {
	repository repository.RepositoryTutor
}

func NewServiceTutor(repository repository.RepositoryTutor) *serviceTutor {
	return &serviceTutor{repository: repository}
}

func (s *serviceTutor) GetTutorByUsername(tutorGet *schemas.SchemaGetTutor) (*schemas.SchemaTutor, error) {
	tutor, err := s.repository.GetTutorByUsername(tutorGet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &schemas.SchemaTutor{
		ID:                tutor.ID,
		Username:          tutor.Edges.User.Username,
		Firstname:         tutor.Edges.User.FirstName,
		Lastname:          tutor.Edges.User.LastName,
		Email:             tutor.Edges.User.Email,
		Phone:             tutor.Edges.User.Phone,
		Address:           tutor.Edges.User.Address,
		Birthdate:         tutor.Edges.User.BirthDate,
		Gender:            tutor.Edges.User.Gender,
		ProfilePictureURL: *tutor.Edges.User.ProfilePictureURL,
		Description:       *tutor.Description,
		OmiseBankToken:    *tutor.OmiseBankToken,
		CitizenId:         tutor.CitizenID,
	}, nil
}

func (s *serviceTutor) GetTutors() ([]*schemas.SchemaTutor, error) {
	tutors, err := s.repository.GetTutors()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var tutorResponses []*schemas.SchemaTutor
	for _, tutor := range tutors {
		fmt.Println(tutor.Edges.Schedule.Day3)
		tutorResponses = append(tutorResponses, &schemas.SchemaTutor{
			ID:                tutor.ID,
			Username:          tutor.Edges.User.Username,
			Firstname:         tutor.Edges.User.FirstName,
			Lastname:          tutor.Edges.User.LastName,
			Email:             tutor.Edges.User.Email,
			Phone:             tutor.Edges.User.Phone,
			Address:           tutor.Edges.User.Address,
			Birthdate:         tutor.Edges.User.BirthDate,
			Gender:            tutor.Edges.User.Gender,
			ProfilePictureURL: *tutor.Edges.User.ProfilePictureURL,
			Description:       *tutor.Description,
			OmiseBankToken:    *tutor.OmiseBankToken,
			CitizenId:         tutor.CitizenID,
			// Schedule: schemas.SchemaRawSchedule{
			// 	Sunday:    tutor.Edges.Schedule.Day0,
			// 	Monday:    tutor.Edges.Schedule.Day1,
			// 	Tuesday:   tutor.Edges.Schedule.Day2,
			// 	Wednesday: tutor.Edges.Schedule.Day3,
			// 	Thursday:  tutor.Edges.Schedule.Day4,
			// 	Friday:    tutor.Edges.Schedule.Day5,
			// 	Saturday:  tutor.Edges.Schedule.Day6,
			// },
		})
	}
	return tutorResponses, nil
}

func (s *serviceTutor) CreateTutor(tutorCreate *schemas.SchemaCreateTutor) (*schemas.SchemaTutor, error) {
	tutor, err := s.repository.CreateTutor(tutorCreate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &schemas.SchemaTutor{
		ID:                tutor.ID,
		Username:          tutor.Edges.User.Username,
		Firstname:         tutor.Edges.User.FirstName,
		Lastname:          tutor.Edges.User.LastName,
		Email:             tutor.Edges.User.Email,
		Phone:             tutor.Edges.User.Phone,
		Address:           tutor.Edges.User.Address,
		Birthdate:         tutor.Edges.User.BirthDate,
		Gender:            tutor.Edges.User.Gender,
		ProfilePictureURL: *tutor.Edges.User.ProfilePictureURL,
		Description:       *tutor.Description,
		OmiseBankToken:    *tutor.OmiseBankToken,
		CitizenId:         tutor.CitizenID,
	}, nil
}

func (s *serviceTutor) UpdateTutor(tutorUpdate *schemas.SchemaUpdateTutor) (*schemas.SchemaTutor, error) {
	tutor, err := s.repository.UpdateTutor(tutorUpdate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &schemas.SchemaTutor{
		ID:                tutor.ID,
		Username:          tutor.Edges.User.Username,
		Firstname:         tutor.Edges.User.FirstName,
		Lastname:          tutor.Edges.User.LastName,
		Email:             tutor.Edges.User.Email,
		Phone:             tutor.Edges.User.Phone,
		Address:           tutor.Edges.User.Address,
		Birthdate:         tutor.Edges.User.BirthDate,
		Gender:            tutor.Edges.User.Gender,
		ProfilePictureURL: *tutor.Edges.User.ProfilePictureURL,
		Description:       *tutor.Description,
		OmiseBankToken:    *tutor.OmiseBankToken,
		CitizenId:         tutor.CitizenID,
		Schedule: schemas.SchemaRawSchedule{
			Sunday:    tutor.Edges.Schedule.Day0,
			Monday:    tutor.Edges.Schedule.Day1,
			Tuesday:   tutor.Edges.Schedule.Day2,
			Wednesday: tutor.Edges.Schedule.Day3,
			Thursday:  tutor.Edges.Schedule.Day4,
			Friday:    tutor.Edges.Schedule.Day5,
			Saturday:  tutor.Edges.Schedule.Day6,
		},
	}, nil
}

func (s *serviceTutor) DeleteTutor(tutorDelete *schemas.SchemaDeleteTutor) error {
	err := s.repository.DeleteTutor(tutorDelete)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *serviceTutor) UpdateTutorSchedule(scheduleUpdate *schemas.SchemaUpdateSchedule) (*schemas.SchemaRawSchedule, error) {
	schedule, err := s.repository.UpdateSchedule(scheduleUpdate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &schemas.SchemaRawSchedule{
		Sunday:    schedule.Day0,
		Monday:    schedule.Day1,
		Tuesday:   schedule.Day2,
		Wednesday: schedule.Day3,
		Thursday:  schedule.Day4,
		Friday:    schedule.Day5,
		Saturday:  schedule.Day6,
	}, nil
}

func (s *serviceTutor) GetTutorSchedule(scheduleRequest *schemas.SchemaGetSchedule) (*schemas.SchemaRawSchedule, error) {
	schedule, err := s.repository.GetSchedule(scheduleRequest)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &schemas.SchemaRawSchedule{
		Sunday:    schedule.Day0,
		Monday:    schedule.Day1,
		Tuesday:   schedule.Day2,
		Wednesday: schedule.Day3,
		Thursday:  schedule.Day4,
		Friday:    schedule.Day5,
		Saturday:  schedule.Day6,
	}, nil
}
