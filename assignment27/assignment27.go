/*Extend (26) to use URL classes instead of a browser to download a file. Rest of the
functionalities will be the same as the previous problem.*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//function to write the content to the file
func writeToFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	parsedUrl, err := url.Parse("https://pkg.go.dev/std")
	if err != nil {
		fmt.Println("Error while parsing the url!!")
	}

	response, err := http.Get(parsedUrl.String())
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	filename := "first.html"
	writeToFile(filename, body)
	fmt.Println(filename)
}
