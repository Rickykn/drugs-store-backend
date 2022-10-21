package handlers

import "github.com/Rickykn/drugs-store-backend.git/services"

type Handler struct {
	userService services.UserService
	drugService services.DrugService
}

type HandlerConfig struct {
	UserService services.UserService
	DrugService services.DrugService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		userService: c.UserService,
		drugService: c.DrugService,
	}
}
