package services

import (
	"net/http"

	"github.com/Gemba-Kaizen/menumeister-authservice/internal/models"
	repository "github.com/Gemba-Kaizen/menumeister-authservice/internal/repository/merchant"
	"github.com/Gemba-Kaizen/menumeister-authservice/pkg/pb"
	"github.com/Gemba-Kaizen/menumeister-authservice/pkg/services"
)

type AuthService struct {
	MerchantRepo *repository.MerchantRepository
	Jwt services.JwtWrapper
}

// Check email exists, create merchant if not exists
func (s *AuthService) RegisterMerchant(email string, password string) (*pb.RegisterResponse, error) {
	var merchant models.Merchant

	merchant.Email = email
	merchant.Password = services.HashPassword(password);

	err := s.MerchantRepo.CreateMerchant(&merchant)

	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error: "e-mail already taken",
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

// Check that email exists, check if password correct, generate token
func (s *AuthService) Login(email string, password string) (*pb.LoginResponse, error) {
  merchant, err := s.MerchantRepo.GetMerchantByEmail(email); if err!= nil {
		return &pb.LoginResponse{
      Status: http.StatusNotFound,
      Error: "merchant not found",
    }, nil
	}

	match := services.CheckPassword(password, merchant.Password)

	if !match {
		return &pb.LoginResponse{
      Status: http.StatusUnauthorized,
      Error: "merchant not found",
    }, nil
	}

	token, _ := s.Jwt.GenerateToken(*merchant)

	return &pb.LoginResponse{
		Status: http.StatusOK,
  	Token: token,
	}, nil
}

func (s *AuthService) ValdiateSession(token string) (*pb.ValidateMerchantResponse, error){
	claims, err := s.Jwt.ValidateToken(token)

	if err != nil {
		return &pb.ValidateMerchantResponse{
			Status: http.StatusBadRequest,
      Error: err.Error(),
		}, nil
	}

	merchant, err := s.MerchantRepo.GetMerchantByEmail(claims.Email);
	
	if err != nil {
		return &pb.ValidateMerchantResponse{
      Status: http.StatusNotFound,
      Error: "merchant not found",
    }, nil
	}

	return &pb.ValidateMerchantResponse{
		Status: http.StatusOK,
		MerchantId: merchant.Id,
	}, nil
}