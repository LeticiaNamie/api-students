package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQLite: %s", err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {
	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Msgf("Failed to create student:")

		return result.Error
	}

	log.Info().Msg("Create student!")

	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}
	err := s.DB.Find(&students).Error

	return students, err
}
