package students

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/controller/services"
)

func Delete(c *gin.Context) {
	repository := database.GetRepository()
	deleteStudentService := services.NewDeleteStudentService(repository)
	id := c.Param("id")

	response := deleteStudentService.Execute(id)

	c.JSON(int(response.Status), response)
}
