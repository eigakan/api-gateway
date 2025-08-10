package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/eigakan/api-gateway/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type myClaims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

type AuthMiddleware struct {
	jwtConf config.JwtConfig
}

func NewAuthMiddleware(jwtConf config.JwtConfig) *AuthMiddleware {
	return &AuthMiddleware{jwtConf: jwtConf}
}

func (m *AuthMiddleware) extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("No jwt token provided")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("Wrong header format")
	}

	return jwtToken[1], nil
}

func (m *AuthMiddleware) isTokenValid(jwtToken string) string {
	token, err := jwt.ParseWithClaims(jwtToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.jwtConf.Secret), nil
	})

	if err != nil {
		log.Println("Parse error:", err)
		return ""
	}

	claims, ok := token.Claims.(*myClaims)

	if !ok {
		return ""
	}

	if token.Valid {
		return claims.Login
	}

	return ""
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := m.extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		login := m.isTokenValid(jwtToken)

		if login == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
		}

		c.Set("login", login)
		c.Next()
	}
}
