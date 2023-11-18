package controllers

import (
	"fmt"
	"net/http"
)

// HomeHandler manipula a rota "/".
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo à minha API Go!")
}
