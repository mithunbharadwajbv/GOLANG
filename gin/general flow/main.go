package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mithun/gin_example/controller"
	"github.com/mithun/gin_example/service"
)

var (
	personService    service.PersonService       = service.New()
	personController controller.PersonController = controller.New(personService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, personController.FindAllPerson())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, personController.SavePerson(ctx))
	})

	server.Run(":8080")
}
