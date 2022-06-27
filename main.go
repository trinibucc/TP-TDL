package main

import (
	"TP-TDL/estadio"
	"fmt"
)

func main() {

	var estadioInstancia = estadio.NuevoEstadio()

	var seccion = estadio.NuevaSeccion("Platea", 5, 6000)

	estadioInstancia.AgregarSeccion(seccion)

	menu := `¿Qué deseas hacer?
			[1] -- Ingrese la cantidad de entradas que desea comprar		
			[2] -- Salir

			----->	`
	//[2] -- Mostrar Tabla -> Verificar.
	//[3] -- Actualizar Compra ya realizada /*Cambiar monto o titular de la operacion
	//[4] -- Eliminar Operacion
	//[5] -- Salir

	var opcion int
	var entradas int64
	var efectuar string
	fmt.Print(menu)
	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		fmt.Print("Ingrese cantidad de entradas:")
		fmt.Scanln(&entradas)
		solicitud := seccion.SolicitarEntradas(entradas)
		//El rechazo de una compra equivaldria a actualizar la tabla y el estado de esa transaccion/compra
		if !solicitud {
			fmt.Print("solicitud rechazada\n")
		} else {
			fmt.Print(("solicitud aceptada\n"))
		}
		//En el caso de que la compra sea aceptada (cumple las condiciones), se procederia a efectuar
		//la compra, seria algo que se resuelve de manera interna (no como en este caso)
		fmt.Print("Ingrse 'S si desea comprar o de lo contrario ingrese 'N: ")
		fmt.Scanln(&efectuar)
		if efectuar == "S" && solicitud {
			seccion.RealizarCompra(entradas)
			fmt.Println("compra realizada")
		} else {
			fmt.Print("No se puede efectuar porque su compra ha sido rechazada")
		}
	case 2:
		fmt.Println("Fin programa")
	}

}
