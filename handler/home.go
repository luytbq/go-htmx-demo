package handler

import (
	"net/http"

	"github.com/luytbq/ksjwf/types"
	"github.com/luytbq/ksjwf/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	authenticatedUser, ok := r.Context().Value(types.AUTHENTICATED_USER_KEY).(types.AuthenticatedUser)
	if ok == false {
		authenticatedUser = types.AuthenticatedUser{}
	}
	return home.Index(authenticatedUser).Render(r.Context(), w)
}
