package function // Paquete function

// Al estar dentro del mismo paquete, no es necesario importar "fmt" ni "errors"
// ya que ya se importaron en function.go y se compilan juntos por que son del mismo paquete

// Patron de diseño Factory
// Una función que retorna otra función basada en un parámetro de entrada
func FactoryOperation(op Operation) func(float64, float64) float64 {
	switch op {
		case SUM:
			return sum
		case SUBTRACTION:
			return sub
		case MULTIPLICATION:
			return mul
		case DIVISION:
			return div
	}
	return nil
}

// Funciones para cada operación
func sum(x, y float64) float64 {
	return x + y
}

func sub(x, y float64) float64 {
	return x - y
}

func mul(x, y float64) float64 {
	return x * y
}

func div(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	return x / y
}