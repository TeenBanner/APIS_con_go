package handler

import (
	"encoding/json"
	"net/http"
	"res-API/authorization"
	"res-API/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "Estructuta mal estructurada", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}
	if !isLoginValid(&data) {
		resp := newResponse(Error, "usuario o contrasena no validos", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}
	datatoken := map[string]string{"token": token}
	resp := newResponse(Message, "OK", datatoken)
	responseJSON(w, http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ed.team" && data.Password == "123456"
}
