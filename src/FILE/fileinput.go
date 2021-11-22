package main

import (
	"fmt"
	"os"
)

func main() {
	os.Stdin, _ = os.OpenFile("input.txt", os.O_RDONLY, 0666)

	// if err != nil {
	// 	panic(err)
	// }

	data := make([]byte, 100)
	count, err := os.Stdin.Read(data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %d bytes: %q\n", count, data[:count])

}
