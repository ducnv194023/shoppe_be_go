package errors

import (
    "fmt"
    "net/http"
)

// AppError định nghĩa cấu trúc lỗi chung cho ứng dụng
type AppError struct {
    Code    int    `json:"code"`    // Mã lỗi HTTP
    Message string `json:"message"` // Thông báo lỗi cho người dùng
    Err     error  `json:"-"`      // Lỗi gốc (ẩn khỏi response)
}

func (e *AppError) Error() string {
    return fmt.Sprintf("code: %d, message: %s, error: %v", e.Code, e.Message, e.Err)
}

// Lỗi cơ sở dữ liệu
func NewDBError(err error) *AppError {
    return &AppError{
        Code:    http.StatusInternalServerError,
        Message: "Lỗi cơ sở dữ liệu",
        Err:     err,
    }
}

// Lỗi validation
func NewValidationError(message string) *AppError {
    return &AppError{
        Code:    http.StatusBadRequest,
        Message: message,
    }
}

// Lỗi không tìm thấy tài nguyên
func NewNotFoundError(resource string) *AppError {
    return &AppError{
        Code:    http.StatusNotFound,
        Message: fmt.Sprintf("Không tìm thấy %s", resource),
    }
}

// Lỗi xác thực
func NewAuthenticationError(message string) *AppError {
    return &AppError{
        Code:    http.StatusUnauthorized,
        Message: message,
    }
}