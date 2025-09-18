package main

import "fmt"

/*
	Se debe ir ingresando valores por consola hasta que se agrega
	cero y finaliza el programa, todos los valores deben ser numericos.

	En caso de agregar un valor que no sea numerico el programa lo ignorara y seguira
	pidiendo valores normalmente

	Todos los valores ingresados por consola se deben agregar a un array de tipo
	numerico e imprimirlo por pantalla al finalizar
*/

func main() {
	// Definir el map
	myArray := make(map[int]string)

	// Agregar valores al map key-value
	myArray[10] = "notebook"
	myArray[15] = "tv"
	myArray[21] = "heladera"
	myArray[27] = "monitor"
	myArray[35] = "camara"

	var key int // Variable para almacenar la entrada del usuario
	var values []string // Slice para almacenar los valores encontrados

	// Bucle para pedir entradas al usuario
	for {
		fmt.Print("Ingrese un código (o 0 para salir): ")
		fmt.Scanf("%d", &key) // Leer la entrada del usuario

		// Salir del bucle si el usuario ingresa 0
		if key == 0 {
			break
		}
		// Buscar el valor en el map y agregarlo al slice
		if value, exists := myArray[key]; exists {
			values = append(values, value) // Agregar el valor encontrado al slice
		} else {
			values = append(values, "No encontrado")
		}
	}


	fmt.Println("Los valores del array son:", values)
}