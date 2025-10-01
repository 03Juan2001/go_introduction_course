package main

import (
	"fmt"
	"os"
)

/*
	El paquete os proporciona una plataforma independiente para interactuar con el sistema operativo.
	Algunas de las características más comunes incluyen:
	- Manipulación de archivos y directorios
	- Gestión de procesos
	- Manejo de variables de entorno
	- Entrada y salida estándar
	-Metódos comunes:
		os.Open(): Abre un archivo para lectura.
		os.Create(): Crea un archivo para escritura.
		os.Remove(): Elimina un archivo o directorio.
		os.Rename(): Cambia el nombre de un archivo o directorio.
		os.Stat(): Obtiene información sobre un archivo o directorio.
		os.Chdir(): Cambia el directorio de trabajo actual.
		os.Getwd(): Obtiene el directorio de trabajo actual.
		os.Getenv(): Obtiene el valor de una variable de entorno.
		os.Setenv(): Establece el valor de una variable de entorno.
		os.Exit(): Termina el programa con un código de estado específico.
*/

func main() {

	file, err := os.Open("miFile.txt") // Abre un archivo para lectura
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1) // Termina el programa con un código de estado 1 (error)
	}

	data := make([]byte, 100)     // Crea un slice de bytes para almacenar los datos leídos
	count, err := file.Read(data) // Lee hasta 100 bytes del archivo
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1) // Termina el programa con un código de estado 1 (error)
	}

	fmt.Printf("Read %d bytes: %q\n", count, data[:count]) // Imprime la cantidad de bytes leídos y su contenido

	// Acceder a variables de entorno
	path := os.Getenv("USERNAME") // Obtiene el valor de la variable de entorno USERNAME
	fmt.Println(path)             // Imprime el valor de la variable de entorno

	os.Setenv("MY_VAR", "my_value") // Establece una nueva variable de entorno
	value := os.Getenv("MY_VAR")    // Obtiene el valor de la nueva variable de entorno
	fmt.Println(value)              // Imprime el valor de la nueva variable de entorno

	dir, err := os.Getwd() // Obtiene el directorio de trabajo actual
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1) // Termina el programa con un código de estado 1 (error)
	}
	fmt.Println("Current directory:", dir)

	os.Setenv("ENV_EXIST", "")         // Establece una variable de entorno
	env1 := os.Getenv("ENV_EXIST")     // Obtiene el valor de la variable de entorno
	env2 := os.Getenv("ENV_NOT_EXIST") // Intenta obtener una variable de entorno que no existe

	fmt.Println("ENV_EXIST:", env1)     // Imprime el valor de la variable de entorno existente
	fmt.Println("ENV_NOT_EXIST:", env2) // Imprime el valor de la variable de entorno no existente (cadena vacía)

	env1, ok1 := os.LookupEnv("ENV_EXIST")     // Busca la variable de entorno y devuelve un booleano indicando si existe
	env2, ok2 := os.LookupEnv("ENV_NOT_EXIST") // Busca la variable de entorno que no existe

	fmt.Println("ENV_EXIST:", env1, ok1)     // Imprime el valor y el estado de la variable de entorno existente
	fmt.Println("ENV_NOT_EXIST:", env2, ok2) // Imprime el valor y el estado de la variable de entorno no existente

	os.Setenv("DB_USERNAME", "juan")      // Establece una variable de entorno para el nombre de usuario de la base de datos
	os.Setenv("DB_PASSWORD", "1234")      // Establece una variable de entorno para la contraseña de la base de datos
	os.Setenv("DB_HOST", "localhost")     // Establece una variable de entorno para el host de la base de datos
	os.Setenv("DB_PORT", "5432")          // Establece una variable de entorno para el puerto de la base de datos
	os.Setenv("DB_NAME", "mi_base_datos") // Establece una variable de entorno para el nombre de la base de datos

	// ${VAR} o $VAR
	// ${VAR:-default} o $VAR:-default
	// ${VAR:+alt} o $VAR:+alt
	// ${VAR:?err} o $VAR:?err
	dbURL := os.ExpandEnv("postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME") // Construye la URL de conexión a la base de datos utilizando las variables de entorno
	fmt.Println("Database URL:", dbURL)                                                      // Imprime la URL de conexión a la base de datos
}
