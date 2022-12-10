package students

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/golang-project-pattern/api/model"
	"github.com/golang-project-pattern/api/shared/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(c *gin.Context) {
	var student model.Student
	c.BindJSON(&student)

	count, _ := database.DB.CountDocuments(context.TODO(), bson.D{{Key: "name", Value: student.Name}})
	if count != 0 {
		c.JSON(400, controller.NewResponse("Username already registered", "error"))
		return
	}

	if len(student.Name) < 10 {
		c.JSON(400, controller.NewResponse("Username must have more than 10 caracters.", "error"))
	}
	student.ID = uuid.GetId()
	insertOneResult, err := database.DB.InsertOne(context.TODO(), student, &options.InsertOneOptions{})

	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, insertOneResult)
}
