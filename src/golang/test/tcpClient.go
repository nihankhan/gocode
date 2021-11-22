package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Test to send: ")

		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n")

		msg := bufio.NewReader(conn)

		status, _ := msg.ReadString('\n')

		fmt.Print("Message from server: " + status)
	}
}
