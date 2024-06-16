package api

import (
	"fmt"
	"net/http"

	"github.com/LeticiaNamie/api-students/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/studapi.Echonts", createStudent)
	api.Echo.GET("/students/:id", getStudent)
	api.Echo.PUT("/students/:id", updateStudent)
	api.Echo.DELETE("/students/:id", deleteStudent)
}

// Handler
func getStudents(c echo.Context) error {
	students, err := db.GetStudents()

	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	return c.JSON(http.StatusOK, students)
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
	return c.String(http.StatusOK, "Get student: "+id)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update student: "+id)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Delete student: "+id)
}
