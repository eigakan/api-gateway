package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/eigakan/nats-shared/model"
	"github.com/eigakan/nats-shared/topics"
	"github.com/gin-gonic/gin"
)

type RegisterHttpRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
	Login    string `json:"login" binding:"required,min=3,max=20"`
}

type RegisterHttpResponseDTO struct {
	Ok bool `json:"ok"`
}

func (h *UserHanlders) Register(c *gin.Context) {
	var reqDto RegisterHttpRequestDTO
	if err := c.ShouldBindJSON(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.nc.Request(
		topics.UserCreate,
		dto.CreateUserRequestDTO{
			Login:    reqDto.Login,
			Password: reqDto.Password,
			Email:    reqDto.Email,
		},
		2*time.Second,
	)
	// Timeout or no responders or something
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	var resDto model.NatsResponse[dto.CreateUserResponseDTO]
	if err := json.Unmarshal(res.Data, &resDto); err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	if !resDto.Status {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	c.JSON(http.StatusOK, RegisterHttpResponseDTO{Ok: true})
}
