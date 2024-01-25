package handler

import (
	"encoding/json"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Metodo no permitido"}`))
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "El o la persona No tiene una estructua correcta"}`))
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "error al crear la persona"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type": "message", "message": "persona Creada Correctamente"}`))

}

func (p *person) getALL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Metodo no permitido"}`))
		return
	}

	resp, err := p.storage.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "error al Obtener todas las persona"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "error al convertir el slice en json"}`))
		return
	}

}

func (p *person) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Metodo no permitido"}`))
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Id no valido"}`))
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "El o la persona No tiene una estructua correcta"}`))
		return
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "No se pudo actualizar la persona "}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message_type": "message", "message": "persona Actualizada Correctamente"}`))

}
