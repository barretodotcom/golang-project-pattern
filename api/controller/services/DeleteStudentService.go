package services

import (
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/interfaces"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteStudentService struct {
	repository interfaces.IDatabase
}

func NewDeleteStudentService(db interfaces.IDatabase) interfaces.IDeleteStudentInterface {

	return &DeleteStudentService{
		repository: db,
	}

}

func (d *DeleteStudentService) Execute(studentId string) controller.Response {

	studentExists := d.repository.CountStudents(bson.D{{Key: "id", Value: studentId}})

	if studentExists == 0 {
		return controller.NewResponse("There is no student with this id.", 400)
	}
	d.repository.DeleteOneById(studentId)

	return controller.NewResponse("Student deleted with success.", 201)
}
