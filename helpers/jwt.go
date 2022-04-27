package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := parseToken.SignedString([]byte(secretKey))

	return token
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errorResponse := errors.New("sign in to proceed")
	header := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(header, "Bearer")

	if !bearer {
		return nil, errorResponse
	}

	stringToken := strings.Split(header, " ")[1]

	parseToken, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

	_, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok && !parseToken.Valid {
		return nil, errors.New("invalid token")
	}

	return parseToken.Claims.(jwt.MapClaims), nil
}
