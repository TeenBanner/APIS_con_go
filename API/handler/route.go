package handler

import (
	"net/http"
	"res-API/middleware"
)

// routePerson
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/v1/persons/update", middleware.Log(middleware.Authentication(h.Update)))
	mux.HandleFunc("/v1/persons/get-all", middleware.Log(middleware.Authentication(h.getALL)))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(middleware.Authentication(h.Delete)))
	mux.HandleFunc("/v1/persons/Get-By-Id", middleware.Log(middleware.Authentication(h.GetById)))
}

// ROuteLogin
func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/v2/login", h.login)
}
