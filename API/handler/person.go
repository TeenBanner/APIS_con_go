package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"res-API/model"
	"strconv"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Metodo no Permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := newResponse(Error, "La estructura persona no esta correcta", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		response := newResponse(Error, "Error al crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Persona Creada Correctamente", data)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) getALL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Error al obtener todas la personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Personas obtenidas correctamente", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Metodo no Permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "Id no valido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := newResponse(Error, "la persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		response := newResponse(Error, "No se pudo actualizar la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "persona creada correctamente", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "Id no valido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "No se pudo borrar la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "se pudo borrar la persona correctamente", nil)
	responseJSON(w, http.StatusOK, response)

}

func (p *person) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response := newResponse(Error, "El id no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}

	Person, err := p.storage.GetByID(ID)

	if err != nil {
		response := newResponse(Error, "El id no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}

	response := newResponse(Message, "se Obtuvo la persona por su ID satisfactoriamente", Person)
	responseJSON(w, http.StatusOK, response)

}
