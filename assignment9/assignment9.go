/*Extend (7) to accept statistical operation instead of “proceed”. Valid values for “<operation>”
are count (to count number of valid numbers), mean to calculate arithmetic mean, min to
calculate minimum value (minimum of all numbers input), max for maximum value (maximum
of all numbers input). [Switch & functions, 4 hours]*/
package main

import (
	"fmt"
	"strconv"
)

//Function for calculating the maximum number from an array
func max_number(array []int) int {
	// sort.Ints(array)
	// // fmt.Println(array)
	// num := array[len(array)-1]
	// return num
	maximum := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > maximum {
			maximum = array[i]
		}
	}
	return maximum

}

//Function to calculating the minimum number from an array
func min_number(array []int) int {
	// sort.Ints(array)
	// result := array[0]
	// return result
	minimum := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < minimum {
			minimum = array[i]
		}
	}
	return minimum

}

//Function to calculating the number of element in the array
func count_number(array []int) {
	fmt.Println("Number of valid numbers are:", len(array))

}

//Function to calculating the mean of an elements
func calculate_mean(array []int) {
	var total int
	for i := 0; i < len(array); i++ {
		total = total + array[i]

	}
	avarage := total / len(array)
	fmt.Println("Mean of numbers is:", avarage)

}
func main() {
	var num string
	var array []int
	fmt.Print("Enter the number:") //as a string input
	for {
		fmt.Scan(&num)
		if num == "proceed" {
			break //user enter proceed then stop scanning the input from the user
		}
		num1, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Enter the valid number!!") //number is not valid
		}
		array = append(array, num1)
		fmt.Println(array)

	}
	var operation string
	fmt.Println("1)-Max\n2)-Min\n3)-Count\n4)-Mean")
	fmt.Print("Choose a operation you want to perform::")
	fmt.Scan(&operation)
	switch operation {
	case "max":
		result := max_number(array)
		fmt.Println("Maximum number is :", result)
	case "min":
		result1 := min_number(array)
		fmt.Println("Minimum number is:", result1)

	case "count":
		count_number(array)

	case "mean":
		calculate_mean(array)

	default:
		fmt.Println("Enter valid operation!!")

	}
}
