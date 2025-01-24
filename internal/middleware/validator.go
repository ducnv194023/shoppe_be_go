package validator

import (
    "github.com/go-playground/validator/v10"
    "regexp"
    "unicode"
)

var validate *validator.Validate

func InitValidator() {
    validate = validator.New()
    
    // Register custom validators
    validate.RegisterValidation("password", validatePassword)
    validate.RegisterValidation("fullname", validateName)
    validate.RegisterValidation("phone", validatePhone)
}

func GetValidator() *validator.Validate {
    return validate
}

// Base validators
func validatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    
    var (
        hasUpper   = false
        hasLower   = false
        hasNumber  = false
        hasSpecial = false
    )
    
    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsNumber(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }
    
    return hasUpper && hasLower && hasNumber && hasSpecial
}

func validateName(fl validator.FieldLevel) bool {
    nameRegex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
    return nameRegex.MatchString(fl.Field().String())
}

func validatePhone(fl validator.FieldLevel) bool {
    phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
    return phoneRegex.MatchString(fl.Field().String())
}