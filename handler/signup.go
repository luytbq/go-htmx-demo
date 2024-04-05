package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/luytbq/ksjwf/types"
	"github.com/luytbq/ksjwf/view/signup"
)

func HandleSignupIndex(w http.ResponseWriter, r *http.Request) error {
	return signup.SignupIndex(types.SignupCredentials{}, types.SignupError{}).Render(r.Context(), w)
}

func HandleSignupPost(w http.ResponseWriter, r *http.Request) error {
	signupCredential, signupError, err := validateSignup(r.FormValue("email"), r.FormValue("password"), r.FormValue("confirmPassword"))

	slog.Info("HandleSignupPost", "email", r.FormValue("email"), "password", r.FormValue("password"), "confirmPassword", r.FormValue("confirmPassword"))
	slog.Info("SignupCredential", "email", signupCredential.Email)
	slog.Info("SignupError", "email", signupError.EmailError, "password", signupError.PasswordError, "confirmPassword", signupError.ConfirmPasswordError)

	if err != nil {
		return signup.SignupIndex(signupCredential, signupError).Render(r.Context(), w)
	}

	w.Header().Set("Hx-Redirect", "/signup/redirect")
	return signup.SignupIndex(signupCredential, signupError).Render(r.Context(), w)
}

func validateSignup(email, password, confirmPassword string) (signinCreds types.SignupCredentials, signupError types.SignupError, err error) {
	signinCreds = types.SignupCredentials{Email: email}

	// check if email is empty
	if len(email) == 0 {
		msg := "Email is required"
		signupError = types.SignupError{EmailError: msg}
		err = fmt.Errorf(msg)
		return
	}

	// check if password is empty
	if len(password) == 0 {
		msg := "Password is required"
		signupError = types.SignupError{PasswordError: msg}
		err = fmt.Errorf(msg)
		return
	}

	signinCreds.Password = password

	// check if password and confirm password are the same
	if password != confirmPassword {
		msg := "Confirm password does not match"
		signupError = types.SignupError{ConfirmPasswordError: msg}
		err = fmt.Errorf(msg)
		return
	}

	return
}

func HandleSignupRedirectIndex(w http.ResponseWriter, r *http.Request) error {
	return signup.SignupRedirectIndex().Render(r.Context(), w)
}
