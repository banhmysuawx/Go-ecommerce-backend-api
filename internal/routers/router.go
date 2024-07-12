package routers

import (
	"github.com/banhmysuawx/Go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user/:id", controller.NewUserController().GetUserByID)
	}

	return r
}

