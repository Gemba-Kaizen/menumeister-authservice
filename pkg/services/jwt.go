package services

import (
	"errors"
	"time"

	"github.com/Gemba-Kaizen/menumeister-authservice/internal/models"
	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// Object to sign, used to identify user
type JwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
}

func (w *JwtWrapper) GenerateToken(merchant models.Merchant) (signedToken string, err error) {
	claims := &JwtClaims{
		Id:    merchant.Id,
		Email: merchant.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(w.ExpirationHours) * time.Hour).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
    return nil, err
  }

	claims, ok := token.Claims.(*JwtClaims)

	if !ok {
		return nil, errors.New("could not parse claims")
	}

	if claims.ExpiresAt <= time.Now().Local().Unix() {
		return nil, errors.New("yoken expired")
	}

	return claims, nil
}
