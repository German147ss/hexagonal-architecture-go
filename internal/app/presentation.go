package app

import "errors"

// PresentationSlide representa una diapositiva dentro de una presentación con su orden individual.
type PresentationSlide struct {
	Slide
	Order int `json:"order"`
}

// Presentation representa una presentación que contiene diapositivas en un orden específico.
type Presentation struct {
	ID            int                 `json:"id"`
	Slides        []PresentationSlide `json:"slides"`
	BackgroundURL string              `json:"background_url"`
}

// Validate realiza validaciones básicas en una presentación.
func (p *Presentation) Validate() error {
	if len(p.Slides) == 0 {
		return errors.New("la presentación no puede estar vacía")
	}

	// Puedes agregar más validaciones según tus necesidades.

	return nil
}
