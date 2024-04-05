package handler

import (
	"net/http"

	"github.com/luytbq/ksjwf/types"
	"github.com/luytbq/ksjwf/view/signup"
)

func HandleSignupIndex(w http.ResponseWriter, r *http.Request) error {
	return signup.SignupIndex(types.SignupCredentials{}, types.SignupError{}).Render(r.Context(), w)
}
