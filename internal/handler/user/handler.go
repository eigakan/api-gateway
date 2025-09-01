package user

import (
	"github.com/eigakan/api-gateway/internal/pkg"
	nats_client "github.com/eigakan/nats-shared/client"
)

type UserHanlders struct {
	nc  *nats_client.Client
	Jwt *pkg.Jwt
}

func NewUserHandlers(nc *nats_client.Client, j *pkg.Jwt) *UserHanlders {
	return &UserHanlders{nc: nc, Jwt: j}
}
