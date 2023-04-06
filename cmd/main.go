package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Gemba-Kaizen/menumeister-authservice/config"
	"github.com/Gemba-Kaizen/menumeister-authservice/internal/db"
	repository "github.com/Gemba-Kaizen/menumeister-authservice/internal/repository/merchant"
	api "github.com/Gemba-Kaizen/menumeister-authservice/pkg/api/auth"
	"github.com/Gemba-Kaizen/menumeister-authservice/pkg/pb"
	services "github.com/Gemba-Kaizen/menumeister-authservice/pkg/services/auth"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	// Init DB
	h := db.Init(c.DBUrl)

	// Init authService
	authService := &services.AuthService{
		MerchantRepo: &repository.MerchantRepository{H: &h},
	}

	// Init handlers
	authHandler := &api.AuthHandler{AuthService: authService}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed at listen: ", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	grpcService := grpc.NewServer()

	// Register each handler endpoint to grpc Server
	pb.RegisterAuthServiceServer(grpcService, authHandler)
	// pb.RegisterService2ServiceServer(grpcServer, service2Handler)

	if err := grpcService.Serve(lis); err != nil {
		log.Fatalln("Failed at serve: ", err)
	}
}
