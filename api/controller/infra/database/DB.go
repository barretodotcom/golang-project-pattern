package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-project-pattern/api/interfaces"
	"github.com/golang-project-pattern/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Collection

type Repository struct {
	db *mongo.Collection
}

func ConnectStudents() {
	mongoEnv := []string{
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
	}
	uri := fmt.Sprintf(`mongodb+srv://%s:%s@%s`, mongoEnv[0], mongoEnv[1], mongoEnv[2])

	conn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err.Error())
	}

	database := conn.Database("students")

	collection := database.Collection("students")

	db = collection
}

func GetRepository() interfaces.IDatabase {
	repository := &Repository{
		db: db,
	}
	return repository
}

func (r *Repository) CountStudents(filter bson.D) int64 {
	count, _ := r.db.CountDocuments(context.TODO(), filter)

	return count
}

func (r *Repository) InsertStudent(student model.Student, options *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return r.db.InsertOne(context.TODO(), student, options)
}

func (r *Repository) FindOneById(id string) *mongo.SingleResult {
	return r.db.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}})
}

func (r *Repository) DeleteOneById(id string) (*mongo.DeleteResult, error) {
	return r.db.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}}, &options.DeleteOptions{})
}

func (r *Repository) FindAllStudents() (*mongo.Cursor, error) {
	return r.db.Find(context.TODO(), bson.D{}, &options.FindOptions{})
}

func (r *Repository) UpdateStudentById(id string, params bson.D) {
	r.db.FindOneAndUpdate(context.TODO(), bson.D{{Key: "id", Value: id}}, params)
}
