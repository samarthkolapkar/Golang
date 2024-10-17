package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url, err := http.Get("https://golang.org/pkg")
	if err != nil {
		log.Fatal("Error while getting the response from the url!!")
	}
	defer url.Body.Close()
	bs := make([]byte, 1024)
	url.Body.Read(bs)
	fmt.Println(string(bs))
	ioutil.WriteFile("second.txt", bs, 0644)
}
