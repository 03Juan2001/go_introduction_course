package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("Type: %T, bytes %d, bits %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	// Definir un array
	// var nombre [tamaño]tipo
	// El tamaño es parte del tipo
	// El tamaño debe ser una constante entera positiva
	var myArrayVar [6]int
	fmt.Println(myArrayVar)
	fmt.Printf("Type: %T, bytes %d, bits %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)	

	// Definir e inicializar un array
	myArrayVar1 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar1)

	// Asignar valores a indices específicos
	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9

	fmt.Println("Size: ", len(myArrayVar))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	/*
		Slices
		- Son una abstracción sobre arrays
		- Permiten trabajar con subarrays
		- Son dinámicos (pueden crecer y decrecer)
		- Se definen con la sintaxis: var nombre []tipo
		- El tamaño no es parte del tipo
		- No se pueden asignar valores a índices específicos
	*/
	var mySliceVar []int
	fmt.Println(mySliceVar)
	fmt.Printf("Size: %d, value: %v\n", len(mySliceVar), mySliceVar) // len() devuelve el número de elementos del slice

	mySliceVar = append(mySliceVar, 10, 20, 30, 40, 50) // append() añade un elemento al final del slice
	fmt.Printf("Size: %d, value: %v\n", len(mySliceVar), mySliceVar)

	mySlice := myArrayVar[2:4] // Crear un slice a partir de un array (slice del índice 2 al 4, sin incluir el 4)
	fmt.Println(mySlice)
	fmt.Printf("Size: %d, value: %v\n", len(mySlice), mySlice) // len() devuelve el número de elementos del slice

	fmt.Println(&myArrayVar[2]) // Dirección del elemento 2 del array
	fmt.Println(&mySlice[0])    // Dirección del elemento 0 del slice (apunta al mismo lugar que el elemento 2 del array)

	fmt.Println(myArrayVar) // Array original

	// Modificar el slice
	mySlice[0] = 80
	mySlice[1] = 90
	fmt.Println(myArrayVar)
	fmt.Println(myArrayVar[:4]) // Array original hasta el índice 4 (sin incluir el 4)
	fmt.Println(myArrayVar[2:]) // Array original desde el índice 2 (incluyendo el 2)

	slice := make([]int, 3) // Crear un slice con make(tipo, tamaño, capacidad opcional)
	fmt.Println(slice)
	fmt.Printf("Size: %d, value: %v\n", len(slice), slice)

	/*
		Resumen:
		- Los Arrays son de tamaño fijo y el tamaño es parte del tipo
		- Los Slices son dinámicos y no tienen tamaño fijo
		- Los Slices son una abstracción sobre Arrays y permiten trabajar con subarrays
		- Los Slices se crean a partir de Arrays o con la función make()

		- Make: Se utiliza para crear slices, mapas y canales. Para slices, make([]T, length, capacity) crea un slice del tipo T con una longitud y capacidad específicas.
		- Append: Se utiliza para añadir elementos a un slice. La función append(slice, elements...) devuelve un nuevo slice con los elementos añadidos.
	*/
}