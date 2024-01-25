// ports
package http_movie_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"labora_movies/internal/app"
)

// MovieHandler maneja las solicitudes HTTP relacionadas con las películas.
type MovieHandler struct {
	MovieService app.MovieService
}

// CreateMovie maneja la creación de una nueva película.
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie app.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
		return
	}

	err = h.MovieService.CreateMovie(movie)
	if err != nil {
		http.Error(w, "Error al crear la película", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateMovie maneja la actualización de una película existente.
func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie app.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
		return
	}

	err = h.MovieService.UpdateMovie(movie)
	if err != nil {
		http.Error(w, "Error al actualizar la película", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ParseMovieID extrae el ID de película de la URL.
func ParseMovieID(r *http.Request) (int, error) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		return 0, app.ErrInvalidID
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, app.ErrInvalidID
	}

	return id, nil
}
