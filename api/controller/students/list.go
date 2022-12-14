package students

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/controller/services"
)

func List(c *gin.Context) {
	repository := database.GetRepository()
	listStudentsService := services.NewListStudentsService(repository)

	response := listStudentsService.Execute()

	c.JSON(int(response.Status), response)
}
