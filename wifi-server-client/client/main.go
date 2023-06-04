package main

import (
	"bufio"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	handleError(err)
	for {
    message,err := bufio.NewReader(conn).ReadString('\n');
		handleError(err)
		println("Received message:",message)
	}
}
