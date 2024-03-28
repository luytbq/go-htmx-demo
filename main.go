package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"gihub.com/luytbq/ksjwf/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")
	port := os.Getenv("PORT_NUMBER")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/*", handlers.Test)

	fs := http.FileServer(http.Dir("/home/luytbq/Documents/projects/go-htmx-demo/public/"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	slog.Info("Server is running on port " + port)
	http.ListenAndServe(":"+port, router)
}
