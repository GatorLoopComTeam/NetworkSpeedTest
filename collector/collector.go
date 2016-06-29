package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.101:" + os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	s := 0
	r := 0
	for {
		fmt.Fprintf(conn, "250.5000\n")
		s = s + 1
		fmt.Println("Packets Sent: ", s)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if message == "ack\n" {
			r = r + 1
			fmt.Println("Packets Received: ", r)
		}
	}
}
