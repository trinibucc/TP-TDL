package main

import (
	"bufio"        // Leer líneas incluso si tienen espacios
	"database/sql" //Interactuar con bases de datos
	"fmt"          //Para los mensajes
	_ "mysql"      // La librería que nos permite conectar a MySQL. Descargada de github.com/go-sql-driver/mysql (Yo clone el repo)
	"os"           // El búfer, para leer desde la terminal con os.Stdin
	"strconv"      //Para conversión
)

type Cliente struct {
	Nombre     string
	Id, Compra int
}

func main() {

	mostrar := make(chan string)
	menu := `¿Qué deseas hacer?
[1] -- Agregar Compra
[2] -- Mostrar Tabla -> Verificar. 
[3] -- Actualizar Compra ya realizada /*Cambiar monto o titular de la operacion */
[4] -- Eliminar Operacion
[5] -- Salir
----->	`
	var opcion int
	var cliente Cliente
	for opcion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&opcion)
		scanner := bufio.NewScanner(os.Stdin)
		switch opcion {
		case 1:
			fmt.Println("Ingresa el nombre:")
			if scanner.Scan() {
				cliente.Nombre = scanner.Text()
			}
			fmt.Println("Ingrese el monto de la compra:")
			if scanner.Scan() {
				var compra = scanner.Text()
				cliente.Compra, _ = strconv.Atoi(compra) //La conversion puede dar un error y por eso el "_"
			}
			/*¿El ID debería ser generado de manera random? -> Id consecutivos*/
			err := insertar(cliente) //err = error
			if err != nil {
				fmt.Printf("Error insertando: %v", err)
			} else {
				fmt.Println("Insertado correctamente")
			}
		case 2:
			go mostrarClientes(mostrar) //Le paso el channel
			/*clientes, err := obtenerClientes()
			if err != nil {
				fmt.Printf("Error obteniendo contactos: %v", err)
			} else {

				for _, cliente := range clientes {
					fmt.Println("====================")
					fmt.Printf("Nombre: %s\n", cliente.Nombre)
					fmt.Printf("Id: %d\n", cliente.Id)
					fmt.Printf("Compra: %d\n", cliente.Compra)
				}
			}*/
		case 3:
			/* ¿El ID al actualizar una compra lo mantenemos o le generamos uno nuevo?
			Opcion 1 fmt.Println("Ingresa el id:")
			Opcion 2 generarlo de manera random */
			fmt.Println("Ingrese el ID del comprador:")
			if scanner.Scan() {
				idString := scanner.Text()
				cliente.Id, _ = strconv.Atoi(idString)

			}
			fmt.Println("Ingresa el nuevo nombre:")
			if scanner.Scan() {
				cliente.Nombre = scanner.Text()
			}
			fmt.Println("Ingresa el nuevo valor de compra:")
			if scanner.Scan() {
				var compra = scanner.Text()
				cliente.Compra, _ = strconv.Atoi(compra)
			}
			err := actualizar(cliente)
			if err != nil {
				fmt.Printf("Error actualizando: %v", err)
			} else {
				fmt.Println("Actualizado correctamente")
			}
		case 4:
			fmt.Println("Ingresa el ID de la operacion que desea eliminar:")
			fmt.Scanln(&cliente.Id)
			err := eliminar(cliente)
			if err != nil {
				fmt.Printf("Error eliminando: %v", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		}
	}
}

func eliminar(cliente Cliente) error {
	db, err := obtenerBaseDeDatos() /*Aca deberiaoms obtener la DB*/
	if err != nil {
		return err
	}
	defer db.Close() /*Cerrar la conexion de la base de datos*/

	sentenciaAEjecutar, err := db.Prepare("DELETE FROM Compradores WHERE ID = ?") /*Comando de SQL*/
	if err != nil {
		return err
	}
	defer sentenciaAEjecutar.Close()

	_, err = sentenciaAEjecutar.Exec(cliente.Id)
	if err != nil {
		return err
	}
	return nil
}

func insertar(cliente Cliente) (err error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos la setencia para la base de datos
	sentenciaAEjecutar, err := db.Prepare("INSERT INTO Compradores (Nombre, Id, Compra) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer sentenciaAEjecutar.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaAEjecutar.Exec(cliente.Nombre, cliente.Id, cliente.Compra)
	if err != nil {
		return err
	}
	return nil
}

func mostrarClientes(mostrar chan<- string){
	clientes, err := obtenerClientes()
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
	}
	select{
	case result <- response:
		for _, cliente := range clientes {
					fmt.Println("====================")
					fmt.Printf("Nombre: %s\n", cliente.Nombre)
					fmt.Printf("Id: %d\n", cliente.Id)
					fmt.Printf("Compra: %d\n", cliente.Compra)
	}
}


func obtenerClientes() ([]Cliente, error) { //Depende si usamos el struct o la clase de CLIENTE
	clientes := []Cliente{} //Areglo de clientes
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT Id, Nombre, Compra FROM Compradores")

	if err != nil {
		return nil, err
	}

	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var cliente Cliente

	// Recorrer todas las filas
	for filas.Next() {
		err = filas.Scan(&cliente.Id, &cliente.Nombre, &cliente.Compra)
		// Verificamos si tenemos algun error
		if err != nil {
			return nil, err
		}

		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func actualizar(cliente Cliente) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	/*Volvemos al problema del ID. Yo aca lo plantie que si actualizas una compra el Id se
	mantiene. Creo que sería lo mejor para no generar uno nuevo*/
	sentenciaAEjecutar, err := db.Prepare("UPDATE Compradores SET Nombre = ?, Compra = ? WHERE Id = ?")
	if err != nil {
		return err
	}
	defer sentenciaAEjecutar.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaAEjecutar.Exec(cliente.Nombre, cliente.Compra, cliente.Id)
	return err

}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "Espectadores"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*La parte más importante es en donde hacemos la conexión. Ahí especificamos la contraseña, el host, el usuario y el nombre de la base de datos.
El host siempre es localhost, pero puede que sea otro, igualmente se puede cambiar el puerto.

Los errores más comunes que pueden aparecer:
- Cuando especificamos mal la IP o el puerto: Error conectando: dial tcp 127.0.0.1:3306: connectex: No se puede establecer una conexión ya que el equipo de destino denegó expresamente dicha conexión.
- Si no ponemos el usuario o la contraseña correcta: Error conectando: Error 1044: Access denied for user ”@’localhost’ to database ‘agenda’
- En caso de que no exista la base de datos: Error conectando: Error 1049: Unknown database ‘agenda’*/
