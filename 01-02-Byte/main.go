package main

import (
	"fmt"
)

func main() {
	// Ejemplo de uso de bytes
	var A byte = 'A' // byte es un alias para uint8
	var a byte = 'a'
	var R byte = 82
	var s byte = 115

	vector := []byte{65, 97, 82, 115}

	fmt.Println()
	fmt.Println("value in Ascii: ")
	fmt.Println(A)
	fmt.Println(a)
	fmt.Println(R)
	fmt.Println(s)

	fmt.Println(vector)

	fmt.Println()
	fmt.Println("value in String: ")
	fmt.Println(string(A))
	fmt.Println(string(a))
	fmt.Println(string(R))
	fmt.Println(string(s))
}