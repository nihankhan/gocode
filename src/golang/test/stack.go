package main

import (
	"fmt"
)

func main() {
	var tt int
	var input []rune

	fmt.Scan(&tt)

	for i := 0; i < tt; i++ {

		fmt.Scan(input)

		if isValid(input) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isValid(input []rune) bool {
	var stack []rune
	var last rune

	for i := 0; i < len(input); i++ {
		if input[i] == '(' || input[i] == '{' || input[i] == '[' {
			stack = append(stack, input[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			last = stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			if last == '(' && input[i] != ')' {
				return false
			}

			if last == '{' && input[i] != '}' {
				return false
			}

			if last == '[' && input[i] != ']' {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
