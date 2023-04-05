package services

import (
	"github.com/Gemba-Kaizen/example-microservice-gRPC/internal/models"
	"github.com/Gemba-Kaizen/example-microservice-gRPC/internal/repository/ping"
)

// PingService contains the business logic for the Ping service.
type PingService struct {
	PingRepo *repository.PingRepository
}

// GetPingAndAdd returns the Ping model with the given ID and adds "edited From Microservice" to its Message field.
func (s *PingService) GetPingAndAdd(id uint64) (*models.Ping, error) {
	// Get the Ping model from the repository with the given ID.
	ping, err := s.PingRepo.FindById(id)
	if err != nil {
		// Return nil and the error if the repository returned an error.
		return nil, err
	}
	// Add "edited From Microservice" to the Ping model's Message field. - aka business logic
	ping.Message += "edited From Microservice"

	// Return the updated Ping model and nil error.
	return ping, nil
}

// CreatePingAndAdd creates a new Ping model with the given data and adds "added when creating in micro service" to its Message field.
func (s *PingService) CreatePingAndAdd(ping *models.Ping) error {
	// Add "added when creating in micro service" to the Ping model's Message field.
	ping.Message += "added when creating in micro service"

	// Create the Ping model in the repository.
	err := s.PingRepo.CreatePing(ping)
	if err != nil {
		// Return the error if the repository returned an error.
		return err
	}

	// Return nil if the Ping model was created successfully.
	return nil
}