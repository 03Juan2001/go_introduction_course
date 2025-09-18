package main

import (
	"fmt"
	"go_introduction_course/06-Funciones/function" // Importar el paquete function
)

func main() {
	// Llamar a la función Add del paquete function
	result := function.Add(2, 3) // Asignar el valor retornado a una variable
	fmt.Println("Result:", result)
	function.RepeatString(3, "Hello, Go!") // Llamar a la función RepeatString del paquete function

	fmt.Println("----- Calculator -----")
	// Usar la función Calc del paquete function
	v, err := function.Calc(function.SUM, 3, 6)
	fmt.Println("Sum: ", v, " Error: ", err)

	v, err = function.Calc(function.SUBTRACTION, 3, 6)
	fmt.Println("Subtraction: ", v, " Error: ", err)

	v, err = function.Calc(function.DIVISION, 3, 0)
	fmt.Println("Division: ", v, " Error: ", err)

	v, err = function.Calc(function.MULTIPLICATION, 3, 6)
	fmt.Println("Multiplication: ", v, " Error: ", err)

}


