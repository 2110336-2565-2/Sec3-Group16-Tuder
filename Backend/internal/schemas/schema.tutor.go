package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaTutor struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
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

type SchemaGetTutor struct {
	Username string `json:"username"`
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
	ID                uuid.UUID `json:"id"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Phone             string    `json:"phone"`
	Address           string    `json:"address"`
	ProfilePictureURL string    `json:"profile_picture_URL"`
	Birthdate         time.Time `json:"birthdate"`
	Gender            string    `json:"gender"`
	Description       string    `json:"description"`
	OmiseBankToken    string    `json:"omise_bank_token"`
}

type SchemaDeleteTutor struct {
	ID uuid.UUID `json:"id"`
}
