package handler

import (
	"log/slog"
	"net/http"
)

func Make(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			slog.Error("Internal Server Error:", "path", r.URL.Path, "err", err)
		}
	}
}
