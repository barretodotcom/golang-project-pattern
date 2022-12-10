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

func List(c *gin.Context) {
	DB := database.DB
	var allStudents []model.Student
	students, err := DB.Find(context.TODO(), bson.D{{}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponse(err.Error(), "error"))
		return
	}
	err = students.All(context.TODO(), &allStudents)

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponse(err.Error(), "error"))
		return
	}

	c.JSON(http.StatusOK, allStudents)
}
