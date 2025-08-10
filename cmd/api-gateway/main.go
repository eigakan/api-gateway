package main

import (
	"fmt"

	"github.com/eigakan/api-gateway/config"
	"github.com/eigakan/api-gateway/internal/handler/user"
	"github.com/eigakan/api-gateway/internal/router"
	nats_client "github.com/eigakan/nats-shared/client"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()

	nc, err := nats_client.NewClient(config.Nats.Host, config.Nats.Port)
	if err != nil {
		fmt.Printf("Error connecting to NATS: %v\n", err)
	}
	defer nc.Drain()

	r := gin.Default()

	userHandlers := user.NewUserHandlers(nc, config.Jwt)
	router.NewUserRouter(userHandlers).RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", config.Http.Port))
}
