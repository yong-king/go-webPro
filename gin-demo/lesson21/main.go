package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID int64 `gorm:"primaryKey"`
	// 下面这种会造成当传入 “” 和 0 时也会是默认值
	//Name string `gorm:"default:'dlrb'"`
	//Age  int64  `gorm:"default:8"`
	// 方法一
	//Name *string `gorm:"default:'dlrb'"`
	//Age  *int64  `gorm:"default:8"`
	// 方法二
	Name sql.NullString `gorm:"default:'dlrb'"`
	Age  sql.NullInt64  `gorm:"default:8"`
}

func main() {
	db, err := gorm.Open("mysql", "root:youngking98@tcp(127.0.0.1:3307)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	u := User{}
	//u := User{Name: new(string), Age: new(int64)}
	//u := &User{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: true}}
	fmt.Println(db.NewRecord(&u)) // 查看主键是否存在
	db.Create(&u)
	fmt.Println(db.NewRecord(&u))
}
