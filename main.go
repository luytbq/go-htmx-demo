package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"gihub.com/luytbq/ksjwf/handler"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT_NUMBER")
	staticPath := os.Getenv("STATIC_PATH")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handler.Make(handler.HandleHomeIndex))

	fs := http.FileServer(http.Dir(staticPath))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	slog.Info("Server is running:", "port", port)
	slog.Info("File Server:", "dir", staticPath)
	http.ListenAndServe(":"+port, router)
}

func initEverything() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}
