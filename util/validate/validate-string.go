package validate

import "regexp"

const (
	emptyString         = ""
	emailEmptyMsg       = "Email must not be empty"
	emailInvalidMsg     = "Email is not valid"
	passwordEmptyMsg    = "Password must not be empty"
	passwordTooShortMsg = "Password must be at least 8 characters"
)

func ValidateEmail(email string) (ok bool, msg string) {
	if len(email) == 0 {
		msg = emailEmptyMsg
		return
	}
	regexpEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !regexpEmail.MatchString(email) {
		msg = emailInvalidMsg
		return
	}
	ok = true
	return
}

func ValidatePassword(password string) (ok bool, msg string) {
	if len(password) == 0 {
		msg = passwordEmptyMsg
		return
	}

	regexpPassword := regexp.MustCompile(`^.{8,}$`)
	if !regexpPassword.MatchString(password) {
		msg = passwordTooShortMsg
		return
	}
	ok = true
	return
}
