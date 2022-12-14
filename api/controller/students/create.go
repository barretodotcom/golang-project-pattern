package students

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/controller/services"
	"github.com/golang-project-pattern/api/model"
)

func Create(c *gin.Context) {
	var student model.Student
	c.BindJSON(&student)

	studentRepository := database.GetRepository()
	service := services.NewCreateStudentService(studentRepository)

	response := service.Execute(student)

	c.JSON(int(response.Status), response)
}
