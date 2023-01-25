package mysql

import (
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"todolist_mux/internal/configuration"
)

func ConnectDB() (*sql.DB, error) {

	var config configuration.Config
	db, err := os.Open("config.json")
	if err != nil {
		logrus.Fatal(err.Error())
	}

	defer db.Close()

	jsonParse := json.NewDecoder(db)
	if err = jsonParse.Decode(&config); err != nil {
		logrus.Fatal(err.Error())
	}

	return nil, err
}
