package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"todolist_mux/internal/database/mysql"
)

func main() {
	router := mux.NewRouter()

	db, err := mysql.ConnectDB()
	if err != nil {
		logrus.Error(err)
	}
	defer db.Close()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("HELLO GUYS")
	})

	log.Fatal(http.ListenAndServe(":3000", router))
}
