/*Write a program to accept two dates (any of the supported period) and print an output
whether date1 and date2 are equal, date1 is earlier than date2 or date1 is later than date2*/

package main

import (
	"fmt"
	"time"
)

//This function will check the date is equal,comes after the date1 or before the date1
func check_date(date1, date2 time.Time) (output string) {
	if date1.Equal(date2) {
		output = "Both dates are equal"
	} else if date1.Before(date2) {
		output = "Date 1 is earlier than Date 2."
	} else {
		output = "Date 1 is later than Date 2."
	}
	return output
}

func main() {
	var first_date string
	var second_date string
	fmt.Println("Enter the date as per the format(YYYY-MM-DD)")
	fmt.Printf("Enter the first date:")
	fmt.Scan(&first_date)

	date1, err := time.Parse("2006-01-02", first_date)
	if err != nil {
		fmt.Println("Error while parsing the date!!")

	}

	fmt.Printf("Enter the second date:")
	fmt.Scan(&second_date)

	date2, err := time.Parse("2006-01-02", second_date)
	// fmt.Println(date1, date2)
	if err != nil {
		fmt.Println("Error while parsing the date!!")

	}

	check_date(date1, date2)

}
