package interfaces

import "github.com/golang-project-pattern/api/controller"

type IDeleteStudentInterface interface {
	Execute(string) controller.Response
}
