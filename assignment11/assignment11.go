/*Extend (9) to support “countodd” and “counteven” operations to respectively print number of
times odd numbers and number of even numbers found within the list. [Expressions, 2
hours]*/

package main

import (
	"fmt"
	"strconv"
)

//function for counting the number of even and odd number count in an array
func count_even_odd(number_array []int) (int, int) {
	var count_even int = 0
	var count_odd int = 0
	for i := 0; i < len(number_array); i++ {
		if number_array[i]%2 == 0 {
			count_even += 1
		} else {
			count_odd += 1
		}
	}
	return count_even, count_odd
}
func main() {
	var number string
	var number_array []int
	fmt.Println("Enter the number:")
	for {
		fmt.Scan(&number)
		if number == "proceed" {
			break
		}
		number, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Enter the valid number!!")
		}
		number_array = append(number_array, number)
		// fmt.Println(number_array)
	}
	result, _ := count_even_odd(number_array)
	fmt.Println("Number of even numbers in an array is:", result)
	fmt.Println("Number of odd numbers in an array is:", result)

}
