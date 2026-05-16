package main

import (
	"net/http"
	"os"

	"github.com/CesarSotnas/simple-crud-chi/src/controller"
	"github.com/CesarSotnas/simple-crud-chi/src/internal/paths"
	"github.com/CesarSotnas/simple-crud-chi/src/repository"
	"github.com/CesarSotnas/simple-crud-chi/src/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	server := os.Getenv("PORT")

	// dependency injection
	repository := repository.NewGetChampionsRotationsRepository()
	service := service.NewChampionsRotationsSerivce(repository)
	controller := controller.NewChampionsRotationsController(service)

	//routes
	r.Get(paths.GetChampionsRotationPath, controller.GetChampionsController)

	http.ListenAndServe(server, r)
}
