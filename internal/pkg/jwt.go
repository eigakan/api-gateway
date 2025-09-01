package pkg

import (
	"fmt"

	"time"

	"github.com/eigakan/api-gateway/config"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	config *config.JwtConfig
}

func NewJwt(config *config.JwtConfig) *Jwt {
	return &Jwt{
		config: config,
	}
}

type Claims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

func (j *Jwt) makeJwtMapClaims(userID uint) jwt.MapClaims {
	return jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * time.Duration(j.config.ExpHours)).Unix(),
	}
}

func (j *Jwt) Generate(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.makeJwtMapClaims(userId))
	return token.SignedString([]byte(j.config.Secret))
}

func (j *Jwt) Verify(token string) (uint, bool) {
	parsed, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return 0, false
	}

	claims, ok := parsed.Claims.(*Claims)

	if !ok {
		return 0, false
	}

	if parsed.Valid {
		return claims.UserID, true
	}

	return 0, false
}
