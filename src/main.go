package main

import (
	"fmt"
	"net/http"

	"api-golang/src/api/routes"
)

func main() {
	r := routes.NewRouter()

	port := 8080
	fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
