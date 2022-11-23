package core

import (
	"sort"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

type invalidField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validation Errors is Validation errors message formatting.
// array of object errors

func ValidationErrors(validationErrors validation.Errors) []invalidField {
	var errors []invalidField
	var fields []string

	for field := range validationErrors {
		fields = append(fields, field)
	}
	sort.Strings(fields)

	for _, field := range fields {
		errors = append(errors, invalidField{
			Field:   strings.ToLower(field),
			Message: validationErrors[field].Error(),
		})
	}

	return errors
}
