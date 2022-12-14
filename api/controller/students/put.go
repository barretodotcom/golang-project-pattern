package students

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/controller/services"
	"github.com/golang-project-pattern/api/model"
)

func Put(c *gin.Context) {
	repository := database.GetRepository()
	putStudentService := services.NewPutStudentService(repository)
	var studet model.Student
	studentId := c.Param("id")
	c.BindJSON(&studet)

	response := putStudentService.Execute(studet, studentId)

	c.JSON(int(response.Status), response)
}
