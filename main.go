package main

import (
	"crudgolang/controllers"
	"crudgolang/db"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := NewRouter()

	fmt.Println("starting web server at http://localhost:7000/")
	e.Start(":7000")
}

func NewRouter() *echo.Echo {
	e := echo.New()

	// Initialize main database
	db.Db = db.Connect()

	// e.GET("/student/:id", controllers.GetStudent)
	// e.GET("/student", controllers.GetAllStudent)
	// e.POST("/student", controllers.CreateStudent)
	// e.PUT("/student/:id", controllers.UpdateStudent)
	// e.DELETE("/student/:id", controllers.DeleteStudent)

	e.GET("/data/:id", controllers.GetData)
	e.GET("/data", controllers.GetAllData)
	e.GET("/file", controllers.Reportcsv)

	return e
}
