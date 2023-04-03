package helpers

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_SECRET")

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	// creates a new token with the specified signing method and claims.
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// creates and returns a complete, signed JWT.
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}
