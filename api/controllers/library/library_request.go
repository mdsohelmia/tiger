package library

import validation "github.com/go-ozzo/ozzo-validation"

type CreateLibraryRequest struct {
	Title string `json:"title"`
}

func (request CreateLibraryRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required),
	)
}
