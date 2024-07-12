package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "guest")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"Array": []string{"a", "b", "c"},
	})
}