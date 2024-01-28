package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:     email,
		Passsword: password,
	}
	data := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error en marshal de login %v", err)
	}

	resp := httpClient(http.MethodPost, url, "", data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("No se pudo leer el body Error: %v ", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Se esperaba un Statuscode OK se obtuvo %v, respuesta %s", resp.StatusCode, string(body))
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)

	if err != nil {
		log.Fatalf("Error en el unmarshal del body")
	}
	fmt.Println(string(body))

	return dataResponse
}
