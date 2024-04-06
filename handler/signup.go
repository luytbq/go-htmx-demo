package handler

import (
	"fmt"
	"net/http"

	"github.com/luytbq/ksjwf/types"
	"github.com/luytbq/ksjwf/util/validate"
	"github.com/luytbq/ksjwf/view/signup"
)

func HandleSignupIndex(w http.ResponseWriter, r *http.Request) error {
	return signup.SignupIndex(types.SignupCredentials{}, types.SignupError{}).Render(r.Context(), w)
}

func HandleSignupPost(w http.ResponseWriter, r *http.Request) error {
	signupCredential, signupError, err := validateSignup(r.FormValue("email"), r.FormValue("password"), r.FormValue("confirmPassword"))

	if err != nil {
		return signup.SignupForm(signupCredential, signupError).Render(r.Context(), w)
	}

	w.Header().Set("Hx-Redirect", "/signup/redirect")
	return nil
}

func validateSignup(email, password, confirmPassword string) (signupCreds types.SignupCredentials, signupError types.SignupError, err error) {
	signupCreds = types.SignupCredentials{Email: email}

	if ok, msg := validate.ValidateEmail(email); !ok {
		signupError.EmailError = msg
		err = fmt.Errorf(msg)
		return
	}

	if ok, msg := validate.ValidatePassword(password); !ok {
		signupError.PasswordError = msg
		err = fmt.Errorf(msg)
		return
	}

	signupCreds.Password = password

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
