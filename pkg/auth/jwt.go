package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJWT(email string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "tixmaster",
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
	})

	jwt, err := claims.SignedString([]byte("itsasecretcuh"))
	if err != nil {
		return "", err
	}

	return jwt, nil
}
