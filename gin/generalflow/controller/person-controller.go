package controller

import (
	"github.com/mithun/generalflow/entity"
	"github.com/mithun/generalflow/service"
)

type PersonController interface {
	SavePerson(entity.Person) error
}

type personController struct {
	service service.PersonService
}

func NewPersonController(service service.PersonService) PersonController {
	return &personController{
		service: service,
	}
}

func (controller *personController) SavePerson(person entity.Person) error {
	controller.service.SavePerson(person)
	return nil
}
