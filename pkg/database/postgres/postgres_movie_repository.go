// adaptador
package postgres

import (
	"database/sql"
	"fmt"

	"labora_movies/internal/app"

	_ "github.com/lib/pq"
)

// PostgresMovieRepository es una implementación de MovieRepository para PostgreSQL.
type PostgresMovieRepository struct {
	db *sql.DB
}

// NewPostgresMovieRepository crea una nueva instancia de PostgresMovieRepository.
func NewPostgresMovieRepository(connectionStr string) *PostgresMovieRepository {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	// Asegurarse de que la base de datos esté accesible
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresMovieRepository{db: db}
}

// Save guarda una nueva película en la base de datos.
func (r *PostgresMovieRepository) Save(movie app.Movie) error {
	_, err := r.db.Exec("INSERT INTO movies (title, year, rating) VALUES ($1, $2, $3)",
		movie.Title, movie.Year, movie.Rating)

	if err != nil {
		return fmt.Errorf("error al insertar la película en la base de datos: %v", err)
	}

	return nil
}

// Update actualiza una película existente en la base de datos.
func (r *PostgresMovieRepository) Update(movie app.Movie) error {
	_, err := r.db.Exec("UPDATE movies SET title=$1, year=$2, rating=$3 WHERE id=$4",
		movie.Title, movie.Year, movie.Rating, movie.ID)

	if err != nil {
		return fmt.Errorf("error al actualizar la película en la base de datos: %v", err)
	}

	return nil
}
