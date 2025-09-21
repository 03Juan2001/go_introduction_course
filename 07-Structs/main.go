package main

import (
	"encoding/json"
	"fmt"
	"go_introduction_course/07-Structs/structsInterface/structs"
	"go_introduction_course/07-Structs/structsInterface/vehicles"
)

func main() {
	// Uso de structs
	// Una struct es una colección de campos
	// Se accede a los campos con el operador punto (.)
	var p1 structs.Product

	// Asignar valores a los campos
	p1.ID = 1
	p1.Name = "Laptop"
	p1.Type = structs.Type{ID: 1, Code: "ELEC", Description: "Electronics"}
	p1.Price = 999.99
	p1.Tags = []string{"computer", "portable", "technology"}

	// Otra forma de crear una struct es declarando una variable e inicializando los campos uno por uno
	p2 := structs.Product{}
	p2.ID = 2
	p2.Name = "Smartphone"
	p2.Type = structs.Type{ID: 1, Code: "ELEC", Description: "Electronics"}
	p2.Price = 499.99
	p2.Tags = []string{"mobile", "touchscreen", "technology"}

	// Otra forma de crear una struct es usando una literal de struct
	// Los valores se asignan en el orden en que se definen los campos en la struct
	// Si no se especifica un campo, se asigna el valor cero del tipo (0 para int, "" para string, 0.0 para float, etc.)
	p3 := structs.Product{3, "Smartwatch", structs.Type{}, 0, 299.99, nil}

	// Otra forma de crear una struct es usando una literal de struct con nombres de campo
	// Esto permite asignar los campos en cualquier orden y omitir campos si se desea
	p4 := structs.Product{
		ID:    4,
		Count: 5,
		Name:  "Tablet",
		// Asignación por nombre de campo
		Type: structs.Type{
			ID:          1,
			Code:        "ELEC",
			Description: "Electronics",
		},
		Price: 199.99,
		Tags:  []string{"portable", "touchscreen"},
	}

	// Serializar struct a JSON
	// La serialización convierte la struct en una cadena JSON
	// Se usa la función json.Marshal que devuelve un slice de bytes y un error
	// Se pueden usar las etiquetas definidas en los campos para personalizar la serialización
	v, err := json.Marshal(p4)
	fmt.Println(string(v), err)

	// Imprimir structs
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)

	// Llamar a un método de la struct
	fmt.Println("Total Price:", p4.TotalPrice())
	p4.SetName("New Tablet")
	fmt.Println("Updated Name:", p4.Name)
	p4.AddTag("new", "sale")
	fmt.Println("Updated Tags:", p4.Tags)
	fmt.Println()

	// Uso de struct anidada
	fmt.Println("--- Struct Anidada ---")

	fmt.Println("Productos:")
	p5 := structs.Product{
		ID:    5,
		Name:  "Headphones",
		Type:  structs.Type{ID: 2, Code: "AUDIO", Description: "Audio"},
		Count: 2,
		Price: 149.99,
		Tags:  []string{"music", "sound"},
	}
	fmt.Println(p5)
	fmt.Println("Total Price:", p5.TotalPrice())
	fmt.Println()

	p6 := structs.Product{
		ID:    6,
		Name:  "Camera",
		Type:  structs.Type{ID: 3, Code: "PHOTO", Description: "Photography"},
		Count: 1,
		Price: 799.99,
		Tags:  []string{"camera", "photography"},
	}
	fmt.Println(p6)
	fmt.Println("Total Price:", p6.TotalPrice())
	fmt.Println()

	// Uso de struct Car
	car := structs.NewCar(12341)
	car.AddProduct(p5, p6)
	fmt.Println(car)
	fmt.Println(("Total products:"), len(car.Products))
	fmt.Println("Total Price:", car.TotalPrice())
	fmt.Println()

	// Nota: En Go, las structs son tipos de valor, lo que significa que cuando se asignan a una nueva variable o se pasan a una función,
	// se crea una copia de la struct. Para evitar copias innecesarias, se pueden usar punteros a structs.

	fmt.Println("--- Interfaces ---")
	// Una interface es un conjunto de métodos que un tipo debe implementar
	// Una struct puede implementar una interface si define todos los métodos de la interface
	// No es necesario declarar explícitamente que una struct implementa una interface
	// Esto permite un diseño más flexible y desacoplado
	fmt.Println()

	fmt.Println("Vehicles:")
	carV := vehicles.Car{Time: 120}
	fmt.Println(carV.Distance())

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "BUS"}

	var d float64
	for _, v := range vArray {
		fmt.Println("Vehicle:", v)
		vehicle, err := vehicles.NewVehicle(v, 400)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Distance: %.2f\n", vehicle.Distance())
		d += vehicle.Distance()
		fmt.Println()
	}
	fmt.Printf("Total Distance: %.2f\n", d)
}
