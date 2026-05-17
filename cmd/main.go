package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CesarSotnas/simple-crud-chi/internal/controller"
	"github.com/CesarSotnas/simple-crud-chi/internal/repository"
	"github.com/CesarSotnas/simple-crud-chi/internal/routes"
	"github.com/CesarSotnas/simple-crud-chi/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	godotenv.Load()

	server := ":" + os.Getenv("PORT")

	// dependency injection
	fmt.Println("Iniciando injeção de dependências")
	repository := repository.NewGetChampionsRotationsRepository()
	service := service.NewChampionsRotationsSerivce(repository)
	controller := controller.NewChampionsRotationsController(service)

	fmt.Println("Iniciando rotas")
	//routes
	r.Get(routes.GetChampionsRotationPath, controller.GetChampionsController)

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(server, r)
}
