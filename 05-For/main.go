package main

import (
	"fmt"
)

func main() {
	/*
		For:
		- Estructura de control para repetir un bloque de código
		- Se puede usar con diferentes condiciones
		- Se puede usar con mapas y slices
	*/

	// Declarar e inicializar una variable
	sum := 0

	// Bucle For: Contar hasta 10
	fmt.Println("Bucle For: Contar hasta 10")
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)

	sum = 1

	// Bucle For: Detenerse cuando sum sea menor a 1000
	fmt.Println("Bucle For: Detenerse cuando sum sea menor a 1000")
	for sum < 1000 {
		sum++
	}
	fmt.Println(sum)

	sum = 0
	// Bucle For: Bucle infinito con break
	fmt.Println("Bucle For: Bucle infinito con break")
	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	// Declarar e inicializar un array
	arr := []int{1, 2, 3, 1, 2, 3}

	// Recorrer un array con for range
	// range devuelve el índice y el valor
	fmt.Println("Recorrer un array con for range")
	for i := range arr {
		fmt.Println("Index: ", i, "Value: ", arr[i])
	}

	fmt.Println()

	// Recorrer un array con for range (ignorar el índice)
	fmt.Println("Recorrer un array con for range (ignorar el índice)")
	for _, v := range arr {
		fmt.Println("Value: ", v)
	}

	fmt.Println()

	// Recorrer un mapa con for range
	maps2 := map[string]float64{
		"A": 12.3,
		"B": 23.1,
		"C": 34,
	}

	// range devuelve la clave y el valor
	fmt.Println("Recorrer un mapa con for range")
	for key, value := range maps2 {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	fmt.Println()

	// Recorrer mapa con array
	fmt.Println("Recorrer mapa con array")
	map3 := map[string][]int{
		"A": nil,
		"B": {2, 34, 1, 2, 4},
		"C": {4, 5, 3, 2, 1},
	}

	// Ciclo For anidado
	// El primer range devuelve la clave y el slice
	for key, value := range map3 {
		fmt.Println("Key: ", key)
		// El segundo range recorre el slice
		for i := range value {
			fmt.Println("Value: ", value[i])
		}
		fmt.Println()
	}
	fmt.Println()

	/*
		Resumen:
		- El bucle For es la única estructura de control de bucle en Go
		- Se puede usar con diferentes condiciones
		- Se puede usar con mapas y slices
		- El bucle For puede ser infinito si no se especifica una condición de salida
		- El bucle For puede tener múltiples variables de control
		- El bucle For puede tener múltiples condiciones de salida
		- El bucle For puede tener un bloque de código vacío
		- El bucle For puede tener un bloque de código vacío y una condición de salida
		- El bucle For puede tener un bloque de código vacío y múltiples condiciones de salida
		- Range: Se usa para recorrer arrays, slices, mapas y strings
		- Range devuelve dos valores: el índice y el valor
		- Se puede ignorar el índice usando el identificador de espacio en blanco _
	*/

}
