package apiRequest

type TodolistUpdateStatusRequest struct {
	Id     int  `json:"id"`
	Status bool `json:"status"`
}
