package mysql

import (
	"database/sql"
)

func ConnectDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "alwi09"
	dbPass := "alwiirfani11"
	dbName := "rest_full_api_todolist"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db, err
}
