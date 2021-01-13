package main

import (
	"github.com/mithun/generalflow/collections"
	"github.com/mithun/generalflow/controller"
	"github.com/mithun/generalflow/helper"
	"github.com/mithun/generalflow/service"
)

var (
	personCollection collections.PersonCollection = collections.NewPersonCollection()
	personService    service.PersonService        = service.NewPersonService(personCollection)
	personController controller.PersonController  = controller.NewPersonController(personService)
)

func main() {

	// helper.MongoHelper()

	helper.Temp()
}

func temp() (string, error) {
	
}
