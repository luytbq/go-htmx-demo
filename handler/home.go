package handler

import (
	"net/http"

	"gihub.com/luytbq/ksjwf/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
