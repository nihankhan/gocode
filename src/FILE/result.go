package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	roll  int
	num   int
	marks []int
)

func main() {
	file, err := os.Open("notes.txt")

	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		//fmt.Scanf("%d %d", &roll, &num)
		fmt.Println(s.Text())

		//fmt.Printf("%v %v\n", roll, num)

		//marks[roll-1] += num;
	}
}
