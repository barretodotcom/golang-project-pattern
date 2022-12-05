package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-project-pattern/api/model"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, model.Students)

}
