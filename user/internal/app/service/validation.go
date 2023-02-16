package service

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email adress")
)

func validateEmail(email string) error {
	regex := regexp.MustCompile(`^[a-z0-9]+@[a-z.-]+\.[a-z]{2,}$`)
	if !regex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}
