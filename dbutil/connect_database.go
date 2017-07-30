package dbutil

import (
	"fmt"

	"github.com/dcrosby42/picfinder/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDatabase(cfg config.DbConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("%s:%s@(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	return sqlx.Connect("mysql", connStr)
}
