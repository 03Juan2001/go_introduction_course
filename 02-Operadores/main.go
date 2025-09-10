package main

import "fmt"

func main() {
	yearsOld := 32

	// Operadores
	fmt.Println("Operadores")
	fmt.Println(yearsOld > 30) // false
	fmt.Println(yearsOld < 32) // falso
	fmt.Println(yearsOld <= 32) // true
	fmt.Println(yearsOld >= 40) // false
	fmt.Println(yearsOld == 32) // true
	fmt.Println(yearsOld != 24) // true

	fmt.Println()
	// Operadores Lógicos
	fmt.Println("Operadores Lógicos")
	fmt.Println("AND")
	fmt.Println(yearsOld < 32 || yearsOld == 32) // true
	fmt.Println(yearsOld < 32 || yearsOld == 33) // true
	fmt.Println(yearsOld < 40 || yearsOld == 33) // true

	fmt.Println()
	fmt.Println("OR")
	fmt.Println(yearsOld < 32 && yearsOld == 32) // false
	fmt.Println(yearsOld < 32 && yearsOld == 33) // false
	fmt.Println(yearsOld < 40 && yearsOld == 32) // true

	fmt.Println()
	fmt.Println("NOT")
	fmt.Println(!(yearsOld < 32))	 // true
	fmt.Println(!(yearsOld > 32))	 // true
	fmt.Println(!(yearsOld == 32))	 // false

	// Condicionales
	fmt.Println()
	fmt.Println("Condicionales")
	yearsOld = 20
	if yearsOld > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOld)
	} else {
		fmt.Printf("%d is lower than 18\n", yearsOld)
	}

	boolValue := true
	if boolValue {
		fmt.Println("Bool value is true")
	} else {
		fmt.Println("Bool value is false")
	}

	if value := true; value {
		fmt.Println("Value is true")
	}

	number := 1
	if number == 1 {
		fmt.Println("Number is 1")
	} else if number == 2 {
		fmt.Println("Number is 2")
	} else if number == 3 {
		fmt.Println("Number is 3")
	}

	switch number := 2; number {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
	default:
		fmt.Println("Number is not 1, 2 or 3")
	}

	switch {
	case number > 0:
		fmt.Println("Number is positive")
	case number < 0:
		fmt.Println("Number is negative")
	default:
		fmt.Println("Number is zero")
	}	
}
