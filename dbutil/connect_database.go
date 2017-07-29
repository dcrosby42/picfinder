package dbutil

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDatabase() (*sqlx.DB, error) {
	username := "picfinder"
	password := "picfinder"
	dbname := "picfinder"
	connStr := fmt.Sprintf("%s:%s@(localhost:3306)/%s", username, password, dbname)
	return sqlx.Connect("mysql", connStr)
}
