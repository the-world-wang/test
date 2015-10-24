package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	db, _ := sql.Open("mysql", "root:123456@/go")
	stmt, _ := db.Prepare("insert into student(age ,name) values(?,?)")
	stmt.Exec(23, "wanghao")

	rows, _ := db.Query("select * from student")
	for rows.Next() {
		var id int
		var age int
		var name string
		rows.Scan(&id, &age, &name)
		fmt.Println(id, age, name)
	}
}
