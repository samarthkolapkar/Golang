package main

import "fmt"

func main() {
	// var slice1 []int
	slice1 := make([]int, 0)
	var roll_no int
	for {
		fmt.Scan(&roll_no)
		if roll_no == 0 {
			break
		}
		slice1 = append(slice1, roll_no)
	}
	fmt.Println(slice1)
	for i := range slice1 {
		if slice1[i] == 5 {
			fmt.Println("found")
		} else {
			fmt.Println("Not Found!!")
			break
		}
	}

	// slice1[2] = 40
	// slice1 = append(slice1, 76)
	// slice1 = append(slice1, 43)
	// fmt.Println(slice1)
}
