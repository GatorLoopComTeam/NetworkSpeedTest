package main

import (
	"net"
	"fmt"
	"bufio"
	"time"
	"os"
)

func checkErr(err error) {
	return
}


func main() {
	//listen to incoming connection
	ln, err := net.Listen("tcp", ":" + os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
	}
  defer conn.Close()

	r := 0
  s := 0
	startTime := time.Now()

	//continue to receive messages indefinitely
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
      r = r + 1
      fmt.Println("Packets received: ", r)
			_, err := conn.Write([]byte("ack\n"))
			if err != nil {
				fmt.Println(err)
			}
			s = s + 1
      fmt.Println("Packets sent: ", s)
      fmt.Println("Total Time: ", time.Since(startTime).String())
		}
	}

}
