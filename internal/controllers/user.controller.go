package controllers

import (
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/responses"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	user_service *services.UserService
}

// controller -> service -> repository -> models -> database
func NewUserController() *UserController {
	return &UserController{
		user_service: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	responses.SuccessResponse(ctx, responses.ResCodeSuccess, []string{"dinh", "nguyen", "quang"})
}
