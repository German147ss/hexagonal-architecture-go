package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"labora_movies/internal/app"

	_ "github.com/lib/pq"
)

// PostgresSlideRepository es una implementación de SlideRepository para PostgreSQL.
type PostgresSlideRepository struct {
	db *sql.DB
}

// NewPostgresSlideRepository crea una nueva instancia de PostgresSlideRepository.
func NewPostgresSlideRepository(connectionStr string) *PostgresSlideRepository {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	// Asegurarse de que la base de datos esté accesible
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresSlideRepository{db: db}
}

// Save guarda una nueva diapositiva en la base de datos.
func (r *PostgresSlideRepository) Save(slide app.Slide) error {
	_, err := r.db.Exec("INSERT INTO slides (image_url, title, description, external_url, youtube_url, format, color) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		slide.ImageURL, slide.Title, slide.Description, slide.ExternalURL, slide.YouTubeURL, slide.Format, slide.Color)

	if err != nil {
		return fmt.Errorf("error al insertar la diapositiva en la base de datos: %v", err)
	}

	return nil
}

// Update actualiza una diapositiva existente en la base de datos.
func (r *PostgresSlideRepository) Update(slide app.Slide) error {
	_, err := r.db.Exec("UPDATE slides SET image_url=$1, title=$2, description=$3, external_url=$4, youtube_url=$5, format=$6, color=$7 WHERE id=$8",
		slide.ImageURL, slide.Title, slide.Description, slide.ExternalURL, slide.YouTubeURL, slide.Format, slide.Color, slide.ID)

	if err != nil {
		return fmt.Errorf("error al actualizar la diapositiva en la base de datos: %v", err)
	}

	return nil
}

// SaveBulk guarda varias diapositivas en la base de datos.
func (r *PostgresSlideRepository) SaveBulk(slides []app.Slide) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO slides (image_url, title, description, external_url, youtube_url, format, color) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return fmt.Errorf("error al preparar la consulta para insertar diapositivas: %v", err)
	}
	defer stmt.Close()

	for _, slide := range slides {
		_, err := stmt.Exec(slide.ImageURL, slide.Title, slide.Description, slide.ExternalURL, slide.YouTubeURL, slide.Format, slide.Color)
		if err != nil {
			return fmt.Errorf("error al insertar la diapositiva en la base de datos: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error al confirmar la transacción: %v", err)
	}

	return nil
}

// Save guarda una nueva presentación en la base de datos.
func (r *PostgresSlideRepository) SavePresentation(presentation app.Presentation) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO presentations (background_url) VALUES ($1) RETURNING id", presentation.BackgroundURL)
	if err != nil {
		return fmt.Errorf("error al insertar la presentación en la base de datos: %v", err)
	}

	presentationID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID de la presentación insertada: %v", err)
	}

	stmt, err := tx.Prepare("INSERT INTO presentation_slides (presentation_id, slide_id, slide_order) VALUES ($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("error al preparar la consulta para insertar diapositivas en la presentación: %v", err)
	}
	defer stmt.Close()

	for _, slide := range presentation.Slides {
		_, err := stmt.Exec(presentationID, slide.ID, slide.Order)
		if err != nil {
			return fmt.Errorf("error al insertar la diapositiva en la presentación: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error al confirmar la transacción: %v", err)
	}

	return nil
}

// GetByID obtiene una presentación por su ID de la base de datos.
func (r *PostgresSlideRepository) GetByID(id int) (*app.Presentation, error) {
	var presentation app.Presentation

	row := r.db.QueryRow("SELECT background_url FROM presentations WHERE id = $1", id)
	err := row.Scan(&presentation.BackgroundURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no se encontró una presentación con ID %d", id)
		}
		return nil, fmt.Errorf("error al obtener la presentación de la base de datos: %v", err)
	}

	rows, err := r.db.Query("SELECT s.id, s.image_url, s.title, s.description, s.external_url, s.youtube_url, s.format, s.color, ps.slide_order FROM slides s INNER JOIN presentation_slides ps ON s.id = ps.slide_id WHERE ps.presentation_id = $1 ORDER BY ps.slide_order", id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las diapositivas de la presentación: %v", err)
	}
	defer rows.Close()

	var slides []app.PresentationSlide
	for rows.Next() {
		var slide app.PresentationSlide
		err := rows.Scan(&slide.ID, &slide.ImageURL, &slide.Title, &slide.Description, &slide.ExternalURL, &slide.YouTubeURL, &slide.Format, &slide.Color, &slide.Order)
		if err != nil {
			return nil, fmt.Errorf("error al escanear la fila de diapositiva: %v", err)
		}
		slides = append(slides, slide)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre las filas de diapositivas: %v", err)
	}

	presentation.ID = id
	presentation.Slides = slides

	return &presentation, nil
}

func (r *PostgresSlideRepository) SaveSlidesAndCreatePresentation(slides []app.Slide, backgroundURL string) error {
	// Guardar las diapositivas en la base de datos
	if err := r.SaveBulk(slides); err != nil {
		return fmt.Errorf("error al guardar las diapositivas: %v", err)
	}

	// Crear una nueva presentación con las diapositivas guardadas
	newPresentation := app.Presentation{
		Slides:        make([]app.PresentationSlide, len(slides)),
		BackgroundURL: backgroundURL,
	}

	for i, slide := range slides {
		newPresentation.Slides[i] = app.PresentationSlide{
			Slide: slide,
			Order: i + 1, // Asigna el orden secuencial
		}
	}

	// Guardar la nueva presentación en la base de datos
	if err := r.SavePresentation(newPresentation); err != nil {
		return fmt.Errorf("error al crear la nueva presentación: %v", err)
	}

	return nil
}
