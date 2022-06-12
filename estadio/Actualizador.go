package estadio

type Actualizador struct {
}

func (actualizador *Actualizador) ActualizarEstado(disponible Disponible, entradasDisponibles int64) {

	if entradasDisponibles > 0 {
		disponible.seccion.ModificarCapacidad(entradasDisponibles, disponible)
	} else {
		disponible.seccion.ModificarCapacidad(0, Agotado{disponible.seccion})
	}

}
