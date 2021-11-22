package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Static/")))
	//http.Handle("/", http.FileServer(http.Dir("/.Static")))

	fmt.Println("Http Server Running...")

	log.Fatal(http.ListenAndServe(":8005", nil))
}
