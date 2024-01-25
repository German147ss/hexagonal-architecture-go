// core
package app

import "errors"

// MovieRepository define las operaciones que deben ser implementadas por un repositorio de películas.
type MovieRepository interface {
	Save(movie Movie) error
	Update(movie Movie) error
}

// MovieService contiene la lógica de negocio para las películas.
type MovieService struct {
	repo MovieRepository
}

// NewMovieService crea una nueva instancia de MovieService.
func NewMovieService(repo MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

// CreateMovie crea una nueva película.
func (s *MovieService) CreateMovie(movie Movie) error {
	// Realizar validaciones antes de crear la película
	if err := movie.Validate(); err != nil {
		return err
	}

	// Guardar la película en el repositorio
	return s.repo.Save(movie)
}

// UpdateMovie actualiza una película existente.
func (s *MovieService) UpdateMovie(movie Movie) error {
	// Realizar validaciones antes de actualizar la película
	if err := movie.Validate(); err != nil {
		return err
	}

	// Actualizar la película en el repositorio
	return s.repo.Update(movie)
}

// ErrInvalidID se utiliza para indicar un ID de película inválido.
var ErrInvalidID = errors.New("ID de película inválido")
