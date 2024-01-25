package main

import (
	"log"
	"net/http"

	http_movie_handler "labora_movies/api/http"
	"labora_movies/internal/app"
	"labora_movies/pkg/database/postgres"
)

func main() {
	// Configurar y conectar la base de datos PostgreSQL
	postgresRepo := postgres.NewPostgresMovieRepository("postgres://alfred:4lfr3d@localhost:5431/labora?sslmode=disable")

	// Configurar el servicio de películas con el repositorio PostgreSQL
	movieService := app.NewMovieService(postgresRepo)

	// Configurar el controlador HTTP para las películas
	movieHandler := &http_movie_handler.MovieHandler{
		MovieService: *movieService,
	}

	// Configurar rutas de la API HTTP
	http.HandleFunc("/movies/create", movieHandler.CreateMovie)
	http.HandleFunc("/movies/update", movieHandler.UpdateMovie)

	// Iniciar el servidor HTTP
	log.Panic(http.ListenAndServe(":8080", nil))
}
