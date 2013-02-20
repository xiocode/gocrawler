package gocrawler

import (
	"database/sql"
	"fmt"
	. "gocrawler/utils"
	//_ "github.com/Go-SQL-Driver/MySQL"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/xio?charset=utf8")
	CheckErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT tb_golang_test SET title=?,departname=?,created=?")
	CheckErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	CheckErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	CheckErr(err)

	res, err = stmt.Exec(id)
	CheckErr(err)

	affect, err = res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)

	db.Close()

}

