package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(-42) // Convierte un entero a una cadena
	fmt.Println(s)

	s = strconv.FormatBool(true) // Convierte un booleano a una cadena
	fmt.Println(s)
	s = strconv.FormatFloat(3.14159, 'f', 2, 64) // Convierte un float a una cadena con formato
	fmt.Println(s)
	s = strconv.FormatInt(12345, 10) // Convierte un entero a una cadena en base 10
	fmt.Println(s)
	s = strconv.FormatUint(12345, 10) // Convierte un entero sin signo a una cadena en base 10
	fmt.Println(s)

	b, err := strconv.ParseBool("sarasa") // Convierte una cadena a un booleano
	fmt.Println(b, err)

	f, err := strconv.ParseFloat("3.14159", 64) // Convierte una cadena a un float
	fmt.Println(f, err)

	in, err := strconv.ParseInt("-42", 10, 64) // Convierte una cadena a un entero en base 10
	fmt.Println(in, err)

	u, err := strconv.ParseUint("42", 10, 64) // Convierte una cadena a un entero sin signo en base 10
	fmt.Println(u, err)

}
