package collection

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CollectionCreateRequest struct {
	Name      string `json:"name"`
	LibraryId string `json:"library_id"`
	ParentId  string `json:"parent_id,omitempty"`
}

func (request CollectionCreateRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.LibraryId, validation.Required, is.UUID),
		validation.Field(&request.ParentId, is.UUID),
	)
}
