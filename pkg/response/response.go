package response

import (
	"github.com/gin-gonic/gin"
	
	"github.com/ducnv194023/shoppe_be_go/pkg/errors"

)

type Response struct {
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Code    int         `json:"code"`
    Success bool        `json:"success"`
}

func ResponseSuccess(c *gin.Context, code int, data interface{}) {
    c.JSON(code, Response{
        Data:    data,
        Code:    code,
        Success: true,
    })
}

func ResponseError(c *gin.Context, err *errors.AppError) {
    c.JSON(err.Code, Response{
        Error:   err.Message,
        Code:    err.Code, 
        Success: false,
    })
}