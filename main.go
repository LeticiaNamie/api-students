package main

import (
	"fmt"
	"net/http"

	"github.com/LeticiaNamie/api-students/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}

func createStudent(c echo.Context) error {
	student := db.Student{}

	if err := c.Bind(&student); err != nil {
		fmt.Println("Error to bind data")

		return err
	}

	if err := db.AddStudent(student); err != nil {
		fmt.Println("Error to create student")

		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	fmt.Println("Create student!")

	return c.JSON(http.StatusOK, student)
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Student: "+id)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update student: "+id)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Delete student: "+id)
}
