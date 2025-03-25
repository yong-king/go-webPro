package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

//type User struct {
//	gorm.Model // 默认字段
//	Name       string
//}

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` // 唯一
	Role         string  `gorm:"size:255"`                       // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"`                // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                 // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`                     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                              // 忽略本字段
}

type Animal struct {
	AnimalID uint   `gorm:"primarykey"`    // 主键
	Name     string `gorm:"column:mingzi"` // 修改列名
}

// 修改表明
func (Animal) TableName() string {
	return "dong_wus"
}

func main() {
	dsn := "root:youngking98@tcp(127.0.0.1:3307)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,   // V2让表名不加 s
			TablePrefix:   "ssm_", // v2 添加前缀
			// 之前的user表不会被删除，而是新建一个ssm_user
		},
	})
	if err != nil {
		panic(err)
	}
	//db.SingularTable(true) //Gorm V1中让表名不加s
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "prefix_" + defaultTableName // 添加前缀
	// }   //Gorm V1中自定义表名

	// 自动迁移
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})                   // 会创建一个dong_wus的表
	db.Table("classmates").AutoMigrate(&User{}) // 创建一个classmates表，用users的字段

}
