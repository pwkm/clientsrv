package utils

import (
	"errors"
	"net/mail"

	"github.com/go-passwd/validator"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidPassword(password string) bool {
	passwordValidator := validator.New(validator.MinLength(5, errors.New(" to short")), validator.MaxLength(10, errors.New("to long")))
	err := passwordValidator.Validate(password)
	return err == nil
}
