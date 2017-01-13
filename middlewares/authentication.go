package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	config "github.com/spf13/viper"
	"net/http"
)

const Bearer = "Bearer"

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header().Get("Authorization")
			l := len(Bearer)
			he := echo.NewHTTPError(http.StatusUnauthorized)

			if len(auth) > l+1 && auth[:l] == Bearer {
				t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {

					// Always check the signing method
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// Return the key for validation
					return []byte(config.GetString("jwt_signature")), nil
				})
				if err == nil && t.Valid {
					// Store token claims in echo.Context
					c.Set("claims", t.Claims)
					return nil
				}
			}
			return he
		}
	}
}
