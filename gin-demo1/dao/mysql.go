package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init
	"time"
)

var DB *sql.DB

func InitMySQL() (err error) {
	dsn := "root:youngking98@tcp(127.0.0.1:3307)/db1"
	// 这里并不会连接，而是检查连接中的参数是否正确
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 尝试连接数据库
	err = DB.Ping()
	if err != nil {
		fmt.Printf("connect to mysql failed, err:%v\n", err)
		return err
	}

	// 设置数据库配置
	DB.SetMaxOpenConns(100)                 // 最大连接数
	DB.SetMaxIdleConns(10)                  // 最大空闲连接数
	DB.SetConnMaxLifetime(time.Second * 60) // 最大空闲时间
	return nil
}
