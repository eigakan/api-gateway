package user

import (
	"github.com/nats-io/nats.go"
)

type UserHanlders struct {
	nc *nats.Conn
}

func NewUserHandlers(nc *nats.Conn) *UserHanlders {
	return &UserHanlders{nc: nc}
}
