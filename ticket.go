package main

import (
	// Leer líneas incluso si tienen espacios
	"bufio"
	"fmt" //Para los mensajes
	"os"
	"strings"
	"time"
	//Para conversión
	// "database/sql" Interactuar con bases de datos
)

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~¿Podría ser una clase?~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
type Ticket struct {
	FechaReserva   string
	FechaConcierto string
	Id             int
	seccion        string
}
type Cliente struct {
	Nombre     string
	Id, Compra int
}

func cargar() {}
func fechaConcierto(fechaC string) string {
	return fechaC
}

func id(c int) int {

	return c + 1
}

func seccion() {}

func main() {

	menu := `¿Qué deseas hacer?
[1] -- Agregar Compra
[2] -- Mostrar Tabla -> Verificar. 
[3] -- Actualizar Compra ya realizada /*Cambiar monto o titular de la operacion */
[4] -- Eliminar Operacion
[5] -- Salir
----->	`
	var opcion int

	for opcion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&opcion)
		scanner := bufio.NewScanner(os.Stdin)
		var cliente Cliente

		switch opcion {
		case 1:
			var compra = tarifa()
			cliente.Compra = compra

			fmt.Print("Ingresa el nombre: ")

			if scanner.Scan() {
				cliente.Nombre = scanner.Text()
			}

			//fmt.Println("El monto de la compra es:")
			//if scanner.Scan() {
			//	//var compra = scanner.Text()
			//	//cliente.Compra, _ = strconv.Atoi(compra) //La conversion puede dar un error y por eso el "_"
			//	cliente.Compra = compra

			//}
			fmt.Println("")
			fmt.Println("----------------------------------------------------")
			fmt.Println(strings.ToUpper(cliente.Nombre)+" el monto de la compra es: $", cliente.Compra)
			fmt.Println("----------------------------------------------------")
			var estado_compra bool
			estado_compra = confirmacion()

			leyendaCompra(estado_compra)
			time.Sleep(3 * time.Second)
			//}
			// ver formato de listas para insertar cliente !!!!!!!!!!!!!!
			//	/*¿El ID debería ser generado de manera random? -> Id consecutivos*/
			//	err := insertar(cliente) //err = error
			//	if err != nil {
			//		fmt.Printf("Error insertando: %v", err)
			//	} else {
			//		fmt.Println("Insertado correctamente")
			//	}
			//case 2:
			//	clientes, err := obtenerClientes()
			//	if err != nil {
			//		fmt.Printf("Error obteniendo contactos: %v", err)
			//	} else {
			//		for _, clientes := range clientes {
			//			fmt.Println("====================")
			//			fmt.Printf("Nombre: %s\n", cliente.Nombre)
			//			fmt.Printf("Id: %d\n", cliente.Id)
			//			fmt.Printf("Id: %d\n", cliente.Compra)
			//		}
			//	}
			//case 3:
			//	/* ¿El ID al actualizar una compra lo mantenemos o le generamos uno nuevo?
			//	Opcion 1 fmt.Println("Ingresa el id:")
			//	Opcion 2 generarlo de manera random */
			//	fmt.Scanln(&cliente.Id)
			//	fmt.Println("Ingresa el nuevo nombre:")
			//	if scanner.Scan() {
			//		cliente.Nombre = scanner.Text()
			//	}
			//	fmt.Println("Ingresa el nuevo valor de compra:")
			//	if scanner.Scan() {
			//		var compra = scanner.Text()
			//		cliente.Compra, _ = strconv.Atoi(compra)
			//	}
			//	err := actualizar(cliente)
			//	if err != nil {
			//		fmt.Printf("Error actualizando: %v", err)
			//	} else {
			//		fmt.Println("Actualizado correctamente")
			//	}
			//case 4:
			//	fmt.Println("Ingresa el ID de la operacion que desea eliminar:")
			//	fmt.Scanln(&cliente.Id)
			//	err := eliminar(cliente)
			//	if err != nil {
			//		fmt.Printf("Error eliminando: %v", err)
			//	} else {
			//		fmt.Println("Eliminado correctamente")
			//	}
			//}
		}
	}
}

//------------- funcioens tarifas ------------------
func tarifa() int {
	//var seccion int

	mostrarTarifas()
	tarifa := seleccionTarifa()
	return tarifa
}

func mostrarTarifas() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("----- Tarifas -----")
	fmt.Println("Secciones:")
	fmt.Println("1: $ 5.000")
	fmt.Println("2: $ 8.000")
	fmt.Println("3: $ 15.000")
	fmt.Println("")
	fmt.Println("")

}

func seleccionTarifa() int {
	var seccion int
	var tarifa int
	//var opcion int

	for seccion != 1 && seccion != 2 && seccion != 3 {
		fmt.Print("Ingresar opcion: ")

		fmt.Scanln(&seccion) //guarda la seccion

		switch seccion {
		case 1:
			tarifa = 5000
		case 2:
			tarifa = 8000
		case 3:
			tarifa = 15000
		}
	}
	return tarifa
}

func confirmacion() bool {

	fmt.Print("Desea confirmar la operacion? s/n: ")
	var opcion string
	var estado_compra bool

	fmt.Scanln(&opcion) //guarda la seccion

	switch opcion {
	case "s":
		estado_compra = true
	case "n":
		estado_compra = false

	}
	return estado_compra
}

func guardarCompra() {
	//falta implementar
}

func leyendaCompra(estado_compra bool) {
	fmt.Println("")
	fmt.Println("----------------------------------------------------")
	if estado_compra == false {
		fmt.Println("******* La compra no fue realizada. Gracias  ******** ")
	} else {
		guardarCompra()
		fmt.Println(" *****   Felicitades!! Su compra se realizo exitosamente!  ********** ")
	}
	fmt.Println("----------------------------------------------------")
	fmt.Println("")
}

//-----------------------------------------------------

//func eliminar(cliente Cliente) error {
//	err := obtenerBaseDeDatos() /*Aca deberiaoms obtener la DB*/
//	db := obtenerBaseDeDatos()
//	if err != nil {
//		return err
//	}
//	defer db.Close() /*Cerrar la conexion de la base de datos*/
//
//	sentenciaAEjecutar, err := db.Prepare("DELETE FROM Compradores WHERE ID = ?") /*Comando de SQL*/
//	if err != nil {
//		return err
//	}
//	defer sentenciaAEjecutar.Close()
//
//	_, err = sentenciaAEjecutar.Exec(cliente.Id)
//	if err != nil {
//		return err
//	}
//	return nil
//
//
//func insertar(cliente Cliente) (err error) {
//	db, err := obtenerBaseDeDatos()
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	// Preparamos la setencia para la base de datos
//	sentenciaAEjecutar, err := db.Prepare("INSERT INTO Compradores (Nombre, Id, Compra) VALUES(?, ?, ?)")
//	if err != nil {
//		return err
//	}
//	defer sentenciaAEjecutar.Close()
//	// Ejecutar sentencia, un valor por cada '?'
//	_, err = sentenciaAEjecutar.Exec(cliente.Nombre, cliente.Id, cliente.Compra)
//	if err != nil {
//		return err
//	}
//	return nil
//
//
//func obtenerClientes() ([]Cliente, error) { /*Depende si usamos el struct o la clase de CLIENTE*/
//	clientes := []Cliente{} //Areglo de clientes
//	db, err := obtenerBaseDeDatos()
//	if err != nil {
//		return nil, err
//	}
//	defer db.Close()
//	filas, err := db.Query("SELECT Id, Nombre, Compra FROM Compradores")
//
//	if err != nil {
//		return nil, err
//	}
//
//	defer filas.Close()
//
//	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
//	var cliente Cliente
//
//	// Recorrer todas las filas
//	for filas.Next() {
//		err = filas.Scan(&cliente.Id, &cliente.Nombre, &cliente.Compra)
//		// Verificamos si tenemos algun error
//		if err != nil {
//			return nil, err
//		}
//
//		clientes = append(clientes, cliente)
//	}
//
//	return clientes, nil
//
//
//func actualizar(cliente Cliente) error {
//	db, err := obtenerBaseDeDatos()
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	/*Volvemos al problema del ID. Yo aca lo plantie que si actualizas una compra el Id se
//	mantiene. Creo que sería lo mejor para no generar uno nuevo*/
//	sentenciaAEjecutar, err := db.Prepare("UPDATE Compradores SET Nombre = ?, Compra = ? WHERE Id = ?")
//	if err != nil {
//		return err
//	}
//	defer sentenciaAEjecutar.Close()
//	// Pasar argumentos en el mismo orden que la consulta
//	_, err = sentenciaAEjecutar.Exec(cliente.Nombre, cliente.Id, cliente.Compra)
//	return err
//
//
//
