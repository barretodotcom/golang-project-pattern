package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

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

	DB = collection
}
