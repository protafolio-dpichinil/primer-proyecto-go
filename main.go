package main

import (
	"fmt"
	"log"
	"net/http"

	"primer-proyecto-go/handlers"
)

// handler requerido para el path "/"
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Nos aseguramos de que solo respondemos a peticiones GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "El servicio esta activo")
}

func main() {
	// Registramos nuestra función `helloHandler` para que maneje las peticiones a la ruta "/".
	//http.HandleFunc("/hello", helloHandler)
	//http.HandleFunc("/goodbye", goodbyeHandler)
	http.HandleFunc("/", healthHandler)

	// Creamos una instancia de nuestro router de usuarios.
	userRouter := handlers.NewUserRouter()
	// Todas las peticiones que comiencen con "/user/" serán manejadas por userRouter.
	http.Handle("/user/", userRouter)

	// Creamos una instancia de nuestro router de clientes.
	clientRouter := handlers.NewClientRouter()
	// Todas las peticiones que comiencen con "/client/" serán manejadas por clientRouter.
	http.Handle("/client/", clientRouter)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	// Iniciamos el servidor en el puerto 8080. log.Fatal se ejecutará si hay un error al iniciar.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
