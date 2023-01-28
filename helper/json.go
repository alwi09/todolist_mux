package helper

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		logrus.Error(err, "Read from request body failed")
	}

	return err
}

func WriteFromRequestBody(w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		logrus.Error(err, "Write from request body failed")

	}

	return err
}
