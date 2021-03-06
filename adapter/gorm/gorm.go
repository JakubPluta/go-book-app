package gorm

import (
	"fmt"
	"myapp/config"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


func New(conf *config.Config) (*gorm.DB, error) {
	cfg := &mysql.Config{
        Net:                  "tcp",
        Addr:                 fmt.Sprintf("%v:%v", conf.Db.Host, conf.Db.Port),
        DBName:               conf.Db.DbName,
        User:                 conf.Db.Username,
        Passwd:               conf.Db.Password,
        AllowNativePasswords: true,
        ParseTime:            true,
	}
	return gorm.Open("mysql", cfg.FormatDSN())
}