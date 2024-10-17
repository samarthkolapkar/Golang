package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Function to count the occurrences of a word in a file
func wordOccurrence(filename, searchWord string) (int, error) {
	var wordCount int

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening the file: %v", err)
	}
	defer file.Close()

	// Create a buffer to read each line
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && n == 0 {
			break
		}
		// Split each line into words
		lines := strings.Split(string(buffer[:n]), "\n")
		for _, line := range lines {
			words := strings.Fields(line)
			for _, word := range words {
				if strings.ToLower(word) == strings.ToLower(searchWord) {
					wordCount++
				}
			}
		}
	}

	return wordCount, nil
}

// Function to write data to a file
func writeToFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	fmt.Printf("Data written to file: %s\n", filename)
	return nil
}

func main() {
	// Check if filename argument is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <filename> <search_word>")
		return
	}

	filename := os.Args[1]
	searchWord := os.Args[2]

	url := "https://pkg.go.dev/std"
	// Fetch data from URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while getting the URL:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while reading response body:", err)
		return
	}

	// Write data to file
	if err := writeToFile(filename, body); err != nil {
		fmt.Println("Error while writing to file:", err)
		return
	}

	// Count occurrences of search word in file
	count, err := wordOccurrence(filename, searchWord)
	if err != nil {
		fmt.Println("Error while counting word occurrences:", err)
		return
	}
	fmt.Printf("Number of occurrences of '%s' in file '%s': %d\n", searchWord, filename, count)
}
