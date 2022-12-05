package api

import "github.com/golang-project-pattern/api/controller/students"

func (e *Service) SetupRoutes() {
	studentsGroup := e.Group("/students")

	studentsGroup.GET("/list", students.List)
	studentsGroup.POST("/create", students.Create)
	studentsGroup.DELETE("/delete/:id", students.Delete)
	studentsGroup.PUT("/update/:id", students.Put)
}
