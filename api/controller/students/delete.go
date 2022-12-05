package students

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/model"
)

func Delete(c *gin.Context) {

	var newStudents []model.Student
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	for _, student := range model.Students {
		if student.ID != id {
			newStudents = append(newStudents, student)

		}
	}

	model.Students = newStudents

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"id":      id,
	})
}
