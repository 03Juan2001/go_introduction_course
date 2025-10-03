package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	User User `json:"data"`
}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Year    int    `json:"year"`
	Color   string `json:"color"`
	Pantone time.Time `json:"pantone_value"`
}

type UserCreate struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Job  string `json:"job"`
	CreatedAt string `json:"createdAt"`
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
	user, _ := GetUser("2")

	// Imprime los detalles del usuario obtenido
	fmt.Printf("Usuario obtenido: %+v\n", user)
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Año: %d\n", user.Year)
	fmt.Printf("Nombre completo: %s\n", user.Name)
	fmt.Printf("Color: %s\n", user.Color)
	fmt.Printf("Pantone: %s\n", user.Pantone)

	fmt.Println("----- Crear un usuario -----")

	// Crear un nuevo usuario
	newUser, err := CreateUser("morpheus", "leader")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Nuevo usuario creado: %+v\n", newUser)
	fmt.Printf("ID: %s\n", newUser.ID)
	fmt.Printf("Nombre: %s\n", newUser.Name)
	fmt.Printf("Trabajo: %s\n", newUser.Job)
	fmt.Printf("Creado en: %s\n", newUser.CreatedAt)
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

func Post(url string, data interface{}) ([]byte, error) {
	b, err := json.Marshal(data) // Convierte los datos a JSON
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b)) // Crea una nueva solicitud POST
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json") // Establece el encabezado Content-Type

	client := &http.Client{}
	resp, err := client.Do(req) // Realiza la solicitud
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body) // Lee el cuerpo de la respuesta
	if err != nil {
		return nil, err
	}
	
	return responseBody, nil
}

func CreateUser(name, job string) (*UserCreate, error) {
	user := &UserCreate{
		Name:    name,
		Job:     job,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	req , err := Post(fmt.Sprintf("%s/users", baseURL), user)
	if err != nil {
		return nil, err
	}
	if  err := json.Unmarshal(req, &user); err != nil {
		return nil, err
	}

	return user, nil
}