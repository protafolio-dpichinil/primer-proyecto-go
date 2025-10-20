package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler responde a las peticiones HTTP en la ruta que le asignemos.
// Recibe un http.ResponseWriter para escribir la respuesta y un http.Request con la información de la petición.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Nos aseguramos de que solo respondemos a peticiones GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "¡Hola desde el microservicio Go!")
}

func main() {
	// Registramos nuestra función `helloHandler` para que maneje las peticiones a la ruta "/".
	http.HandleFunc("/", helloHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	// Iniciamos el servidor en el puerto 8080. log.Fatal se ejecutará si hay un error al iniciar.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
