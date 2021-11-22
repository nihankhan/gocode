package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/home/nihan/Nihan/MySQL/MySQL Tutorial for Beginners [Full Course]_HIGH.webm")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Printf("File size : %.2f MB\n", float64(stat.Size()/1024/1024))
	fmt.Println("File Name :", stat.Name())
	fmt.Println("File Mode :", stat.Mode())
	fmt.Println("File Last Modified :", stat.ModTime())
	fmt.Println("Is Directory :", stat.IsDir())

}
