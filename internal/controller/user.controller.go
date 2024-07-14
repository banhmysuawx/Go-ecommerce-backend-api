package controller

import (
	"github.com/banhmysuawx/Go-ecommerce-backend-api/internal/service"
	"github.com/banhmysuawx/Go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.ErrorResponse(c, 20001, "Invalid ID")
		return
	}
	user, err := uc.userService.GetInfoUserByID(id)
	if err != nil {
		response.ErrorResponse(c, 20002, "User not found")
		return
	}
	response.SuccessResponse(c, 20000, user)
}

