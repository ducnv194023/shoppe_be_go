package errors

import (
    "fmt"
    "net/http"
)

type AppError struct {
    Code    int    `json:"code"`    
    Message string `json:"message"` 
    Err     error  `json:"-"`      
}

func (e *AppError) Error() string {
    return fmt.Sprintf("code: %d, message: %s, error: %v", e.Code, e.Message, e.Err)
}

func NewBadRequestError(message string) *AppError {
    return &AppError{
        Code:    http.StatusBadRequest,
        Message: message,
    }
}

func NewUnauthorizedError(message string) *AppError {
    return &AppError{
        Code:    http.StatusUnauthorized,
        Message: message,
    }
}

func NewInternalError(err error) *AppError {
    return &AppError{
        Code:    http.StatusInternalServerError,
        Message: "Internal server error",
        Err:     err,
    }
}