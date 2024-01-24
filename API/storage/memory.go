package storage

import (
	"fmt"
	"res-API/model"
)

// Memory
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// create
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

// Actualiza un Usuario
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}

	m.Persons[ID] = *person

	return nil
}

// Elimina una persona
func (m *Memory) Delete(ID int) error {

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}

	delete(m.Persons, ID)

	return nil
}

func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID %d %w", ID, model.ErrIDPersonDoesNotExist)
	}

	return person, nil
}

func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}
