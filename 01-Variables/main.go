package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		Declaracion de variables (Var).

		- Tipos de datos:
			- Int: Tipo de dato entero con signo.
			- uint: Tipo de dato entero sin signo.
			- string: Tipo de dato para cadenas de texto.
			- bool: Tipo de dato booleano (true/false).
	*/

	var myIntVar int // Declaracion de variable entera
	myIntVar = -42
	fmt.Println("My Integer Variable:", myIntVar)

	var myUintVar uint // Declaracion de variable entera sin signo
	myUintVar = 42
	fmt.Println("My Unsigned Integer Variable:", myUintVar)

	var myStringVar string // Declaracion de variable string
	myStringVar = "Hola, mundo!"
	fmt.Println("My String Variable:", myStringVar)

	var myBoolVar bool // Declaracion de variable booleana
	myBoolVar = true
	fmt.Println("My Boolean Variable:", myBoolVar)

	// Nota: El operador & se utiliza para obtener la direccion de memoria de la variable.
	fmt.Println("My variables are of types:", &myIntVar, &myUintVar, &myStringVar, &myBoolVar)

	// Delaracion de una varible en una sola linea.
	var myOtherIntVar int = 100
	fmt.Println("My Other Integer Variable:", myOtherIntVar)

	myOtherIntVar2 := 200 // Declaracion corta de variable (Solo dentro de funciones)
	fmt.Println("My Other Integer Variable 2:", myOtherIntVar2)

	/*
		Constantes (Const).

		- Las constantes son valores que no pueden cambiar durante la ejecucion del programa.
		- Se declaran con la palabra clave "const".
		- Las constantes deben ser inicializadas en el momento de su declaracion.
	*/
	fmt.Println("--------------------------- Constantes --------------------------------------")

	const myFistConst = "a122"
	fmt.Println("My const is: ", myFistConst)

	const myIntConst int = 12
	fmt.Println("My const is: ", myIntConst)

	const myStringConst string = "Jan"
	fmt.Println("My const is: ", myStringConst)

	fmt.Println("-----------------------------Espacios De Memoria--------------------------------------")
	fmt.Println()

	/*
		Metadatos:
			-%d: Para numeros enteros
			-%f: Para numeros flotantes
			-%s: Para los strings
			-%t: Para una tabulacion
			-%n: Para un salto d linea.
			-%T: Para mostrar tipo de datp
	*/

	fmt.Println("-----------------------------INT--------------------------------------")
	// Valores de memoria en lo que respecta a valorees entros (Si permite valores neegativos)
	var my8BitsIntVar int8
	fmt.Printf("Int defalut value: %d\n", my8BitsIntVar)
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)
	fmt.Println()

	fmt.Println("-----------------------------UINT--------------------------------------")
	// Valores de memoria en lo qu respecta a valorees UInt (No permite valors negativos)
	var my8BitsUIntVar uint8
	fmt.Printf("Uint default value: %d\n", my8BitsUIntVar)
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my8BitsUIntVar, unsafe.Sizeof(my8BitsUIntVar), unsafe.Sizeof(my8BitsUIntVar)*8)

	var my16BitsuIntVar uint16
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my16BitsuIntVar, unsafe.Sizeof(my16BitsuIntVar), unsafe.Sizeof(my16BitsuIntVar)*8)

	var my32BitsuIntVar uint32
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my32BitsuIntVar, unsafe.Sizeof(my32BitsuIntVar), unsafe.Sizeof(my32BitsuIntVar)*8)

	var my64BitsuIntVar uint64
	fmt.Printf("type: %T, byte: %d, bits: %d\n", my64BitsuIntVar, unsafe.Sizeof(my64BitsuIntVar), unsafe.Sizeof(my64BitsuIntVar)*8)
	fmt.Println()

	fmt.Println("-----------------------------Float--------------------------------------")
	// Valor de moemria float
	var myFloat32Var float32
	var myFloat64Var float64
	fmt.Printf("Float Default Value: %f\n", myFloat32Var)
	fmt.Printf("type: %T, byte: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)
	fmt.Printf("type: %T, byte: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)
	fmt.Println()

	fmt.Println("-----------------------------String--------------------------------------")
	var myStringVar23 string
	fmt.Printf("String default value: %s\n", myStringVar23)
	myStringVar45 := `My first program
	in
	gholang`
	fmt.Printf("The variable value is: %s\n", myStringVar45)

	{
		fmt.Println("-----------------------------Conversion de variables--------------------------------------")
		// Conversion de valor float a string
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatString := fmt.Sprintf("%f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatString, floatString)

		// Conversion de valor int
		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStringVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStringVar, intStringVar)
	}
}
