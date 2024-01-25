package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	http_movie_handler "labora_movies/api/http"
	"labora_movies/internal/app"
	"labora_movies/pkg/database/postgres"

	"github.com/joho/godotenv"
)

func init() {
	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtener las credenciales de la base de datos desde variables de entorno
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("La variable de entorno DB_URL no está configurada")
	}

	// Configurar y conectar la base de datos PostgreSQL
	postgresRepo := postgres.NewPostgresMovieRepository(dbURL)

	// Configurar el servicio de películas con el repositorio PostgreSQL
	movieService := app.NewMovieService(postgresRepo)

	// Configurar el controlador HTTP para las películas
	movieHandler := &http_movie_handler.MovieHandler{
		MovieService: *movieService,
	}

	// Configurar rutas de la API HTTP
	http.HandleFunc("/movies/create", movieHandler.CreateMovie)
	http.HandleFunc("/movies/update", movieHandler.UpdateMovie)
}

func main() {
	fmt.Println("Iniciando el servidor HTTP en http://localhost:8080")
	// Iniciar el servidor HTTP
	log.Panic(http.ListenAndServe(":8080", nil))
}
