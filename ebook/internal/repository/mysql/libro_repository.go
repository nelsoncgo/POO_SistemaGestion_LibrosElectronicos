package mysql

import (
	"database/sql"
)

type LibroRepository struct {
	db *sql.DB
}

func NewLibroRepository(
	db *sql.DB,
) *LibroRepository {

	return &LibroRepository{
		db: db,
	}
}

func (r *LibroRepository) Crear(
	libro *domain.Libro,
) error {

	query := `
	INSERT INTO libros
	(
		titulo,
		isbn,
		descripcion,
		formato,
		disponible
	)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		libro.Titulo(),
		libro.ISBN(),
		libro.Descripcion(),
		libro.Formato(),
		libro.Disponible(),
	)

	return err
}





