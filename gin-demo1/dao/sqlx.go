package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBs *sqlx.DB

func InitSqlx() (err error) {
	// 连接
	dsn := "root:youngking98@tcp(127.0.0.1:3307)/db1?charset=utf8mb4&parseTime=true&loc=Local"
	DBs, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("conncet failed, err:%v\n", err)
		return err
	}

	DBs.SetMaxOpenConns(20)
	DBs.SetMaxIdleConns(5)
	return nil
}
