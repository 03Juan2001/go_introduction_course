package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
	El paquete log proporciona una forma sencilla de registrar mensajes de log en la consola o en archivos.
	Algunas de las características más comunes incluyen
	- Registro de mensajes con diferentes niveles de severidad (Info, Warning, Error)
	- Formateo de mensajes de log con información adicional (fecha, hora, archivo, línea)
	- Redirección de la salida de log a diferentes destinos (consola, archivo, red)
	-Metódos comunes:
		log.Print(): Registra un mensaje de log con nivel de severidad "Info".
		log.Printf(): Registra un mensaje de log con formato y nivel de severidad "Info".
		log.Println(): Registra un mensaje de log con nivel de severidad "Info" y añade un salto de línea.
		log.Fatal(): Registra un mensaje de log con nivel de severidad "Error" y termina el programa.
		log.Fatalf(): Registra un mensaje de log con formato y nivel de severidad "Error", y termina el programa.
		log.Panic(): Registra un mensaje de log con nivel de severidad "Error" y provoca un pánico.
		log.Panicf(): Registra un mensaje de log con formato y nivel de severidad "Error", y provoca un pánico.
		log.SetOutput(): Establece el destino de la salida de log (por defecto es la consola).
		log.SetFlags(): Establece las banderas para el formato de los mensajes de log (por defecto es la fecha y hora).
*/

func main() {

	// Crear o abrir un archivo para registrar los logs
	log.Println("Starting the application...")

	// log.Fatal("Fatal error occurred, terminating the application.") // Registra un mensaje de log y termina el programa

	// log.Panic("Panicking due to a severe error!") // Registra un mensaje de log y provoca un pánico

	log.Print("Application running smoothly.") // Registra un mensaje de log con nivel de severidad "Info"

	log.Print("Another info message.") // Registra otro mensaje de log con nivel de severidad "Info"
	date := time.Now()

	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano())) // Crear o abrir un archivo para registrar los logs
	if err != nil {
		panic(err)
	}

	l := log.New(file, "Success: ", log.LstdFlags|log.Lshortfile) // Crear un nuevo logger con el archivo como destino y un prefijo personalizado
	l.Println("Logging to a file...1")
	l.Println("Logging to a file...2")
	l.Println("Logging to a file...3")
	l.Println("Logging to a file...4")

	l.SetPrefix("Error: ") // Cambiar el prefijo del logger
	l.Println("Logging to a file...1")
	l.Println("Logging to a file...2")
	l.Println("Logging to a file...3")
	l.Println("Logging to a file...4")
}
