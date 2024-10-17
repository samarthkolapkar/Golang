// Write a program to calculate the average of all elements in an array.
package main

import "fmt"

func main() {
	var n int
	var total int = 0
	var average float64
	fmt.Println("Enter the size of an array:")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])

	}
	for i := 0; i < n; i++ {
		total = total + arr[i]
	}
	average = float64(total) / float64(n)
	fmt.Println("Average is:", average)

}
