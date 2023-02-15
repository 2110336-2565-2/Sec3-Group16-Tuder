package utils

import (
	"github.com/go-passwd/validator"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digit     = "1234567890"
	nonAlph   = "(~!@#$%^&*_-+=`|\\(){}[]:;\"'<>,.?/)"
)

func PasswordValidator(username, email, firstname, lastname string) *validator.Validator {

	// REMARK: other alphabetic utf-8 char is not permitted
	passwordValidator := validator.New(
		validator.MinLength(8, nil),
		validator.MaxLength(16, nil),
		validator.ContainsOnly(lowercase+uppercase+digit+nonAlph, nil),
		validator.ContainsAtLeast(lowercase, 1, nil),
		validator.ContainsAtLeast(uppercase, 1, nil),
		validator.ContainsAtLeast(digit, 1, nil),
		validator.Similarity([]string{
			username, email, firstname, lastname,
		}, nil, nil),
	)
	return passwordValidator
}
