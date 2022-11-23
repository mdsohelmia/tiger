package video

import validation "github.com/go-ozzo/ozzo-validation"

type VideoCreateRequest struct {
	Title    string `json:"name"`
	Duration string `json:"duration"`
	// array of objects
	Captions []Caption `json:"captions"`
}

type Caption struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (c Caption) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.Start, validation.Required),
		validation.Field(&c.End, validation.Required),
	)
}

func (v VideoCreateRequest) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Title, validation.Required),
		validation.Field(&v.Duration, validation.Required),
		validation.Field(&v.Captions, validation.Required),
	)
}
