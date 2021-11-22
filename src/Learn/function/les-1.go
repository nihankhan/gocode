package main

import (
	"fmt"
)

func main() {
	var numOne, numTwo int

	fmt.Println("Enter first number: ")
	fmt.Scan(&numOne)

	fmt.Println("Enter second number: ")
	fmt.Scan(&numTwo)

	result, nOne, _ := addNumbers(numOne, numTwo)

	fmt.Printf("Result: %d =  %d\n", nOne, result)
}

func addNumbers(numOne, numTwo int) (addedResult int, nOne int, nTwo int) {
	addedResult = numOne + numTwo

	return
}
