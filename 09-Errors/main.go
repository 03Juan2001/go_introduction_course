package main

import (
	"errors"
	"fmt"
	"go_introduction_course/09-Errors/operator"
)

func main() {

	var err error
	err = errors.New("this is an error")
	fmt.Println(err)         // Imprime el error
	fmt.Println(err.Error()) // Devuelve en forma de string el error

	err2 := fmt.Errorf("this is another error: %d", 42) // Crear un error con formato
	fmt.Println(err2)                                   // Imprime el error
	fmt.Println(err2.Error())                           // Devuelve en forma de string el error

	defer func() {
		fmt.Println("In main defer")
		if r := recover(); r != nil {
			fmt.Println("There arent any results", r)
			fmt.Println("Recovered in", r)
		}
	}()
	fmt.Println(err2)

	Z := operator.Div(3, 4) // Llamada a la función Div que puede causar un pánico

	fmt.Println("Result:")
	fmt.Println("Z is:", Z)
}
