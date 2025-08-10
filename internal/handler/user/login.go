package user

import (
	"encoding/json"
	"net/http"
	"time"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/eigakan/nats-shared/model"
	"github.com/eigakan/nats-shared/topics"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginHttpRequestDTO struct {
	Login    string `json:"login" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=5"`
}

type LoginHttpResponseDTO struct {
	Token string `json:"token"`
}

func (h *UserHanlders) makeJwtClaim(login string) jwt.MapClaims {
	return jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * time.Duration(h.JwtConf.ExpHours)).Unix(),
	}

}

func (h *UserHanlders) Login(c *gin.Context) {
	var reqDto LoginHttpRequestDTO
	if err := c.ShouldBindJSON(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error whiel parsuing request"})
		return
	}

	res, err := h.nc.Request(
		topics.UserCheckPassword,
		dto.CheckPasswordRequestDTO{
			Login:    reqDto.Login,
			Password: reqDto.Password,
		},
		2*time.Second,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth error"})
		return
	}

	var resDto model.NatsResponse[dto.CheckPasswordResponseDTO]
	if err := json.Unmarshal(res.Data, &resDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error1"})
		return
	}

	if !resDto.Data.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error2"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, h.makeJwtClaim(reqDto.Login))
	tokenStr, _ := token.SignedString([]byte(h.JwtConf.Secret))

	c.JSON(http.StatusOK, LoginHttpResponseDTO{
		Token: tokenStr,
	})
}
