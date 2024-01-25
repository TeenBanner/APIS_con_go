package model

type Community struct {
	Name string
}

type Communities []Community

type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
	FirstName   string      `json:"first_name"`
}

type Persons []Person
