package common

import (
	"regexp"
)

// isEmailValid is to validate email
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
