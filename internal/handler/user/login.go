package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/eigakan/nats-shared/model"
	"github.com/eigakan/nats-shared/topics"
	"github.com/gin-gonic/gin"
)

type LoginHttpRequestDTO struct {
	Login    string `json:"login" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=5"`
}

func (h *UserHanlders) Login(c *gin.Context) {
	var reqDto LoginHttpRequestDTO
	if err := c.ShouldBindJSON(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	np := dto.LoginRequestDTO{
		Login:    reqDto.Login,
		Password: reqDto.Password,
	}

	natsData, err := json.Marshal(np)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	res, err := h.nc.Request(topics.UserLoginTopic, natsData, 2*time.Second)
	// Timeout or no responders or something
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var resDto model.NatsResponse[dto.LoginResponseDTO]
	if err := json.Unmarshal(res.Data, &resDto); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if resDto.Status {
		c.JSON(http.StatusOK, resDto)
	} else {
		c.JSON(http.StatusBadRequest, resDto)
	}
}
