package main

import (
	"net"
	"fmt"
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func checkErr(err error) {
	return
}

func main() {
	//connect to maria database
        db, err := sql.Open("mysql", "pi1:gavin@/gatorloop?charset=utf8")
        checkErr(err)
        defer db.Close()

        var version string
        db.QueryRow("SELECT VERSION()").Scan(&version)
        fmt.Println("Connected to:", version)

	//listen to incoming connection
	ln, err := net.Listen("tcp", ":30000")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
	}

	counter := 0
	startTime := time.Now()
	//continue to receive messages indefinitely
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			fmt.Println(counter)
			fmt.Println("time since start: " + time.Since(startTime).String())
			i, err := conn.Write([]byte("ack\n"))
			if err != nil {
				fmt.Println(err)
			}
			
			stmt, err := db.Prepare("INSERT speed SET value=?, time=NOW()")
                        checkErr(err)

                        i, err = strconv.Atoi(message)
                        res, err := stmt.Exec(i)
                        checkErr(err)
                        _, err = res.LastInsertId()
                        checkErr(err)

			counter = counter + 1
		}
	}
}
