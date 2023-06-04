package main

import (
	"net"
  "fmt"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	handleError(err)
}

type Client struct {
  client_id int
  conn net.Conn
}


func sendMessage(client_id int,conn net.Conn,coordinates [4]string){
  conn.Write([]byte(coordinates[client_id]))
}

func broadCast(listOfClients []Client,coordinates [4]string){
  for _,client := range listOfClients {
    fmt.Println(coordinates[client.client_id])
    go sendMessage(client.client_id,client.conn,coordinates)
  }
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
  var listOfClients []Client
  var coordinates = [...]string{
    "1,1,1",
    "2,2,2",
    "3,3,3",
    "4,4,4,4",
  }
	handleError(err)
  fmt.Println("TCP server running at port 8080")
	defer listener.Close()
  counter := 0
	for {
		conn, err := listener.Accept()
    fmt.Println("Connected to client: number of clients ",len(listOfClients))
    listOfClients = append(listOfClients,Client{
      client_id: counter,
      conn: conn,
    })
    counter+=1
    if(len(listOfClients) == 4){
      broadCast(listOfClients,coordinates)
    }
		handleError(err)
    go do(conn)
	}
}
