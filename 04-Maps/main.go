package main

import (
	"fmt"
)

func main() {
	/*
		Maps:
		- Colección de pares clave-valor
		- Las claves son únicas
		- No tienen un orden específico
	*/

	m1 := make(map[int]string) // Crear un map con claves de tipo int y valores de tipo string
	// Otra forma de crear un map
	// m1 := map[int]string{}

	// Añadir elementos al map
	m1[1] = "A"
	m1[2] = "B"
	m1[3] = "C"

	fmt.Println(m1)
	fmt.Println(m1[2]) // Acceder al valor de la clave 2

	// Eliminar un par clave-valor
	delete(m1, 2) // Eliminar el par clave-valor con clave 2
	fmt.Println(m1)

	// Modificar el valor de una clave
	m1[1] = "A2"
	fmt.Println(m1[1])

	m1[7] = ""
	fmt.Println(m1[7])  // Devuelve el valor vacío
	fmt.Println(m1[99]) // Devuelve el valor vacío

	// Comprobar si una clave existe
	v, ok := m1[7] // ok es true si la clave existe, false si no existe
	fmt.Println(v, ok)

	_, ok = m1[99] // ok es true si la clave existe, false si no existe
	fmt.Println(v, ok)

	// Instanciar los mapas manera 2
	m2 := map[int]string{
		4: "D",
		5: "E",
		6: "F",
	}

	fmt.Println(m2)

	/*
		- Resumen: 
		- Los Maps son colecciones de pares clave-valor
		- Las claves son únicas y no tienen un orden específico
		- Se crean con la función make() o con la sintaxis literal
		- Se accede a los valores mediante las claves
		- Se pueden añadir, modificar y eliminar pares clave-valor
		- Se puede comprobar si una clave existe
		- Sirven para representar estructuras de datos como diccionarios, tablas hash, etc.
	*/
}
