package db

import (
	"log"

	"github.com/Gemba-Kaizen/example-microservice-gRPC/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err!= nil {
    log.Fatalln((err))
  }

	// Migrate db for every model you have
	db.AutoMigrate(&models.Ping{})

	return Handler{db}
}