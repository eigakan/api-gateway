package user

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"

// 	dto "github.com/eigakan/nats-shared/dto/user"
// 	"github.com/eigakan/nats-shared/model"
// 	"github.com/eigakan/nats-shared/topics"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// type MeHttpRequestDTO struct {
// 	Login    string `json:"login" binding:"required,min=3,max=20"`
// 	Password string `json:"password" binding:"required,min=5"`
// }

// type MeHttpResponseDTO struct {
// 	Token string `json:"token"`
// }

// func (h *UserHanlders) Login(c *gin.Context) {
// 	var reqDto LoginHttpRequestDTO
// 	if err := c.ShouldBindJSON(&reqDto); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Error whiel parsuing request"})
// 		return
// 	}

// 	np := dto.CheckPasswordRequestDTO{
// 		Login:    reqDto.Login,
// 		Password: reqDto.Password,
// 	}

// 	res, err := h.nc.Request(topics.UserCheckPassword, np, 2*time.Second)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth error"})
// 		return
// 	}

// 	var resDto model.NatsResponse[dto.CheckPasswordResponseDTO]
// 	if err := json.Unmarshal(res.Data, &resDto); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	if !resDto.Data.Valid {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, h.makeJwtClaim(reqDto.Login))
// 	tokenStr, _ := token.SignedString(h.jwtConf.Secret)

// 	c.JSON(http.StatusOK, LoginHttpResponseDTO{
// 		Token: tokenStr,
// 	})
// }
