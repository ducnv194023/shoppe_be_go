package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Message: message,
		Data: data,
	})	
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Message: message,
	})
}