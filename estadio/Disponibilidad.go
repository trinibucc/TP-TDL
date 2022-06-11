package estadio

type Disponibilidad interface {
	ChequearDisponibilidad(int64) bool
	EfectuarCompra(int64)
}
