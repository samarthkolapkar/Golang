package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//it ensure that every time it will generate new number
	var start_number int
	var last_number int
	fmt.Printf("Enter the starting number:")
	fmt.Scan(&start_number)

	fmt.Printf("Enter the ending number:")
	fmt.Scan(&last_number)

	result := rand.Intn(last_number-start_number+1) + start_number
	// result1 := randInRange(start_number, last_number)
	//last_number-start_number+1) + start_number- This makes sure that random number fall in the range

	fmt.Println(result)
}

// package main

// import "fmt"

// func find_max(arr []int) *int {
// 	max := arr[0]
// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i] > max {
// 			max = arr[i]
// 		}
// 	}
// 	return &max
// }
// func main() {
// 	var arr []int
// 	var number_of_element int

// 	fmt.Printf("Enter the size of array:")
// 	fmt.Scan(&number_of_element)
// 	fmt.Printf("Enter the element of array:")
// 	arr = make([]int, number_of_element)
// 	for i := 0; i < number_of_element; i++ {
// 		fmt.Scan(&arr[i])
// 	}
// 	result := find_max(arr)
// 	// if err != nil {
// 	// 	fmt.Println("Error")
// 	// }
// 	fmt.Println(result)
// 	// fmt.Println(&result)

// }
