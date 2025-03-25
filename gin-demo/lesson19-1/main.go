package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:youngking98@tcp(127.0.0.1:3307)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	//u1 := UserInfo{1, "迪丽热巴", "女", "跳舞"}
	//u2 := UserInfo{2, "章若楠", "女", "唱歌"}

	// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)

	// 查询记录
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%+v\n", u)

	// 更新记录
	db.Model(&u).Where("id = ?", 1).Update("hobby", "干饭")
	// 删除
	db.Delete(&u)
}
