package error_handling

import (
	"net/http"
	"todolist_mux/model/domain"
)

func ErrorHandlingInternalServerError(err error) domain.Response {

	apiResponse := domain.Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Data:    err,
	}
	return apiResponse
}

func ErrorHandlingStatusNotFound(err error) domain.Response {

	apiResponse := domain.Response{
		Status:  http.StatusNotFound,
		Message: "Not found",
		Data:    err,
	}
	return apiResponse
}

func ErrorHandlingBadRequest(err error) domain.Response {

	apiResponse := domain.Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Data:    err,
	}
	return apiResponse
}
