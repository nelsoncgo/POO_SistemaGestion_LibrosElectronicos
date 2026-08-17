package domain

// Entidad
type Libro struct {
	id          int
	titulo      string
	isbn        string
	descripcion string
	formato     string
	disponible  bool
}

// Constructor

func NewLibro(
	id int,
	titulo string,
	isbn string,
	descripcion string,
	formato string,
) *Libro {

	return &Libro{
		id:          id,
		titulo:      titulo,
		isbn:        isbn,
		descripcion: descripcion,
		formato:     formato,
		disponible:  true,
	}
}

// Getters

func (l *Libro) ID() int {
	return l.id
}

func (l *Libro) Titulo() string {
	return l.titulo
}

func (l *Libro) ISBN() string {
	return l.isbn
}

func (l *Libro) Disponible() bool {
	return l.disponible
}

// Metodos 


func (l *Libro) MarcarDisponible() {
	l.disponible = true
}


func (l *Libro) MarcarNoDisponible() {
	l.disponible = false
}















