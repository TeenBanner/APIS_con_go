package handler

import "net/http"

// routePerson
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", h.create)
	mux.HandleFunc("/v1/persons/update", h.Update)
	mux.HandleFunc("/v1/persons/get-all", h.getALL)
}
