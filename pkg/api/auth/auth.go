package api

import (
	"context"

	pb "github.com/Gemba-Kaizen/menumeister-authservice/pkg/pb"
	"github.com/Gemba-Kaizen/menumeister-authservice/pkg/services/auth"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	AuthService *services.AuthService
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	res, err := h.AuthService.RegisterMerchant(req.Email, req.Password)
	return res, err
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := h.AuthService.Login(req.Email, req.Password)
	return res, err
}

func (h *AuthHandler) ValidateSession(ctx context.Context, req *pb.ValidateMerchantRequest) (*pb.ValidateMerchantResponse, error) {
	res, err := h.AuthService.ValdiateSession(req.Token)
	return res, err
}
