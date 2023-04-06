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

func (h *AuthHandler) RegisterMerchant(ctx context.Context, req *pb.RegisterMerchantRequest) (*pb.RegisterMerchantResponse, error) {
	res, err := h.AuthService.RegisterMerchant(req.Email, req.Password)
	return res, err
}

func (h *AuthHandler) LoginMerchant(ctx context.Context, req *pb.LoginMerchantRequest) (*pb.LoginMerchantResponse, error) {
	res, err := h.AuthService.LoginMerchant(req.Email, req.Password)
	return res, err
}

func (h *AuthHandler) ValidateMercant(ctx context.Context, req *pb.ValidateMerchantRequest) (*pb.ValidateMerchantResponse, error) {
	res, err := h.AuthService.ValdiateSession(req.Token)
	return res, err
}
