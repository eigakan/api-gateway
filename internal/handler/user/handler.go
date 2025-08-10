package user

import (
	nats_client "github.com/eigakan/nats-shared/client"
)

type UserHanlders struct {
	nc *nats_client.Client
}

func NewUserHandlers(nc *nats_client.Client) *UserHanlders {
	return &UserHanlders{nc: nc}
}
