package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	http_slide_handler "labora_movies/api/http"
	"labora_movies/internal/app"
	"labora_movies/pkg/database/pg"

	"github.com/joho/godotenv"
)

// Infrastructure
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

	pgRepo := pg.NewPostgresSlideRepository(dbURL)

	slideService := app.NewSlideService(pgRepo)
	// Configurar el servicio de películas con el repositorio PostgreSQL

	// Configurar el controlador HTTP para las películas
	slideHandler := &http_slide_handler.SlideHandler{
		Service: slideService,
	}

	// Configurar rutas de la API HTTP
	http.HandleFunc("/slides/create", slideHandler.CreateSlideHandler)
	http.HandleFunc("/slides/bulk", slideHandler.SaveBulkHandler)
	http.HandleFunc("/slides/create_presentation", slideHandler.SaveSlidesAndCreatePresentation)

}

func main() {
	fmt.Println("Iniciando el servidor HTTP en http://localhost:8080")
	// Iniciar el servidor HTTP
	log.Panic(http.ListenAndServe(":8080", nil))
}
