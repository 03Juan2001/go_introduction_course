package main

import "fmt"

func main() {
	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println("Value of i:", i)                  // Valor de i es 34
	fmt.Println("Address of i:", &i)               // Dirección de memoria de i
	fmt.Println("Value of iP (address of i):", iP) // Valor de iP es la dirección de memoria de i
	fmt.Println("Value pointed by iP:", *iP)       // Valor apuntado por iP es 34

	*iP = 55
	fmt.Println("Updated value of i:", i) // Valor de i es 55

	myVar := 30
	fmt.Printf("addrs: %p, values: %d\n", &myVar, myVar)
	increment(myVar)
	fmt.Println(myVar)
	incrementP(&myVar)
	fmt.Println(myVar)
}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementP(val *int) {
	fmt.Println(val)
	*val++
}
