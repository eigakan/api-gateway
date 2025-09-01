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

type LoginHttpRequestDTO struct {
	Login    string `json:"login" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=5"`
}

type LoginHttpResponseDTO struct {
	Token string `json:"token"`
}

func (h *UserHanlders) Login(c *gin.Context) {
	var reqDto LoginHttpRequestDTO
	if err := c.ShouldBindJSON(&reqDto); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Login or password not valid"))
		return
	}

	res, err := h.nc.Request(
		topics.UserGetByPassword,
		dto.GetUserByPasswordRequestDTO{
			Login:    reqDto.Login,
			Password: reqDto.Password,
		},
		2*time.Second,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	var resDto model.NatsResponse[dto.GetUserByPasswordResponseDTO]
	if err := json.Unmarshal(res.Data, &resDto); err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	if !resDto.Status {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
		return
	}

	token, err := h.Jwt.Generate(resDto.Data.ID)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New(""))
	}

	c.JSON(http.StatusOK, LoginHttpResponseDTO{
		Token: token,
	})
}
