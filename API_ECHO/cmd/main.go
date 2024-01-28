package main

import (
	//propios
	"log"

	//core
	"res-API/authorization"
	"res-API/handler"
	"res-API/storage"

	// \DEPENDENCIAS
	"github.com/labstack/echo/v4"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")

	if err != nil {
		log.Fatalf("no se pudo cargar los certifiados: %v", err)
	}

	store := storage.NewMemory()

	e := echo.New()

	// mux := http.NewServeMux()

	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	log.Println("Servidor corriendo en http://127.0.0.1:8080")
	err = e.Start(":3000")
	if err != nil {
		log.Printf("error en el servidor %v\n", err)
	}
}
