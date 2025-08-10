package user

import (
	"github.com/eigakan/api-gateway/config"
	nats_client "github.com/eigakan/nats-shared/client"
)

type UserHanlders struct {
	nc      *nats_client.Client
	JwtConf config.JwtConfig
}

func NewUserHandlers(nc *nats_client.Client, jwtConf config.JwtConfig) *UserHanlders {
	return &UserHanlders{nc: nc, JwtConf: jwtConf}
}
