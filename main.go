package main

import (
	"github.com/golang-project-pattern/api"
)

func main() {
	service := api.NewService()
	service.Start()
}
