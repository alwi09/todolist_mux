package apiRequest

type TodolistUpdateRequest struct {
	Id          int    `json:"id"`
	Title       string `validate:"required,min=2" json:"title"`
	Description string `validate:"required,min=2" json:"description"`
}
