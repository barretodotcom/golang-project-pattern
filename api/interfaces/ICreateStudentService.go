package interfaces

import (
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/model"
)

type ICreateStudentService interface {
	Execute(model.Student) controller.Response
}
