package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/luytbq/ksjwf/types"
	"github.com/luytbq/ksjwf/view/login"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return login.LoginIndex(types.LoginCredentials{}, types.LoginError{}).Render(r.Context(), w)
}

func HandleLoginPost(w http.ResponseWriter, r *http.Request) error {
	slog.Info("HandleLoginPost", "email", r.FormValue("email"))

	email := r.FormValue("email")
	password := r.FormValue("password")

	if loginCredential, loginError, err := validateLogin(email, password); err != nil {
		return login.LoginIndex(loginCredential, loginError).Render(r.Context(), w)
	}

	if r.FormValue("email") != "luyt@gmail.com" {
		return login.LoginIndex(
			types.LoginCredentials{Email: r.FormValue("email")},
			types.LoginError{EmailError: "Email not found"},
		).Render(r.Context(), w)
	}
	if r.FormValue("password") != "test" {
		return login.LoginIndex(
			types.LoginCredentials{Email: r.FormValue("email")},
			types.LoginError{PasswordError: "Incorrect email/password"},
		).Render(r.Context(), w)
	}

	return login.LoginIndex(types.LoginCredentials{}, types.LoginError{}).Render(r.Context(), w)
}

func validateLogin(email, password string) (loginCredential types.LoginCredentials, loginError types.LoginError, err error) {
	if len(email) == 0 {
		loginCredential = types.LoginCredentials{Email: email}
		msg := "Email is required"
		loginError = types.LoginError{EmailError: msg}
		err = fmt.Errorf(msg)
		return
	}

	if len(password) == 0 {
		loginCredential = types.LoginCredentials{Email: email}
		msg := "Password is required"
		loginError = types.LoginError{PasswordError: msg}
		err = fmt.Errorf(msg)
		return
	}
	return
}
