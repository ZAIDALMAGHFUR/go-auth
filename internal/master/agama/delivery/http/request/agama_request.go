package request

type AgamaRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
}
