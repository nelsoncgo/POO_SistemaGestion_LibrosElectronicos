package services

import (
	"ebook/internal/domain"
	"ebook/internal/interfaces"
)

type BibliotecaService struct {
	repo interfaces.LibroRepository
}

func NewBibliotecaService(
	repo interfaces.LibroRepository,
) *BibliotecaService {

	return &BibliotecaService{
		repo: repo,
	}
}

func (s *BibliotecaService) AgregarLibro(
	libro *domain.Libro,
) error {

	return s.repo.Crear(libro)
}

func (s *BibliotecaService) BuscarLibro(
	id int,
) (*domain.Libro, error) {

	return s.repo.BuscarPorID(id)
}



