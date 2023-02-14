package schemas

import (
	"time"
)

type SchemaRegister struct {
	Username        string    `json:"username,omitempty"`
	Password        string    `json:"password,omitempty"`
	Email           string    `json:"email,omitempty"`
	ConfirmPassword string    `json:"confirm_password,omitempty"`
	Firstname       string    `json:"firstname,omitempty"`
	Lastname        string    `json:"lastname,omitempty"`
	Address         string    `json:"address,omitempty"`
	Phone           string    `json:"phone,omitempty"`
	Birthdate       time.Time `json:"birthdate,omitempty"`
	Gender          string    `json:"gender,omitempty"`
}

type SchemaRegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
