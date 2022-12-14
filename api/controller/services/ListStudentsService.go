package services

import (
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/interfaces"
	"github.com/golang-project-pattern/api/model"
)

type ListStudentsService struct {
	Repository interfaces.IDatabase
}

func NewListStudentsService(repository interfaces.IDatabase) interfaces.IListStudentsService {
	return &ListStudentsService{
		Repository: repository,
	}
}

func (ls *ListStudentsService) Execute() controller.Response {

	var allStudents []model.Student

	allStudentsCursor, err := ls.Repository.FindAllStudents()

	allStudentsCursor.Decode(&allStudents)

	if err != nil {
		controller.NewResponse("Internal server error", 400, err.Error())
	}

	return controller.NewResponse("All students found.", 200, allStudents)

}
