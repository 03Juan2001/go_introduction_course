package main

import (
	"fmt"
	"go_introduction_course/06-Function/function" // Course es el nombre del módulo (go.mod)
)

func main() {
	// Usar funciones del paquete function importado
	fmt.Println(function.Add(3, 5))
	function.RepeatString(3, "Juan")
	fmt.Println()

	fmt.Println("----- Calculator -----")
	v, err := function.Calc(function.SUM, 3, 6)
	// Basicamente en Go no existen las excepciones, por lo que los errores se manejan devolviendo un valor de error
	// y comprobando si es nil o no
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Sum: ", v)
	}

	v, err = function.Calc(function.SUBTRACTION, 3, 6)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Subtraction: ", v)
	}

	v, err = function.Calc(function.DIVISION, 3, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Division: ", v)
	}

	v, err = function.Calc(function.MULTIPLICATION, 3, 6)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Multiplication: ", v)
	}
	fmt.Println()

	fmt.Println("----- Split -----")
	// Funciones con múltiples valores de retorno
	// La función Split divide un entero en dos partes y las devuelve
	// Se pueden capturar ambos valores de retorno o solo uno de ellos usando el identificador _
	x, y := function.Split(20)
	fmt.Printf("Split 20 in %d and %d\n", x, y) // fmt.Printf("Value1: ", x, " Value2: ", y)
	fmt.Println()

	fmt.Println("----- Variadic Functions -----")
	// Funciones con número variable de argumentos
	// La función MSum suma todos los números que se le pasan como argumentos
	// Se define con ...float64, lo que indica que puede recibir un número variable de argumentos de tipo float64
	v = function.MSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("MSum: ", v)
	fmt.Println()

	// Función MOperation que realiza una operación (suma, resta, multiplicación, división) sobre un número variable de argumentos
	// y devuelve el resultado y un error si ocurre alguno
	fmt.Println("----- MOperation -----")
	v, err = function.MOperation(function.SUM, 2, 7, 1)
	fmt.Println("MOperation SUM: ", v, " Error: ", err)

	v, err = function.MOperation(function.SUBTRACTION, 2, 1, 3, 2, 1)
	fmt.Println("MOperation SUBTRACTION: ", v, " Error: ", err)

	v, err = function.MOperation(function.MULTIPLICATION, 2, 1, 3, 2, 1)
	fmt.Println("MOperation MULTIPLICATION: ", v, " Error: ", err)

	v, err = function.MOperation(function.DIVISION, 2, 1, 0, 2, 1)
	fmt.Println("MOperation DIVISION: ", v, " Error: ", err)
	fmt.Println()

	fmt.Println("----- Factory -----")
	fn := function.FactoryOperation(function.SUM)
	fmt.Println("Factory SUM: ", fn(3, 4))

	fn = function.FactoryOperation(function.SUBTRACTION)
	fmt.Println("Factory SUBTRACTION: ", fn(3, 4))

	fn = function.FactoryOperation(function.MULTIPLICATION)
	fmt.Println("Factory MULTIPLICATION: ", fn(3, 4))

	fn = function.FactoryOperation(function.DIVISION)
	fmt.Println("Factory DIVISION: ", fn(3, 4))
}
