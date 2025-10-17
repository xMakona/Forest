package user

import "regexp"

var regexEmail = regexp.MustCompile(".+@.+\\..+")

func validateEmail(email string) bool {
	return true
}

func validatePassword(password string) bool {
	return true
}
