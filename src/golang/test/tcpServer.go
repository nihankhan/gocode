package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Start server.....")

	lsn, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	defer lsn.Close()

	con, err := lsn.Accept()

	if err != nil {
		panic(err)
	}

	for {
		msg, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Println("Message Received:", msg)
	}

}
