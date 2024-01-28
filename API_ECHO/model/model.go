package model

import "errors"

var (
	// person can't be nil
	ErrPersonCanNotBeNil = errors.New("la persona no puede ser nula")
	// la persona no esxiste
	ErrIDPersonDoesNotExist = errors.New("El Id de la persona no puede ser nulo")
)
