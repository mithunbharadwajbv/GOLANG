package helper

import (
	"github.com/mithun/generalflow/collections"
	"github.com/mithun/generalflow/controller"
	"github.com/mithun/generalflow/entity"
	"github.com/mithun/generalflow/service"
)

var (
	personCollection collections.PersonCollection = collections.NewPersonCollection()
	personService    service.PersonService        = service.NewPersonService(personCollection)
	personController controller.PersonController  = controller.NewPersonController(personService)
)

func MongoHelper() {
	person1 := entity.GetSamplePerson("1")
	personController.SavePerson(person1)
}
