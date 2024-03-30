package handler

import (
	"log/slog"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	slog.Info("Test endpoint is called")
	w.Write([]byte("Hello, World!"))
}
