package main

import (
	"fmt"
)

func main() {
	var a string = "Hello"

	addNumber(&a)

	fmt.Println(a)
}

func addNumber(a *string) {
	*a += " Nihan"
}
