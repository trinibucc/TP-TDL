package estadio

type Seccion struct {
	nombreSeccion       string
	EntradasDisponibles int64
	Precio              int64
	EstadoSeccion       Disponibilidad
}

func NuevaSeccion(nombre string, entradas int64, precio int64) Seccion {
	seccion := Seccion{
		nombreSeccion:       nombre,
		EntradasDisponibles: entradas,
		Precio:              precio,
		EstadoSeccion:       nil,
	}
	if entradas > 0 {
		seccion.EstadoSeccion = Disponible{seccion: seccion}
	} else {
		seccion.EstadoSeccion = Agotado{seccion: seccion}
	}
	return seccion
}

func (seccion *Seccion) ModificarCapacidad(entradas int64, nuevoEstado Disponibilidad) {
	seccion.EntradasDisponibles = entradas
	seccion.EstadoSeccion = nuevoEstado
}

func (seccion *Seccion) SolicitarEntradas(entradas int64) bool {

	solicitud := seccion.EstadoSeccion.ChequearDisponibilidad(entradas)
	return solicitud
}

func (seccion *Seccion) RealizarCompra(entradas int64) {

	seccion.EstadoSeccion.EfectuarCompra(entradas)
}
