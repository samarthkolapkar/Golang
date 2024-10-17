// package main
package main

import (
	"fmt"
	"time"
)

func calculateDifference(date1, date2 time.Time) (years, days, minutes int) {
	// Calculate the difference
	difference := date2.Sub(date1)

	// Calculate years
	years = int(difference.Hours() / 24 / 365)

	// Calculate remaining days after considering the years
	remainingDays := difference.Hours()/24 - float64(years*365)

	// Calculate days
	days = int(remainingDays)

	// Calculate remaining minutes after considering the years and days
	remainingMinutes := difference.Minutes() - float64(years*365*24*60) - float64(days*24*60)

	// Calculate minutes
	minutes = int(remainingMinutes)

	return years, days, minutes
}

func main() {
	var firstDate, secondDate string
	fmt.Println("Enter the dates in the format YYYY-MM-DD")
	fmt.Printf("Enter the first date: ")
	fmt.Scan(&firstDate)
	date1, err := time.Parse("2006-01-02", firstDate)
	if err != nil {
		fmt.Println("Error while parsing the first date!")
		return
	}

	fmt.Printf("Enter the second date: ")
	fmt.Scan(&secondDate)
	date2, err := time.Parse("2006-01-02", secondDate)
	if err != nil {
		fmt.Println("Error while parsing the second date!")
		return
	}

	years, days, minutes := calculateDifference(date1, date2)
	fmt.Printf("Difference: %d years, %d days, %d minutes\n", years, days, minutes)
}
