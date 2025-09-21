package vehicles

import "fmt"

// Definición de una interfaz Vehicle con un método Distance
type Vehicle interface {
	Distance() float64
}

// Definición de structs que implementan la interfaz Vehicle
const (
	CarVehicle        = "CAR"
	MotorcycleVehicle = "MOTORCYCLE"
	TruckVehicle      = "TRUCK"
	BusVehicle        = "BUS"
)

// Cada struct tiene un campo Time que representa el tiempo en minutos
type Car struct {
	Time int
}
type Motorcycle struct {
	Time int
}
type Truck struct {
	Time int
}
type Bus struct {
	Time int
}

// Implementación del método Distance para cada struct
// La distancia se calcula como velocidad * tiempo (en horas)
// Velocidades asumidas:
// Car: 100 km/h
// Motorcycle: 120 km/h
// Truck: 80 km/h
// Bus: 50 km/h
func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

func (m *Motorcycle) Distance() float64 {
	return 120 * (float64(m.Time) / 60)
}

func (t *Truck) Distance() float64 {
	return 80 * (float64(t.Time) / 60)
}

func (b *Bus) Distance() float64 {
	return 50 * (float64(b.Time) / 60)
}

// Función para crear un vehículo basado en el tipo
func NewVehicle(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	case BusVehicle:
		return &Bus{Time: time}, nil
	default:
		return nil, fmt.Errorf("tipo de vehículo desconocido: %s", v)
	}
}
