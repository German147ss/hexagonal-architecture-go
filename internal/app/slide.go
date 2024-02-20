package app

import "errors"

// Slide representa una diapositiva.
type Slide struct {
	ID          int    `json:"id"`
	ImageURL    string `json:"image_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ExternalURL string `json:"external_url"`
	YouTubeURL  string `json:"youtube_url"`
	Format      string `json:"format"`
	Color       string `json:"color"`
}

// Validate realiza validaciones básicas en una diapositiva.
func (s *Slide) Validate() error {
	if s.ImageURL == "" {
		return errors.New("la URL de la imagen no puede estar vacía")
	}

	if s.Title == "" {
		return errors.New("el título de la diapositiva no puede estar vacío")
	}

	if s.Format != "landscape" && s.Format != "portrait" {
		return errors.New("el formato de la diapositiva debe ser 'landscape' o 'portrait'")
	}

	// Puedes agregar más validaciones según tus necesidades.

	return nil
}
