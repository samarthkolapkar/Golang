package main

import (
	"fmt"
	"os"
	"strings"
)

//this function will ignore the common words
func ignore_commonwords(word string) bool {

	common_words := map[string]bool{"a": true, "an": true, "the": true}
	return common_words[word] || len(word) == 1
}

//this funnction will count the occurance of each word
func ocurrance_words(read []byte) {
	words := strings.Fields(strings.ToLower(string(read)))
	wordCount := make(map[string]int)
	for _, word := range words {
		if !ignore_commonwords(word) {
			wordCount[word]++
		}

	}

	// Print the result
	// fmt.Println("Word Occurrences")
	for word, count := range wordCount {
		fmt.Printf("%s %d\n", word, count)
	}
}

//This function is used to read the file
func readfile(filename string) {
	read, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while readin the file")
	}
	ocurrance_words(read)
}
func main() {
	filename := os.Args[1]
	readfile(filename)
	//common_words := []string{"a", "an", "the"}

}
