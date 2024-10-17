//Implement a program to check if an array is sorted in ascending order.
package main

import (
	"fmt"
	"sort"
)

func sorted(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			//if any element in the array is greater than its preceeding element then return false
			return 0
		}
	}
	return 1
}
func main() {
	arr := []int{2, 3, 5, 1, 5}
	sort.Ints(arr)
	result := sorted(arr)
	if result == 1 {
		fmt.Println("Sorted array!!")
	} else {
		fmt.Println("Not sorted array!!")
	}
	// if issorted(arr) {
	// 	fmt.Println("Sorted array!!")
	// } else {
	// 	fmt.Println("Not sorted!!")
	// }
}
