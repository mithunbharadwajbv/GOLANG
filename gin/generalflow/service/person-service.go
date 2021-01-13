package service

import (
	"github.com/mithun/generalflow/collections"
	"github.com/mithun/generalflow/entity"
)

type PersonService interface {
	SavePerson(entity.Person) error
}

type personservice struct {
	collection collections.PersonCollection
}

func NewPersonService(collection collections.PersonCollection) PersonService {
	return &personservice{
		collection: collection,
	}
}

func (service *personservice) SavePerson(person entity.Person) error {
	service.collection.SavePerson(person)
	return nil
}
