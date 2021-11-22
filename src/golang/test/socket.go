package main

import (
	"fmt"
	"net"
)

func main() {
	lsn, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	defer lsn.Close()

	for {
		conn, err := lsn.Accept()

		if err != nil {
			panic(err)
		}

		fmt.Println("New Connection")

		go handler(conn)

	}
}

func handler(conn net.Conn) {
	for {
		data := "Hello Client"
		conn.Write([]byte(data))
	}
	conn.Close()
}
