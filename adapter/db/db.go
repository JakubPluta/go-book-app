package db

import (
	"database/sql"
	"fmt"
	"myapp/config"

	"github.com/go-sql-driver/mysql"
)


func New(config *config.Config) (*sql.DB, error){
	cfg := &mysql.Config{
		Net: "tcp",
		Addr : fmt.Sprintf("%v:%v", config.Db.Host, config.Db.Port),
		DBName: config.Db.DbName,
		   User:                 config.Db.Username,
        Passwd:               config.Db.Password,
		     AllowNativePasswords: true,
        ParseTime:            true,
	}
	return sql.Open("mysql", cfg.FormatDSN())
} 