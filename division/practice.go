//Question-Division of two number if the denominator is zero then it will throw an error

package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Denominator must not be zero!!")
	}
	return a / b, nil
}
func main() {
	var numerator int
	var denominator int
	fmt.Scan(&numerator)
	fmt.Scan(&denominator)
	result, err := division(numerator, denominator)
	if err != nil {
		fmt.Println("Error-", err)
		return
	}
	fmt.Println("Division:", result)
}
