package model

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Students []Student = []Student{
	{ID: 1, Name: "Amanda"},
	{ID: 2, Name: "Lucas"},
}
