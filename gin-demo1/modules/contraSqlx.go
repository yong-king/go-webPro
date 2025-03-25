package modules

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gin-demo/dao"
	"github.com/jmoiron/sqlx"
	"strings"
)

type User1 struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

func QueryRaw() {
	sqlStr := "select id, name, age from user where id = ?"
	var u User1
	err := dao.DBs.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("query success, id : %d, name : %s\n", u.ID, u.Name)
}

func QueryMultiRows() {
	sqlStr := "select id, name, age from user where id > ?"
	var u []User1
	err := dao.DBs.Select(&u, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for _, u := range u {
		fmt.Printf("query success, id : %d, name : %s\n", u.ID, u.Name)
	}
}

func InsertRow() {
	sqlStr := "insert into user(name, age) values (?, ?)"
	ret, err := dao.DBs.Exec(sqlStr, "迪丽热巴", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Printf("inserted %d rows\n", affected)
}

func UpdateRow() {
	sqlStr := "update user set name=?, age=? where id=?"
	ret, err := dao.DBs.Exec(sqlStr, "章若楠", 18, 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	fmt.Printf("inserted %d rows\n", affected)
}

func DeleteRow() {
	sqlStr := "delete from user where id=?"
	ret, err := dao.DBs.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	fmt.Printf("deleted %d rows\n", affected)
}

func InsertUser() (err error) {
	sqlStr := "insert into user(name, age) values (:name, :age)"
	_, err = dao.DBs.NamedExec(sqlStr, map[string]interface{}{
		"name": "yk",
		"age":  18,
	})
	return
}

func NameQuery() {
	sqlStr := "select * from user where name=:name"
	rows, err := dao.DBs.NamedQuery(sqlStr, map[string]interface{}{"name": "yk"})
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User1
		err = rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("name:%s, age:%d\n", u.Name, u.Age)
	}

	u := User1{Name: "yk"}
	rows, err = dao.DBs.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User1
		err = rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("name:%s, age:%d\n", u.Name, u.Age)
	}
}

func Tansaction() (err error) {
	tx, err := dao.DBs.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Printf("rollback failed, err:%v\n", err)
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Printf("commit tx successfully\n")
		}
	}()
	sqlStr1 := "Update user set age=20 where id=?"

	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}

// 批量插入
// 自己插入
func BunchInsertUSer(users []*User1) error {
	// 存放（？， ？) 的切片
	valuesString := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		valuesString = append(valuesString, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 执行拼接
	sqlStr := fmt.Sprintf("insert into user(name, age) values %s",
		strings.Join(valuesString, ","))
	_, err := dao.DBs.Exec(sqlStr, valueArgs...)
	return err
}

// 使用sqlx.In实现批量插入
// 前提是需要我们的结构体实现driver.Valuer接口：
func (u User1) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

func BunchInsertUSer2(users []interface{}) error {
	// 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	query, args, _ := sqlx.In("insert into user(name, age) values (?), (?), (?)", users...)
	fmt.Println(query)
	fmt.Println(args)
	_, err := dao.DBs.Exec(query, args...)
	return err
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User1) error {
	_, err := dao.DBs.NamedExec("insert into user (name, age) values (:name, :age)", users)
	return err
}

func QueryByIDs(ids []int) (users []User1, err error) {
	query, args, err := sqlx.In("select name, age from user where id in (?) ORDER BY FIND_IN_SET(id, ?) ", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定。
	// 重新生成对应数据库的查询语句（如PostgreSQL 用 `$1`, `$2` bindvar）
	query = dao.DBs.Rebind(query)
	err = dao.DBs.Select(&users, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = dao.DBs.Rebind(query)

	err = dao.DBs.Select(&users, query, args...)
	return
}
