package main

import "fmt"
func main() {
	license := true
	age := 18

	if age >= 18 && license {
		fmt.Println("Puedes seguir avanzando")
	}else if license == false && age >= 18 || age < 18 && license == true {
		fmt.Println("No puedes seguir avanzando")
	}

}