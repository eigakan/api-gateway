package router

import "github.com/eigakan/api-gateway/internal/handler/user"

type UserRouter struct {
	uh *user.UserHanlders
}

func NewUserRouter(uh *user.UserHanlders) *UserRouter {
	return &UserRouter{uh: uh}
}

func (ur *UserRouter) RegisterRoutes() {
	// Here you would typically use a router to register the user handlers
	// For example:
	// r.POST("/register", ur.uh.Register)
}
