package students

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Put(c *gin.Context) {
	var student model.Student
	var newStudentParameters model.Student
	id := c.Param("id")

	studentExists := database.DB.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if studentExists.Err() == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadGateway, controller.NewResponse("Student doesn't exists.", "error"))
		return
	}

	err := c.BindJSON(&newStudentParameters)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponse("Invalid request body: Not a JSON.", "error"))
		return
	}

	studentExists.Decode(&student)

	student.Name = newStudentParameters.Name

	database.DB.FindOneAndUpdate(context.TODO(), bson.D{{Key: "id", Value: student.ID}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "id", Value: student.ID},
				{Key: "name", Value: student.Name},
			},
			},
		})

	c.JSON(http.StatusNoContent, controller.NewResponse("Updated.", "success."))
	return
}
