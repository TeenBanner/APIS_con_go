package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:3000"

func main() {
	lc := loginClient(url+"/v2/login", "contacto@ed.team", "123456")
	// fmt.Printf("%v", lc)

	person := Person{
		Name:        "alexys Lozada",
		Age:         40,
		Communities: []Community{Community{Name: "EDteam"}, Community{Name: "golang Backend"}},
	}

	gr := createPerson(url+"/v1/persons", lc.Data.Token, &person)
	fmt.Println(gr)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return response
}
