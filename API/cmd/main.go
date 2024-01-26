package main

import (
	"log"
	"net/http"
	"res-API/handler"
	"res-API/storage"
)

func main() {
	store := storage.NewMemory()

	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	log.Println("Servidor corriendo en http://127.0.0.1:3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Printf("error en el servidor %v\n", err)
	}
}
