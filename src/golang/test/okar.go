package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://okar.ml")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response Status : ", resp.Status)

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(data))
}
