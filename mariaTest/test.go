package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	return
}

func main() {
	db, err := sql.Open("mysql", "pi1:gavin@/gatorloop?charset=utf8")
	checkErr(err)
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	stmt, err := db.Prepare("INSERT speed SET value=?, time=?")
	checkErr(err)
	
	res, err := stmt.Exec(240.5000, "2016-06-07 01:33:20")
	checkErr(err)
	
	id, err := res.LastInsertId()
	checkErr(err)
	
	fmt.Println(id)
}
