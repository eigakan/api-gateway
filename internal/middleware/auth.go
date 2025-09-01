package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/eigakan/api-gateway/internal/pkg"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwt *pkg.Jwt
}

func NewAuthMiddleware(jwt *pkg.Jwt) *AuthMiddleware {
	return &AuthMiddleware{jwt: jwt}
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

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := m.extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New(""))
			return
		}

		userId, ok := m.jwt.Verify(jwtToken)

		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Invalid JWT token"))
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
