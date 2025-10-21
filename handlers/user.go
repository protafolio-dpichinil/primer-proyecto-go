package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// UserRouter es un manejador que enruta las peticiones de /user a otros manejadores.
type UserRouter struct {
	routes map[string]http.HandlerFunc
}

// NewUserRouter crea y configura un nuevo router para los usuarios.
func NewUserRouter() *UserRouter {
	router := &UserRouter{
		routes: make(map[string]http.HandlerFunc),
	}
	// Aquí registramos las sub-rutas para los usuarios.
	router.routes["list"] = userListHandler
	router.routes["get"] = userGetHandler
	return router
}

// ServeHTTP hace que UserRouter implemente la interfaz http.Handler.
// Este método es el "controlador de tráfico" para todas las rutas bajo /user/.
func (ur *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/user/")

	// Buscamos el manejador correspondiente en nuestro mapa de rutas.
	if handler, ok := ur.routes[path]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func userListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Respuesta desde /user/list: Aquí deberá implementarse la lista de todos los usuarios.")
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Respuesta desde /user/get: Aquí deberá implementarse la devolucion de los datos de un usuario específico.")
}
