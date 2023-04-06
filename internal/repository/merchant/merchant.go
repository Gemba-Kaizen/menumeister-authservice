package repository

import (
	"github.com/Gemba-Kaizen/menumeister-authservice/internal/db"
	"github.com/Gemba-Kaizen/menumeister-authservice/internal/models"
)

type MerchantRepository struct {
	H *db.Handler
}

func (r *MerchantRepository) CreateMerchant(merchant *models.Merchant) error{
	if result := r.H.DB.Create(merchant); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MerchantRepository) GetMerchantByEmail(email string) (*models.Merchant, error) {
	var merchant models.Merchant

	if result := r.H.DB.Where(&models.Merchant{Email: email}).First(&merchant); result.Error!= nil {
		return nil, result.Error
	}
	return &merchant, nil
}