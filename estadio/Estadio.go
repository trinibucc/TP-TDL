package estadio

type Estadio struct {
	secciones []Seccion
}

func NuevoEstadio() *Estadio {
	estadio := &Estadio{}
	estadio.secciones = make([]Seccion, 0)
	return estadio
}

func (estadio Estadio) AgregarSeccion(nuevaSeccion Seccion) {
	estadio.secciones = append(estadio.secciones, nuevaSeccion)

}

func ObtenerSeccion(nombre string, estadio Estadio) Seccion {
	var i = 0
	for i := 0; i < len(estadio.secciones); i++ {
		if estadio.secciones[i].nombreSeccion == nombre {
			return estadio.secciones[i]
		}
	}

	return estadio.secciones[i]
}
