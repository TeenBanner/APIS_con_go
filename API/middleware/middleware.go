package middleware

import (
	"log"
	"net/http"
	"res-API/authorization"
	"time"
)

type handler func(http.ResponseWriter, *http.Request)

// Middlewarlog Imprime las solicitudes, tiempo y ruta en consola
func Log(f handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Peticion %q, Metodo: %q", r.URL.Path, r.Method)
		startTime := time.Now()

		defer func() {
			duration := time.Since(startTime)
			log.Printf("Time: %q", duration)
		}()

		f(w, r)
	}
}

// forbidden: handler para falla de Autorizacion
func forbbiden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autoizacion"))
}

// Verifica si El token es valido
func Authentication(f handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbbiden(w, r)
			return
		}
		f(w, r)
	}
}
