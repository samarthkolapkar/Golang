//Create a function in Go to reverse an array.
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr) - 1
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i] = arr[n-i], arr[i]

	}
	fmt.Println(arr)

}
