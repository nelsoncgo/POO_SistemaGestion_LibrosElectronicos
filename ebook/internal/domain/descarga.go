package domain

import "time"

 // Entidad Descarga

type Descarga struct {
	id        int
	libroID   int
	usuarioID int
	fecha     time.Time
}
