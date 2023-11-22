package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL驱动
)

func main() {
	// 打开数据库连接
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/shop")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 检查数据库连接
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the database")

	// 执行查询并处理结果
	rows, err := db.Query("SELECT id, username FROM user where id= ?", 1)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 遍历查询结果
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}

	//  使用事务
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}

	res, err := tx.Exec("UPDATE user SET username = ? WHERE id = ?", "changed11", 1)
	affected, err := res.RowsAffected()
	fmt.Println(affected)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	_, err = tx.Exec("DELETE FROM user WHERE id = ?", 2)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}

}
