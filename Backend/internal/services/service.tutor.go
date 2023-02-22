package services

import (
	repository "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/repositorys"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type ServiceTutor interface {
	GetTutorByUsername(tutorGet *schemas.SchemaGetTutor) (*schemas.SchemaTutor, error)
	GetTutors() ([]*schemas.SchemaTutor, error)
	CreateTutor(tutorCreate *schemas.SchemaCreateTutor) (*schemas.SchemaTutor, error)
	UpdateTutor(tutorUpdate *schemas.SchemaUpdateTutor) (*schemas.SchemaTutor, error)
	DeleteTutor(tutorDelete *schemas.SchemaDeleteTutor) error
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
		return nil, err
	}
	var tutorResponses []*schemas.SchemaTutor
	for _, tutor := range tutors {
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
		})
	}
	return tutorResponses, nil
}

func (s *serviceTutor) CreateTutor(tutorCreate *schemas.SchemaCreateTutor) (*schemas.SchemaTutor, error) {
	tutor, err := s.repository.CreateTutor(tutorCreate)
	if err != nil {
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
	user, tutor, err := s.repository.UpdateTutor(tutorUpdate)
	if err != nil {
		return nil, err
	}
	return &schemas.SchemaTutor{
		ID:        tutor.ID,
		Username:  user.Username,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		Birthdate: user.BirthDate,
		Gender:    user.Gender,
		//ProfilePictureURL: *user.ProfilePictureURL,
		Description:    *tutor.Description,
		OmiseBankToken: *tutor.OmiseBankToken,
		CitizenId:      tutor.CitizenID,
	}, nil
}

func (s *serviceTutor) DeleteTutor(tutorDelete *schemas.SchemaDeleteTutor) error {
	err := s.repository.DeleteTutor(tutorDelete)
	if err != nil {
		return err
	}
	return nil
}
