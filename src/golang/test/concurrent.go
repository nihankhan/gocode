package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)

	go countFruit("Apple", ch)

	ch <- "go man go"

	time.Sleep(time.Second * 5)

	fmt.Println(<-ch)
}

func countFruit(name string, ch chan string) {
	data := <-ch

	fmt.Println(name)
	fmt.Println(data)

	ch <- "stopped"
}
