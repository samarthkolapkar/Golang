package main

import (
	"fmt"
)

func main() {
	var rows int
	var cols int
	var char string
	// var matrix [][]int
	fmt.Printf("Enter the character to be display:")
	fmt.Scan(&char)
	fmt.Printf("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Printf("Enter the number of columns:")
	fmt.Scan(&cols)
	//print the character in rows and columns enter by user
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%s ", char)
		}
		fmt.Println()
	}
	// matrix := make([][]int, rows)
	// for i := range matrix {
	// 	matrix[i] = make([]int, cols)
	// }
	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < cols; j++ {
	// 		fmt.Printf("Enter the number at [%d][%d]:", i, j)
	// 		fmt.Scan(&matrix[i][j])
	// 	}
	// }

	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < cols; j++ {
	// 		fmt.Printf("%d ", matrix[i][j])
	// 	}
	// 	fmt.Println()
	// }

}
