// core
package app

import "errors"

// Movie representa una película.
type Movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

// Validate realiza validaciones básicas en una película.
func (m *Movie) Validate() error {
	if m.Title == "" {
		return errors.New("el título de la película no puede estar vacío")
	}

	if m.Year < 1900 {
		return errors.New("año de película inválido")
	}

	if m.Rating < 0 || m.Rating > 10 {
		return errors.New("calificación de película fuera de rango (0-10)")
	}

	return nil
}
