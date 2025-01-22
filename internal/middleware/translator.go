package validator

import (
    "fmt"
    "strings"
    "github.com/go-playground/validator/v10"
)

func TranslateValidationErrors(err error) string {
    var errMsgs []string
    
    validatorErrs, ok := err.(validator.ValidationErrors)
    if !ok {
        return err.Error()
    }

    for _, e := range validatorErrs {
        errMsgs = append(errMsgs, getErrorMsg(e))
    }
    
    return strings.Join(errMsgs, ", ")
}

func getErrorMsg(e validator.FieldError) string {
    switch e.Tag() {
    case "required":
        return fmt.Sprintf("%s is required", e.Field())
    case "email":
        return "Invalid email format"
    case "min":
        return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
    case "max":
        return fmt.Sprintf("%s must not exceed %s characters", e.Field(), e.Param())
    case "password":
        return "Password must contain uppercase, lowercase, number and special character"
    case "name":
        return "Name can only contain letters and spaces"    
    case "phone":
        return "Invalid phone number format"
    default:
        return fmt.Sprintf("%s is not valid", e.Field())
    }
}