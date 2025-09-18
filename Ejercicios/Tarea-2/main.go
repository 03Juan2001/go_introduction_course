package main
 
import "fmt"
 
 
func main() {
 
    array := [5]int{4, 2, 5, 6, 7}
    
    // Realizar la funcionalidad
	for i := range array { // Recorrer el array
		// array[i] += 20	   // Otra forma de recorrer el array
		array[i] = array[i] + 20 // Sumar 20 a cada elemento
	}
 
    fmt.Println("Los valores del array son: ", array)
}