package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Define el puerto en el que el servidor escuchará
	r := chi.NewRouter()

	// Configura los manejadores para las rutas
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got /users")
		io.WriteString(w, "This is my user endpoint!\n")
	})

	r.Get("/courses", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got /courses")
		io.WriteString(w, "This is my course endpoint!\n")
	})

	// Inicia el servidor HTTP
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

// Manejador para la ruta /users
// func getUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("got /users")
// 	io.WriteString(w, "This is my user endpoint!\n")
// }

// // Manejador para la ruta /courses
// func getCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("got /courses")
// 	io.WriteString(w, "This is my course endpoint!\n")
// }