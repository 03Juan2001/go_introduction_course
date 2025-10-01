package main

import (
	"log"
	"os"
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

	file, err1 := os.Create("app.log") // Crear o abrir un archivo para registrar los logs
	if err1 != nil {
		log.Fatal("Error creating log file:", err1) // Registra un mensaje de log y termina el programa
	}
	defer file.Close() // Asegurarse de cerrar el archivo al final

	log.SetOutput(file) // Redirigir la salida de log al archivo
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Establecer el formato de los mensajes de log

}
