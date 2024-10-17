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

//Search the candidate from the map
func Search_candidate(candidates_list map[string]bool, search_name string) bool {

	for value := range candidates_list {
		if strings.ToLower(search_name) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

//This function is get the list of user and store it into map
func getUseInput(candiates_list map[string]bool) {
	fmt.Println("Enter the names of the candidates:")
	for {
		var candidate_name string
		fmt.Scan(&candidate_name)
		//if user will enter the name in upper case convert it into lower case
		if strings.ToLower(candidate_name) == "done" {
			break
		}
		//convert candidate names to lower and append it to candidates list
		candiates_list[strings.ToLower(candidate_name)] = true

	}
	return
}

func main() {
	var search_name string
	candiates_list := make(map[string]bool)

	getUseInput(candiates_list)
	fmt.Printf("Enter the candidate name you want to search:")
	fmt.Scan(&search_name)
	result := Search_candidate(candiates_list, strings.ToLower(search_name))
	if result == true {
		fmt.Printf("<%s> Exist\n", search_name)
	} else {
		fmt.Printf("<%s> not Exist\n", search_name)
	}

}
