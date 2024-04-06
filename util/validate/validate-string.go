package validate

import "regexp"

func ValidateEmail(email string) (ok bool, msg string) {
	if len(email) == 0 {
		msg = "Email must not be empty"
		return
	}
	regexpEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !regexpEmail.MatchString(email) {
		msg = "Email is not valid"
		return
	}
	ok = true
	return
}

func ValidatePassword(password string) (ok bool, msg string) {
	if len(password) == 0 {
		msg = "Password must not be empty"
		return
	}

	regexpPassword := regexp.MustCompile(`^.{8,}$`)
	if !regexpPassword.MatchString(password) {
		msg = "Password must be at least 8 characters"
		return
	}
	ok = true
	return
}
