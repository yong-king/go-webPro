package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

type Result struct {
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:youngking98@tcp(127.0.0.1:3307)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	// 添加
	//u1 := User{Name: "jinzhu", Age: 10}
	//u2 := User{Name: "ykk", Age: 20}
	//db.Create(&u1)
	//db.Create(&u2)

	// 查询
	// 一般查询
	//var u User
	//u2 := new(User)
	//db.First(&u) // 根据主键查询第一条，主键必须为int
	//fmt.Printf("%+v\n", u)
	//db.Take(&u2) // 根据主键随机获取一条记录
	//fmt.Printf("%+v\n", u2)
	//db.Last(&u)
	//fmt.Printf("%+v\n", u)// 根据主键获取最后一条记录
	//var u3 []User
	//// 获取所有记录
	//db.Find(&u3)
	//for _, v := range u3 {
	//	fmt.Println("%+v\n", v)
	//}
	//db.First(&u, 2)  // 查询指定的某条记录(仅当主键为整型时可用)
	//fmt.Printf("%+v\n", u)

	// 获取符合条件的第一条
	//db.Where("name = ?", "ykk").First(&u)
	//fmt.Printf("%+v\n", u.Age)

	// 获取符合条件的所有
	//db.Where("name = ?", "ykk").Find(&u3)
	//for _, u3 := range u3 {
	//	fmt.Printf("%+v\n", u3.Age)
	//}

	// 筛选出name != ykk 的字段
	//var u User
	//db.Where("name <> ?", "ykk").Find(&u)
	//fmt.Printf("%+v\n", u)

	// 类似于y的
	//var user []User
	//db.Where("name Like ?", "%y%").Find(&user)
	//for _, user := range user {
	//	println(user.Name)
	//}

	//db.Where("updated_at > ?", "2025-03-15").Find(&user)
	////// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	//fmt.Printf("%+v\n", user)

	//db.Where("created_at BETWEEN ? AND ?", "2025-03-15", "2025-03-17").Find(&user)
	//
	// struct
	//db.Where(&User{Age: 10}).Find(&user)
	//fmt.Printf("%+v\n", user)

	//db.Where(map[string]interface{}{"name": "ykk"}).Find(&user)
	//fmt.Printf("%+v\n", user)

	// 主键切片
	//db.Where([]int64{3, 4}).Find(&user)
	//fmt.Printf("%+v\n", user)

	// "" 和 0 不会被当作查询条件，除非指针或者sql.NullInt。。。指针或实现 Scanner/Valuer 接口来避免这个问题

	// 排除
	//db.Not("name", "jinzhu").First(&user)
	//fmt.Println(user)

	// or条件
	//db.Where("name = 'jinzhu'").Or(User{Age: 20}).Find(&user)
	//fmt.Println(user)

	// 内联条件
	//db.Find(&user, "name = ?", "jinzhu")
	//fmt.Println(user)

	// 为查询 SQL 添加额外的 SQL 操作
	//FOR UPDATE 作用：
	//锁定查询到的行，防止其他事务修改它（行级锁）。
	//适用于并发控制，确保只有当前事务可以修改数据，其他事务必须等待锁释放。
	//FOR UPDATE 确保查询到的行在当前事务提交之前不会被其他事务修改。
	//db.Set("gorm:query_option", "FOR UPDATE").First(&user, 3)
	//// SELECT * FROM users WHERE id = 10 FOR UPDATE;

	// FirstOrInit
	// 如果找到了，就将查询到的数据赋值给 user。
	// 如果找不到，则初始化一个 user 结构体，并填充 Name: "Jinzhu"，但不会插入数据库。
	//var user User
	//db.FirstOrInit(&user, User{Name: "dddd"})

	//如果记录未找到，将使用参数初始化 struct. attrs
	//db.Where(User{Name: "dlrb"}).Attrs(User{Age: 25}).FirstOrInit(&user)
	//fmt.Println(user)
	// 不管记录是否找到，都将参数赋值给 struct. assign
	//db.Where(User{Name: "kkk"}).Assign(User{Age: 18}).FirstOrInit(&user)
	//fmt.Println(user)
	//FirstOrCreate  获取匹配的第一条记录, 否则根据给定的条件创建一个新的记录 (仅支持 struct 和 map 条件)
	//db.FirstOrCreate(&user, User{Name: "ykk"})
	//fmt.Println(user)
	// attrs   未找到，插入，找到了，不插入，都给结构体
	// assign	未找到，插入，找到了，更新，都给结构体
	//db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
	//db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)

	// Select，指定你想从数据库中检索出的字段，默认会选择全部字段。
	//db.Select("id").Where("name = ?", "ykk").First(&user)
	//fmt.Println(user)

	// db.Order("age desc, name").Find(&users)
	////// SELECT * FROM users ORDER BY age desc, name;
	//
	//// 多字段排序
	//db.Order("age desc").Order("name").Find(&users)
	////// SELECT * FROM users ORDER BY age desc, name;
	//
	//// 覆盖排序
	//db.Order("age desc").Find(&users1).Order("age", true).Find(&users2)
	////// SELECT * FROM users ORDER BY age desc; (users1)
	////// SELECT * FROM users ORDER BY age; (users2)

	// 子查询
	//db.Where("amount > ?", db.Table("oreders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	//db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));

	// 最大记录
	//db.Limit(3).Find(&user)
	//fmt.Println(user)

	// -1 取消 Limit 条件
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	//// SELECT * FROM users LIMIT 10; (users1)
	//// SELECT * FROM users; (users2)

	//Offset，指定开始返回记录前要跳过的记录数
	//db.Offset(3).Find(&users)
	//// SELECT * FROM users OFFSET 3;

	// -1 取消 Offset 条件
	//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	//// SELECT * FROM users OFFSET 10; (users1)
	//// SELECT * FROM users; (users2)

	// Count，该 model 能获取的记录总数。
	//db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count)
	////// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
	////// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)
	//
	//db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	////// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)
	//
	//db.Table("deleted_users").Count(&count)
	////// SELECT count(*) FROM deleted_users;
	//
	//db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	////// SELECT count( distinct(name) ) FROM deleted_users; (count)

	//Pluck，查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
	var user User
	var ages []string
	db.Find(&user).Pluck("name", &ages)
	for _, age := range ages {
		fmt.Println(age)
	}

	//var user User
	//var result Result
	//Scan，扫描结果至一个 struct.
	//db.Table("users").Select("name, age").Where("name = ?", "ykk").Scan(&result)
	//fmt.Println(result)
	//// 原生 SQL
	//db.Raw("SELECT name, age FROM users WHERE name = ?", "ykk").Scan(&result)

}
