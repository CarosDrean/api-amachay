package middleware

import (
	"errors"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(userResult models.ClaimUser) (string, error) {
	claims := models.Claim{
		ClaimUser: userResult,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "Admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(t string) (models.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &models.Claim{}, verifyFunction)
	if err != nil {
		return models.Claim{}, err
	}
	if !token.Valid {
		return models.Claim{}, errors.New("token no válido")
	}

	claim, ok := token.Claims.(*models.Claim)
	if !ok {
		return models.Claim{}, errors.New("no se pudo obtener los claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return publicKey, nil
}
