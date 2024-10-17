/*Write a program to take the names of candidates as input. Prompt user to keep entering new
names till user enters “done”. Once a list of names are accepted, prompt the user for a
search pattern (regex). Output shall be a list of all names where the search pattern exists.
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
)

//This function is for searching the name of the candidate
func Search_name(candidate_slice []string, regExp string) bool {
	// matching_count := 0
	// //complies into the object for further operations here matching operation is perform
	// regExp1, err := regexp.Compile(regExp)
	// //if there error while compiling it will print the error
	// if err != nil {
	// 	fmt.Println("Error while compiling the regular expression!!")
	// }
	// for _, name := range candidate_slice {
	// 	if regExp1.MatchString(name) {
	// 		matching_count += 1
	// 	}
	// }
	// return matching_count
	////complies into the object for further operations here matching operation is perform
	regExp1, err := regexp.Compile(regExp)
	if err != nil {
		fmt.Println("Error while compiling regular expression:", err)

	}

	for _, name := range candidate_slice {
		if regExp1.MatchString(name) {
			// fmt.Printf("<%s> found\n", regExp)
			return true

		}
		// fmt.Printf("<%s> not found\n", regExp)

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
	var regExp string
	fmt.Printf("Enter the regular expression:")
	fmt.Scan(&regExp)
	if Search_name(candidate_slice, regExp) {
		fmt.Printf("<%s> Found", regExp)
	} else {
		fmt.Printf("<%s> not found\n", regExp)
	}

}
