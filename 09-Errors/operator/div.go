package operator

import "fmt"

func Div(a, b float64) float64 {
	if b == 0 {
		panic(fmt.Sprintf("Division by zero: a=%f, b=%f", a, b))
	}
	return a / b
}
 