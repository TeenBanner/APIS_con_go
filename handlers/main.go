package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/HolaMundo", HolaMundo)

	http.HandleFunc("/saludar", saludar)
	http.HandleFunc("/despedirse", despedirse)

	http.ListenAndServe(":8080", nil)
	fmt.Println("escuchando en http://127.0.0.1:8080")
	defer CerrarServidor()
}
