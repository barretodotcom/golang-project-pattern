package students

import (
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
	repository := database.GetRepository()

	count := repository.CountStudents(bson.D{{Key: "name", Value: student.Name}})
	if count != 0 {
		c.JSON(400, controller.NewResponse("Username already registered", "error"))
		return
	}

	if len(student.Name) < 10 {
		c.JSON(400, controller.NewResponse("Username must have more than 10 caracters.", "error"))
	}
	student.ID = uuid.GetId()
	insertOneResult, err := repository.InsertStudent(student, &options.InsertOneOptions{})

	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, insertOneResult)
}
