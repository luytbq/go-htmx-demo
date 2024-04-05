package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/luytbq/ksjwf/handler"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT_NUMBER")
	staticPath := os.Getenv("STATIC_PATH")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	setupRouter(router)

	fs := http.FileServer(http.Dir(staticPath))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	slog.Info("Server is running:", "port", port)
	slog.Info("File Server:", "dir", staticPath)
	http.ListenAndServe(":"+port, router)
}

func setupRouter(router *chi.Mux) {
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))
	router.Post("/login", handler.Make(handler.HandleLoginPost))
	router.Get("/signup", handler.Make(handler.HandleSignupIndex))
}

func initEverything() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}
