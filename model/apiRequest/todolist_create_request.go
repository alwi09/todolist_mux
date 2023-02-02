package apiRequest

type TodolistCreateRequest struct {
	Title       string `validate:"required,min=2" json:"title"`
	Description string `validate:"required,min=2" json:"description"`
}
