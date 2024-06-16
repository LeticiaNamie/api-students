package api

import (
	"fmt"
	"net/http"

	"github.com/LeticiaNamie/api-students/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/studapi.Echonts", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

// Handler
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()

	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}

	if err := c.Bind(&student); err != nil {
		fmt.Println("Error to bind data")

		return err
	}

	if err := api.DB.AddStudent(student); err != nil {
		fmt.Println("Error to create student")

		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	fmt.Println("Create student!")

	return c.JSON(http.StatusOK, student)
}

func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Get student: "+id)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update student: "+id)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Delete student: "+id)
}
