package main

import "database/sql"

type Employee struct {
    EmployeeID   int `json:"id"`
    FullName string `json:"fullname"`
	Birthday string `json:"birthday"`
	Gender string `json:"gender"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    interface{} `json:"data"`
    Message interface{} `json:"message"`
}

type Event struct {
    EventID   int `json:"id"`
    Name string `json:"name"`
	EventDate string `json:"eventDate"`
}

// repository represent the repository model
type repository struct {
	db *sql.DB
}

// Repository represent the repositories
type IHandler interface {
	GetEmployees()
}