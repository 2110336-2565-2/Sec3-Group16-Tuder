package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaTutor struct {
	ID                uuid.UUID         `json:"id"`
	Username          string            `json:"username"`
	Email             string            `json:"email"`
	Firstname         string            `json:"firstname"`
	Lastname          string            `json:"lastname"`
	Phone             string            `json:"phone"`
	Address           string            `json:"address"`
	Birthdate         time.Time         `json:"birthdate"`
	Gender            string            `json:"gender"`
	ProfilePictureURL string            `json:"profile_picture_URL"`
	Description       string            `json:"description"`
	OmiseBankToken    string            `json:"omise_bank_token"`
	CitizenId         string            `json:"citizen_id"`
	Schedule          SchemaRawSchedule `json:"raw_schedule"`
}

type SchemaGetTutorResponse struct {
	ID                uuid.UUID         `json:"id"`
	Username          string            `json:"username"`
	Email             string            `json:"email"`
	Firstname         string            `json:"firstname"`
	Lastname          string            `json:"lastname"`
	Phone             string            `json:"phone"`
	Address           string            `json:"address"`
	Birthdate         time.Time         `json:"birthdate"`
	Gender            string            `json:"gender"`
	ProfilePictureURL string            `json:"profile_picture_URL"`
	Description       string            `json:"description"`
	OmiseBankToken    string            `json:"omise_bank_token"`
	CitizenId         string            `json:"citizen_id"`
	Schedule          SchemaRawSchedule `json:"raw_schedule"`
}

type SchemaGetTutor struct {
	Username string `json:"username"`
}

type SchemaGetTutorScheduleByCourseId struct {
	Course_id uuid.UUID `json:"course_id"`
}
type SchemaGetTutorByID struct {
	ID uuid.UUID `json:"id"`
}

type SchemaCreateTutor struct {
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Email             string    `json:"email"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Phone             string    `json:"phone"`
	Address           string    `json:"address"`
	Birthdate         time.Time `json:"birthdate"`
	Gender            string    `json:"gender"`
	ProfilePictureURL string    `json:"profile_picture_URL"`
	Description       string    `json:"description"`
	OmiseBankToken    string    `json:"omise_bank_token"`
	CitizenId         string    `json:"citizen_id"`
}

type SchemaUpdateTutor struct {
	Username       string    `json:"username,omitempty"` // TODO This must be removed when jwt is completely function
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	ProfilePicture []byte    `json:"new_profile_picture"`
	Birthdate      time.Time `json:"birthdate"`
	Gender         string    `json:"gender"`
	Description    string    `json:"description"`
	OmiseBankToken string    `json:"omise_bank_token"`
	Email          string    `json:"email"`
	Schedule       Schedule  `json:"schedule"`
}

type SchemaDeleteTutor struct {
	ID uuid.UUID `json:"id"`
}

type SchemaGetReviews struct {
	Username string `json:"username"`
}

type SchemaGetReviewsResponse struct {
	Firstname   string            `json:"tutor_firstname"`
	Lastname    string            `json:"tutor_lastname"`
	TotalScore  float32           `json:"total_score"`
	TotalReview int               `json:"total_review"`
	Reviews     []*ReviewResponse `json:"reviews"`
}

type SchemaGetCourses struct {
	Username string `json:"username"`
}
