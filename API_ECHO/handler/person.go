package handler

import (
	//"encoding/json"
	//"errors"
	"errors"
	"net/http"

	"res-API/model"

	"strconv"

	"github.com/labstack/echo/v4"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {
	data := model.Person{}
	err := c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "Respuesta mal estructurada", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Error al crear la persona"})
	}
	response := newResponse(Message, "Data Creada", data)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) getALL(c echo.Context) error {
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "No se pudo obtener la data", nil)
		c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Personas obtenidas correctamente", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) Update(c echo.Context) error {
	IDsrt := c.Param("id")

	ID, err := strconv.Atoi(IDsrt)

	if err != nil {
		response := newResponse(Error, "el id no es valido", err)
		c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "Error", "message": "Esctructura incorrecta"})
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "error al actualizar la persona"})
	}
	response := newResponse(Message, "Data Modificada satisfactoriamente", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) Delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "Id no valido", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "No se pudo borrar la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "se pudo borrar la persona correctamente", ID)
	return c.JSON(http.StatusOK, response)
}

func (p *person) GetById(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := newResponse(Error, "El id no existe", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	Person, err := p.storage.GetByID(ID)

	if err != nil {
		response := newResponse(Error, "El id no existe", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	response := newResponse(Message, "se Obtuvo la persona por su ID satisfactoriamente", Person)
	return c.JSON(http.StatusOK, response)
}
