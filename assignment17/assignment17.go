package main

import (
	"fmt"
	"time"
)

//This function will calculate the difference between the dates in specified format
func calculate_difference(date1, date2 time.Time) (years, days, minutes int) {
	// var difference string

	difference := date2.Sub(date1)
	years = int((difference.Hours() / 24) / 365)

	remainingDays := difference.Hours()/24 - float64(years*365)

	// Calculate days
	days = int(remainingDays)

	// Calculate remaining minutes after considering the years and days
	remainingMinutes := difference.Minutes() - float64(years*365*24*60) - float64(days*24*60)
	minutes = int(remainingMinutes)
	return years, days, minutes
}
func main() {
	// currentTime := time.Now().UTC()
	var first_date string
	var second_date string
	fmt.Println("Enter the date as a format (YYYY-MM-DD)")
	//First date
	fmt.Printf("Enter the first date:")
	fmt.Scan(&first_date)
	date1, err := time.Parse("2006-01-02", first_date)
	if err != nil {
		fmt.Println("Error while parsing the date!")
	}
	//Second date
	fmt.Printf("Enter the second date:")
	fmt.Scan(&second_date)
	date2, err := time.Parse("2006-01-02", second_date)
	if err != nil {
		fmt.Println("Error while parsing the date!")
	}
	// fmt.Println(currentTime)
	years, days, minutes := calculate_difference(date1, date2)
	fmt.Printf("Differnce between date is %d years %d days %d minutes\n", years, days, minutes)
}
