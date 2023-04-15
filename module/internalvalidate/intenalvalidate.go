package internalvalidate

import (
	"errors"
	"regexp"
)

func ValidatePhone(phone string) error {
	rgx := regexp.MustCompile(`^08[1-9][0-9]{10,13}$`)
	check := rgx.MatchString(phone)
	if !check {
		return errors.New("error validate phone number")
	}

	return nil
}

func ValidatePassword(pwd string) error {
	rgx := regexp.MustCompile(`^(.*[A-Z].*)$`)
	check := rgx.MatchString(pwd)
	if !check {
		return errors.New("error validate password must contain uppercase")
	}

	rgx = regexp.MustCompile(`^(.*[1-9].*)$`)
	check = rgx.MatchString(pwd)
	if !check {
		return errors.New("error validate password must contain number")
	}

	return nil
}
