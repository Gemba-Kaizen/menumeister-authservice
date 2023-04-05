package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Gemba-Kaizen/example-microservice-gRPC/config"
	"github.com/Gemba-Kaizen/example-microservice-gRPC/internal/db"
	repository "github.com/Gemba-Kaizen/example-microservice-gRPC/internal/repository/ping"
	api "github.com/Gemba-Kaizen/example-microservice-gRPC/pkg/api/ping"
	pb "github.com/Gemba-Kaizen/example-microservice-gRPC/pkg/pb"
	services "github.com/Gemba-Kaizen/example-microservice-gRPC/pkg/services/ping"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	// Init DB
	h := db.Init(c.DBUrl)

	// Init pingService
	pingService := &services.PingService{
		PingRepo: &repository.PingRepository{H: &h},
	}

	// Init handlers
	pingHandler := &api.PingHandler{PingService: pingService}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed at listen: ", err)
	}

	fmt.Println("Ping Svc on", c.Port)

	grpcService := grpc.NewServer()

	// Register each handler endpoint to grpc Server
	pb.RegisterPingServiceServer(grpcService, pingHandler)
	// pb.RegisterService2ServiceServer(grpcServer, service2Handler)

	if err := grpcService.Serve(lis); err != nil {
		log.Fatalln("Failed at serve: ", err)
	}
}
