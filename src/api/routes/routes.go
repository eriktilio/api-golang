package routes

import (
	"api-golang/src/api/controllers"

	"github.com/gorilla/mux"
)

// NewRouter cria e retorna um novo roteador.
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")

	return r
}
