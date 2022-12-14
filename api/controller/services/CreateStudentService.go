package services

import (
	"github.com/golang-project-pattern/api/controller"
	"github.com/golang-project-pattern/api/interfaces"
	"github.com/golang-project-pattern/api/model"
	"github.com/golang-project-pattern/api/shared/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateStudentService struct {
	Repository interfaces.IDatabase
}

func NewCreateStudentService(repository interfaces.IDatabase) interfaces.ICreateStudentService {
	return &CreateStudentService{
		Repository: repository,
	}
}

func (c *CreateStudentService) Execute(student model.Student) controller.Response {
	studentExists := c.Repository.CountStudents(bson.D{{Key: "name", Value: student.Name}})

	if studentExists > 0 {
		return controller.NewResponse("User already exists.", 400)
	}

	uuid := uuid.GetId()

	student.ID = uuid

	_, err := c.Repository.InsertStudent(student, &options.InsertOneOptions{})

	if err != nil {
		return controller.NewResponse("Internal server error, try again", 401)
	}

	return controller.NewResponse("Created.", 200)

}
