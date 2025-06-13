package pkg

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func StructErrorMap(err error) map[string]string {
	errors := map[string]string{}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			errors[fieldErr.Field()] = validationMessage(fieldErr)
		}
	}

	return errors
}

func ValidateStruct(s interface{}) (bool, map[string]string) {
	err := validate.Struct(s)
	if err != nil {
		return false, StructErrorMap(err)
	}
	return true, nil
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "bool":
		return "must be a boolean value"
	case "string":
		return "must be a string"
	case "int":
		return "must be an integer"
	case "float":
		return "must be a float"
	case "time":
		return "must be a valid time"
	case "min":
		return "must be at least " + fe.Param() + " characters"
	case "max":
		return "must be at most " + fe.Param() + " characters"
	case "unique":
		return "must contain unique values"
	case "date":
		return "must be a valid date"
	case "datetime":
		return "must be a valid datetime"
	default:
		return "is invalid"
	}
}
