package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Data struct {
	User User `json:"data"`
}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Year    int    `json:"year"`
	Color   string `json:"color"`
	Pantone string `json:"pantone_value"`
}

const (
	baseURL = "https://reqres.in/api"
)

func main() {
	// Realiza una solicitud GET a una API pública
	req, err := GetReqExample(fmt.Sprintf("%s/users/2", baseURL))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(req))

	fmt.Println("----- Otra forma de hacer la solicitud GET -----")
	// Otra forma de hacer una solicitud GET
	req2, err := Get(fmt.Sprintf("%s/users/2", baseURL))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(req2))

	fmt.Println("----- Obtener un usuario -----")
	// Obtener un usuario específico
	user, err := GetUser("2")
	if err != nil {
		log.Fatal(err)
	}
	// Imprime los detalles del usuario obtenido
	fmt.Printf("Usuario obtenido: %+v\n", user)
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Año: %d\n", user.Year)
	fmt.Printf("Nombre completo: %s\n", user.Name)
	fmt.Printf("Color: %s\n", user.Color)
	fmt.Printf("Pantone: %s\n", user.Pantone)
}

/* Ejemplo básico de una solicitud GET */
func GetReqExample(url string) ([]byte, error) {
	resp, err := http.Get(url) // Realiza una solicitud GET
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()            // Asegura que el cuerpo de la respuesta se cierre
	body, err := io.ReadAll(resp.Body) // Lee el cuerpo de la respuesta
	if err != nil {
		return nil, err
	}

	return body, nil
}

/* Ejemplo avanzado de una solicitud GET */
func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil) // Crea una nueva solicitud GET
	if err != nil {
		return nil, err
	}

	clien := &http.Client{}    // Crea un cliente HTTP
	resp, err := clien.Do(req) // Realiza la solicitud
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

/* Obtener un usuario específico de la API */
func GetUser(userID string) (*User, error) {
	req, err := Get(fmt.Sprintf("%s/users/%s", baseURL, userID))
	if err != nil {
		return nil, err
	}
	var dataResponse Data

	if err := json.Unmarshal(req, &dataResponse); err != nil {
		return nil, err
	}

	return &dataResponse.User, nil
}

