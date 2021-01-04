package service

import (
	"github.com/mithun/gin_example/entity"
)

type PersonService interface {
	SavePerson(entity.Person) entity.Person
	FindAllPerson() []entity.Person
}

type personService struct {
	persons []entity.Person
}

func New() PersonService {
	return &personService{
		persons: []entity.Person{},
	}
}

func (service *personService) SavePerson(person entity.Person) entity.Person {
	service.persons = append(service.persons, person)
	return person
}

func (service *personService) FindAllPerson() []entity.Person {
	return service.persons
}
