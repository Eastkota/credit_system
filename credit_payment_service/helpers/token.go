package helpers

import (
	"credit_system/auth_service/model"
	
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(id string, duration time.Duration, clientSecret string, owner *model.StoreOwner) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"token_id": id,
			"exp":      time.Now().Add(duration),
			"owner":     owner,
			"issuer":  "iBestTechnologies",	
		})
	tokenString, err := token.SignedString([]byte(clientSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
