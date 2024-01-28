package handler

import (
	// "net/http"
	middlewares "res-API/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// / routePerson
func RoutePerson(e *echo.Echo, storage Storage) {

	h := newPerson(storage)
	persons := e.Group("/v1/persons")
	// middleware
	e.Use(middleware.Recover() /*middleware.Logger()*/)
	persons.Use(middlewares.Authentication, middlewares.Log)

	// crear persona
	persons.POST("", h.create)
	// eliminar persona
	persons.DELETE("/:id", h.Delete)
	// actualizar persona
	persons.PUT("/:id", h.Update)
	// obtener personas
	persons.GET("", h.getALL)
	// obtener persona por ID
	persons.GET("/:id", h.GetById)

}

// ROuteLogin
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v2/login", h.login)
}
