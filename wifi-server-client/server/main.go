package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Client struct {
	client_id int
	conn      net.Conn
}

func sendMessage(client_id int, conn net.Conn, coordinates [4]string) {
	fmt.Println([]byte(coordinates[client_id]))
	num, err := conn.Write([]byte(coordinates[client_id]))
	fmt.Println("num: ", num)
	handleError(err)
}

func broadCast(listOfClients []Client, coordinates [4]string) {
	fmt.Println("broadcasting")
	fmt.Println(listOfClients)
	for _, client := range listOfClients {
		fmt.Println(coordinates[client.client_id])
		sendMessage(client.client_id, client.conn, coordinates)
	}
}

func handleCommand(commandChannel chan string, coordinates [4]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your input")
	message, err := reader.ReadString('\n')
	handleError(err)

	message = strings.TrimSpace(message)

	if message == "broadcast" {
		broadCast(listOfClients, coordinates)
	} else {
		id, err := strconv.Atoi(message)
		handleError(err)

		var client Client
		for _, c := range listOfClients {
			if c.client_id == id {
				client = c
			}
		}

		sendMessage(id, client.conn, coordinates)
	}
}

var listOfClients []Client

func main() {
	listener, err := net.Listen("tcp", ":8080")
	var coordinates = [...]string{
		"1,1,1\n",
		"2,2,2\n",
		"3,3,3\n",
		"4,4,4,4\n",
	}
	handleError(err)
	fmt.Println("TCP server running at port 8080")
	defer listener.Close()

	counter := 0
	commandChannel := make(chan string)

	go func() {
		for {
			handleCommand(commandChannel, coordinates)
		}
	}()

	for {
		conn, err := listener.Accept()
		newClient := Client{
			client_id: counter,
			conn:      conn,
		}
		listOfClients = append(listOfClients, newClient)
		fmt.Println("Connected to client: number of clients ", len(listOfClients))
		counter += 1
		handleError(err)
	}
}
