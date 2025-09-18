package main

import "fmt"

/*
	Se debe ir ingresando valores por consola hasta que se agrega
	cero y finaliza el programa.

	Todos los valores ingresados por consola se deben agregar a un array
	e imprimirlo por pantalla al finalizar
*/

func main() {
	// Definir
	array := []int{}
	var num int

	// Realizar la funcionalidad
	// Como lo realicee yo!!
	for i := 0; ; i++ {
		fmt.Scan(&num) // &: Es para indicar que es una referencia
		if num == 0 {
			break
		}
		array = append(array, num)
	}

	// Como lo realizo el instructor!!
	var array1 []string
	fmt.Println("Seleccione valores, se sale con cero")

	for {
		var value string
		fmt.Scanf("%s", &value)

		if value == "0" {
			break
		}

		array1 = append(array1, value)
	}

	fmt.Println("Los valores del array son: ", array)
	fmt.Println("Los valores del array1 son: ", array1)
}
