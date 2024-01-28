package middlewares

import (
	"log"
	"net/http"
	"res-API/authorization"
	"time"

	"github.com/labstack/echo/v4"
)

// Middlewarlog Imprime las solicitudes, tiempo y ruta en consola

// forbidden: handler para falla de Autorizacion

// Verifica si El token es valido
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		_, err := authorization.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "No esta Autorizado"})
		}
		return f(c)
	}
}

func Log(f echo.HandlerFunc) echo.HandlerFunc {
	startTime := time.Now()
	return func(c echo.Context) error {
		duration := time.Since(startTime)

		log.Printf("Peticion a: %q metodo: %q %q", c.Request().URL.RequestURI(), c.Request().Method, duration)

		return f(c)
	}
}
