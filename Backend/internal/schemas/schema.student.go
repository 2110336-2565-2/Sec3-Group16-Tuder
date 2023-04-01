package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaStudent struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Phone             string    `json:"phone"`
	Address           string    `json:"address"`
	Birthdate         time.Time `json:"birthdate"`
	Gender            string    `json:"gender"`
	ProfilePictureURL *string   `json:"profile_picture_URL"`
}

type SchemaGetStudent struct {
	Username string `json:"username"`
}

type SchemaCreateStudent struct {
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
}

type SchemaUpdateStudent struct {
	Username          string    `json:"username"` // TODO This must be removed when jwt is completely function
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Email			  string 	`json:"email"`
	Phone             string    `json:"phone"`
	Address           string    `json:"address"`
	ProfilePicture []byte    `json:"new_profile_picture"`
	Birthdate         time.Time `json:"birthdate"`
	Gender            string    `json:"gender"`
}

type SchemaDeleteStudent struct {
	ID uuid.UUID `json:"id"`
}
