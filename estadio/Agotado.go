package estadio

type Agotado struct {
	seccion Seccion
}

func (agotado Agotado) ChequearDisponibilidad(entradasSolicitadas int64) bool {

	//enviar error de no hay entradas suficientes
	return false
}

func (agotado Agotado) EfectuarCompra(entrada int64) {}
