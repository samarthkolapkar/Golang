/*Write a program to accept a date and print whether the date falls in a leap year. Accept a
date in any format supported in one of the previous problems.*/
package main

import (
	"fmt"
	"time"
)

//This function will check the date is falling in leap year or not
func leap_or_not(date time.Time) (output string) {
	//Separate the year from the date
	year := date.Year()

	//fmt.Println(year)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		output = "it's a leap year!!"
	} else {
		output = "it is not a leap year!!"
	}
	return output
}
func main() {
	var inputdate string

	fmt.Printf("Enter the date:")
	fmt.Scan(&inputdate)
	date, err := time.Parse("2006-01-02", inputdate)
	if err != nil {
		fmt.Println("Error while parsing the date!!")
	}
	leap_or_not(date)
}
