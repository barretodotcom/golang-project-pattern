package interfaces

import "github.com/golang-project-pattern/api/controller"

type IListStudentsService interface {
	Execute() controller.Response
}
