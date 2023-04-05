package repository

import (
	"github.com/Gemba-Kaizen/example-microservice-gRPC/internal/db"
	"github.com/Gemba-Kaizen/example-microservice-gRPC/internal/models"
)

type PingRepository struct {
	H *db.Handler
}

func (r *PingRepository) FindById(id uint64) (*models.Ping, error) {
	var ping models.Ping
	if result := r.H.DB.First(&ping, id); result.Error!= nil {
		return nil, result.Error
	}

	return &ping, nil
}

func (r *PingRepository) CreatePing(ping *models.Ping) error{
	if result := r.H.DB.Create(ping); result.Error != nil {
		return result.Error
	}
	
	return nil
}