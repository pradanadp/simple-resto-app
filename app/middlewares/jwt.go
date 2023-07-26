package middlewares

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/pradanadp/simple-resto-app/app/config"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT),
		SigningMethod: "HS256",
	})
}

func GenerateToken(customerID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["customerID"] = customerID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT))
}

func ExtractToken(e echo.Context) (string, error) {
	customer := e.Get("customer").(*jwt.Token)
	if customer.Valid {
		claims := customer.Claims.(jwt.MapClaims)
		customerID := claims["customerID"].(string)
		return customerID, nil
	}
	return "", errors.New("failed to extract jwt-token")
}
