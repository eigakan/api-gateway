package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/eigakan/nats-shared/model"
	nats_model "github.com/eigakan/nats-shared/model"
	"github.com/eigakan/nats-shared/topics"
	"github.com/gin-gonic/gin"
)

type MeHttpRequestDTO struct {
	Login    string `json:"login" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=5"`
}

type MeHttpResponseDTO struct {
	nats_model.User
}

func (h *UserHanlders) Me(c *gin.Context) {
	var reqDto LoginHttpRequestDTO
	if err := c.ShouldBindJSON(&reqDto); err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	userIdstr := c.GetString("userId")
	if userIdstr == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	userId, err := strconv.Atoi(userIdstr)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	res, err := h.nc.Request(topics.UserGet, dto.GetUserRequestDTO{UserID: uint(userId)}, 2*time.Second)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	var resDto model.NatsResponse[dto.GetUserResponseDTO]
	if err := json.Unmarshal(res.Data, &resDto); err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	if !resDto.Status {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	c.JSON(http.StatusOK, MeHttpResponseDTO{User: resDto.Data.User})
}
