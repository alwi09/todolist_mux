package error_handling

import (
	"net/http"
	"todolist_mux/helper"
	"todolist_mux/model/domain"
)

func ErrorHandlingInternalServerError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := domain.Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Data:    err,
	}

	w.WriteHeader(apiResponse.Status)
	helper.WriteFromRequestBody(w, apiResponse)
}

func ErrorHandlingStatusNotFound(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := domain.Response{
		Status:  http.StatusNotFound,
		Message: "Not found",
		Data:    err,
	}

	w.WriteHeader(apiResponse.Status)
	helper.WriteFromRequestBody(w, apiResponse)
}

func ErrorHandlingBadRequest(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := domain.Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Data:    err,
	}

	w.WriteHeader(apiResponse.Status)
	helper.WriteFromRequestBody(w, apiResponse)

}
