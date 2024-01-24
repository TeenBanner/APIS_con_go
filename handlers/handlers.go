package main

import (
	"fmt"
	"net/http"
)

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}

func despedirse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Adios")
}

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}
