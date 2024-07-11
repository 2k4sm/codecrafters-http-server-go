package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)

	request, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading data from connection : ", err.Error())
	}

	requestSlice := strings.Split(request, " ")

	if requestSlice[1] != "/" {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
