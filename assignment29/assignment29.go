/*Write a program to accept a filename from command line argument, read a file and print the
number of times each word occurs in a file. Perform case insensitive match while counting
the occurrence of each word. [Hashmap & File operation & String operation, 6 hours]
 Extend the above program to ignore common words (e.g. “the”, “a”, “an”, …) and single letter
words (e.g. “I”).*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// func ReadFile(urlresponse http.Response) {
// 	body, err := ioutil.ReadAll(urlresponse)
// }

//This function is used to count the word occurance in the file
func word_occurance(filename string, search_word string, commonwords []string) (int, error) {
	var word_count int
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening the file")
	}
	defer file.Close()

	// for word := range file {
	// 	if word == search_word {
	// 		word_count += 1
	// 	}
	// }
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && n == 0 {
			break
		}
		lines := strings.Split(string(buffer[:n]), "\n")
		for _, line := range lines {
			words := strings.Fields(line)
			for _, word := range words {
				if !ignore_commonwords(word, commonwords) && len(word) > 1 && strings.ToLower(word) == strings.ToLower(search_word) {
					word_count += 1
				}
			}
		}
	}

	return word_count, nil
}

//this function is for to ignore the common words in the file
func ignore_commonwords(word string, commonwords []string) bool {
	for _, common := range commonwords {
		if word == common {
			return true
		}
	}
	return false
}

//this function is used printing the output in the file
func output(filename string) {
	fmt.Println(filename)
}

//This fucntion is used to write the content of the page into the file
func writeTofile(filename string, data []byte) error {
	//take the filename as a command line argument
	// filename := os.Args[1]
	//write the content of the page into the file
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error while writing into the file!!")

	}
	output(filename)
	return nil

}

//this function is used to get the url and parse the body of the response url and
// func geturl(Url string) []byte {
// 	urlresponse, err := http.Get(Url)
// 	if err != nil {
// 		fmt.Println("Error while getting the url")
// 	}
// 	defer urlresponse.Body.Close()

// 	body, err := ioutil.ReadAll(urlresponse.Body)
// 	return body

// }
func main() {
	Url := "https://pkg.go.dev/std"
	var search_word string

	filename := os.Args[1]

	urlresponse, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error while getting the url")
	}
	defer urlresponse.Body.Close()

	body, err := ioutil.ReadAll(urlresponse.Body)
	//writeTofile function call
	writeTofile(filename, body)
	fmt.Printf("Enter the word you want search:")
	fmt.Scan(&search_word)
	//word_occurance function call
	commonwords := []string{"a", "an", "the"}
	result, err := word_occurance(filename, search_word, commonwords)
	fmt.Printf("Count of %s in the file is %d\n", search_word, result)

}
