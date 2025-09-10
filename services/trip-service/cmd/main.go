package main

import (
	"fmt"
	"log"
	"net/http"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8083")
)

func main() {

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.AllowContentType("application/json"))
	inmemRepo := repository.NewInmemRepository()

	httpHandler := &h.HttpHandler{
		Sevice: service.NewService(inmemRepo),
	}

	fmt.Println("Trip Service is running...")

	mux.Post("/preview", httpHandler.HandlePreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
