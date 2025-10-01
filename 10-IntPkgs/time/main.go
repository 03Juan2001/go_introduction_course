package main

import (
	"fmt"
	"time"
)

/*
	El paquete time proporciona funcionalidad para medir y mostrar el tiempo.
	Algunas de las características más comunes incluyen:
	- Obtener la hora actual
	- Formatear fechas y horas
	- Medir intervalos de tiempo

	-Metódos comunes:
		time.Now(): Devuelve la hora actual.
		time.Date(): Crea una fecha y hora específica.
		time.Sleep(): Pausa la ejecución durante un período de tiempo especificado.
		time.Parse(): Analiza una cadena de texto y la convierte en un valor de tiempo.
		time.Format(): Formatea un valor de tiempo en una cadena de texto según un formato especificado.
		time.Add(): Añade una duración a un valor de tiempo.
*/

func main() {

	p := fmt.Println
	now := time.Now()

	p(now) // Imprime la fecha y hora actual

	then := time.Date(2020, 11, 10, 23, 0, 0, 0, time.UTC)
	p(then) // Imprime una fecha y hora específica

	then = then.Add(time.Hour * 1) // Añade 1 hora a la fecha "then"
	p(then)                        // Imprime la nueva fecha y hora

	p(then.Hour())   // Imprime la hora de la fecha "then"
	p(then.Minute()) // Imprime los minutos de la fecha "then"
	p(then.Second()) // Imprime los segundos de la fecha "then"
	p(then.Year())   // Imprime el año de la fecha "then"
	p(then.Month())
	p(then.Day())      // Imprime el día del mes de la fecha "then"
	p(then.Location()) // Imprime la zona horaria de la fecha "then"
	p(then.Weekday())  // Imprime el día de la semana de la fecha "then"

	p(then.Before(now)) // Comprueba si "then" es antes que "now"
	p(then.After(now))  // Comprueba si "then" es después que "now"
	p(then.Equal(now))  // Comprueba si "then" es igual a "now"

	diff := now.Sub(then)       // Calcula la diferencia entre "now" y "then"
	p(diff)                     // Imprime la diferencia de tiempo
	p(diff.Hours())             // Imprime la diferencia en horas
	p(diff.Minutes())           // Imprime la diferencia en minutos
	p(diff.Seconds())           // Imprime la diferencia en segundos
	p(diff.Nanoseconds())       // Imprime la diferencia en nanosegundos
	p(diff.Milliseconds())      // Imprime la diferencia en milisegundos
	p(diff.String())            // Imprime la diferencia en formato string
	p(diff.Truncate(time.Hour)) // Trunca la diferencia al valor más cercano de la hora

	p(now.Format(time.RFC3339))
}
