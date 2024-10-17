/*Write a program to prompt for two whole positive numbers -- “num1” and “num2”. Print
multiplication table for the number e.g. for num1=3 and num2=20, output will be “3 * 1 = 3\n3
* 2 = 6\n … \n3 * 20 = 60”. [For loop, 6 hours]
*/
package main

import (
	"fmt"
)

//function to check the number is whole
func checkwhole(num1 int, num2 int) error {
	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Number is not a whole number!!")
	}
	return nil
}

//function to print the multiplication table
func multiplicationtable(num1 int, num2 int) {
	for i := 1; i <= num2; i++ {
		fmt.Printf("%d*%d = %d\n", num1, i, num1*i)
	}
	return

}
func main() {
	var num1 int
	var num2 int
	fmt.Print("Enter the first number::")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number:")
	fmt.Scan(&num2)
	err := checkwhole(num1, num2)
	if err == nil {
		multiplicationtable(num1, num2)

	}
}
