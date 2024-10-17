//Implement a Go function to find the maximum element in an array.
package main

import (
	"fmt"
	"sort"
)

func findmax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
func main() {
	var n int
	fmt.Println("Enter the size of an array:")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])

	}
	max := findmax(arr)
	fmt.Println("Maximum element in an array is:", max)
	sort.Ints(arr)
	fmt.Printf("Maximum element in an array is:%d\n", arr[n-1])
}
