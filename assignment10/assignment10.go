/*Extend (8) to support “sort” operation. Use an in-built function call for sorting numbers. [Core
Classes & array operations, 6 hours]*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

//Function to sort the numbers
func sort_the_numbers(array_number []int) []int {
	sort.Ints(array_number)

	return array_number

}
func main() {
	var number string
	var array_number []int
	fmt.Print("Enter the number:")
	for {
		fmt.Scan(&number)
		if number == "proceed" {
			break
		}
		number, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Enter the valid number!!")
		}
		array_number = append(array_number, number)
	}
	result := sort_the_numbers(array_number)
	fmt.Println(result)

}
