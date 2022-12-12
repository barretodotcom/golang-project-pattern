package interfaces

import (
	"github.com/golang-project-pattern/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDatabase interface {
	CountStudents(bson.D) int64
	InsertStudent(model.Student, *options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOneById(string) *mongo.SingleResult
	DeleteOneById(string) (*mongo.DeleteResult, error)
	FindAllStudents() (*mongo.Cursor, error)
	UpdateStudentById(string, bson.D)
}
