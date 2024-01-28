package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/", saludar)
	// grupo de rutas nos ayuda a evitar escribir toda la direccion de la ruta usar middleware.
	// e.POST("/crear", crear)
	// e.GET("/consulta", consultar)
	// e.PUT("/actualizar", actualizar)
	// e.DELETE("/Eliminar", eliminar)
	// creamos un grupo de personas
	persons := e.Group("/persons")
	// creamos una ruta apartir del grupo
	persons.POST("", crear)
	// en vez de usar query params de manera tradicional
	// http://localhost/v1/persons?id=10 podemos hacerlo mas semantico http://localhost/v1/persons/10
	// esto podemos hacerlo en echo asi
	// podemos evitar usar el middleware en la ruta asi
	persons.Use(middlewareLogPersonas) // hacemos que el grupo de rutas use este middleware
	persons.POST("", crear)

	persons.GET("/:id", consultar)
	persons.PUT("/:id", actualizar)
	persons.DELETE("/:id", eliminar)

	// ejemplo controlar panic
	persons.GET("/dividir", dividir)

	e.Start("192.168.5.206:8080")

}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "Hola Mundo"})
}

// middleware
func middlewareLogPersonas(f echo.HandlerFunc) echo.HandlerFunc { // recibe y retorna una funcion de tipo echo.HandlerFunc
	return func(c echo.Context) error { // retorna una funcion con un contexto de echo
		startTime := time.Now()
		log.Println("peticion echa a personas")
		defer func() {
			duration := time.Since(startTime)
			log.Printf("duration: %q", duration)
		}()
		return f(c) // y retorna la ejecucion de la funcion f con el contexto de c
	}
}
func dividir(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)

	if f >= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": "no se puede dividir entre 0"})
	}

	r := 3000 / f

	return c.String(http.StatusOK, strconv.Itoa(r))
}

func crear(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"Message": "data Creada"})
}

func consultar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Consultado"+id)
}

func actualizar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "data Actualizada"})
}

func eliminar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "data Eliminada"})
}
