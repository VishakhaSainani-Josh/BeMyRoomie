package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "BeMyRoomie",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err

	}
}
