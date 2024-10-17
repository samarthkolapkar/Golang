/*Extend (26) to accept URL as a command line argument instead of a hardcoded URL within
the program.*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//This function is to write the content of webpage to the file
func writeToFile(filename string, data []byte) error {
	// Write data to file with read-write permissions for the owner and read-only permissions for others
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	Url := os.Args[1]
	response, err := http.Get(Url) //used to load the url
	if err != nil {
		fmt.Println("Error:", err)

	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	filename := "test.html"
	writeToFile(filename, body)
	fmt.Println(filename)

}
