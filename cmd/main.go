package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"todolist_mux/internal/database/mysql"
	"todolist_mux/internal/handlers"
	"todolist_mux/model/apiRequest"
)

func main() {
	router := mux.NewRouter()

	db, err := mysql.ConnectDB()
	if err != nil {
		logrus.Error(err)
	}
	defer db.Close()

	handler := handlers.NewTodolistHandlerImpl(db)

	// CREATE
	router.HandleFunc("/api/createTodolist", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateTodo(w, r, &apiRequest.TodolistCreateRequest{})
	}).Methods(http.MethodPost)
	// FIND ALL
	router.HandleFunc("/api/findAllTodolist", func(w http.ResponseWriter, r *http.Request) {
		handler.FindAllTodo(w, r)
	}).Methods(http.MethodGet)
	// FIND BY ID
	router.HandleFunc("/api/findByIdTodolist/{todolistId}", func(w http.ResponseWriter, r *http.Request) {
		handler.FindByIdTodo(w, r)
	}).Methods(http.MethodGet)
	// UPDATE
	router.HandleFunc("/api/updateTodolist/{todolistId}", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateTodo(w, r, &apiRequest.TodolistUpdateRequest{})
	}).Methods(http.MethodPut)
	// UPDATE STATUS
	router.HandleFunc("/api/updateStatusTodolist/{todolistId}", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateStatusTodo(w, r, &apiRequest.TodolistUpdateStatusRequest{})
	}).Methods(http.MethodPut)
	// DELETE
	router.HandleFunc("/api/deleteTodolist/{todolistId}", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteTodo(w, r)
	}).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":1234", router))
}
