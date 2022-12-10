package students

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Delete(c *gin.Context) {

	id := c.Param("id")
	DB := database.DB
	var student model.Student

	studentExists := DB.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}})

	if studentExists == nil {
		c.JSON(http.StatusUnprocessableEntity, controller.NewResponse("Student not found.", "error"))
		return
	}

	studentExists.Decode(&student)

	_, err := DB.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: student.ID}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponse("Could not delete student, try again.", "error"))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponse("Student deleted with success.", "success"))
}
