package middleware

import (
    "fmt"
    "net/http"

	"github.com/ducnv194023/shoppe_be_go/pkg/errors"
	"github.com/ducnv194023/shoppe_be_go/pkg/response"


)

// Middleware xử lý panic và lỗi
func ErrorHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                appErr := &errors.AppError{
                    Code:    http.StatusInternalServerError,
                    Message: "Lỗi hệ thống",
                    Err:     fmt.Errorf("%v", err),
                }
                response.RespondWithError(w, appErr)
            }
        }()
        next.ServeHTTP(w, r)
    })
}