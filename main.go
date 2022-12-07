package main

import (
	"log"

	"github.com/golang-project-pattern/api"
	"github.com/golang-project-pattern/api/controller/infra/database"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

}

func main() {
	database.ConnectStudents()
	service := api.NewService()
	service.Start()
}
