package modules

import (
	"fmt"
	"gin-demo/dao"
)

// 这里的首字母可以不用大写，Scan 方法会根据 字段的指针 进行赋值，而不是通过 Go 反射匹配字段名
type User struct {
	id   int    `db:"id"`
	age  int    `db:"age"`
	name string `db:"name"`
}

// 查询单行
func QueryRowDemo() {
	sqlStr := "select id, name, age from user where id = ?"
	var user User
	// 这里后面必须带有Scan，不然会导致后面的无法查询,除非设置最大连接数
	err := dao.DB.QueryRow(sqlStr, 1).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Println("queryRowDemo err:", err)
		return
	}
	fmt.Println("query success!", user)
}

// 查询多行
func QueryMuitiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := dao.DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Println("queryMuitiRowDemo err:", err)
		return
	}
	// 持有连接必须关闭
	defer rows.Close()

	// 	循环获取数据
	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			fmt.Println("queryMuitiRowDemo err:", err)
			return
		}
		fmt.Println("queryMuitiRowDemo:", user)
	}
}

// 插入
func InsertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := dao.DB.Exec(sqlStr, "zhangsan", 18)
	if err != nil {
		fmt.Println("insertRowDemo err:", err)
		return
	}
	id, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Println("insertRowDemo err:", err)
		return
	}
	fmt.Println("insertRowDemo:", id)
}

// 更新数据
func UpdateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := dao.DB.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func DeleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := dao.DB.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// sql语句预处理
// 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
// 避免SQL注入问题。
func PrepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例
func PrepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

// sql注入
func SqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name = '%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u User
	err := dao.DB.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Println("query success", u)
}

// 事务
func TransactionDemo() {
	// 开始事务
	tx, err := dao.DB.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}
	defer tx.Rollback()

	sqlStr1 := "insert into user(name, age) values (?,?)"
	ret1, err := tx.Exec(sqlStr1, "wangwu", 18)
	if err != nil {
		tx.Rollback()
		fmt.Printf("insert failed111, err:%v\n", err)
		return
	}
	affected1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("insert failed222, err:%v\n", err)
		return
	}

	sqlStr2 := "update user set age=? where id=?"
	ret2, err := tx.Exec(sqlStr2, 20, 6)
	if err != nil {
		tx.Rollback()
		fmt.Printf("update failed333, err:%v\n", err)
		return
	}
	affected2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("udate failed444, err:%v\n", err)
		return
	}

	if affected1 == 1 && affected2 == 1 {
		// 提交事务
		fmt.Println("提交事务")
		tx.Commit()
	} else {
		tx.Rollback()
		fmt.Println("回滚了")
	}

	fmt.Println("exex tran success.")
}
