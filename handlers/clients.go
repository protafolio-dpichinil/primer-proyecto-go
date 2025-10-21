package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// ClientRouter es un manejador que enruta las peticiones de /clients a otros manejadores.
type ClientRouter struct {
	routes map[string]http.HandlerFunc
}

// NewClientRouter crea y configura un nuevo router para los clientes.
func NewClientRouter() *ClientRouter {
	router := &ClientRouter{
		routes: make(map[string]http.HandlerFunc),
	}
	router.routes["list"] = clientListHandler
	router.routes["get"] = clientGetHandler

	return router
}

// ServeHTTP hace que ClientRouter implemente la interfaz http.Handler.
func (cr *ClientRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/client/")

	if handler, ok := cr.routes[path]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func clientListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Respuesta desde /client/list: Aquí tienes la lista de todos los clientes.")
}

func clientGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Respuesta desde /client/get: Aquí tienes los datos de un cliente específico.")
}
