package router

import (
	"github.com/eigakan/api-gateway/internal/handler/user"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	uh *user.UserHanlders
}

func NewUserRouter(uh *user.UserHanlders) *UserRouter {
	return &UserRouter{uh: uh}
}

func (ur *UserRouter) RegisterRoutes(r *gin.Engine) *gin.Engine {
	// Auth routes
	ag := r.Group("/auth")
	ag.POST("/register", ur.uh.Register)
	ag.POST("/login", ur.uh.Login)

	return r
}
