package handler

import (
	"net/http"
	"res-API/authorization"
	"res-API/model"

	"github.com/labstack/echo/v4"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {
	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := newResponse(Error, "Estructuta mal estructurada", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if !isLoginValid(&data) {
		resp := newResponse(Error, "usuario o contrasena no validos", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	datatoken := map[string]string{"token": token}
	resp := newResponse(Message, "OK", datatoken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ed.team" && data.Password == "123456"
}
