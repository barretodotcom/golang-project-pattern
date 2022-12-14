package interfaces

import (
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/model"
)

type IPutStudentService interface {
	Execute(model.Student, string) controller.Response
}
