package api

import (
	"context"
	"net/http"

	"github.com/Gemba-Kaizen/example-microservice-gRPC/internal/models"
	"github.com/Gemba-Kaizen/example-microservice-gRPC/pkg/services/ping"
	pb "github.com/Gemba-Kaizen/example-microservice-gRPC/pkg/pb"
)

type PingHandler struct {
	pb.UnimplementedPingServiceServer
	PingService *services.PingService
}

// GetPing handles the gRPC GetPing request.
func (h *PingHandler) GetPing(ctx context.Context, req *pb.GetPingRequest) (*pb.GetPingResponse, error) {
	// Call the PingService to retrieve the ping and add a message. - call wtv services to do business logic
	ping, err := h.PingService.GetPingAndAdd(uint64(req.UserId))

	// If there was an error, return a gRPC error response.
	if err != nil {
		return &pb.GetPingResponse{
			Status: http.StatusConflict,
      Error: err.Error(),
		}, nil
	}

	// No error, create a GetPingData message.
	data := &pb.GetPingData{
		UserId: ping.UserId,
		Message: ping.Message,
	}

	// Return a GetPingResponse with the data.
	return &pb.GetPingResponse{
		Status: http.StatusCreated,
		Data:     data,
	}, nil
}

// CreatePing handles the gRPC CreatePing request.
func (h *PingHandler) CreatePing(ctx context.Context, req *pb.CreatePingRequest) (*pb.CreatePingResponse, error) {
	// Create a new Ping with the message from the gRPC request.
	ping := &models.Ping{Message: req.Message}

	// Call the PingService to create the ping and add a message.
	if err := h.PingService.CreatePingAndAdd(ping); err != nil {
		// If there was an error, return a gRPC error response.
		return &pb.CreatePingResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// No error, create a CreatePingResponse with the ID.
	return &pb.CreatePingResponse{
		Status: http.StatusCreated,
		Id: ping.UserId,
	}, nil
}
