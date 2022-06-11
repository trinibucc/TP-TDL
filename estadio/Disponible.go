package estadio

type Disponible struct {
	seccion      Seccion
	actualizador *Actualizador
}

func (disponible Disponible) ChequearDisponibilidad(entradas int64) bool {

	var disponibles = disponible.seccion.EntradasDisponibles - entradas
	if disponibles < 0 {
		return false
	} else {
		return true
	}

}

func (disponible Disponible) EfectuarCompra(entradas int64) {

	var entradasDisponibles = disponible.seccion.EntradasDisponibles - entradas

	disponible.actualizador.ActualizarEstado(disponible, entradasDisponibles)

}
