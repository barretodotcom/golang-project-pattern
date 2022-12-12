package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/model"
)

func Delete(c *gin.Context) {

	id := c.Param("id")
	repository := database.GetRepository()
	var student model.Student

	studentExists := repository.FindOneById(id)

	if studentExists == nil {
		c.JSON(http.StatusUnprocessableEntity, controller.NewResponse("Student not found.", "error"))
		return
	}

	studentExists.Decode(&student)

	_, err := repository.DeleteOneById(student.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponse("Could not delete student, try again.", "error"))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponse("Student deleted with success.", "success"))
}
