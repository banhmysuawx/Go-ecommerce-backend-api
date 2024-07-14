package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"` // 20001: Success, 20003: Param invalid, 500: Error
	Message string   `json:"message"` // Message
	Data interface{} `json:"data"` // Data

}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:   code,
		Message: GetMsg(code),
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:   code,
		Message: GetMsg(code),
		Data: data,
	})
}