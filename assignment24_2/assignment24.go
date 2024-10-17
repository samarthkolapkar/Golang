/*Write a program to take the names of candidates as input. Prompt user to keep entering new
names till user enters “done”. Once a list of names are accepted, prompt the user for a
name. Output shall be “<name> exists” or “<name> does not exist”. A name exists if the
name exactly matches one of the names provided earlier. Use case insensitive match for
comparison.*/

package main

import (
	"fmt"
	"strings"
)

//This function is for searching the name of the candidate
func Search_name(candidate_slice []string, search_name string) bool {
	// var search_name string
	// fmt.Scan(&search_name)
	for _, name := range candidate_slice {
		if strings.ToLower(name) == strings.ToLower(search_name) {
			return true
		}
	}
	return false
}

//This function is used to get the userinput and store it into the slice
func getUserInput() []string {
	// candidate_slice = make([]string)
	var candidate_slice []string
	var candidate_name string
	fmt.Println("Enter the names of candidates:")
	for {
		fmt.Scan(&candidate_name)
		if strings.ToLower(candidate_name) == "done" {
			break
		}
		candidate_slice = append(candidate_slice, candidate_name)

	}
	return candidate_slice
}

func main() {
	var candidate_slice = getUserInput()
	// var candidate_name string
	var search_name string
	// getUserInput(candidate_slice, candidate_name)
	fmt.Println("Enter the name of candidate::")
	fmt.Scan(&search_name)
	// result := Search_name(candidate_slice, search_name)
	if Search_name(candidate_slice, search_name) {
		fmt.Printf("<%s> exist\n", search_name)
	} else {
		fmt.Printf("<%s> not exist\n", search_name)
	}

}
