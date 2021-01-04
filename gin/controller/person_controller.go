package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mithun/gin_example/entity"
	"github.com/mithun/gin_example/service"
)

type PersonController interface {
	SavePerson(ctx *gin.Context) entity.Person
	FindAllPerson() []entity.Person
}

type controller struct {
	service service.PersonService
}

func New(service service.PersonService) PersonController {
	return &controller{
		service: service,
	}
}

func (c *controller) SavePerson(ctx *gin.Context) entity.Person {
	var person entity.Person
	ctx.BindJSON(&person)
	c.service.SavePerson(person)
	return person
}

func (c *controller) FindAllPerson() []entity.Person {
	return c.service.FindAllPerson()
}
