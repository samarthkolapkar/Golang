package main

import "fmt"

func add(x int, y int) int {
	out := x + y
	return out
}
func main() {
	var num1 int = 4
	var num2 int = 5
	// var result int = num1 + num2
	result := add(num1, num2)
	fmt.Println(result)
}
