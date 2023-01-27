package handlers

import (
	"net/http"
	"todolist_mux/model/apiRequest"
)

type TodolistHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request, request apiRequest.TodolistCreateRequest) error
	FindAllTodo(w http.ResponseWriter, r *http.Request) error
	FindByIdTodo(w http.ResponseWriter, r *http.Request) error
	UpdateTodo(w http.ResponseWriter, r *http.Request, request apiRequest.TodolistUpdateRequest) error
	UpdateStatusTodo(w http.ResponseWriter, r *http.Request, request apiRequest.TodolistUpdateStatusRequest) error
	DeleteTodo(w http.ResponseWriter, r *http.Request) error
}
