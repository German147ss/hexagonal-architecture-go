package http_slide_handler

import (
	"encoding/json"
	"net/http"

	"labora_movies/internal/app"
)

// SlideHandler maneja las solicitudes relacionadas con las diapositivas.
type SlideHandler struct {
	Service *app.SlideService
}

// NewSlideHandler crea una nueva instancia de SlideHandler.
func NewSlideHandler(service *app.SlideService) *SlideHandler {
	return &SlideHandler{Service: service}
}

// CreateSlideHandler maneja las solicitudes para crear una nueva diapositiva.
func (h *SlideHandler) CreateSlideHandler(w http.ResponseWriter, r *http.Request) {
	var slide app.Slide
	if err := json.NewDecoder(r.Body).Decode(&slide); err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateSlide(slide); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// SaveBulkHandler maneja las solicitudes para guardar un conjunto de diapositivas.
func (h *SlideHandler) SaveBulkHandler(w http.ResponseWriter, r *http.Request) {
	var slides []app.Slide
	if err := json.NewDecoder(r.Body).Decode(&slides); err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	if err := h.Service.SaveBulk(slides); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SlideHandler) SaveSlidesAndCreatePresentation(w http.ResponseWriter, r *http.Request) {
	var slides []app.Slide
	if err := json.NewDecoder(r.Body).Decode(&slides); err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	backgroundURL := r.URL.Query().Get("background_url")
	if backgroundURL == "" {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	if err := h.Service.SaveSlidesAndCreatePresentation(slides, backgroundURL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
