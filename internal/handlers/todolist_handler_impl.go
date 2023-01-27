package handlers

import (
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todolist_mux/error_handling"
	"todolist_mux/helper"
	"todolist_mux/model/apiRequest"
	"todolist_mux/model/domain"
)

type TodolistHandlerImpl struct {
	DB *sql.DB
}

func NewTodolistHandlerImpl(DB *sql.DB) *TodolistHandlerImpl {
	return &TodolistHandlerImpl{
		DB: DB,
	}
}

func (handler *TodolistHandlerImpl) CreateTodo(w http.ResponseWriter, r *http.Request, request *apiRequest.TodolistCreateRequest) error {
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		logrus.Error(err)
		helper.PanicIfError(err)
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		error_handling.ErrorHandlingBadRequest(w, err)
		logrus.Error(err)
		return nil
	}

	exec, err := handler.DB.Exec("INSERT INTO todolist_mux(title,description) VALUES (?,?)", request.Title, request.Description)
	if err != nil {
		error_handling.ErrorHandlingInternalServerError(w, err)
		logrus.Error(err)
	}

	id, err := exec.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return nil
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Create todolist Successfully " + strconv.FormatInt(id, 10),
		Data:    nil,
	}

	logrus.Info("Create todolist successfully")

	err = helper.WriteFromRequestBody(w, apiResponse)
	if err != nil {
		logrus.Error(err)
	}

	return nil
}

func (handler *TodolistHandlerImpl) FindAllTodo(w http.ResponseWriter, r *http.Request) error {
	var todolist apiRequest.ApiResponse
	var sliceTodolist []apiRequest.ApiResponse

	rows, err := handler.DB.Query("SELECT id,title,description,status FROM todolist_mux")
	if err != nil {
		error_handling.ErrorHandlingInternalServerError(w, err)
		logrus.Error(err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(&todolist.Id, &todolist.Title, &todolist.Description, &todolist.Status)
		if err != nil {
			logrus.Error(err)
		} else {
			sliceTodolist = append(sliceTodolist, todolist)
		}
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    sliceTodolist,
	}

	logrus.Info("Find all Successfully")

	helper.WriteFromRequestBody(w, apiResponse)

	return nil
}

func (handler *TodolistHandlerImpl) FindByIdTodo(w http.ResponseWriter, r *http.Request) error {
	var todolist domain.Todolist
	var arrTodolist []domain.Todolist

	params := mux.Vars(r)
	todolistId := params["todolistId"]

	var count int
	if err := handler.DB.QueryRow("SELECT COUNT(*) FROM todolist_mux WHERE id=?", todolistId).Scan(&count); err != nil {
		error_handling.ErrorHandlingInternalServerError(w, err)
		logrus.Error("Failed to check Todolist ", err)
		return nil
	}

	// jika id tidak sama dengan id yang di kirim dalam request
	if count == 0 {
		error_handling.ErrorHandlingStatusNotFound(w, errors.New("id not found in database"))
		logrus.Error(errors.New("Id not found"))
		return nil
	}

	// get todolist dari database berdasarkan id
	rows, err := handler.DB.Query("SELECT id,title,description,status FROM todolist_mux WHERE id=?", todolistId)
	if err != nil {
		error_handling.ErrorHandlingInternalServerError(w, err)
		logrus.Error("Error get id ", err)
		return nil
	}

	// perulangan sebanyak data yang di cari di database
	for rows.Next() {
		err = rows.Scan(&todolist.Id, &todolist.Title, &todolist.Description, &todolist.Status)
		if err != nil {
			logrus.Fatal(err)
		} else {
			arrTodolist = append(arrTodolist, todolist)
		}
	}

	apiResponse := domain.Response{
		Status:  http.StatusOK,
		Message: "Get todolist by id successfully",
		Data:    arrTodolist,
	}

	logrus.Info("Success get todolist by id")

	helper.WriteFromRequestBody(w, apiResponse)

	return nil
}

func (handler *TodolistHandlerImpl) UpdateTodo(w http.ResponseWriter, r *http.Request, request apiRequest.TodolistUpdateRequest) (domain.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *TodolistHandlerImpl) UpdateStatusTodo(w http.ResponseWriter, r *http.Request, request apiRequest.TodolistUpdateStatusRequest) (domain.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *TodolistHandlerImpl) DeleteTodo(w http.ResponseWriter, r *http.Request, todolistId int) error {
	//TODO implement me
	panic("implement me")
}
