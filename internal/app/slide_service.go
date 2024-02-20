package app

import (
	"errors"
)

// SlideRepository define las operaciones que deben ser implementadas por un repositorio de diapositivas.
type SlideRepository interface {
	Save(slide Slide) error
	Update(slide Slide) error
	SavePresentation(presentation Presentation) error
	SaveBulk(slides []Slide) error
	GetByID(id int) (*Presentation, error)
	SaveSlidesAndCreatePresentation(slides []Slide, backgroundURL string) error
}

// SlideService contiene la lógica de negocio para las diapositivas.
type SlideService struct {
	repo SlideRepository
}

// NewSlideService crea una nueva instancia de SlideService.
func NewSlideService(repo SlideRepository) *SlideService {
	return &SlideService{repo: repo}
}

// CreateSlide crea una nueva diapositiva.
func (s *SlideService) CreateSlide(slide Slide) error {
	// Realizar validaciones antes de crear la diapositiva
	if err := slide.Validate(); err != nil {
		return err
	}

	// Guardar la diapositiva en el repositorio
	return s.repo.Save(slide)
}

// CreatePresentation crea una nueva presentación.
func (s *SlideService) SavePresentation(presentation Presentation) error {
	// Realizar validaciones antes de crear la presentación
	if err := presentation.Validate(); err != nil {
		return err
	}

	// Guardar la presentación en el repositorio
	return s.repo.SavePresentation(presentation)
}

// SaveBulk guarda un conjunto de diapositivas en la base de datos.
func (s *SlideService) SaveBulk(slides []Slide) error {
	// Realizar validaciones antes de guardar las diapositivas
	if len(slides) == 0 {
		return errors.New("no se pueden guardar diapositivas vacías")
	}

	// Guardar las diapositivas en el repositorio
	return s.repo.SaveBulk(slides)
}

func (s *SlideService) SaveSlidesAndCreatePresentation(slides []Slide, backgroundURL string) error {
	// Realizar validaciones antes de obtener la presentación
	if len(slides) == 0 {
		return errors.New("no se pueden guardar diapositivas vacías")
	}

	// Obtener la presentación del repositorio
	return s.repo.SaveSlidesAndCreatePresentation(slides, backgroundURL)
}
