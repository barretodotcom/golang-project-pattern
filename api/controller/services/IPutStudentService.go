package services

import (
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/interfaces"
	"github.com/golang-project-pattern/api/model"
	"go.mongodb.org/mongo-driver/bson"
)

type PutStudentService struct {
	Repository interfaces.IDatabase
}

func NewPutStudentService(repository interfaces.IDatabase) interfaces.IPutStudentService {
	return &PutStudentService{
		Repository: repository,
	}
}

func (ps *PutStudentService) Execute(student model.Student, studentId string) controller.Response {
	studentsLength := ps.Repository.CountStudents(bson.D{{Key: "id", Value: studentId}})

	if studentsLength == 0 {
		return controller.NewResponse("Student does not exists", 400)
	}

	ps.Repository.UpdateStudentById(student.ID, bson.D{
		{Key: "name", Value: student.Name},
	})

	return controller.NewResponse("Student updated with success", 201)

}
