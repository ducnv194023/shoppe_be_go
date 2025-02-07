package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ducnv194023/shoppe_be_go/pkg/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log error
				fmt.Printf("panic occurred: %v\n", err)

				appErr := &errors.AppError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
					Err:     fmt.Errorf("%v", err),
				}

				c.JSON(appErr.Code, gin.H{
					"error":   appErr.Message,
					"details": appErr.Err.Error(),
					"code":    appErr.Code,
					"success": false,
				})

				c.Abort()
			}
		}()

		c.Next()
	}
}
