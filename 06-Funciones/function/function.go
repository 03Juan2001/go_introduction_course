package function

import (
	"errors"
	"fmt"
)

/**
 * Las funciones en Go se definen con la palabra clave "func"
 * El nombre de la función debe comenzar con una letra mayúscula si se quiere exportar (hacer pública)
 * Los parámetros se definen con el tipo después del nombre.
 * Se pueden definir múltiples parámetros del mismo tipo agrupándolos.
 * Se pueden definir múltiples valores de retorno.
 * Si una función no retorna ningún valor, no es necesario especificarlo.
 */
 
type Operation int // Definir un nuevo tipo llamado Operation basado en int

const (
	SUM Operation = iota // Iota es un contador que empieza en 0 y se incrementa en 1 en cada línea
	SUBTRACTION
	MULTIPLICATION
	DIVISION
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Printf("Hola %s, ¿cómo estás?\n", value)
	}
}

// Calc realiza una operación matemática básica entre dos números
// Devuelve el resultado y un booleano que indica si hubo un error (true = error, false = no error)
func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUBTRACTION:
		return x - y, nil
	case DIVISION:
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	case MULTIPLICATION:
		return x * y, nil
	default:
		return 0, errors.New("unknown operation")
	}
}
