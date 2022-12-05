package students

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/model"
)

func Put(c *gin.Context) {
	var student model.Student

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
	}
	c.BindJSON(&student)

	for i, v := range model.Students {
		if v.ID == id {
			student.ID = v.ID
			model.Students[i] = student
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update made with success.",
		"status":  "ok",
	})
	return
}
