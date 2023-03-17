package utils

import "github.com/go-playground/validator/v10"

type ErrorMessageEmpty struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMessageEmpty(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "oneof":
		return "these fields should be user and admin"
	}

	return "Unknown error"
}
