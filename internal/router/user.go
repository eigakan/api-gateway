package router

import (
	"github.com/eigakan/api-gateway/internal/handler/user"
	"github.com/eigakan/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	uh *user.UserHanlders
}

func NewUserRouter(uh *user.UserHanlders) *UserRouter {
	return &UserRouter{uh: uh}
}

func (ur *UserRouter) RegisterRoutes(r *gin.Engine) *gin.Engine {
	ag := r.Group("/auth")
	ag.Use(middleware.NewResponseMiddleware().Handler())
	// Protected routes
	prg := ag.Group("/")
	prg.Use(middleware.NewAuthMiddleware(ur.uh.Jwt).Handler())

	prg.GET("/me", ur.uh.Me)

	// Auth routes
	ag.POST("/register", ur.uh.Register)
	ag.POST("/login", ur.uh.Login)

	return r
}
