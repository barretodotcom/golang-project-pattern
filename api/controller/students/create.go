package students

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/model"
)

func Create(c *gin.Context) {

	var student model.Student

	c.BindJSON(&student)

	for _, v := range model.Students {
		if v.Name == student.Name {
			c.JSON(400, gin.H{
				"message": "This name is already used.",
				"status":  "error",
			})
			return
		}
	}
	student.ID = len(model.Students) + 1
	model.Students = append(model.Students, student)

	c.JSON(201, gin.H{
		"message": "ok",
		"status":  "created",
	})
}
