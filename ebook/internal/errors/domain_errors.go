package errors

import "errors"

// Control de errores

var (
	ErrLibroNoEncontrado =
		errors.New("libro no encontrado")

	ErrLibroNoDisponible =
		errors.New("libro no disponible")

	ErrISBNDuplicado =
		errors.New("isbn duplicado")

	ErrUsuarioNoEncontrado =
		errors.New("usuario no encontrado")
)
